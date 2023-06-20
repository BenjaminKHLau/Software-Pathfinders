package main

import (
	"github.com/benjaminkhlau/go-crud/controllers"
	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Admin and authentication routes
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignUp)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.PUT("/admin", middleware.RequireAuth, controllers.SetAdmin)
	r.PUT("/admin/:id", middleware.RequireAuth, controllers.SetAdminStatus)

	// Post Content routes
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.GET("/posts/:id", controllers.PostsShow)

	// CRUD routes for Path
	r.GET("/paths", controllers.PathsAll)
	r.GET("/paths/:id", controllers.PathsShow)
	r.POST("/paths", middleware.RequireAuth, controllers.PathsCreate)
	r.PUT("/paths/:id", controllers.PathsUpdate)
	r.DELETE("/paths/:id", controllers.PathsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
