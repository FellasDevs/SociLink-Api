package postrepository

import (
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetDeletedPost(post *models.Post, userId *uuid.UUID, db *gorm.DB) error {
	query := db.Preload(clause.Associations)

	query = query.Where("visibility = ?", authtypes.Public)
	query = query.Where("deleted = ?", true)

	if userId != nil {
		utils.UseJoinPostsAndFriendships(query)

		query = query.Or("posts.user_id = ?", userId)
		query = query.Where("deleted = ?", true)

		query = query.Or("visibility = ?", authtypes.Friends)
		query = query.Where("deleted = ?", true)
		utils.UseAreUserAndPostOwnerFriends(query, *userId)
	}

	result := query.First(&post)

	return result.Error
}
