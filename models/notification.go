package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique;"`
	UserID      uuid.UUID `gorm:"type: UUID; not null"`
	User        User
	Text        string    `gorm:"type: VARCHAR(50); not null"`
	SubjectID   uuid.UUID `gorm:"type: UUID"`
	SubjectType string    `gorm:"type: VARCHAR(15)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//func (notification Notification) toResponseDto() dto.NotificationResponseDto {
//	return dto.NotificationResponseDto{
//		Text:        notification.Text,
//		SubjectID:   notification.SubjectID,
//		SubjectType: notification.SubjectType,
//		CreatedAt:   notification.CreatedAt,
//	}
//}
