package models

import (
	"github.com/google/uuid"
	"time"
)

type Friendship struct {
	ID        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique"`
	UserID    uuid.UUID `gorm:"type: UUID; not null; check: user_id <> friend_id;"`
	FriendID  uuid.UUID `gorm:"type: UUID; not null"`
	User      User
	Friend    User
	Accepted  bool `gorm:"not null; default: false"`
	Pending   bool `gorm:"not null; default: true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
