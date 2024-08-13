package routes

import (
	"github.com/W5DEV/go-hp-trivia-api/controllers"
	"github.com/W5DEV/go-hp-trivia-api/middleware"
	"github.com/gin-gonic/gin"
)

type SourcesRouteController struct {
	sourcesController controllers.SourcesController
}

func NewRouteSourcesController(sourcesController controllers.SourcesController) SourcesRouteController {
	return SourcesRouteController{sourcesController}
}

func (pc *SourcesRouteController) SourcesRoute(rg *gin.RouterGroup) {

	router := rg.Group("sources")
	
	router.GET("/", pc.sourcesController.FindSources)
	router.GET("/topic", pc.sourcesController.FindSourcesByTopic)
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.sourcesController.CreateSources)
	router.PUT("/:sourcesId", pc.sourcesController.UpdateSources)
}