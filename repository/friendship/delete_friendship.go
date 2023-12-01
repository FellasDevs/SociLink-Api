package friendshiprepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func DeleteFriendship(friendship *models.Friendship, db *gorm.DB) error {
	result := db.Delete(&friendship)

	return result.Error
}
