package friendshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllFriendships(userId uuid.UUID, db *gorm.DB) ([]models.Friendship, error) {
	var friendships []models.Friendship

	query := db.Preload(clause.Associations)

	query = query.Where("(user_id = ? OR friend_id = ?) AND accepted = ?", userId, userId, true)

	result := query.Find(&friendships)

	return friendships, result.Error
}
