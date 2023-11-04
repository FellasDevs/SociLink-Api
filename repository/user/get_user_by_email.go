package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUserByEmail(email string, db *gorm.DB) (models.User, error) {
	var user models.User

	result := db.Where(models.User{Email: email}).First(&user)

	return user, result.Error
}
