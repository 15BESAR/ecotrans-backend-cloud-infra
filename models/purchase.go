package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	PurchaseID  string    `json:"purchaseId" gorm:"primaryKey"`
	VoucherID   string    `json:"voucherId"`
	UserID      string    `json:"userId"`
	BuyDate     time.Time `json:"buyDate" binding:"required"`
	BuyQuantity int       `json:"buyQuantity" binding:"required"`
}

type PurchaseReceipt struct {
	Purchase
	VoucherStockRemaining int `json:"voucherStockRemaining"`
	UserPointsRemaining   int `json:"userPointsRemaining"`
}

func (purchase *Purchase) BeforeCreate(tx *gorm.DB) error {
	purchase.PurchaseID = uuid.New().String()
	return nil
}
