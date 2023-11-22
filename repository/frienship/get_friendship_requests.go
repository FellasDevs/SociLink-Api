package frienshiprepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendshipRequests(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Friendship, error) {
	var friendships []models.Friendship

	query := db.Preload(clause.Associations).Where(models.Friendship{FriendID: userId, Pending: true})

	utils.UsePagination(query, pagination)

	result := query.Order("created_at desc").Find(&friendships)

	return friendships, result.Error
}
