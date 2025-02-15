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

func (r *CatRepository) GetAllCatsLimited() ([]models.Cat, error) {
	var cats []models.Cat
	result := r.db.Find(&cats)
	return cats, result.Error
}

func (r *CatRepository) CreateCat(cat *models.Cat) error {
	return r.db.Create(cat).Error
}

func (r *CatRepository) UpdateSalary(id string, salary float64) error {
	result := r.db.Exec("UPDATE cats SET salary = ? WHERE id = ?", salary, id)
	return result.Error
}

func (r *CatRepository) DeleteCat(id uint) error {
	return r.db.Delete(&models.Cat{}, id).Error
}
