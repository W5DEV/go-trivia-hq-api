package controllers

import (
	"math/rand"
	"net/http"
	"strconv"
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
		Completed: 		payload.Completed,
		QuestionOrigin: payload.QuestionOrigin,
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
		Completed: 		payload.Completed,
		QuestionOrigin: payload.QuestionOrigin,
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
func (pc *QuestionsController) FindQuestions(ctx *gin.Context) {


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

// Get Random Questions Handler
func (pc *QuestionsController) FindRandomQuestions(ctx *gin.Context) {
    // Step 1: Extract the number from the query
    countStr := ctx.Query("count")
    count, err := strconv.Atoi(countStr)
    if err != nil || count <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid count specified"})
        return
    }

    // Step 2: Fetch all questions
    var questions []models.Questions
    results := pc.DB.Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    // Step 3: Random selection
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) {
        questions[i], questions[j] = questions[j], questions[i]
    })

    // If the count is greater than the number of available questions, adjust it
    if count > len(questions) {
        count = len(questions)
    }

    selectedQuestions := questions[:count]

    // Step 4: Return the selected questions
    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(selectedQuestions), "data": selectedQuestions})
}

// Get Questions By Difficulty Handler
func (pc *QuestionsController) FindQuestionsByDifficulty(ctx *gin.Context) {
    difficulty := ctx.Query("difficulty")
    if difficulty == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No difficulty specified"})
        return
    }

    var questions []models.Questions
    if err := pc.DB.Where("difficulty = ?", difficulty).Find(&questions).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
        return
    }

    // Shuffle questions
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

    // If count is provided, parse it and limit the number of questions returned
    if countStr, exists := ctx.GetQuery("count"); exists {
        count, err := strconv.Atoi(countStr)
        if err != nil || count <= 0 {
            ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid count specified"})
            return
        }
        if count < len(questions) {
            questions = questions[:count]
        }
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get Questions By question_origin Handler
func (pc *QuestionsController) FindQuestionsByQuestionOrigin(ctx *gin.Context) {
    question_origin := ctx.Query("question_origin")
    if question_origin == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No question_origin specified"})
        return
    }

    var questions []models.Questions
    if err := pc.DB.Where("question_origin = ?", question_origin).Find(&questions).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
        return
    }

    // Shuffle questions
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

    // If count is provided, parse it and limit the number of questions returned
    if countStr, exists := ctx.GetQuery("count"); exists {
        count, err := strconv.Atoi(countStr)
        if err != nil || count <= 0 {
            ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid count specified"})
            return
        }
        if count < len(questions) {
            questions = questions[:count]
        }
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get Questions By Tag Handler
func (pc *QuestionsController) FindQuestionsByTag(ctx *gin.Context) {
    tag := ctx.Query("tag")
    if tag == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No tag specified"})
        return
    }

    var questions []models.Questions
    // Corrected query for finding questions by tag as a single string within a jsonb array
    if err := pc.DB.Where("tags @> ?", `["`+tag+`"]`).Find(&questions).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
        return
    }

    // Shuffle questions if count is provided
    if countStr, exists := ctx.GetQuery("count"); exists {
        count, err := strconv.Atoi(countStr)
        if err != nil || count <= 0 {
            ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid count specified"})
            return
        }

        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

        if count < len(questions) {
            questions = questions[:count]
        }
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}