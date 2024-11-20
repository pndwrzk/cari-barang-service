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
	RetrieveProduct() ([]*dto.ResponseGetProduct, error)
	RetrieveProductByCategory(slug string) error
	StoreProduct(bodyRequest dto.RequestBodyProduct) error

	Destroyproduct(id uint) error

	ModifyProduct(id uint, bodyRequest dto.RequestBodyProduct) error

	ModifyProductStatus(id uint, bodyRequest dto.RequestUpdateStatusProduct) error
}

type productUsecase struct {
	repository repository.ProductRepository
}

// ModifyProduct implements ProductUsecase.
func (usecase *productUsecase) ModifyProduct(id uint, bodyRequest dto.RequestBodyProduct) error {
	tx := usecase.repository.BeginTransaction()
	if tx == nil {
		return errors.New(constants.MESSAGE_ADD_PRODUCT_ERROR)
	}
	product := entity.Product{
		ID:          id,
		Name:        bodyRequest.Name,
		Slug:        utils.GenerateSlug(bodyRequest.Name),
		Description: bodyRequest.Description,
	}
	newProduct, err := usecase.repository.InsertProduct(tx, product)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := usecase.repository.DeleteProductCategory(tx, newProduct.ID); err != nil {
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

// ModifyProductStatus implements ProductUsecase.
func (usecase *productUsecase) ModifyProductStatus(id uint, bodyRequest dto.RequestUpdateStatusProduct) error {
	data, err := usecase.repository.ReadProductById(id)
	if err != nil {
		return err
	}

	if data.IsActive == bodyRequest.Status {
		return errors.New("")
	}
	data.IsActive = bodyRequest.Status

	if err = usecase.repository.UpdateProduct(*data); err != nil {
		return err
	}
	return nil
}

// Destroyproduct implements ProductUsecase.
func (usecase *productUsecase) Destroyproduct(id uint) error {
	data, err := usecase.repository.ReadProductById(id)
	if err != nil {
		return err
	}
	if data.IsActive == 1 {
		return errors.New(constants.MESSAGE_UPDATE_STATUS_ERROR)
	}
	return usecase.repository.DeleteProduct(*data)
}

// RetrieveProductByCategory implements ProductUsecase.
func (usecase *productUsecase) RetrieveProductByCategory(slug string) error {
	panic("unimplemented")
}

// RetrieveProduct implements ProductUsecase.
func (usecase *productUsecase) RetrieveProduct() ([]*dto.ResponseGetProduct, error) {
	data, err := usecase.repository.ReadProduct()
	if err != nil {
		return nil, err
	}
	return data, nil

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
