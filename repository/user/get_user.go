package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetUser(user *models.User, db *gorm.DB) error {
	result := db.Preload(clause.Associations).First(&user)

	return result.Error
}
