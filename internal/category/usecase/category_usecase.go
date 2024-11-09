package usecase

import "github.com/pndwrzk/cari-barang-service/internal/category/repository"

type CategoryUseCase interface {
	AddCategory()
}

type categoryUseCase struct {
	repository repository.CategoryRepository
}

func NewCategortyUseCase(repository repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{repository}
}

// AddCategory implements CategoryUseCase.
func (c *categoryUseCase) AddCategory() {
	panic("unimplemented")
}
