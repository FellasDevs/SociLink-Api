package database

import (
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	db, err := GetDbConnection()

	return db, err
}
