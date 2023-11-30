package commentrepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteComment(commentId uuid.UUID, db *gorm.DB) error {
	result := db.Delete(&models.Comment{ID: commentId})

	return result.Error
}
