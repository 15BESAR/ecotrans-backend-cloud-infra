package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Voucher struct {
	VoucherId   string `json:"voucherId" gorm:"primaryKey"`
	PartnerId   string `json:"partnerId" binding:"required"`
	VoucherName string `json:"voucherName" binding:"required"`
	VoucherDesc string `json:"voucherDesc" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"imageUrl" binding:"required"`
	Stock       string `json:"stock" binding:"required"`
	Price       string `json:"price" binding:"required"`
}

func (voucher *Voucher) BeforeCreate(tx *gorm.DB) error {
	voucher.VoucherId = uuid.New().String()
	return nil
}
