package db

import (
	"log"
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error has occured while reading the .env file")
	}
}

func DBConnection() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error has occured while trying to connect to the database")
	}

	db.AutoMigrate(&models.Issue{}, &models.User{})

	DB = db

	log.Print("Connected to the database")
}
