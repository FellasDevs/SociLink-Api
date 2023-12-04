package notification

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetNotifications(userId uuid.UUID, db *gorm.DB) ([]models.Notification, error) {
	var notifications []models.Notification

	result := db.Where("user_id = ?", userId).Find(&notifications)

	return notifications, result.Error
}
