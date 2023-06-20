package controllers

import (
	"net/http"

	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	user, exists := c.Get("user")
	person := user.(models.User).ID
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	// Get data off req body
	var body struct {
		Title    string
		Body     string
		AuthorID uint
		PathID   uint
	}

	c.Bind(&body)

	// Find Path
	var path models.Path
	initializers.DB.Where("ID = ?", body.PathID).First(&path)
	post := models.Post{Title: body.Title, Body: body.Body, AuthorID: person, PathID: body.PathID, Author: user.(models.User), Paths: path}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"things": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
