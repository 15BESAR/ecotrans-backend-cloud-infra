package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Voucher struct {
	VoucherId   string `json:"voucherId" gorm:"primaryKey"`
	PartnerId   string `json:"partnerId"`
	VoucherName string `json:"voucherName"`
	VoucherDesc string `json:"voucherDesc"`
	Category    string `json:"category"`
	ImageUrl    string `json:"imageUrl"`
	Stock       string `json:"stock"`
	Price       string `json:"price"`
}

func (voucher *Voucher) BeforeCreate(tx *gorm.DB) error {
	voucher.VoucherId = uuid.New().String()
	return nil
}
