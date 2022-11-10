package main

import (
	"log"
	"os"

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
	router := gin.New()
	router.Use(gin.Recovery())

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
