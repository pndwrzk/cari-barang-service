package usecase

import "github.com/pndwrzk/cari-barang-service/internal/product/repository"

type ProductUsecase interface {
}

type productUsecase struct {
	repository repository.ProductRepository
}

// NewProductUsecase creates a new instance of ProductUsecase.
func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return &productUsecase{repository}
}
