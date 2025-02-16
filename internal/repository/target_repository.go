package repository

import (
	"encoding/json"
	"errors"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"gorm.io/gorm"
)

type TargetRepository struct {
	db *gorm.DB
}

func NewTargetRepository(db *gorm.DB) *TargetRepository {
	return &TargetRepository{db: db}
}

func (r *TargetRepository) MarkTargetAsCompleted(targetID uint) error {
	result := r.db.Exec("UPDATE targets SET completed = ? WHERE id = ?", true, targetID)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *TargetRepository) UpdateTargetNotes(targetID uint, notes []string) error {
	notesJSON, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	result := r.db.Exec(`
		UPDATE targets 
		SET notes = ?, updated_at = NOW() 
		WHERE id = ? and deleted_at IS NULL
		AND completed = FALSE 
		AND mission_id IN (SELECT id FROM missions WHERE completed = FALSE)
	`, notesJSON, targetID)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *TargetRepository) AddTargetToMission(target *models.Target) error {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(*) FROM targets 
		WHERE mission_id = ? AND completed = FALSE AND deleted_at IS NULL
	`, target.MissionID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count >= 3 {
		return errors.New("mission already has the maximum number of active targets")
	}

	return r.db.Create(target).Error
}

func (r *TargetRepository) DeleteTarget(targetID uint) error {
	result := r.db.Where("id = ? AND completed = false and deleted_at IS NULL", targetID).Delete(&models.Target{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
