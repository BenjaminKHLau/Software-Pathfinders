package main

import (
	"os"

	"github.com/benjaminkhlau/go-crud/controllers"
	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r.Use(cors.Default())
	// r.Use(corsMiddleware())

	// Admin and authentication routes
	r.POST("/api/login", controllers.Login)
	r.POST("/api/signup", controllers.SignUp)
	r.GET("/api/validate", middleware.RequireAuth, controllers.Validate)
	r.PUT("/api/admin", middleware.RequireAuth, controllers.SetAdmin)
	r.PUT("/api/admin/:id", middleware.RequireAuth, controllers.SetAdminStatus)
	r.POST("/api/logout", controllers.Logout)
	r.GET("/api/users", middleware.RequireAuth, middleware.AdminAccess, controllers.AllUsers)

	// Post Content routes
	// Must have Path to create Post
	r.POST("/api/paths/:pathID/posts", middleware.RequireAuth, middleware.AdminAccess, controllers.PostsCreate)
	r.GET("/api/posts", controllers.PostsAll)
	r.PUT("/api/posts/:id", middleware.RequireAuth, middleware.AdminAccess, controllers.PostsUpdate)
	r.DELETE("/api/posts/:id", middleware.RequireAuth, middleware.AdminAccess, controllers.PostsDelete)
	r.GET("/api/posts/:id", controllers.PostsSingle)
	r.GET("/api/user", middleware.RequireAuth, controllers.PostsOfUser)

	// CRUD routes for Path
	// Make a Path before creating a Post
	r.GET("/api/paths", controllers.PathsAll)
	r.GET("/api/paths/:id", controllers.PathsSingle)
	r.POST("/api/paths", middleware.RequireAuth, middleware.AdminAccess, controllers.PathsCreate)
	r.PUT("/api/paths/:pathID", middleware.RequireAuth, middleware.AdminAccess, controllers.PathsUpdate)
	r.DELETE("/api/paths/:id", middleware.RequireAuth, middleware.AdminAccess, controllers.PathsDelete)

	// CRUD routes for Cohort
	// Make a Path before creating a Cohort
	r.POST("/api/paths/:pathID/cohorts", middleware.RequireAuth, middleware.AdminAccess, controllers.CohortCreate)
	r.PUT("/api/cohorts/:cohortID", middleware.RequireAuth, middleware.AdminAccess, controllers.CohortUpdate)
	r.GET("/api/cohorts/:cohortID/users", controllers.CohortUsers)
	r.GET("/api/cohorts", controllers.CohortsAll)
	r.POST("/api/cohorts/:cohortID/users/:userID", middleware.RequireAuth, middleware.AdminAccess, controllers.AddUserToCohort)
	r.DELETE("/api/cohorts/:cohortID/users/:userID", middleware.RequireAuth, middleware.AdminAccess, controllers.RemoveUserFromCohort)
	r.DELETE("/api/cohorts/:cohortID", middleware.RequireAuth, middleware.AdminAccess, controllers.CohortDelete)
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		// Handle preflight requests
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}

// 		c.Next()
// 	}
// }
