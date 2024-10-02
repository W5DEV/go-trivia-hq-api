package routes

import (
	"github.com/W5DEV/go-hp-trivia-api/controllers"
	"github.com/W5DEV/go-hp-trivia-api/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionsRouteController struct {
	questionsController controllers.QuestionsController
}

func NewRouteQuestionsController(questionsController controllers.QuestionsController) QuestionsRouteController {
	return QuestionsRouteController{questionsController}
}

func (pc *QuestionsRouteController) QuestionsRoute(rg *gin.RouterGroup) {

	router := rg.Group("questions")
	
	router.GET("/", pc.questionsController.FindQuestions)
	router.GET("/:questionsId", pc.questionsController.FindQuestionsById)
	router.GET("/random", pc.questionsController.FindRandomQuestions)
	router.GET("/tag", pc.questionsController.FindQuestionsByTag)
	router.GET("/difficulty", pc.questionsController.FindQuestionsByDifficulty)
	router.GET("/origin", pc.questionsController.FindQuestionsByQuestionOrigin)
	router.GET("/recents", pc.questionsController.FindRecentQuestions)
	router.PUT("/answer", pc.questionsController.SubmitAnswer)
	router.PUT("/like", pc.questionsController.RecordLike)
	router.PUT("/dislike", pc.questionsController.RecordDislike)
	router.GET("/tags", pc.questionsController.FindAllTags)
	router.GET("/popular", pc.questionsController.FindMostPopularQuestions)
	router.GET("/most-liked", pc.questionsController.FindMostLikedQuestions)
	router.GET("/least-answered", pc.questionsController.FindLeastAnsweredQuestions)
	router.GET("/topic", pc.questionsController.FindQuestionsByTopic)
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.questionsController.CreateQuestions)
	router.PUT("/:questionsId", pc.questionsController.UpdateQuestions)
	router.DELETE("/:questionsId", pc.questionsController.DeleteQuestions)
	router.PUT("/topic", pc.questionsController.UpdateTopic)
}