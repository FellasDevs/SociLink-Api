package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID         uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique;"`
	UserID     uuid.UUID `gorm:"->; type: UUID; not null"`
	User       User
	Content    string    `gorm:"not null"`
	Images     []string  `gorm:"type: text[]"`
	Visibility string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time
}
