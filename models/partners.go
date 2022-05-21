package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Partners struct {
	PartnerId   string `json:"partnerId" gorm:"primaryKey"`
	PartnerName string `json:"partnerName"`
}

func (partner *Partners) BeforeCreate(tx *gorm.DB) error {
	partner.PartnerId = uuid.New().String()
	return nil
}
