package controllers

import (
	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PathsCreate(c *gin.Context) {
	var path models.Path
	c.BindJSON(&path)
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

func PathsShow(c *gin.Context) {
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
	initializers.DB.Delete(&models.Path{}, id)
	c.Status(200)
}
