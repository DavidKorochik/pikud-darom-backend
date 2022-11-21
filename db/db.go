package db

import (
	"log"
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func init() {
	initializers.LoadEnvVariables("config.env")
}

func DBConnection() {
	dsn := os.Getenv("DB_URL")
	config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error has occured while trying to connect to the database")
	}

	config.DB.AutoMigrate(&models.Issue{})
	config.DB.AutoMigrate(&models.User{})

	log.Print("Connected to the database")
}
