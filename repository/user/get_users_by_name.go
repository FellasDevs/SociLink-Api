package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUsersByName(search string, db *gorm.DB) ([]models.User, error) {
	var users []models.User

	result := db.Where("name ILIKE ?", "%"+search+"%").Find(&users)

	return users, result.Error
}
