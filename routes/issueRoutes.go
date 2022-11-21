package routes

import (
	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/DavidKorochik/pikud-darom-backend/middleware"
	"github.com/gin-gonic/gin"
)

func IssueRoutes(router *gin.Engine) {
	api := router.Group("/api").Use(middleware.AuthToken)

	api.GET("/issues", controllers.GetAllIssues)
	api.GET("/issues/department", controllers.FilterIssuesByDepartment)
	api.GET("/issues/monitoringSystem", controllers.FilterIssuesByMonitoringSystem)
	api.GET("/issues/issueCause", controllers.FilterIssuesByIssueCause)
	api.POST("/issues", controllers.CreateIssue)
	api.PUT("/issues/:id", controllers.UpdateIssue)
	api.DELETE("/issues/:id", controllers.DeleteIssue)
}
