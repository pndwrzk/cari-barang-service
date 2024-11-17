package migration

import (
	"log"

	"gorm.io/gorm"

	entityCategory "github.com/pndwrzk/cari-barang-service/internal/category/entity"
	entityProduct "github.com/pndwrzk/cari-barang-service/internal/product/entity"
)

func PgMigration(db *gorm.DB) {
	err := db.AutoMigrate(&entityCategory.Category{}, &entityProduct.Product{}, entityProduct.ProductCategory{})
	if err != nil {
		log.Fatalf("Error migrating category table: %v", err)
	}
}
