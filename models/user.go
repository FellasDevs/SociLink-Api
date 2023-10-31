package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"->; type: UUID DEFAULT gen_random_uuid(); not null; unique; primaryKey"`
	Name      string    `gorm:"type: VARCHAR(50); not null; index"`
	Email     string    `gorm:"type: VARCHAR(50); not null; unique"`
	Password  string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}
