package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupMiddleware configures all the middleware for the application.
func SetupMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORS())
	router.Use(BasicAuth())
}

// CORS sets up the CORS middleware.
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // or your frontend host
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// BasicAuth provides a basic authentication middleware.
func BasicAuth() gin.HandlerFunc {
	// Simulated database lookup for user credentials.
	// In production, you should replace this with a proper database or API call.
	requiredUser := "admin"
	requiredPass := "password"

	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()

		if hasAuth && user == requiredUser && pass == requiredPass {
			c.Next()
		} else {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	}
}
