package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Journey struct {
	JourneyID         string    `json:"journeyId" gorm:"primaryKey"`
	UserID            string    `json:"userId" gorm:"primaryKey"`
	StartDate         time.Time `json:"startDate" binding:"required"`
	EndDate           time.Time `json:"endDate" binding:"required"`
	Origin            string    `json:"origin" binding:"required"`
	Destination       string    `json:"destination" binding:"required"`
	DistanceTravelled float64   `json:"distanceTravelled" binding:"required"`
	EmissionSaved     float64   `json:"emissionSaved" binding:"required"`
	Reward            int       `json:"reward" binding:"required"`
}

func (journey *Journey) BeforeCreate(tx *gorm.DB) error {
	journey.journeyID = uuid.New().String()
	return nil
}
