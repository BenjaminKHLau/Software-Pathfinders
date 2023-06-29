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

func CohortUpdate(c *gin.Context) {
	cohortID, _ := strconv.Atoi(c.Param("cohortID"))
	var body struct {
		StartDate string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert string to Time datatype
	date, err := time.Parse("2006-01-02", body.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date needs to be in YYYY-MM-DD format"})
		return
	}

	var cohort models.Cohort
	if err := initializers.DB.First(&cohort, cohortID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cohort not found"})
		return
	}

	cohort.StartDate = date
	if err := initializers.DB.Save(&cohort).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cohort"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cohort": cohort})
}

func CohortUsers(c *gin.Context) {
	cohortID := c.Param("cohortID")

	var cohort models.Cohort
	if err := initializers.DB.Preload("Users").First(&cohort, cohortID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"students": cohort.Users})
}

func CohortsAll(c *gin.Context) {
	var cohorts []models.Cohort
	initializers.DB.Find(&cohorts)
	c.JSON(200, gin.H{
		"cohorts": cohorts,
	})
}

func AddUserToCohort(c *gin.Context) {
	cohortIDStr := c.Param("cohortID")
	userIDStr := c.Param("userID")

	cohortID, err := strconv.ParseUint(cohortIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cohort ID"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var cohort models.Cohort
	if err := initializers.DB.First(&cohort, cohortID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&cohort).Association("Users").Append(&user)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %d added to cohort %d", userID, cohortID)})
}

func RemoveUserFromCohort(c *gin.Context) {
	cohortIDStr := c.Param("cohortID")
	userIDStr := c.Param("userID")

	cohortID, err := strconv.ParseUint(cohortIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cohort ID"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var cohort models.Cohort
	if err := initializers.DB.Preload("Users").First(&cohort, cohortID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&cohort).Association("Users").Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %d removed from cohort %d", userID, cohortID)})
}

func CohortDelete(c *gin.Context) {
	cohortID := c.Param("cohortID")
	var cohort models.Cohort
	initializers.DB.First(&cohort, cohortID)
	initializers.DB.Delete(&models.Cohort{}, cohortID)
	c.JSON(200, gin.H{
		"message": "Cohort ID " + cohortID + " deleted successfully",
	})
}
