package postrepository

import (
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func GetPost(post *models.Post, userId *uuid.UUID, db *gorm.DB) error {
	query := db.Preload(clause.Associations)

	query = query.Where("visibility = ?", authtypes.Public)
	query = query.Where("deleted = ?", false)

	if userId != nil {
		utils.UseJoinPostsAndFriendships(query)

		query = query.Or("posts.user_id = ?", userId)
		query = query.Where("deleted = ?", false)

		query = query.Or("visibility = ?", authtypes.Friends)
		query = query.Where("deleted = ?", false)
		utils.UseAreUserAndPostOwnerFriends(query, *userId)
	}

	result := query.First(&post)

	return result.Error
}
