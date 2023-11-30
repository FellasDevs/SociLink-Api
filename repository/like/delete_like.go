package likerepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func DeleteLike(like *models.Like, db *gorm.DB) error {
	result := db.Where(like).Delete(like)

	return result.Error
}
