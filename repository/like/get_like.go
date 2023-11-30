package likerepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetLike(like *models.Like, db *gorm.DB) error {
	result := db.Where(&like).First(&like)

	return result.Error
}
