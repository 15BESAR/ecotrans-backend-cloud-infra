package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherPurchased struct {
	gorm.Model
	VoucherID   string
	UserID      string
	IsPurchased bool    `default:"false"`
	Voucher     Voucher `gorm:"foreignKey:VoucherID;references:VoucherID"`
	User        User    `gorm:"foreignKey:UserID;references:UserID"`
}

type Voucher struct {
	VoucherID   string `json:"voucherId" gorm:"primaryKey"`
	PartnerID   string `json:"partnerId" binding:"required"`
	PartnerName string `json:"partnerName" binding:"required"`
	VoucherName string `json:"voucherName" binding:"required"`
	VoucherDesc string `json:"voucherDesc" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"imageUrl" default:"https://storage.googleapis.com/voucher-images-2909/jco.jpg"`
	Stock       int    `json:"stock" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

func (voucher *Voucher) BeforeCreate(tx *gorm.DB) error {
	voucher.VoucherID = uuid.New().String()
	return nil
}
