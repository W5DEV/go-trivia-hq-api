package main

import (
	"fmt"
	"log"

	"github.com/W5DEV/go-hp-trivia-api/initializers"
	"github.com/W5DEV/go-hp-trivia-api/models"
)


func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Questions{}, &models.Sources{}, &models.GeneralQuestions{})
	fmt.Println("? Migration complete")
}

