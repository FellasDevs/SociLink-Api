package likerepository

import (
	"SociLinkApi/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCommentLikes(commentId uuid.UUID, db *gorm.DB) ([]models.CommentLike, error) {
	var likes []models.CommentLike
	result := db.Where("comment_id = ?", commentId).Find(&likes)
	return likes, result.Error
}
