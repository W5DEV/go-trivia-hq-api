package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/W5DEV/go-hp-trivia-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SourcesController struct {
	DB *gorm.DB
}

func NewSourcesController(DB *gorm.DB) SourcesController {
	return SourcesController{DB}
}

// Create Sources Handler
func (pc *SourcesController) CreateSources(ctx *gin.Context) {
	var payload *models.CreateSourcesRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newSources := models.Sources{
		Source:       payload.Source,
		Citation:     payload.Citation,
		Active:       payload.Active,
		Topic:        payload.Topic,
		Completed:    payload.Completed,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	result := pc.DB.Create(&newSources)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Sources with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newSources})
}

// Update Sources Handler
func (pc *SourcesController) UpdateSources(ctx *gin.Context) {
	sourcesId := ctx.Param("sourcesId")

	var payload *models.UpdateSources
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedSources models.Sources
	result := pc.DB.First(&updatedSources, "id = ?", sourcesId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No sources exists with that ID"})
		return
	}
	now := time.Now()
	sourcesToUpdate := models.Sources{
		Source:       payload.Source,
		Citation:     payload.Citation,
		Active:       payload.Active,
		Topic:        payload.Topic,
		Completed:    payload.Completed,
		UpdatedAt:    now,
	}

	pc.DB.Model(&updatedSources).Updates(sourcesToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedSources})
}

// Get All Sources Handler
func (pc *SourcesController) FindSources(ctx *gin.Context) {


	var sources []models.Sources
	results := pc.DB.Find(&sources)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(sources), "data": sources})
}

// Get Sources By Topic Handler
func (pc *SourcesController) FindSourcesByTopic(ctx *gin.Context) {
	topic := ctx.Query("topic")

	var sources []models.Sources
	results := pc.DB.Find(&sources, "topic = ?", topic)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(sources), "data": sources})
}