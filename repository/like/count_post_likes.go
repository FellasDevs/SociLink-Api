package likerepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CountPostLikes(postId uuid.UUID, db *gorm.DB) (int64, error) {
	var count int64

	result := db.Model(&models.Like{}).Where("post_id = ?", postId).Count(&count)

	return count, result.Error
}
