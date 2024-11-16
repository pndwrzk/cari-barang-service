package dependencyinjection

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	categoryHandler "github.com/pndwrzk/cari-barang-service/internal/category/delivery/http"
	categoryRepository "github.com/pndwrzk/cari-barang-service/internal/category/repository"
	categoryUsecase "github.com/pndwrzk/cari-barang-service/internal/category/usecase"
	productHandler "github.com/pndwrzk/cari-barang-service/internal/product/delivery/http"
	productRepository "github.com/pndwrzk/cari-barang-service/internal/product/repository"
	productUsecase "github.com/pndwrzk/cari-barang-service/internal/product/usecase"
)

// Dependencyinjection initializes and injects the dependencies for category module.
func Dependencyinjection(db *gorm.DB, app *fiber.App) {
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryUsecase := categoryUsecase.NewCategoryUsecase(categoryRepo)
	categoryHandler.NewCategoryHandler(categoryUsecase).RegisterRoutes(app)

	productRepo := productRepository.NewProductRepository(db)
	productUsecase := productUsecase.NewProductUsecase(productRepo)
	productHandler.NewProductHandler(productUsecase).RegisterRoutes(app)
}
