package responses

import "go_api_learn/models"

type UserListResponse struct {
    Message string        `json:"message"`
    Data    []models.User `json:"data"`
}

type UserResponse struct {
    Message string      `json:"message"`
    Data    models.User `json:"data"`
}