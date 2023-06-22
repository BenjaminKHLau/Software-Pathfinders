package middleware

import (
	"net/http"

	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
)

func AdminAccess(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		c.Abort()
		return
	}

	if !user.(models.User).Admin {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You do not have permission. Need Admin permissions to continue",
		})
		c.Abort()
		return
	}
	c.Next()
}
