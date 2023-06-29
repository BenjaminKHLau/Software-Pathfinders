package controllers

import (
	"net/http"
	"strconv"

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

	// Find Path and pathID
	// var path models.Path
	pathID, _ := strconv.Atoi(c.Param("pathID"))
	// initializers.DB.Where("ID = ?", pathID).First(&path)
	post := models.Post{Title: body.Title, Body: body.Body, UserID: person, PathID: uint(pathID)} //Author: user.(models.User), Paths: path}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsAll(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsSingle(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsOfUser(c *gin.Context) {
	user, exists := c.Get("user")
	person := user.(models.User).Email
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	thisVarHoldsUserPosts := models.User{}
	userQuery := models.User{Email: person}
	initializers.DB.Model(models.User{}).Preload("Posts").Where(&userQuery).Find(&thisVarHoldsUserPosts)
	c.JSON(200, gin.H{
		"posts": thisVarHoldsUserPosts,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	// fmt.Println("CHECK HERE ~~~~~~~~~~~~~", post.UserID)
	// fmt.Println(user.(models.User).ID)
	if post.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "This is not your post",
		})
		return
	}
	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(200, gin.H{
		"message": "Post ID " + id + " deleted successfully",
	})
}
