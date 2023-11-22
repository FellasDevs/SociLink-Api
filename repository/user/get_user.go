package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUser(user *models.User, db *gorm.DB) error {
	result := db.Where(&user).First(&user)

	return result.Error
}
