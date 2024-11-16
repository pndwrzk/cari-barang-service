package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pndwrzk/cari-barang-service/internal/product/usecase"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func (handler *ProductHandler) RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/products", handler.fetchProduct)
}

func (handler *ProductHandler) fetchProduct(app *fiber.Ctx) error {
	return nil
}
func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase}
}
