package models

import (
	"time"
)

type Purchase struct {
	voucherID   string    `json:"voucherId" gorm:"primaryKey;"`
	UserID      string    `json:"userId" gorm:"primaryKey"`
	BuyDate     time.Time `json:"buyDate" binding:"required"`
	BuyQuantity int       `json:"buyQuantity" binding:"required"`
}
