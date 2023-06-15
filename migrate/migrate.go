package main

import (
	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
