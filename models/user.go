package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	UserID          string    `json:"userId" gorm:"primaryKey;default:1"`
	Email           string    `json:"email" gorm:"unique" binding:"required"`
	Username        string    `json:"username" gorm:"unique" binding:"required"`
	Password        string    `json:"password" binding:"required"`
	FirstName       string    `json:"firstName" binding:"required"`
	LastName        string    `json:"lastName" binding:"required"`
	BirthDate       time.Time `json:"birthDate" binding:"required"`
	Age             int       `json:"age"`
	Gender          string    `json:"gender" gorm:"default:null"`
	Job             string    `json:"job" gorm:"default:null"`
	Points          string    `json:"points" gorm:"default:null"`
	VoucherInterest string    `json:"voucherInterest" gorm:"default:null"`
	Domicile        string    `json:"domicile" gorm:"default:null"`
	Education       string    `json:"education" gorm:"default:null"`
	MarriageStatus  string    `json:"marriageStatus" gorm:"default:null"`
	Income          string    `json:"income" gorm:"default:null"`
	Vehicle         string    `json:"vehicle" gorm:"default:null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	println("XXXX BEFORE CREATE XXXX")
	user.UserID = uuid.New().String()
	return nil
}
