package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PartnerLogin struct {
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password" binding:"required"`
}

type Partner struct {
	PartnerID   string    `json:"partnerId" gorm:"primaryKey"`
	PartnerName string    `json:"partnerName" binding:"required"`
	Email       string    `json:"email" gorm:"unique" binding:"required"`
	Username    string    `json:"username" gorm:"unique" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Vouchers    []Voucher `gorm:"foreignKey:PartnerID;references:PartnerID"`
}

func (partner *Partner) BeforeCreate(tx *gorm.DB) error {
	partner.PartnerID = uuid.New().String()
	return nil
}
