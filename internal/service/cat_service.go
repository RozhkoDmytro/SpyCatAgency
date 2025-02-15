package service

import (
	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/repository"
)

type CatService struct {
	repo *repository.CatRepository
}

func NewCatService(repo *repository.CatRepository) *CatService {
	return &CatService{repo: repo}
}

func (s *CatService) GetAllCatsLimited() ([]models.Cat, error) {
	return s.repo.GetAllCatsLimited()
}

func (s *CatService) GetCatDetails(id uint) (*models.Cat, error) {
	return s.repo.GetCatByID(id)
}

func (s *CatService) CreateCat(cat *models.Cat) error {
	return s.repo.CreateCat(cat)
}

func (s *CatService) UpdateCatSalary(id string, salary float64) error {
	return s.repo.UpdateSalary(id, salary)
}

func (s *CatService) DeleteCat(id uint) error {
	return s.repo.DeleteCat(id)
}
