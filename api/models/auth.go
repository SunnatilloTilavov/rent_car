package models

type CustomerLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CustomerLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type CustomerRegisterRequest struct {
	Mail string `json:"mail"`
}



type LoginCustomer struct {
	Login     string  `json:"login"`
	First_name        string  `json:"first_name"`
	Last_name       string  `json:"last_name"`
	Gmail       string  `json:"gmail"`
	Phone     string  `json:"phone"`
	Password string `json:"password"`
	GmailCode string `json:"gmailcode"`
}
