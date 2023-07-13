package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Define a secret key for JWT signing
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func SignUp(c *gin.Context) {
	// Get email/pw off req body
	var body struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
		Phone     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Hash the pw
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	// Create the user
	user := models.User{Email: body.Email, Password: string(hash), FirstName: body.FirstName, LastName: body.LastName, Phone: body.Phone}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"email":     body.Email,
		"firstname": body.FirstName,
		"lastname":  body.LastName,
		"phone":     body.Phone,
		"token":     tokenString,
		// "password": hash,
	})
}

func Login(c *gin.Context) {
	// Get email/pw from the request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Retrieve the user from the database based on the email
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare the provided password with the hashed password from the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // Set the token expiration time

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// Respond with the JWT token
	c.JSON(http.StatusOK, gin.H{
		"profile": user,
		// "email":  body.Email,
		"token": tokenString,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func SetAdminStatus(c *gin.Context) {
	// Check if the user is authenticated
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	// Check if the user has admin authorization
	if user.(models.User).Admin {
		// Get the user ID from the request parameters or body
		userID := c.Param("id")
		// Find the user in the database based on the ID
		var user models.User
		result := initializers.DB.Where("id = ?", userID).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}

		// Set the user's admin status
		user.Admin = !user.Admin

		// Save the updated user in the database
		result = initializers.DB.Save(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to set admin status",
			})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, gin.H{
			"message": "Admin status updated to " + strconv.FormatBool(user.Admin) + " for user ID " + userID,
		})
		return
	}

	// If the user is not authorized as an admin
	c.JSON(http.StatusForbidden, gin.H{
		"error": "User does not have admin authorization",
	})
}

func SetAdmin(c *gin.Context) {
	fmt.Println(c.Get("user"))
	var user models.User
	result := initializers.DB.Where("id = ?", 1).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid User",
		})
		return
	}
	if user.ID == 1 {
		user.Admin = true
		if err := initializers.DB.Save(&user).Error; err != nil {
			// Handle error
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
	}
}

func AllUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}
