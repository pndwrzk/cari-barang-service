package repository

import (
	"gorm.io/gorm"

	"github.com/pndwrzk/cari-barang-service/internal/category/dto"
	"github.com/pndwrzk/cari-barang-service/internal/category/entity"
)

type CategoryRepository interface {
	ReadCategory() ([]*dto.ResponseGetCategory, error)
	CreateCategory(bodyRequest entity.Category) error
	UpdateCategory(bodyRequest entity.Category) error
	ReadByIdCategory(id uint) (*entity.Category, error)
	DeleteByIdCategory(category *entity.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

// DeleteByIdCategory implements CategoryRepository.
func (repository *categoryRepository) DeleteByIdCategory(category *entity.Category) error {
	return repository.db.Delete(&category).Error
}

// ReadByIdCategory implements CategoryRepository.
func (repository *categoryRepository) ReadByIdCategory(id uint) (*entity.Category, error) {
	var category *entity.Category
	if err := repository.db.Table("category").First(&category).Where("id ?", id).Error; err != nil {
		return nil, err
	}
	return category, nil
}

// UpdateCategory implements CategoryRepository.
func (repository *categoryRepository) UpdateCategory(bodyRequest entity.Category) error {
	result := repository.db.Save(&bodyRequest)
	return result.Error
}

// Create implements CategoryRepository.
func (repository *categoryRepository) CreateCategory(bodyRequest entity.Category) error {
	result := repository.db.Create(&bodyRequest)
	return result.Error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) ReadCategory() ([]*dto.ResponseGetCategory, error) {
	var categories []*dto.ResponseGetCategory
	if err := c.db.Table("category").Find(&categories).Where("is_active ?", 1).Error; err != nil {
		return nil, err
	}
	return categories, nil

}
