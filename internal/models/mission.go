package models

import (
	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	CatID     *uint    `json:"cat_id" gorm:"default:NULL"`
	Completed bool     `json:"completed" gorm:"default:false"`
	Targets   []Target `json:"targets" gorm:"foreignKey:MissionID;constraint:OnDelete:CASCADE"`
}
