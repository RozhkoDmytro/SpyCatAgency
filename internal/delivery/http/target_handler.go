package handler

import (
	"net/http"
	"strconv"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/service"
	"github.com/gin-gonic/gin"
)

// TargetHandler обробляє HTTP-запити для цілей
type TargetHandler struct {
	service *service.TargetService
}

// NewTargetHandler створює новий хендлер цілей
func NewTargetHandler(service *service.TargetService) *TargetHandler {
	return &TargetHandler{service: service}
}

// CompleteTargetHandler позначає ціль як завершену
func (h *TargetHandler) CompleteTargetHandler(c *gin.Context) {
	targetID, _ := strconv.Atoi(c.Param("target_id"))

	if err := h.service.CompleteTarget(uint(targetID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to complete target"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target completed"})
}

// UpdateTargetNotesHandler оновлює нотатки цілі
func (h *TargetHandler) UpdateTargetNotesHandler(c *gin.Context) {
	targetID, _ := strconv.Atoi(c.Param("target_id"))

	var req struct {
		Notes []string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.UpdateTargetNotes(uint(targetID), req.Notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update notes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notes updated"})
}

// AddTargetToMissionHandler додає ціль до місії
func (h *TargetHandler) AddTargetToMissionHandler(c *gin.Context) {
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

// DeleteTargetHandler видаляє ціль (якщо вона не завершена)
func (h *TargetHandler) DeleteTargetHandler(c *gin.Context) {
	targetID, _ := strconv.Atoi(c.Param("target_id"))

	if err := h.service.DeleteTarget(uint(targetID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete target"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Target deleted"})
}
