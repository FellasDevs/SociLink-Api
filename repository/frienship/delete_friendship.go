package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteFriendship(friendshipId uuid.UUID, db *gorm.DB) error {
	result := db.Delete(models.Friendship{ID: friendshipId})

	return result.Error
}
