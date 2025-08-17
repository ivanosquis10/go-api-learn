package initializers

import (
	"go_api_learn/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
