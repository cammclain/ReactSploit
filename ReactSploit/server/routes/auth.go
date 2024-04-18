package routes

import (
	"github.com/cammclain/ReactSploit/server/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes initializes the authentication routes
func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		// POST /auth/login - For user login
		authGroup.POST("/login", controllers.Login)

		// POST /auth/register - For new user registration (if applicable)
		authGroup.POST("/register", controllers.Register)

		// POST /auth/logout - For user logout (if you are handling sessions)
		authGroup.POST("/logout", controllers.Logout)

		// Add more authentication related routes here if needed
	}
}
