package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateUser(user *models.User, db *gorm.DB) error {
	result := db.Clauses(clause.Returning{}).Create(&user)

	return result.Error
}
