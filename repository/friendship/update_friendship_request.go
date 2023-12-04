package friendshiprepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func UpdateFriendshipRequest(friendship models.Friendship, db *gorm.DB) error {
	result := db.Save(&friendship)

	return result.Error
}
