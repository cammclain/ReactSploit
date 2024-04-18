package main

import (
	"github.com/cammclain/ReactSploit/server/middleware"
	"github.com/cammclain/ReactSploit/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	middleware.SetupMiddleware(router)
	routes.SetupRouter(router)

	router.Run(":8080")
}
