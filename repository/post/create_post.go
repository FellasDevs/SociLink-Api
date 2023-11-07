package postrepository

import (
	"SociLinkApi/models"

	"gorm.io/gorm"
)

func CreatePost(post *models.Post, db *gorm.DB) error {
	result := db.Create(&post)

	return result.Error
}
