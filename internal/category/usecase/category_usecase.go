package usecase

import (
	"errors"

	"github.com/pndwrzk/cari-barang-service/internal/category/dto"
	"github.com/pndwrzk/cari-barang-service/internal/category/entity"
	"github.com/pndwrzk/cari-barang-service/internal/category/repository"
	"github.com/pndwrzk/cari-barang-service/pkg/utils"
)

type CategoryUsecase interface {
	RetrieveCategory() ([]*dto.ResponseGetCategory, error)
	StoreCategory(requestBody dto.RequestBodyCategory) error
	ModifyCategory(id uint, requestBody dto.RequestBodyCategory) error
	DestroyCategory(id uint) error
}

type categoryUsecase struct {
	repository repository.CategoryRepository
}

// DestroyCategory implements CategoryUsecase.
func (usecase *categoryUsecase) DestroyCategory(id uint) error {
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

// ModifyCategory implements CategoryUsecase.
func (usecase *categoryUsecase) ModifyCategory(id uint, requestBody dto.RequestBodyCategory) error {
	category := entity.Category{
		ID:       id,
		Name:     requestBody.Name,
		ParentID: requestBody.ParentID,
		Slug:     utils.GenerateSlug(requestBody.Name),
	}
	return usecase.repository.UpdateCategory(category)
}

// StoreCategory implements CategoryUsecase.
func (usecase *categoryUsecase) StoreCategory(requestBody dto.RequestBodyCategory) error {
	category := entity.Category{
		Name:     requestBody.Name,
		ParentID: requestBody.ParentID,
		Slug:     utils.GenerateSlug(requestBody.Name),
	}
	return usecase.repository.CreateCategory(category)
}

// RetrieveCategory implements CategoryUsecase.
func (usecase *categoryUsecase) RetrieveCategory() ([]*dto.ResponseGetCategory, error) {
	data, err := usecase.repository.ReadCategory()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// NewCategoryUsecase creates a new instance of CategoryUsecase.
func NewCategoryUsecase(repository repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repository}
}

// GetAllCategories retrieves all active categories.
