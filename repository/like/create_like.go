package likerepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func CreateLike(like *models.Like, db *gorm.DB) error {
	result := db.Create(like)

	return result.Error
}
