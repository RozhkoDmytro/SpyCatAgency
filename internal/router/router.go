package router

import (
	handler "github.com/RozhkoDmytro/SpyCatAgency/internal/delivery/http"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/repository"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	catRepo := repository.NewCatRepository(db)
	catService := service.NewCatService(catRepo)
	catHandler := handler.NewCatHandler(catService)

	r.GET("/cats/:id", catHandler.GetCat)
	r.GET("/cats", catHandler.GetAllCatsHandler)
	r.POST("/cats", catHandler.CreateCat)
	r.PUT("/cats", catHandler.UpdateCat)
	r.DELETE("/cats/:id", catHandler.DeleteCat)

	return r
}
