package main

import (
	"log"
	"net/http"

	"github.com/W5DEV/go-hp-trivia-api/controllers"
	"github.com/W5DEV/go-hp-trivia-api/initializers"
	"github.com/W5DEV/go-hp-trivia-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	QuestionsController      controllers.QuestionsController
	QuestionsRouteController routes.QuestionsRouteController

	SourcesController      controllers.SourcesController
	SourcesRouteController routes.SourcesRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	QuestionsController = controllers.NewQuestionsController(initializers.DB)
	QuestionsRouteController = routes.NewRouteQuestionsController(QuestionsController)

	SourcesController = controllers.NewSourcesController(initializers.DB)
	SourcesRouteController = routes.NewRouteSourcesController(SourcesController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")


	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to the go-hp-trivia-api!"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	QuestionsRouteController.QuestionsRoute(router)
	SourcesRouteController.SourcesRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}