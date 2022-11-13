package main

import (
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	db.DBConnection()

	router := gin.Default()

	routes.IssueRoutes(router)
	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
