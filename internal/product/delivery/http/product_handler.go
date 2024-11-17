package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pndwrzk/cari-barang-service/internal/product/dto"
	"github.com/pndwrzk/cari-barang-service/internal/product/usecase"
	"github.com/pndwrzk/cari-barang-service/pkg/constants"
	"github.com/pndwrzk/cari-barang-service/pkg/response"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func (handler *ProductHandler) RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/products", handler.fetchProduct)
	apiV1.Post("/products", handler.addProduct)
}

func (handler *ProductHandler) fetchProduct(app *fiber.Ctx) error {
	return nil
}

func (handler *ProductHandler) addProduct(app *fiber.Ctx) error {
	var requestBody dto.RequestBodyProduct
	if err := app.BodyParser(&requestBody); err != nil {
		return app.Status(fiber.ErrBadRequest.Code).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_VALIDATION,
			err.Error(),
		))
	}

	if err := handler.usecase.StoreProduct(requestBody); err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_CREATE_ERROR,
			err.Error(),
		))
	}

	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.Created_STATUS, constants.MESSAGE_CREATE_SUCCESS))

}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase}
}
