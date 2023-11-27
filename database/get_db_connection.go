package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDbConnection() (*gorm.DB, error) {
	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	return gorm.Open(postgres.Open(url), &gorm.Config{})
}
