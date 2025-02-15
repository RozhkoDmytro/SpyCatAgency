package service

import (
	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/repository"
)

// TargetService керує логікою цілей
type TargetService struct {
	repo *repository.TargetRepository
}

// NewTargetService створює новий сервіс цілей
func NewTargetService(repo *repository.TargetRepository) *TargetService {
	return &TargetService{repo: repo}
}

// CompleteTarget позначає ціль як завершену
func (s *TargetService) CompleteTarget(targetID uint) error {
	return s.repo.MarkTargetAsCompleted(targetID)
}

// UpdateTargetNotes оновлює нотатки цілі (якщо вона не завершена)
func (s *TargetService) UpdateTargetNotes(targetID uint, notes []string) error {
	return s.repo.UpdateTargetNotes(targetID, notes)
}

// AddTargetToMission додає нову ціль до існуючої місії (якщо місія не завершена)
func (s *TargetService) AddTargetToMission(target *models.Target) error {
	return s.repo.AddTargetToMission(target)
}

// DeleteTarget видаляє ціль (якщо вона не завершена)
func (s *TargetService) DeleteTarget(targetID uint) error {
	return s.repo.DeleteTarget(targetID)
}
