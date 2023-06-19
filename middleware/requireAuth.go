package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/benjaminkhlau/go-crud/initializers"
	"github.com/benjaminkhlau/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	// initializers.DB.First(&user, claims["email"].(string)) // This is incorrect. Correct code on next line
	initializers.DB.Where("email = ?", claims["email"].(string)).First(&user)

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Next()
}
