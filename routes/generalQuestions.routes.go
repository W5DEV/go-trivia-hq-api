package routes

import (
	"github.com/W5DEV/go-hp-trivia-api/controllers"
	"github.com/W5DEV/go-hp-trivia-api/middleware"
	"github.com/gin-gonic/gin"
)

type GeneralQuestionsRouteController struct {
	questionsController controllers.GeneralQuestionsController
}

func NewRouteGeneralQuestionsController(questionsController controllers.GeneralQuestionsController) GeneralQuestionsRouteController {
	return GeneralQuestionsRouteController{questionsController}
}

func (pc *GeneralQuestionsRouteController) GeneralQuestionsRoute(rg *gin.RouterGroup) {

	router := rg.Group("questions")
	
	router.GET("/", pc.questionsController.FindGeneralQuestions)
	router.GET("/:questionsId", pc.questionsController.FindGeneralQuestionsById)
	router.GET("/random", pc.questionsController.FindRandomGeneralQuestions)
	router.GET("/difficulty", pc.questionsController.FindGeneralQuestionsByDifficulty)
	router.PUT("/answer", pc.questionsController.RecordAnswer)
	router.PUT("/like", pc.questionsController.RecordLike)
	router.PUT("/dislike", pc.questionsController.RecordDislike)
	router.GET("/popular", pc.questionsController.FindMostPopularGeneralQuestions)
	router.GET("/most-liked", pc.questionsController.FindMostLikedGeneralQuestions)
	router.GET("/least-answered", pc.questionsController.FindLeastAnsweredGeneralQuestions)
	router.GET("/topic", pc.questionsController.FindGeneralQuestionsByTopic)
	router.GET("/invalid", pc.questionsController.FindInvalidGeneralQuestions)
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.questionsController.CreateGeneralQuestions)
	router.PUT("/:questionsId", pc.questionsController.UpdateGeneralQuestions)
	router.DELETE("/:questionsId", pc.questionsController.DeleteGeneralQuestions)
	router.PUT("/topic", pc.questionsController.UpdateTopic)
}