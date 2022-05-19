package models

type TokenRefresh struct {
	Token string `json:"token"`
}

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}
