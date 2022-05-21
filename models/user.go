package models

import (
	"time"
)

type TokenRefresh struct {
	Token string `json:"token"`
}

type UserLogin struct {
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"unique" json:"email"`
	Username   string    `gorm:"unique" json:"username" binding:"required"`
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	AgeOfBirth time.Time `json:"ageOfBirth" binding:"required"`
	Sex        string    `json:"sex" binding:"required"`
	Address    string    `json:"address" binding:"required"`
	Occupation string    `json:"occupation" binding:"required"`
	Password   string    `json:"password" binding:"required"`
}
