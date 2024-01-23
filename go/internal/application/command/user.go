package command

type AuthenticateUserCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserResult struct {
	AccessToken string `json:"access_token"`
}

type CreateUserCommand struct {
	CompanyID int    `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserResult struct {
	AccessToken string `json:"access_token"`
}
