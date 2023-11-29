package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDbConnection() (*gorm.DB, error) {
	fmt.Println(os.Getenv("DB_STRING"))
	return gorm.Open(postgres.Open(os.Getenv("DB_STRING")), &gorm.Config{})
}
