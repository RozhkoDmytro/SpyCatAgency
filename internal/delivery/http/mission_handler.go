package handler

import (
	"net/http"
	"strconv"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/service"
	"github.com/gin-gonic/gin"
)

type MissionHandler struct {
	service *service.MissionService
}

func NewMissionHandler(service *service.MissionService) *MissionHandler {
	return &MissionHandler{service: service}
}

func (h *MissionHandler) CreateMission(c *gin.Context) {
	var mission models.Mission
	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request "})
		return
	}

	if mission.CatID != nil && *mission.CatID == 0 {
		mission.CatID = nil
	}

	if len(mission.Targets) < 1 || len(mission.Targets) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A mission must have 1-3 targets"})
		return
	}

	if err := h.service.CreateMission(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create mission"})
		return
	}

	c.JSON(http.StatusCreated, mission)
}

func (h *MissionHandler) GetMissionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission ID"})
		return
	}

	mission, err := h.service.GetMissionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
		return
	}

	c.JSON(http.StatusOK, mission)
}

func (h *MissionHandler) GetAllMissions(c *gin.Context) {
	missions, err := h.service.GetAllMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve missions"})
		return
	}
	c.JSON(http.StatusOK, missions)
}

func (h *MissionHandler) AssignCatToMission(c *gin.Context) {
	missionID, _ := strconv.Atoi(c.Param("mission_id"))
	catIDParam := c.Param("cat_id")

	var catID *uint
	if catIDParam != "null" {
		catIDUint, err := strconv.Atoi(catIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cat ID"})
			return
		}
		catIDValue := uint(catIDUint)
		catID = &catIDValue
	}

	err := h.service.AssignCatToMission(uint(missionID), catID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign cat to mission. " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cat assigned to mission"})
}

// CompleteMissionHandler позначає місію як завершену
func (h *MissionHandler) CompleteMission(c *gin.Context) {
	missionID, _ := strconv.Atoi(c.Param("mission_id"))

	if err := h.service.CompleteMission(uint(missionID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to complete mission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mission completed"})
}

func (h *MissionHandler) DeleteMission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteMission(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete mission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mission deleted"})
}
