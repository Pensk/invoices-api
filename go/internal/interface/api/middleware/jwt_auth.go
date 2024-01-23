package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/pensk/invoices-api/internal/domain/repositories"
)

type AuthMiddleware struct {
	userRepo repositories.UserRepository
	logger   *slog.Logger
}

func NewAuthMiddleware(userRepo repositories.UserRepository, logger *slog.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (m *AuthMiddleware) DecodeJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Info("Decoding JWT")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int(claims["userID"].(float64))
			user, err := m.userRepo.GetByID(userID)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			m.logger.Info("User auth", "user_id", user.ID, "company_id", user.CompanyID)

			ctx := context.WithValue(r.Context(), "user_id", user.ID)
			r := r.WithContext(ctx)
			ctx = context.WithValue(r.Context(), "company_id", user.CompanyID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}
