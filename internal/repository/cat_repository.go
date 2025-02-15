package repository

import (
	"errors"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"

	"gorm.io/gorm"
)

type CatRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) *CatRepository {
	return &CatRepository{db: db}
}

func (r *CatRepository) GetCatByID(id uint) (*models.Cat, error) {
	var cat models.Cat
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, errors.New("cat not found")
	}
	return &cat, nil
}

func (r *CatRepository) CreateCat(cat *models.Cat) error {
	return r.db.Create(cat).Error
}

func (r *CatRepository) UpdateCat(cat *models.Cat) error {
	return r.db.Save(cat).Error
}

func (r *CatRepository) DeleteCat(id uint) error {
	return r.db.Delete(&models.Cat{}, id).Error
}
