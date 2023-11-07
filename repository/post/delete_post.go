package postrepository

import (
	"SociLinkApi/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeletePost(postID uuid.UUID, db *gorm.DB) error {
	result := db.Where("id = ?", postID).Delete(&models.Post{})

	return result.Error
}
