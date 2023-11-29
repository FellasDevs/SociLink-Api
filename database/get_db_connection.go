package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDbConnection() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(os.Getenv("DB_STRING")), &gorm.Config{})
}
