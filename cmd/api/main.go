package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/pndwrzk/cari-barang-service/config/database"
	"github.com/pndwrzk/cari-barang-service/config/migration"
	dependencyinjection "github.com/pndwrzk/cari-barang-service/pkg/dependency_injection"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	db := database.DBConnect()
	migration.PgMigration(db)

	dependencyinjection.Dependencyinjection(db, app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := app.Listen(":" + port); err != nil {
		fmt.Println(err)
	}
}
