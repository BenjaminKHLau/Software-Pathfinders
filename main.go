package main

import (
	"github.com/benjaminkhlau/go-crud/controllers"
	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	// r.GET("/", controllers.PostsCreate)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.GET("/posts/:id", controllers.PostsShow)
	r.Run() // listen and serve on 0.0.0.0:8080
}
