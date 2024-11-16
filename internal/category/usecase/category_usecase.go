package usecase

import (
	"errors"

	"github.com/pndwrzk/cari-barang-service/internal/category/dto"
	"github.com/pndwrzk/cari-barang-service/internal/category/entity"
	"github.com/pndwrzk/cari-barang-service/internal/category/repository"
	"github.com/pndwrzk/cari-barang-service/pkg/utils"
)

type CategoryUseCase interface {
	RetrieveCategory() ([]*dto.ResponseGetCategory, error)
	StoreCategory(requestBody dto.RequestBodyCategory) error
	ModifyCategory(id uint, requestBody dto.RequestBodyCategory) error
	DestroyCategory(id uint) error
}

type categoryUseCase struct {
	repository repository.CategoryRepository
}

// DestroyCategory implements CategoryUseCase.
func (usecase *categoryUseCase) DestroyCategory(id uint) error {
	// check status category
	data, err := usecase.repository.ReadByIdCategory(id)
	if err != nil {
		return err
	}
	if data.IsActive == 1 {
		return errors.New("failed to delete because the status category is currently active")
	}
	return usecase.repository.DeleteByIdCategory(data)
}

// ModifyCategory implements CategoryUseCase.
func (usecase *categoryUseCase) ModifyCategory(id uint, requestBody dto.RequestBodyCategory) error {
	category := entity.Category{
		ID:       id,
		Name:     requestBody.Name,
		ParentID: requestBody.ParentID,
		Slug:     utils.GenerateSlug(requestBody.Name),
	}
	return usecase.repository.UpdateCategory(category)
}

// StoreCategory implements CategoryUseCase.
func (usecase *categoryUseCase) StoreCategory(requestBody dto.RequestBodyCategory) error {
	category := entity.Category{
		Name:     requestBody.Name,
		ParentID: requestBody.ParentID,
		Slug:     utils.GenerateSlug(requestBody.Name),
	}
	return usecase.repository.CreateCategory(category)
}

// RetrieveCategory implements CategoryUseCase.
func (usecase *categoryUseCase) RetrieveCategory() ([]*dto.ResponseGetCategory, error) {
	data, err := usecase.repository.ReadCategory()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// NewCategoryUseCase creates a new instance of CategoryUseCase.
func NewCategoryUseCase(repository repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{repository}
}

// GetAllCategories retrieves all active categories.
