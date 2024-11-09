package migration

import (
	"log"

	"gorm.io/gorm"

	entityCategory "github.com/pndwrzk/cari-barang-service/internal/category/entity"
)

func PgMigration(db *gorm.DB) {
	err := db.AutoMigrate(&entityCategory.Category{})
	if err != nil {
		log.Fatalf("Error migrating category table: %v", err)
	}
}
