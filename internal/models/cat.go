package models

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name       string  `json:"name" gorm:"not null"`
	Experience int     `json:"experience" gorm:"not null"`
	Breed      string  `json:"breed" gorm:"not null"`
	Salary     float64 `json:"salary" gorm:"not null"`
}
