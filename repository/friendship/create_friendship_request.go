package friendshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFriendshipRequest(userId uuid.UUID, friendId uuid.UUID, db *gorm.DB) error {
	result := db.Create(&models.Friendship{UserID: userId, FriendID: friendId})

	return result.Error
}
