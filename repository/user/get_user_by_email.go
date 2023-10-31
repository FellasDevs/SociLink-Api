package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUserByEmail(email string, db *gorm.DB) (models.User, error) {
	user := models.User{Email: email}

	result := db.First(&user)

	return user, result.Error
}
