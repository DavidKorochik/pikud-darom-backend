package routes

import (
	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/gin-gonic/gin"
)

func IssueRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/issues", controllers.GetAllIssues)
	api.POST("/issues", controllers.CreateIssue)
	api.PUT("/issues/:id", controllers.UpdateIssue)
	api.DELETE("/issues/:id", controllers.DeleteIssue)
}
