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
	// Must have Path to create Post
	r.POST("/paths/:pathID/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.GET("/posts", controllers.PostsAll)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.PostsUpdate)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.PostsDelete)
	r.GET("/posts/:id", controllers.PostsSingle)
	r.GET("/user", middleware.RequireAuth, controllers.PostsOfUser)

	// CRUD routes for Path
	// Make a Path before creating a Post
	r.GET("/paths", controllers.PathsAll)
	r.GET("/paths/:id", controllers.PathsSingle)
	r.POST("/paths", middleware.RequireAuth, controllers.PathsCreate)
	r.PUT("/paths/:pathID", middleware.RequireAuth, controllers.PathsUpdate)
	r.DELETE("/paths/:id", middleware.RequireAuth, controllers.PathsDelete)

	// CRUD routes for Cohort
	// Make a Path before creating a Cohort
	r.POST("/paths/:pathID/cohort", middleware.RequireAuth, controllers.CohortCreate)
	r.GET("/cohort/:cohortID/users", controllers.CohortUsers)

	r.Run() // listen and serve on 0.0.0.0:8080
}
