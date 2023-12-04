package commentrepository

import (
	"SociLinkApi/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteCommentReply(replyId uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.CommentReply{ID: replyId})

	return result.Error
}
