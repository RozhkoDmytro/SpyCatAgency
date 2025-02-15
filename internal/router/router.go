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

	//  CatHandler
	catRepo := repository.NewCatRepository(db)
	catService := service.NewCatService(catRepo)
	catHandler := handler.NewCatHandler(catService)

	r.GET("/cats/:id", catHandler.GetCat)
	r.GET("/cats", catHandler.GetAllCats)
	r.POST("/cats", catHandler.CreateCat)
	r.PUT("/cats", catHandler.UpdateCat)
	r.DELETE("/cats/:id", catHandler.DeleteCat)

	//  MissionHandler
	missionRepo := repository.NewMissionRepository(db)
	missionService := service.NewMissionService(missionRepo)
	missionHandler := handler.NewMissionHandler(missionService)

	r.POST("/missions", missionHandler.CreateMission)
	r.GET("/missions", missionHandler.GetAllMissions)
	r.GET("/missions/:id", missionHandler.GetMissionByID)
	r.PUT("/missions/:mission_id/cats/:cat_id", missionHandler.AssignCatToMission)
	r.PUT("/missions/:mission_id/complete", missionHandler.CompleteMission)
	r.DELETE("/missions/:id", missionHandler.DeleteMission)

	// TargetHandler
	targetRepo := repository.NewTargetRepository(db)
	targetService := service.NewTargetService(targetRepo)
	targetHandler := handler.NewTargetHandler(targetService)

	r.PUT("/targets/:target_id/complete", targetHandler.CompleteTargetHandler)
	r.PUT("/targets/:target_id/notes", targetHandler.UpdateTargetNotesHandler)
	r.POST("/missions/:mission_id/targets", targetHandler.AddTargetToMissionHandler)
	r.DELETE("/targets/:target_id", targetHandler.DeleteTargetHandler)

	return r
}
