package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Journeys struct {
	journeyID         string    `json:"journeyID" gorm:"primaryKey"`
	userID            string    `json:"journeyID" gorm:"primaryKey"`
	StartDate         time.Time `json:"start_date" binding:"required"`
	EndDate           time.Time `json:"end_date" binding:"required"`
	Origin            string    `json:"origin" binding:"required"`
	Destination       string    `json:"destination" binding:"required"`
	DistanceTravelled float64   `json:"distance_travelled" binding:"required"`
	EmissionSaved     float64   `json:"emission_saved" binding:"required"`
	Reward            int       `json:"reward" binding:"required"`
}

func (journey *Journeys) BeforeCreate(tx *gorm.DB) error {
	journey.journeyID = uuid.New().String()
	return nil
}
