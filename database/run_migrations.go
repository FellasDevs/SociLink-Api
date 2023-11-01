package database

import (
	"SociLinkApi/models"
)

func main() {
	db, err := GetDbConnection()
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}
}
