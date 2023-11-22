package frienshiprepository

import (
	"SociLinkApi/dto"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendships(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) (types.FriendshipListing, error) {
	friendships := types.FriendshipListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	query := db.Preload(clause.Associations)

	query = query.Where("(user_id = ? OR friend_id = ?) AND accepted = ?", userId, userId, true)

	utils.UsePagination(query, &friendships.PaginationResponse)

	result := query.Find(&friendships.Friendships).Scan(&friendships.PaginationResponse)

	return friendships, result.Error
}
