package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique;"`
	UserID    uuid.UUID `gorm:"type: UUID; not null"`
	User      User
	PostID    uuid.UUID `gorm:"type: UUID; not null"`
	Post      Post
	Content   string `gorm:"not null; type: text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
