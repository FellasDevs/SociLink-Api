package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendships(userId uuid.UUID, db *gorm.DB) ([]models.Friendship, error) {
	var friendships []models.Friendship

	query := db.Preload(clause.Associations)

	query = query.Where(models.Friendship{UserID: userId, Accepted: true}).Or(models.Friendship{FriendID: userId, Accepted: true})

	result := query.Find(&friendships)

	return friendships, result.Error
}
