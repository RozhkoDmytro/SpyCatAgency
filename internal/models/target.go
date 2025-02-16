package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Target struct {
	gorm.Model
	MissionID uint            `json:"mission_id" gorm:"not null"`
	Name      string          `json:"name" gorm:"not null"`
	Country   string          `json:"country" gorm:"not null"`
	Notes     StringArrayJSON `json:"notes" gorm:"type:jsonb;default:'[]'"` // JSONB масив
	Completed bool            `json:"completed" gorm:"default:false"`
}

type StringArrayJSON []string

func (s StringArrayJSON) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return json.Marshal(s)
}

func (s *StringArrayJSON) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("StringArrayJSON: Scan source is not []byte")
	}
	return json.Unmarshal(bytes, s)
}
