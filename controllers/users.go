package controllers

import (
	"go_api_learn/initializers"
	"go_api_learn/models"
	"go_api_learn/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct{}

func (uc *UserController) HandleGetAllUsers(c *gin.Context) {

	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := responses.UserListResponse{
		Message: "List of users",
		Data:    users,
	}

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Validate uuid
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	var user models.User

	result := initializers.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	response := responses.UserResponse{
		Message: "User found",
		Data:    user,
	}

	c.JSON(http.StatusOK, response)
}

// func (uc *UserController) CreateUser(c *gin.Context) {
// 	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
// }

// func (uc *UserController) UpdateUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
// }

// func (uc *UserController) DeleteUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
// }