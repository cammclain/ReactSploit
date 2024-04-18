package routes

import (
	"github.com/cammclain/ReactSploit/server/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", controllers.Login)
	return router
}
