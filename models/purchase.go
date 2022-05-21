package models

import (
	"time"
)

type Purchase struct {
	VoucherID   string    `json:"voucherId" gorm:"primaryKey;"`
	UserID      string    `json:"userId" gorm:"primaryKey"`
	BuyDate     time.Time `json:"buyDate" binding:"required"`
	BuyQuantity int       `json:"buyQuantity" binding:"required"`
}

type PurchaseReceipt struct {
	Purchase
	VoucherStockRemaining int `json:"voucherStockRemaining"`
	UserPointsRemaining   int `json:"userPointsRemaining"`
}
