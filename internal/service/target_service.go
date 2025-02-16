package service

import (
	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/repository"
)

// TargetService handles target-related logic
type TargetService struct {
	repo *repository.TargetRepository
}

// NewTargetService creates a new service instance
func NewTargetService(repo *repository.TargetRepository) *TargetService {
	return &TargetService{repo: repo}
}

// CompleteTarget marks a target as completed
func (s *TargetService) CompleteTarget(targetID uint) error {
	return s.repo.MarkTargetAsCompleted(targetID)
}

// UpdateTargetNotes updates the notes for a target
func (s *TargetService) UpdateTargetNotes(targetID uint, notes []string) error {
	return s.repo.UpdateTargetNotes(targetID, notes)
}

func (s *TargetService) AddNoteToTarget(targetID uint, note string) error {
	return s.repo.AddNoteToTarget(targetID, note)
}

// AddTargetToMission adds a new target to an existing mission
func (s *TargetService) AddTargetToMission(target *models.Target) error {
	return s.repo.AddTargetToMission(target)
}

// DeleteTarget deletes a target
func (s *TargetService) DeleteTarget(targetID uint) error {
	return s.repo.DeleteTarget(targetID)
}
