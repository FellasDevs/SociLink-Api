package postrepository

import (
	"SociLinkApi/dto"
	authtypes "SociLinkApi/types/auth"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPostsByUsers(userIds []uuid.UUID, visibility authtypes.Visibility, pagination dto.PaginationRequestDto, db *gorm.DB) (types.PostListing, error) {
	posts := types.PostListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	query := db.Preload(clause.Associations)

	query = query.Where("user_id IN ?", userIds).Where("visibility IN ?", visibility.GetAllowedVisibilities())

	utils.UsePagination(query, &posts.PaginationResponse)

	result := query.Order("created_at desc").Find(&posts.Posts).Scan(&posts.PaginationResponse)

	return posts, result.Error
}
