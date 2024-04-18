package routes

import (
	"github.com/cammclain/ReactSploit/server/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all the route handlers and returns a *gin.Engine instance ready to be run.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Setup the middleware that are common across all routes
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Define a simple test route to check the server is running
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Group routes related to authentication
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
		// Add more authentication routes here if needed
	}

	// Group routes for other functionality as your app grows
	// Example: API operations on Metasploit
	msfRoutes := router.Group("/msf")
	{
		msfRoutes.GET("/status", controllers.CheckMSFStatus)
		// Add more Metasploit related routes here
	}

	return router
}
