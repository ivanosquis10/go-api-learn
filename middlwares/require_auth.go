package middlwares

import (
	"fmt"
	"go_api_learn/initializers"
	"go_api_learn/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserResponse struct {
    ID       string    `json:"id"`
    Email    string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

func RequireAuth(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

	// Validate JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the EXP
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})	
		}

		userID := claims["sub"].(string)

		// Find user
		var user models.User
		result := initializers.DB.Select("id", "email", "created_at", "updated_at").Where("id = ?", userID).First(&user)

		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error al encontrar el usuario"})
			return
		}

		// Crea un objeto de respuesta con solo los campos que quieres
		userResponse := UserResponse{
			ID:       user.ID.String(),
			Email:    user.Email,
			CreatedAt: user.CreatedAt,
		}

		c.Set("user", userResponse)

		c.Next() // Continue to the next handler

	} else {
		log.Printf("Invalid token: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}
