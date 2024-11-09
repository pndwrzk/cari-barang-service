package repository

import "gorm.io/gorm"

type CategoryRepository interface {
	FindAll()
}

type categoryRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// FindAll implements CategoryRepository.
func (c *categoryRepository) FindAll() {
	panic("unimplemented")
}
