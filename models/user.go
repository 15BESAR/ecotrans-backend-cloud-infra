package models

type UserInputRegis struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInputLogin struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenRefresh struct {
	Token string `json:"token"`
}
