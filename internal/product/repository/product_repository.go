package repository

import (
	"gorm.io/gorm"

	"github.com/pndwrzk/cari-barang-service/internal/product/dto"
	"github.com/pndwrzk/cari-barang-service/internal/product/entity"
)

type ProductRepository interface {
	BeginTransaction() *gorm.DB

	ReadProduct() ([]*dto.ResponseGetProduct, error)
	ReadProductById(id uint) (*entity.Product, error)
	InsertProduct(tx *gorm.DB, product entity.Product) (*entity.Product, error)
	InsertProductCategory(tx *gorm.DB, productCategory entity.ProductCategory) error
	DeleteProductCategory(tx *gorm.DB, id uint) error

	DeleteProduct(product entity.Product) error

	UpdateProduct(bodyRequest entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

// DeleteProductCategory implements ProductRepository.
func (repository *productRepository) DeleteProductCategory(tx *gorm.DB, id uint) error {
	return tx.Table("product_category").Where("product_id = ?", id).Error
}

// UpdateProduct implements ProductRepository.
func (repository *productRepository) UpdateProduct(bodyRequest entity.Product) error {
	result := repository.db.Save(&bodyRequest)
	return result.Error
}

// ReadProductById implements ProductRepository.
func (repository *productRepository) ReadProductById(id uint) (*entity.Product, error) {
	var product *entity.Product
	if err := repository.db.Table("product").First(&product).Where("id ?", id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct implements ProductRepository.
func (repository *productRepository) DeleteProduct(product entity.Product) error {
	panic("unimplemented")
}

// ReadProduct implements ProductRepository.
func (repository *productRepository) ReadProduct() ([]*dto.ResponseGetProduct, error) {
	var products []*dto.ResponseGetProduct
	if err := repository.db.Table("product").Find(&products).Where("is_active ?", 1).Error; err != nil {
		return nil, err
	}
	return products, nil

}

// InsertProductCategory implements ProductRepository.
func (repository *productRepository) InsertProductCategory(tx *gorm.DB, productCategory entity.ProductCategory) error {
	return tx.Create(&productCategory).Error
}

// InsertProduct implements ProductRepository.
func (repository *productRepository) InsertProduct(tx *gorm.DB, product entity.Product) (*entity.Product, error) {
	return &product, tx.Save(&product).Error
}

// BeginTransaction implements ProductRepository.
func (repository *productRepository) BeginTransaction() *gorm.DB {
	return repository.db.Begin()
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}
