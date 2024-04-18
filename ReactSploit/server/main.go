package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/ReactSploit/server/middleware"
	"github.com/yourusername/ReactSploit/server/routes"
)

func main() {
	router := gin.Default()
	middleware.SetupMiddleware(router)
	routes.SetupRouter(router)

	router.Run(":8080")
}
