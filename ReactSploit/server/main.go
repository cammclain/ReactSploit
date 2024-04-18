package main

import (
	"github.com/cammclain/ReactSploit/server/middleware"
	"github.com/cammclain/ReactSploit/server/models"
	"github.com/cammclain/ReactSploit/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase() // Initialize database connection
	middleware.SetupMiddleware(router)
	routes.SetupRouter(router)

	router.Run(":8080")
}
