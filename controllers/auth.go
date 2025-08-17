package controllers

import (
	"go_api_learn/initializers"
	"go_api_learn/models"
	"go_api_learn/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleSignUp(c *gin.Context) {
	// Get body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if !utils.ValidateEmail(body.Email) {
		c.JSON(400, gin.H{"error": "invalid email"})
		return
	}

	// hash password
	// passHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "failed to hash password"})
	// 	return
	// }

	// create user
	user := models.User{
		Email:   body.Email,
		Password: body.Password,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	// respond
	c.JSON(201, gin.H{
		"message": "user created successfully",
		"user": map[string]any{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func HandleSignIn(c *gin.Context) {
	// Get body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if !utils.ValidateEmail(body.Email) {
		c.JSON(400, gin.H{"error": "credentials are invalid or missing"})
		return
	}

	// Find user
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "credentials are invalid or missing"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(401, gin.H{"error": "credentials are invalid"})
		return
	}

	// Generate JWT session token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), //  24 hours
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	var cookieTime = 3600 * 24 // 24 hours

	// Generate Cookie with JWT session
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, cookieTime, "", "", false, true)

	c.JSON(200, gin.H{
		"message": "user signed in successfully",
	})
}

func HandleUserSession(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{
		"message": "user session is active",
		"user": user,
	})
}

func HandleLogOut(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(200, gin.H{
		"message": "user logged out successfully",
	})
}