package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pndwrzk/cari-barang-service/internal/product/dto"
	"github.com/pndwrzk/cari-barang-service/internal/product/usecase"
	"github.com/pndwrzk/cari-barang-service/pkg/constants"
	"github.com/pndwrzk/cari-barang-service/pkg/response"
	"github.com/pndwrzk/cari-barang-service/pkg/utils"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func (handler *ProductHandler) RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/products", handler.fetchProduct)
	apiV1.Get("/products/:slug_category/category", handler.fetchProductByCategory)
	apiV1.Post("/products", handler.addProduct)
	apiV1.Put("/products/:id", handler.editProduct)
	apiV1.Delete("/products/:id", handler.removeProduct)
	apiV1.Patch("/product/:id/status", handler.editProductStatus)

}

func (handler *ProductHandler) fetchProduct(app *fiber.Ctx) error {
	var pagging response.Pagination
	data, err := handler.usecase.RetrieveProduct()
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_READ_ERROR,
			err.Error(),
		))
	}
	return app.Status(fiber.StatusOK).JSON(response.FetchAll(
		constants.SUCCESS_STATUS,
		constants.MESSAGE_READ_SUCCESS,
		data,
		pagging,
	))
}

func (handler *ProductHandler) fetchProductByCategory(app *fiber.Ctx) error {
	slugCategory := app.Params("slug_category")
	err := handler.usecase.RetrieveProductByCategory(slugCategory)

	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_READ_ERROR,
			err.Error(),
		))
	}
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

func (handler *ProductHandler) removeProduct(app *fiber.Ctx) error {
	idParam := app.Params("id")
	uintID, err := utils.StrToUint(idParam)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_INVALID_PARMS,
			err.Error(),
		))
	}

	if err = handler.usecase.Destroyproduct(uintID); err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_DELETE_ERROR,
			err.Error(),
		))
	}

	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.SUCCESS_STATUS, constants.MESSAGE_DELETE_SUCCESS))
}

func (handler *ProductHandler) editProduct(app *fiber.Ctx) error {
	idParam := app.Params("id")
	var requestBody dto.RequestBodyProduct
	uintID, err := utils.StrToUint(idParam)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_INVALID_PARMS,
			err.Error(),
		))
	}

	if err := app.BodyParser(&requestBody); err != nil {
		return app.Status(fiber.ErrBadRequest.Code).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_VALIDATION,
			err.Error(),
		))
	}
	if err := handler.usecase.ModifyProduct(uintID, requestBody); err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_UPDATE_ERROR,
			err.Error(),
		))
	}
	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.SUCCESS_STATUS, constants.MESSAGE_DELETE_SUCCESS))
}

func (handler *ProductHandler) editProductStatus(app *fiber.Ctx) error {
	var requestBody dto.RequestUpdateStatusProduct
	idParam := app.Params("id")
	uintID, err := utils.StrToUint(idParam)

	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_INVALID_PARMS,
			err.Error(),
		))
	}

	if err := app.BodyParser(&requestBody); err != nil {
		return app.Status(fiber.ErrBadRequest.Code).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_VALIDATION,
			err.Error(),
		))
	}
	if err = handler.usecase.ModifyProductStatus(uintID, requestBody); err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_UPDATE_ERROR,
			err.Error(),
		))
	}
	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.SUCCESS_STATUS, constants.MESSAGE_UPDATE_SUCCESS))

}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase}
}
