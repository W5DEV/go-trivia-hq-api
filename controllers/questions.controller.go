package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sort"
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
		Type:			payload.Type,
		Tags: 		   	payload.Tags,
        Difficulty:     0,
		CorrectAnswer:  payload.CorrectAnswer,
		Completed: 		payload.Completed,
		QuestionOrigin: payload.QuestionOrigin,
        Topic:          payload.Topic,
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
		Type:     		payload.Type,
		Tags: 		  	payload.Tags,
		CorrectAnswer:  payload.CorrectAnswer,
		Completed: 		payload.Completed,
		QuestionOrigin: payload.QuestionOrigin,
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

// Get Questions Sorted By Most Recently Created Handler
func (pc *QuestionsController) FindRecentQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("limit")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.Questions
    results := pc.DB.Order("created_at desc").Limit(limit).Find(&questions)
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
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
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
    if limit > len(questions) {
        limit = len(questions)
    }

    selectedQuestions := questions[:limit]

    // Step 4: Return the selected questions
    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(selectedQuestions), "data": selectedQuestions})
}

// Get Questions By Difficulty Handler
func (pc *QuestionsController) FindQuestionsByDifficulty(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.Questions
    results := pc.DB.Where("difficulty != 0").Order("difficulty asc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Get Most Liked Questions Handler
func (pc *QuestionsController) FindMostLikedQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.Questions
    results := pc.DB.Order("likes desc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Get Most Popular Questions Handler
func (pc *QuestionsController) FindMostPopularQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.Questions
    results := pc.DB.Where("amount_seen != 0").Order("(likes / amount_seen) desc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// Get Least Answered Questions Handler
func (pc *QuestionsController) FindLeastAnsweredQuestions(ctx *gin.Context) {
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    var questions []models.Questions
    results := pc.DB.Order("amount_seen asc").Limit(limit).Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
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
    //Create an endpoint that uses the same randomizer logic as FindLeastAnsweredQuestions but filters by tag. Include the limiter as well
    limit := 0
    limitString := ctx.Query("count")

    if limitString != "" {
        limit, _ = strconv.Atoi(limitString)
    } else {
        limit = 25
    }

    tag := ctx.Query("tag")
    if tag == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No tag specified"})
        return
    }

    var questions []models.Questions
    if err := pc.DB.Where("tags @> ?", `["`+tag+`"]`).Find(&questions).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
        return
    }

    // Shuffle questions
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

    // Limit the number of quesitons returned based on limit
    if limit < len(questions) && limitString != "all" {
        questions = questions[:limit]
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Record Answer Handler
func (pc *QuestionsController) SubmitAnswer(ctx *gin.Context) {
    questionsId := ctx.Query("questionsId")
    answer := ctx.Query("answer")
    var isCorrect bool

    var question models.Questions
    result := pc.DB.First(&question, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that id exists"})
        return
    }

    answerArray := strings.Split(answer, "|")
    sort.Strings(answerArray)
    sortedAnswer := strings.Join(answerArray, ", ")

    correctAnswerArray := strings.Split(question.CorrectAnswer, "\n")
    sort.Strings(correctAnswerArray)
    sortedCorrectAnswer := strings.Join(correctAnswerArray, ", ")

    if sortedAnswer == sortedCorrectAnswer {
        question.AmountCorrect++
        isCorrect = true
    }

    question.AmountSeen++
    pc.DB.Save(&question)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "correct": isCorrect, "correct_answer": sortedCorrectAnswer, "submitted_answer": sortedAnswer})
}

// Record Like Handler
func (pc *QuestionsController) RecordLike(ctx *gin.Context) {
    questionsId := ctx.Query("questionsId")
    var questions models.Questions
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
func (pc *QuestionsController) RecordDislike(ctx *gin.Context) {
    questionsId := ctx.Query("questionsId")
    var questions models.Questions
    result := pc.DB.First(&questions, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
        return
    }

    questions.Dislikes++
    pc.DB.Save(&questions)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get All Tags Handler
func (pc *QuestionsController) FindAllTags(ctx *gin.Context) {
    var questions []models.Questions
    results := pc.DB.Find(&questions)
    if results.Error != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
        return
    }

    tags := []string{}
    for _, question := range questions {
        var tagSlice []string
        if err := json.Unmarshal(question.Tags, &tagSlice); err != nil {
            // Handle the error here
            continue
        }
        tags = append(tags, tagSlice...)
    }

    // remove duplicates from tags
    tags = removeDuplicates(tags)

    sort.Strings(tags)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": tags})
}

// Function to remove duplicates from a string slice
func removeDuplicates(slice []string) []string {
    encountered := map[string]bool{}
    result := []string{}

    for v := range slice {
        if encountered[slice[v]] {
            // Do not add duplicate element
            continue
        }
        // Add element to map
        encountered[slice[v]] = true
        // Add element to result slice
        result = append(result, slice[v])
    }
    return result
}

// Update Question Topic Handler
func (pc *QuestionsController) UpdateTopic(ctx *gin.Context) {
    questionsId := ctx.Param("questionsId")
    topic := ctx.Query("topic")

    var questions models.Questions
    result := pc.DB.First(&questions, "id = ?", questionsId)
    if result.Error != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No questions with that title exists"})
        return
    }

    questions.Topic = topic
    pc.DB.Save(&questions)

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": questions})
}

// Get Questions By Topic Handler
func (pc *QuestionsController) FindQuestionsByTopic(ctx *gin.Context) {
    topic := ctx.Query("topic")
    if topic == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No topic specified"})
        return
    }

    var questions []models.Questions
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
