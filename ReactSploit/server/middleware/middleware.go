package middleware

import (
	"github.com/gin-gonic/gin"
)

// SetupMiddleware configures all the middleware for the application.
func SetupMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORS())
	router.Use(BasicAuth())
}
