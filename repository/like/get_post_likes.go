package likerepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPostLikes(postId uuid.UUID, db *gorm.DB) ([]models.Like, error) {
	var likes []models.Like

	result := db.Where("post_id = ?", postId).Find(&likes)

	return likes, result.Error
}
