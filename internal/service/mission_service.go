package service

import (
	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/repository"
)

// MissionService керує логікою місій
type MissionService struct {
	repo *repository.MissionRepository
}

// NewMissionService створює новий сервіс місій
func NewMissionService(repo *repository.MissionRepository) *MissionService {
	return &MissionService{repo: repo}
}

func (s *MissionService) CreateMission(mission *models.Mission) error {
	return s.repo.CreateMission(mission)
}

func (s *MissionService) GetAllMissions() ([]repository.MissionWithTargets, error) {
	return s.repo.GetAllMissions()
}

func (s *MissionService) GetMissionByID(id uint) (*repository.MissionWithTargets, error) {
	return s.repo.GetMissionByID(id)
}

func (s *MissionService) AssignCatToMission(missionID uint, catID *uint) error {
	return s.repo.AssignCatToMission(missionID, catID)
}

func (s *MissionService) CompleteMission(missionID uint) error {
	return s.repo.MarkMissionAsCompleted(missionID)
}

func (s *MissionService) DeleteMission(id uint) error {
	return s.repo.DeleteMission(id)
}
