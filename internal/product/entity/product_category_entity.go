package entity

import "github.com/pndwrzk/cari-barang-service/internal/category/entity"

type ProductCategory struct {
	ID             uint             `gorm:"primaryKey;column:id"`
	CategoryID     *uint            `gorm:"column:category_id;onDelete:CASCADE"`
	ProductID      *uint            `gorm:"column:product_id;onDelete:CASCADE"`
	CategoryParent *entity.Category `gorm:"foreignKey:CategoryID;references:ID;constraint:onDelete:CASCADE"`
	ProductParent  *Product         `gorm:"foreignKey:ProductID;references:ID;constraint:onDelete:CASCADE"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}
