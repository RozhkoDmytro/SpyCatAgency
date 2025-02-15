package handler

import (
	"net/http"
	"strconv"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/models"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/service"
	"github.com/gin-gonic/gin"
)

type CatHandler struct {
	service *service.CatService
}

func NewCatHandler(service *service.CatService) *CatHandler {
	return &CatHandler{service: service}
}

func (h *CatHandler) GetCat(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cat ID"})
		return
	}

	cat, err := h.service.GetCatDetails(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cat not found"})
		return
	}

	c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) CreateCat(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := h.service.CreateCat(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create cat"})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

func (h *CatHandler) UpdateCat(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := h.service.UpdateCat(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update cat"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) DeleteCat(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cat ID"})
		return
	}
	if err := h.service.DeleteCat(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete cat"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cat deleted"})
}
