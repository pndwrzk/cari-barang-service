package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/pndwrzk/cari-barang-service/config/database"
	"github.com/pndwrzk/cari-barang-service/config/migration"
	categoryHandler "github.com/pndwrzk/cari-barang-service/internal/category/delivery/http"
	categoryRepository "github.com/pndwrzk/cari-barang-service/internal/category/repository"
	categoryUseCase "github.com/pndwrzk/cari-barang-service/internal/category/usecase"
)

func main() {
	godotenv.Load()

	app := fiber.New()
	db := database.DBConnect()
	migration.PgMigration(db)

	categoryRepository := categoryRepository.NewUserRepository(db)
	categoryUseCase := categoryUseCase.NewCategortyUseCase(categoryRepository)
	categoryHandler.NewCategoryHandler(categoryUseCase).Route(app)
	portApp := fmt.Sprintf(":%s", os.Getenv("PORT"))
	app.Listen(portApp)

}
