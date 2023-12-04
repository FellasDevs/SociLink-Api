package dto

import (
	"github.com/google/uuid"
	"time"
)

type NotificationResponseDto struct {
	Text        string
	SubjectID   uuid.UUID
	SubjectType string
	CreatedAt   time.Time
}
