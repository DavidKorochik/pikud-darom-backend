package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(envFilePath string) {
	err := godotenv.Load(envFilePath)

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
