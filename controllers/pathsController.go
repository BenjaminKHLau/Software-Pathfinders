package controllers

import (
	"fmt"
	"net/http"

	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PathsCreate(c *gin.Context) {
	user, exists := c.Get("user")
	person := user.(models.User).ID
	fmt.Println("HELLOOOOOO ", person)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	var body struct {
		PathName        string
		PathDescription string
	}

	c.Bind(&body)
	path := models.Path{PathName: body.PathName, PathDescription: body.PathDescription}
	result := initializers.DB.Create(&path)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"path": path,
	})
}

func PathsAll(c *gin.Context) {
	var paths []models.Path
	initializers.DB.Find(&paths)
	c.JSON(200, gin.H{
		"paths": paths,
	})
}

func PathsSingle(c *gin.Context) {
	id := c.Param("id")
	var path models.Path
	initializers.DB.First(&path, id)
	c.JSON(200, gin.H{
		"path": path,
	})
}

func PathsUpdate(c *gin.Context) {
	id := c.Param("id")

	var path models.Path
	initializers.DB.First(&path, id)

	c.BindJSON(&path)
	initializers.DB.Save(&path)

	c.JSON(200, gin.H{
		"path": path,
	})
}

func PathsDelete(c *gin.Context) {
	id := c.Param("id")

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	if !user.(models.User).Admin {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You do not have permission",
		})
		return
	}

	initializers.DB.Delete(&models.Path{}, id)
	c.JSON(200, gin.H{
		"message": "successfully deleted",
	})
}
