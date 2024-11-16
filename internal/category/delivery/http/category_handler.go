package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pndwrzk/cari-barang-service/internal/category/dto"
	"github.com/pndwrzk/cari-barang-service/internal/category/usecase"
	"github.com/pndwrzk/cari-barang-service/pkg/constants"
	"github.com/pndwrzk/cari-barang-service/pkg/response"
	"github.com/pndwrzk/cari-barang-service/pkg/utils"
)

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

// Route register all the routes for category module
func (handler *CategoryHandler) RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/categories", handler.fetchCategory)
	apiV1.Post("/categories", handler.addCategory)
	apiV1.Put("/categories/:id", handler.editCategory)
	apiV1.Delete("/categories/:id", handler.removeCategory)

}

func (handler *CategoryHandler) fetchCategory(app *fiber.Ctx) error {
	data, err := handler.usecase.RetrieveCategory()
	var pagging response.Pagination
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_READ_SUCCESS,
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

func (handler *CategoryHandler) addCategory(app *fiber.Ctx) error {
	var requestBody dto.RequestBodyCategory
	if err := app.BodyParser(&requestBody); err != nil {
		return app.Status(fiber.ErrBadRequest.Code).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_VALIDATION,
			err.Error(),
		))
	}
	if err := handler.usecase.StoreCategory(requestBody); err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_CREATE_ERROR,
			err.Error(),
		))
	}

	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.Created_STATUS, constants.MESSAGE_CREATE_SUCCESS))

}

func (handler *CategoryHandler) editCategory(app *fiber.Ctx) error {
	idParam := app.Params("id")
	var requestBody dto.RequestBodyCategory

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

	err = handler.usecase.ModifyCategory(uintID, requestBody)

	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_UPDATE_ERROR,
			err.Error(),
		))
	}

	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.SUCCESS_STATUS, constants.MESSAGE_UPDATE_SUCCESS))
}

func (handler *CategoryHandler) removeCategory(app *fiber.Ctx) error {
	idParam := app.Params("id")
	uintID, err := utils.StrToUint(idParam)
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_INVALID_PARMS,
			err.Error(),
		))
	}
	if err := handler.usecase.DestroyCategory(uintID); err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(response.FailureProcess(
			constants.ERROR_STATUS,
			constants.MESSAGE_DELETE_ERROR,
			err.Error(),
		))
	}

	return app.Status(fiber.StatusCreated).JSON(response.ProcessDataMessageOnly(constants.SUCCESS_STATUS, constants.MESSAGE_DELETE_SUCCESS))

}

func NewCategoryHandler(usecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase}
}
