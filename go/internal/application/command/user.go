package command

type AuthenticateUserCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserResult struct {
	AccessToken string `json:"access_token"`
}
