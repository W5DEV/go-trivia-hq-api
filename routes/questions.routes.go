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
	router.GET("/", pc.questionsController.Findquestions)
	router.GET("/:questionsId", pc.questionsController.FindQuestionsById)
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.questionsController.CreateQuestions)
	router.PUT("/:questionsId", pc.questionsController.UpdateQuestions)
	router.DELETE("/:questionsId", pc.questionsController.DeleteQuestions)
}