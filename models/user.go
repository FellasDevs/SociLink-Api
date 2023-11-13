package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique"`
	Name      string    `gorm:"type: VARCHAR(50); not null; index"`
	Email     string    `gorm:"type: VARCHAR(50); not null; unique"`
	Password  string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Nickname  string    `gorm:"type: VARCHAR(50)"`
	Country   string    `gorm:"type: VARCHAR(50); not null"`
	City      string    `gorm:"type: VARCHAR(50)"`
	Picture   string
	Banner    string
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	Friends   []*User `gorm:"many2many:friendships"`
}
