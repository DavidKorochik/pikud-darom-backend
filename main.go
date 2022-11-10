package main

import (
	"log"
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db.DBConnection()

	router := gin.New()
	router.Use(gin.Recovery())

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
