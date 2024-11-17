package repository

import (
	"gorm.io/gorm"

	"github.com/pndwrzk/cari-barang-service/internal/product/entity"
)

type ProductRepository interface {
	BeginTransaction() *gorm.DB
	InsertProduct(tx *gorm.DB, product entity.Product) (*entity.Product, error)
	InsertProductCategory(tx *gorm.DB, productCategory entity.ProductCategory) error
}

type productRepository struct {
	db *gorm.DB
}

// InsertProductCategory implements ProductRepository.
func (repository *productRepository) InsertProductCategory(tx *gorm.DB, productCategory entity.ProductCategory) error {
	return tx.Create(&productCategory).Error
}

// InsertProduct implements ProductRepository.
func (repository *productRepository) InsertProduct(tx *gorm.DB, product entity.Product) (*entity.Product, error) {
	return &product, tx.Create(&product).Error
}

// BeginTransaction implements ProductRepository.
func (repository *productRepository) BeginTransaction() *gorm.DB {
	return repository.db.Begin()
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}
