package usecase

import (
	"errors"

	"github.com/pndwrzk/cari-barang-service/internal/product/dto"
	"github.com/pndwrzk/cari-barang-service/internal/product/entity"
	"github.com/pndwrzk/cari-barang-service/internal/product/repository"
	"github.com/pndwrzk/cari-barang-service/pkg/constants"
	"github.com/pndwrzk/cari-barang-service/pkg/utils"
)

type ProductUsecase interface {
	StoreProduct(bodyRequest dto.RequestBodyProduct) error
}

type productUsecase struct {
	repository repository.ProductRepository
}

// StoreProduct implements ProductUsecase.
func (usecase *productUsecase) StoreProduct(bodyRequest dto.RequestBodyProduct) error {
	tx := usecase.repository.BeginTransaction()
	if tx == nil {
		return errors.New(constants.MESSAGE_ADD_PRODUCT_ERROR)
	}

	product := entity.Product{
		Name:        bodyRequest.Name,
		Slug:        utils.GenerateSlug(bodyRequest.Name),
		Description: bodyRequest.Description,
	}

	newProduct, err := usecase.repository.InsertProduct(tx, product)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, value := range bodyRequest.CategoryID {
		categoryId := uint(value)
		productCategory := entity.ProductCategory{
			ProductID:  &newProduct.ID,
			CategoryID: &categoryId,
		}
		if err := usecase.repository.InsertProductCategory(tx, productCategory); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// NewProductUsecase creates a new instance of ProductUsecase.
func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return &productUsecase{repository}
}
