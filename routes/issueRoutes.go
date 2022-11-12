package routes

import (
	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/gin-gonic/gin"
)

func IssueRoutes(router *gin.Engine) {
	router.GET("/api/issues", controllers.GetAllIssues)
	router.PUT("/api/issues", controllers.CreateIssue)
	router.PATCH("/api/issues/:id", controllers.UpdateIssue)
	router.DELETE("/api/issues/:id", controllers.DeleteIssue)
}
