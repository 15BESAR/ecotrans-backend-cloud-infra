package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Partner struct {
	PartnerID   string    `json:"partnerId" gorm:"primaryKey"`
	PartnerName string    `json:"partnerName" binding:"required"`
	Vouchers    []Voucher `gorm:"foreignKey:PartnerID;references:PartnerID"`
}

func (partner *Partner) BeforeCreate(tx *gorm.DB) error {
	partner.PartnerID = uuid.New().String()
	return nil
}
