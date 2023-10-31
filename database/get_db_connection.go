package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDbConnection() (*gorm.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")

	return gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
}
