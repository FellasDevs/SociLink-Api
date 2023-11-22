package frienshiprepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendshipRequests(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) (types.FriendshipListing, error) {
	friendships := types.FriendshipListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	query := db.Preload(clause.Associations).Where(models.Friendship{FriendID: userId, Pending: true})

	utils.UsePagination(query, &friendships.PaginationResponse)

	result := query.Order("created_at desc").Find(&friendships.Friendships).Scan(&friendships.PaginationResponse)

	return friendships, result.Error
}
