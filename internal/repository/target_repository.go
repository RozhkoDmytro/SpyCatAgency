package repository

import (
	"encoding/json"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"gorm.io/gorm"
)

// TargetRepository керує запитами до бази даних для цілей
type TargetRepository struct {
	db *gorm.DB
}

// NewTargetRepository створює новий екземпляр репозиторію цілей
func NewTargetRepository(db *gorm.DB) *TargetRepository {
	return &TargetRepository{db: db}
}

// MarkTargetAsCompleted позначає ціль як завершену
func (r *TargetRepository) MarkTargetAsCompleted(targetID uint) error {
	result := r.db.Exec("UPDATE targets SET completed = ? WHERE id = ?", true, targetID)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

// UpdateTargetNotes оновлює нотатки цілі (якщо вона не завершена)
func (r *TargetRepository) UpdateTargetNotes(targetID uint, notes []string) error {
	notesJSON, err := json.Marshal(notes)
	if err != nil {
		return err
	}

	result := r.db.Exec(`
		UPDATE targets 
		SET notes = ?, updated_at = NOW() 
		WHERE id = ? 
		AND completed = FALSE 
		AND mission_id IN (SELECT id FROM missions WHERE completed = FALSE)
	`, notesJSON, targetID)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

// AddTargetToMission додає нову ціль до місії (якщо місія не завершена)
func (r *TargetRepository) AddTargetToMission(target *models.Target) error {
	var mission models.Mission
	if err := r.db.First(&mission, target.MissionID).Error; err != nil {
		return err
	}

	if mission.Completed {
		return gorm.ErrInvalidData
	}

	return r.db.Create(target).Error
}

// DeleteTarget видаляє ціль (тільки якщо вона не завершена)
func (r *TargetRepository) DeleteTarget(targetID uint) error {
	return r.db.Where("id = ? AND completed = false", targetID).Delete(&models.Target{}).Error
}
