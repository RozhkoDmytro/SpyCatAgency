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

func (s *CatService) GetCatDetails(id uint) (*models.Cat, error) {
	return s.repo.GetCatByID(id)
}

func (s *CatService) CreateCat(cat *models.Cat) error {
	return s.repo.CreateCat(cat)
}

func (s *CatService) UpdateCat(cat *models.Cat) error {
	return s.repo.UpdateCat(cat)
}

func (s *CatService) DeleteCat(id uint) error {
	return s.repo.DeleteCat(id)
}
