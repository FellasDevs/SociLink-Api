package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendships(userId uuid.UUID, db *gorm.DB) ([]models.Friendship, error) {
	var friends []models.Friendship

	result := db.Preload(clause.Associations).Where(models.Friendship{UserID: userId, Accepted: true}).Or(models.Friendship{FriendID: userId, Accepted: true}).Find(&friends)

	return friends, result.Error
}
