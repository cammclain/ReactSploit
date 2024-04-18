package controllers

import (
	"net/http"

	"github.com/cammclain/ReactSploit/server/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials services.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	isValid, err := services.AuthenticateWithMSF(credentials)
	if err != nil || !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// You can add JWT token generation here if needed
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
