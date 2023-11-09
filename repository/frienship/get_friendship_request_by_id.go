package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendshipRequestById(friendshipId uuid.UUID, db *gorm.DB) (models.Friendship, error) {
	var friendship models.Friendship

	result := db.Preload(clause.Associations).First(&friendship, friendshipId)

	return friendship, result.Error
}
