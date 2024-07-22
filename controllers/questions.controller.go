package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/W5DEV/go-hp-trivia-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionsController struct {
	DB *gorm.DB
}

func NewQuestionsController(DB *gorm.DB) QuestionsController {
	return QuestionsController{DB}
}

// Create Questions Handler
func (pc *QuestionsController) CreateQuestions(ctx *gin.Context) {
	var payload *models.CreateQuestionsRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newQuestions := models.Questions{
		Question:       payload.Question,
		Answers:        payload.Answers,
		Source:  		payload.Source,
		Type:			payload.Type,
		Tags: 		   	payload.Tags,
		Difficulty:	 	payload.Difficulty,
		CorrectAnswer:  payload.CorrectAnswer,
		CreatedAt:     	now,
		UpdatedAt:     	now,
	}

	result := pc.DB.Create(&newQuestions)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Questions with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newQuestions})
}

// Update Questions Handler
func (pc *QuestionsController) UpdateQuestions(ctx *gin.Context) {
	questionsId := ctx.Param("questionsId")

	var payload *models.UpdateQuestions
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedQuestions models.Questions
	result := pc.DB.First(&updatedQuestions, "id = ?", questionsId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
		return
	}
	now := time.Now()
	questionsToUpdate := models.Questions{
		Question:       payload.Question,
		Answers:        payload.Answers,
		Source:  	 	payload.Source,
		Type:     		payload.Type,
		Tags: 		  	payload.Tags,
		Difficulty:	 	payload.Difficulty,
		CorrectAnswer:  payload.CorrectAnswer,
		CreatedAt:     	now,
		UpdatedAt:    	now,
	}

	pc.DB.Model(&updatedQuestions).Updates(questionsToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedQuestions})
}

// Get Single Questions Handler
func (pc *QuestionsController) FindQuestionsById(ctx *gin.Context) {
	questionsId := ctx.Param("questionsId")

	var questions models.Questions
	result := pc.DB.First(&questions, "id = ?", questionsId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get All Questions Handler
func (pc *QuestionsController) Findquestions(ctx *gin.Context) {


	var questions []models.Questions
	results := pc.DB.Find(&questions)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Delete Questions Handler
func (pc *QuestionsController) DeleteQuestions(ctx *gin.Context) {
	questionsId := ctx.Param("questionsId")

	result := pc.DB.Delete(&models.Questions{}, "id = ?", questionsId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

