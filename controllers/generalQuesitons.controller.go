package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/W5DEV/go-hp-trivia-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GeneralQuestionsController struct {
	DB *gorm.DB
}

func NewGeneralQuestionsController(DB *gorm.DB) GeneralQuestionsController {
	return GeneralQuestionsController{DB}
}

// Create GeneralQuestions Handler
func (pc *GeneralQuestionsController) CreateGeneralQuestions(ctx *gin.Context) {
	var payload *models.CreateGeneralQuestionsRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newGeneralQuestions := models.GeneralQuestions{
		Question:       payload.Question,
        AnswerOne:      payload.AnswerOne,
        AnswerTwo:      payload.AnswerTwo,
        AnswerThree:    payload.AnswerThree,
        AnswerFour:     payload.AnswerFour,
        Difficulty:     0,
		CorrectAnswer:  payload.CorrectAnswer,
        Topic:          payload.Topic,
		CreatedAt:     	now,
		UpdatedAt:     	now,
	}

	result := pc.DB.Create(&newGeneralQuestions)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "GeneralQuestions with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newGeneralQuestions})
}

// Update GeneralQuestions Handler
func (pc *GeneralQuestionsController) UpdateGeneralQuestions(ctx *gin.Context) {
	questionsId := ctx.Param("questionsId")

	var payload *models.UpdateGeneralQuestions
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedGeneralQuestions models.GeneralQuestions
	result := pc.DB.First(&updatedGeneralQuestions, "id = ?", questionsId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
		return
	}
	now := time.Now()
	questionsToUpdate := models.GeneralQuestions{
		Question:       payload.Question,
        AnswerOne:      payload.AnswerOne,
        AnswerTwo:      payload.AnswerTwo,
        AnswerThree:    payload.AnswerThree,
        AnswerFour:     payload.AnswerFour,
		CorrectAnswer:  payload.CorrectAnswer,
        Topic:          payload.Topic,
		UpdatedAt:     	now,
	}

	pc.DB.Model(&updatedGeneralQuestions).Updates(questionsToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedGeneralQuestions})
}

// Get Single GeneralQuestions Handler
func (pc *GeneralQuestionsController) FindGeneralQuestionsById(ctx *gin.Context) {
	questionsId := ctx.Param("questionsId")

	var questions models.GeneralQuestions
	result := pc.DB.First(&questions, "id = ?", questionsId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get All GeneralQuestions Handler
func (pc *GeneralQuestionsController) FindGeneralQuestions(ctx *gin.Context) {


	var questions []models.GeneralQuestions
	results := pc.DB.Find(&questions)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Delete GeneralQuestions Handler
func (pc *GeneralQuestionsController) DeleteGeneralQuestions(ctx *gin.Context) {
	questionsId := ctx.Param("questionsId")

	result := pc.DB.Delete(&models.GeneralQuestions{}, "id = ?", questionsId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Get Random GeneralQuestions Handler
func (pc *GeneralQuestionsController) FindRandomGeneralQuestions(ctx *gin.Context) {
    // Step 1: Extract the number from the query
    countStr := ctx.Query("count")
    count, err := strconv.Atoi(countStr)
    if err != nil || count <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid count specified"})
        return
    }

    // Step 2: Fetch all questions
    var questions []models.GeneralQuestions
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

    selectedGeneralQuestions := questions[:count]

    // Step 4: Return the selected questions
    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(selectedGeneralQuestions), "data": selectedGeneralQuestions})
}

// Get GeneralQuestions By Difficulty Handler
func (pc *GeneralQuestionsController) FindGeneralQuestionsByDifficulty(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.GeneralQuestions
    results := pc.DB.Where("difficulty != 0").Order("difficulty asc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Get Most Liked GeneralQuestions Handler
func (pc *GeneralQuestionsController) FindMostLikedGeneralQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.GeneralQuestions
    results := pc.DB.Order("likes desc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Get Most Popular GeneralQuestions Handler
func (pc *GeneralQuestionsController) FindMostPopularGeneralQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.GeneralQuestions
    results := pc.DB.Where("amount_seen != 0").Order("(likes / amount_seen) desc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Get Least Answered GeneralQuestions Handler
func (pc *GeneralQuestionsController) FindLeastAnsweredGeneralQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.GeneralQuestions
    results := pc.DB.Order("amount_seen asc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Record Answer Handler
func (pc *GeneralQuestionsController) RecordAnswer(ctx *gin.Context) {
    isCorrect := ctx.Query("is_correct")
    questionsId := ctx.Query("questionsId")
    var questions models.GeneralQuestions
    result := pc.DB.First(&questions, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
        return
    } else {
        fmt.Println(isCorrect, ctx.Query("is_correct"))
    }

    if isCorrect == "true" {
        questions.AmountCorrect++
    }
    questions.AmountSeen++

    questions.Difficulty = questions.AmountCorrect / questions.AmountSeen * 100
    pc.DB.Save(&questions)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Record Like Handler
func (pc *GeneralQuestionsController) RecordLike(ctx *gin.Context) {
    questionsId := ctx.Query("questionsId")
    var questions models.GeneralQuestions
    result := pc.DB.First(&questions, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
        return
    }

    questions.Likes++
    pc.DB.Save(&questions)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Record Dislike Handler
func (pc *GeneralQuestionsController) RecordDislike(ctx *gin.Context) {
    questionsId := ctx.Query("questionsId")
    var questions models.GeneralQuestions
    result := pc.DB.First(&questions, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
        return
    }

    questions.Dislikes++
    pc.DB.Save(&questions)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Update Question Topic Handler
func (pc *GeneralQuestionsController) UpdateTopic(ctx *gin.Context) {
    questionsId := ctx.Param("questionsId")
    topic := ctx.Query("topic")

    var questions models.GeneralQuestions
    result := pc.DB.First(&questions, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
        return
    }

    questions.Topic = topic
    pc.DB.Save(&questions)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get GeneralQuestions By Topic Handler
func (pc *GeneralQuestionsController) FindGeneralQuestionsByTopic(ctx *gin.Context) {
    topic := ctx.Query("topic")
    if topic == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No topic specified"})
        return
    }

    var questions []models.GeneralQuestions
    if err := pc.DB.Where("topic = ?", topic).Find(&questions).Error; err != nil {
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

// Get GeneralQuestions by Incorrect CorrectAnswers (Return all questions that have a CorrectAnswer not equal to either AnswerOne, AnswerTwo, AnswerThree or Answer Four )
func (pc *GeneralQuestionsController) FindInvalidGeneralQuestions(ctx *gin.Context) {
    var questions []models.GeneralQuestions
    results := pc.DB.Where("correct_answer != answer_one AND correct_answer != answer_two AND correct_answer != answer_three AND correct_answer != answer_four").Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

