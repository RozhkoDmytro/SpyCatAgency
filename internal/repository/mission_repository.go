package repository

import (
	"encoding/json"
	"errors"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"gorm.io/gorm"
)

type MissionRepository struct {
	db *gorm.DB
}

type MissionWithTargets struct {
	ID        uint            `json:"id"`
	CatID     *uint           `json:"cat_id"`
	Completed bool            `json:"completed"`
	Targets   json.RawMessage `json:"targets"`
}

func NewMissionRepository(db *gorm.DB) *MissionRepository {
	return &MissionRepository{db: db}
}

func (r *MissionRepository) CreateMission(mission *models.Mission) error {
	if mission.CatID != nil {
		var count int64
		err := r.db.Raw(`
            SELECT COUNT(*) FROM missions 
            WHERE cat_id = ? AND completed = FALSE AND deleted_at IS NULL
        `, mission.CatID).Scan(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("this cat is already assigned to another active mission")
		}
	}

	return r.db.Create(mission).Error
}

func (r *MissionRepository) GetAllMissions() ([]MissionWithTargets, error) {
	var missions []MissionWithTargets

	query := `
		SELECT m.id, m.cat_id, m.completed,
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', t.id,
						'name', t.name,
						'country', t.country,
						'notes', t.notes,
						'completed', t.completed
					)
				) FILTER (WHERE t.id IS NOT NULL and t.deleted_at IS NULL), '[]'
			) AS targets
		FROM missions m
		LEFT JOIN targets t ON m.id = t.mission_id
		WHERE m.deleted_at IS NULL
		GROUP BY m.id
	`
	err := r.db.Raw(query).Scan(&missions).Error
	return missions, err
}

func (r *MissionRepository) GetMissionByID(id uint) (*MissionWithTargets, error) {
	var mission MissionWithTargets

	query := `
		SELECT m.id, m.cat_id, m.completed,
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', t.id,
						'name', t.name,
						'country', t.country,
						'notes', t.notes,
						'completed', t.completed
					)
				) FILTER (WHERE t.id IS NOT NULL and t.deleted_at IS NULL), '[]'
			) AS targets
		FROM missions m
		LEFT JOIN targets t ON m.id = t.mission_id
		WHERE m.id = ? and m.deleted_at IS NULL
		GROUP BY m.id
	`

	err := r.db.Raw(query, id).Scan(&mission).Error
	if err != nil {
		return nil, err
	}

	if mission.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &mission, nil
}

func (r *MissionRepository) AssignCatToMission(missionID uint, catID *uint) error {
	var existingMission models.Mission
	if catID != nil {
		err := r.db.Where("cat_id = ? AND completed = FALSE and deleted_at IS NULL", *catID).First(&existingMission).Error
		if err == nil {
			return errors.New("this cat is already assigned to another mission")
		}
	}

	result := r.db.Model(&models.Mission{}).Where("id = ?", missionID).Update("cat_id", catID)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *MissionRepository) MarkMissionAsCompleted(missionID uint) error {
	result := r.db.Exec("UPDATE missions SET completed = ? WHERE id = ?", true, missionID)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *MissionRepository) DeleteMission(id uint) error {
	result := r.db.Where("id = ? AND cat_id IS NULL", id).Delete(&models.Mission{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
