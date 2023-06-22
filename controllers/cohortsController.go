package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CohortCreate(c *gin.Context) {
	pathID, _ := strconv.Atoi(c.Param("pathID"))
	user, exists := c.Get("user")
	admin := user.(models.User).Admin
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	if !admin {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Insufficient permissions. Need Admin permissions to continue",
		})
		return
	}

	var body struct {
		StartDate string
	}
	c.Bind(&body)

	// Convert string to Time datatype
	date, error := time.Parse("2006-01-02", body.StartDate)

	if error != nil {
		fmt.Println(error)
		c.JSON(400, gin.H{
			"error": "Date needs to be in YYYY-MM-DD format",
		})
		return
	}

	cohort := models.Cohort{PathID: uint(pathID), StartDate: date}
	result := initializers.DB.Create(&cohort)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "something terribly wrong has happened",
		})
		return
	}
	c.JSON(200, gin.H{
		"cohort": cohort,
	})

}
