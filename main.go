package main

import (
	"SociLinkApi/database"
	"SociLinkApi/middlewares"
	"SociLinkApi/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Falha em carregar arquivo .env")
	}

	db, err := database.GetDbConnection()
	if err != nil {
		fmt.Println("Falha ao conectar com o banco de dados")
		panic(err)
	}

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--migrate" {
		database.Migrate(db)
		return
	}

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	routes.SetRoutes(router, db)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
