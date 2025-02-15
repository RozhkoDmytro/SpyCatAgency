package repository

import (
	"encoding/json"

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
				) FILTER (WHERE t.id IS NOT NULL), '[]'
			) AS targets
		FROM missions m
		LEFT JOIN targets t ON m.id = t.mission_id
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
				) FILTER (WHERE t.id IS NOT NULL), '[]'
			) AS targets
		FROM missions m
		LEFT JOIN targets t ON m.id = t.mission_id
		WHERE m.id = ?
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
	var result *gorm.DB

	if catID == nil {
		result = r.db.Exec("UPDATE missions SET cat_id = NULL WHERE id = ?", missionID)
	} else {
		result = r.db.Exec("UPDATE missions SET cat_id = ? WHERE id = ?", *catID, missionID)
	}

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
