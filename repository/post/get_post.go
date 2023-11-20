package postrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func GetPost(post *models.Post, db *gorm.DB) error {
	result := db.Preload(clause.Associations).First(&post)

	return result.Error
}
