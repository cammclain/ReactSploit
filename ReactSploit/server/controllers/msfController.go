package controllers

import "github.com/gin-gonic/gin"

// CheckMSFStatus checks the status of the Metasploit server
func CheckMSFStatus(c *gin.Context) {
	// Logic to interact with Metasploit RPC and check status
	c.JSON(200, gin.H{"status": "Metasploit is running"})
}
