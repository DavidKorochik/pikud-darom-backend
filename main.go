package main

import (
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables("config.env")
}

func main() {
	db.DBConnection()

	router := gin.Default()
	port := os.Getenv("PORT")

	routes.IssueRoutes(router)
	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
