package postrepository

import (
	"SociLinkApi/models"

	"gorm.io/gorm"
)

func UpdatePost(post *models.Post, db *gorm.DB) error {
	result := db.Save(&post)

	return result.Error
}
