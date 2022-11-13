package routes

import (
	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/users", controllers.GetAllUsers)
	api.GET("/users/departments", controllers.GetAllUsersDepartments)
	api.POST("/users", controllers.CreateUser)
	api.DELETE("/users/:id", controllers.DeleteUser)
}
