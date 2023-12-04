package likerepository

import (
	"SociLinkApi/models"

	"gorm.io/gorm"
)

func DeleteCommentLike(like *models.CommentLike, db *gorm.DB) error {
	result := db.Where(like).Delete(&models.CommentLike{})
	return result.Error
}
