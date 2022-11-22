package main

import (
	"os"

	"github.com/DavidKorochik/pikud-darom-backend/db"
	_ "github.com/DavidKorochik/pikud-darom-backend/docs"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables("config.env")
}

// If a package that needs to be in use with the CMD doesn't work, we have 2 options:
// 1. Create our project in the GO workspace.
// 2. Run the next command from out terminal: export PATH=$(go env GOPATH)/bin:$PATH.
// Notice that this command will be available only to the specific terminal from where we executed the command.
// If we'll open another terminal, packages that we want to execute from within the CMD won't work.

func main() {
	db.DBConnection()

	router := gin.Default()
	port := os.Getenv("PORT")

	routes.IssueRoutes(router)
	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
