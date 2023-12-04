package models

import (
	"time"

	"github.com/google/uuid"
)

type CommentLike struct {
	ID        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique;"`
	UserID    uuid.UUID `gorm:"type: UUID; not null"`
	User      User
	CommentID uuid.UUID `gorm:"type: UUID; not null"`
	Comment   Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}
