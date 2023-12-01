package notification

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func CreateNotification(notification *models.Notification, db *gorm.DB) error {
	result := db.Create(&notification)

	return result.Error
}
