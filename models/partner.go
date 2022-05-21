package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Partner struct {
	PartnerId   string `json:"partnerId" gorm:"primaryKey"`
	PartnerName string `json:"partnerName" binding:"required"`
}

func (partner *Partner) BeforeCreate(tx *gorm.DB) error {
	partner.PartnerId = uuid.New().String()
	return nil
}
