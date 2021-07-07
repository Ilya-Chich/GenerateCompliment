package main

import (
	"GenerateCompliment/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.New()
	r.GET("/api/v1/compliment", controllers.GenerateCompliment)
	r.Run(os.Getenv("HTTP_PORT"))
}
