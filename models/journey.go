package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Journey struct {
	JourneyID         string    `json:"journeyId" gorm:"primaryKey"`
	UserID            string    `json:"userId" binding:"required" FK`
	StartTime         time.Time `json:"startTime" binding:"required"`
	FinishTime        time.Time `json:"finishTime" binding:"required"`
	Origin            string    `json:"origin" binding:"required"`
	Destination       string    `json:"destination" binding:"required"`
	DistanceTravelled float64   `json:"distanceTravelled" binding:"required"`
	EmissionSaved     float64   `json:"emissionSaved" binding:"required"`
	Reward            int       `json:"reward" binding:"required"`
}

func (journey *Journey) BeforeCreate(tx *gorm.DB) error {
	journey.JourneyID = uuid.New().String()
	return nil
}
