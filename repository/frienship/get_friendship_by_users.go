package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetFriendshipByUsers(userId uuid.UUID, friendId uuid.UUID, db *gorm.DB) (models.Friendship, error) {
	var friendship models.Friendship

	result := db.Where("accepted = ?", true).Where("user_id = ? AND friend_id = ?", userId, friendId).Or("user_id = ? AND friend_id = ?", friendId, userId).First(&friendship)

	return friendship, result.Error
}
