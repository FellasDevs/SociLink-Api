package frienshiprepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AnswerFriendshipRequest(friendshipId uuid.UUID, answer bool, db *gorm.DB) error {
	result := db.Model(models.Friendship{ID: friendshipId}).Updates(models.Friendship{Accepted: answer, Pending: false})

	return result.Error
}
