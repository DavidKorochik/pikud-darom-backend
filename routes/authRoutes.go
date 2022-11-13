package routes

import (
	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.POST("/auth", controllers.LogInUser)
}
