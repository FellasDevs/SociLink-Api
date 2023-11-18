package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUsersByNameOrNickname(search string, db *gorm.DB) ([]models.User, error) {
	var users []models.User

	result := db.Where("name ILIKE ?", "%"+search+"%").Or("nickname ILIKE ?", "%"+search+"%").Find(&users)

	return users, result.Error
}
