package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pensk/invoices-api/internal/application/services"
	"github.com/pensk/invoices-api/internal/infra/db"
	"github.com/pensk/invoices-api/internal/interface/api/handler"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	addr, err := strconv.Atoi(os.Getenv("ADDR"))
	if err != nil {
		logger.Error("invalid port - using default", "error", err.Error())
		addr = 8080
	}

	env := os.Getenv("ENV")
	dsn := os.Getenv("DSN")

	sqldb, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer sqldb.Close()

	ur := db.NewUserRepository(sqldb)

	us := services.NewUserService(ur)

	router := chi.NewRouter()
	api := chi.NewRouter()

	router.Mount("/api", api)

	handler.NewUserHandler(api, us)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", addr),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
