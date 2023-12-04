package likerepository

import (
	"SociLinkApi/models"

	"gorm.io/gorm"
)

func GetCommentLike(like *models.CommentLike, db *gorm.DB) error {
	result := db.Where(&like).First(&like)
	return result.Error
}
