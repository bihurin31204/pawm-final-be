// middleware/authenticate.go
package middleware

import (
	"log"
	"net/http"
	"strings"
	"vlab-backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_jwt_secret")

func Authenticate(c *gin.Context) {

	log.Println("JWT Secret:", string(jwtKey))

	authHeader := c.GetHeader("Authorization")
	log.Println("Authorization Header:", authHeader) // Logging header for debugging
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	log.Println("Token String:", tokenString) // Logging token for debugging
	claims := &handlers.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}
