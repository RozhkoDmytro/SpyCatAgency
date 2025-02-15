package handler

import (
	"net/http"
	"strconv"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TargetHandler struct {
	service *service.TargetService
}

func NewTargetHandler(service *service.TargetService) *TargetHandler {
	return &TargetHandler{service: service}
}

func (h *TargetHandler) CompleteTarget(c *gin.Context) {
	targetID, _ := strconv.Atoi(c.Param("target_id"))

	if err := h.service.CompleteTarget(uint(targetID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to complete target"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target completed"})
}

func (h *TargetHandler) UpdateTargetNotes(c *gin.Context) {
	targetID, _ := strconv.Atoi(c.Param("target_id"))

	var req struct {
		Notes []string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if len(req.Notes) == 0 || len(req.Notes) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Notes cannot be empty or > 3"})
		return
	}

	if err := h.service.UpdateTargetNotes(uint(targetID), req.Notes); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Target not found. Or mission | target is completed"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update notes"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notes updated"})
}

func (h *TargetHandler) AddTargetToMission(c *gin.Context) {
	missionID, _ := strconv.Atoi(c.Param("mission_id"))

	var target models.Target
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	target.MissionID = uint(missionID)

	if err := h.service.AddTargetToMission(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add target"})
		return
	}

	c.JSON(http.StatusCreated, target)
}

func (h *TargetHandler) DeleteTarget(c *gin.Context) {
	targetID, _ := strconv.Atoi(c.Param("target_id"))

	if err := h.service.DeleteTarget(uint(targetID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete target"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target deleted"})
}
