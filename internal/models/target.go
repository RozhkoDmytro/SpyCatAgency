package models

import (
	"gorm.io/gorm"
)

// Target represents a target in a mission
type Target struct {
	gorm.Model
	MissionID uint   `json:"mission_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	Country   string `json:"country" gorm:"not null"`
	Notes     string `json:"notes"`
	Completed bool   `json:"completed" gorm:"default:false"`
}
