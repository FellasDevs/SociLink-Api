package commentrepository

import (
	"SociLinkApi/models"

	"gorm.io/gorm"
)

func CreateCommentReply(reply *models.CommentReply, db *gorm.DB) error {
	result := db.Create(&reply)

	return result.Error
}
