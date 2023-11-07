package postrepository

import (
	"SociLinkApi/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPost(id uuid.UUID, db *gorm.DB) (models.Post, error) {
	var post models.Post

	result := db.First(&post, id)

	return post, result.Error
}
