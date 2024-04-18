package middleware

import (
	"net/http"

	"ReactSploit/server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			challenge(c)
			return
		}

		var foundUser models.User
		result := models.DB.Where("username = ?", user).First(&foundUser)
		if result.Error != nil {
			challenge(c)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(pass)); err != nil {
			challenge(c)
			return
		}

		c.Next()
	}
}

func challenge(c *gin.Context) {
	c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
