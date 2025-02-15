package repository

import (
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
	return r.db.Model(&models.Target{}).Where("id = ?", targetID).Update("completed", true).Error
}

// UpdateTargetNotes оновлює нотатки цілі (якщо вона не завершена)
func (r *TargetRepository) UpdateTargetNotes(targetID uint, notes []string) error {
	return r.db.Model(&models.Target{}).
		Where("id = ? AND completed = false", targetID).
		Update("notes", notes).
		Error
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
