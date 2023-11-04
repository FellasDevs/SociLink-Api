package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User, db *gorm.DB) error {
	result := db.Create(&user)

	return result.Error
}
