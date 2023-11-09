package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendshipRequests(userId uuid.UUID, db *gorm.DB) ([]models.Friendship, error) {
	var friendships []models.Friendship

	result := db.Preload(clause.Associations).Where(models.Friendship{FriendID: userId, Pending: true}).Find(&friendships)

	return friendships, result.Error
}
