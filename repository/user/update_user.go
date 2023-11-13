package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func UpdateUser(user *models.User, db *gorm.DB) error {
	result := db.Save(&user)

	return result.Error
}
