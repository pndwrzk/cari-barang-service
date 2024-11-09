package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pndwrzk/cari-barang-service/internal/category/usecase"
)

type CategoryHandler struct {
	usecase usecase.CategoryUseCase
}

// Route implements CategoryHandler.
func (c *CategoryHandler) Route(app *fiber.App) {
	app.Get("/categories", c.FindAll) // Definisikan rute di dalam grup
}

// FindAll implements CategoryHandler.
func (c *CategoryHandler) FindAll(app *fiber.Ctx) error {
	return app.Status(500).SendString("Error fetching users")
}

func NewCategoryHandler(usecase usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{usecase}
}
