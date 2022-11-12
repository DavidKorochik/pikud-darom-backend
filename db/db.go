package db

import (
	"log"
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error has occured while reading the .env file")
	}
}

func DBConnection() {
	dsn := os.Getenv("DB_URL")
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error has occured while trying to connect to the database")
	}

	config.DB.AutoMigrate(&models.Issue{}, &models.User{})

	log.Print("Connected to the database")
}
