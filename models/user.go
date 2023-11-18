package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique"`
	Name      string    `gorm:"type: VARCHAR(50); not null; index"`
	Email     string    `gorm:"type: VARCHAR(50); not null; unique"`
	Nickname  string    `gorm:"type: VARCHAR(50); not null; unique"`
	Password  string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Country   string    `gorm:"type: VARCHAR(50)"`
	City      string    `gorm:"type: VARCHAR(50)"`
	Picture   string
	Banner    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Friends   []*User `gorm:"many2many:friendships"`
}
