package frienshiprepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendships(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Friendship, error) {
	var friendships []models.Friendship

	query := db.Preload(clause.Associations)

	query = query.Where(models.Friendship{UserID: userId, Accepted: true}).Or(models.Friendship{FriendID: userId, Accepted: true})

	utils.UsePagination(query, pagination)

	result := query.Find(&friendships)

	return friendships, result.Error
}
