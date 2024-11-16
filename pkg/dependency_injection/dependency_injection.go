package dependencyinjection

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	categoryHandler "github.com/pndwrzk/cari-barang-service/internal/category/delivery/http"
	categoryRepository "github.com/pndwrzk/cari-barang-service/internal/category/repository"
	categoryUseCase "github.com/pndwrzk/cari-barang-service/internal/category/usecase"
)

// Dependencyinjection initializes and injects the dependencies for category module.
func Dependencyinjection(db *gorm.DB, app *fiber.App) {
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler.NewCategoryHandler(categoryUseCase).Route(app)
}
