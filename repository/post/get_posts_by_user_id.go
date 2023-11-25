package postrepository

import (
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPostsByUserId(userId uuid.UUID, posts *types.PostListing, db *gorm.DB) error {
	query := db.Preload("User")

	query = query.Where("user_id = ?", userId)

	utils.UsePagination(query, "posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at", &posts.PaginationResponse)

	result := query.Order("created_at desc").Find(&posts.Posts).Scan(&posts.PaginationResponse)

	return result.Error
}
