package likerepository

import (
	"SociLinkApi/models"

	"gorm.io/gorm"
)

func CreateCommentLike(like *models.CommentLike, db *gorm.DB) error {
	result := db.Create(&like)
	return result.Error
}
