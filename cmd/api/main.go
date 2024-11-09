package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/pndwrzk/cari-barang-service/config/database"
	categoryHandler "github.com/pndwrzk/cari-barang-service/internal/category/delivery/http"
	categoryRepository "github.com/pndwrzk/cari-barang-service/internal/category/repository"
	categoryUseCase "github.com/pndwrzk/cari-barang-service/internal/category/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	db := database.DBConnect()
	categoryRepository := categoryRepository.NewUserRepository(db)
	categoryUseCase := categoryUseCase.NewCategortyUseCase(categoryRepository)
	categoryHandler.NewCategoryHandler(categoryUseCase).Route(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting Fiber server: ", err)
	}

}
