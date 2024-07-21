package routes

import (
	"github.com/W5DEV/go-hp-trivia-api/controllers"
	"github.com/W5DEV/go-hp-trivia-api/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}

