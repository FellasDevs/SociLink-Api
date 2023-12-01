package friendshiprepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFriendship(friendship *models.Friendship, db *gorm.DB) error {
	result := db.Preload(clause.Associations).First(&friendship)

	return result.Error
}
