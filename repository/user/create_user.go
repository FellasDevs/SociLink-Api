package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func CreateUser(user models.User, db *gorm.DB) (models.User, error) {
	result := db.Create(&user)

	return user, result.Error
}
