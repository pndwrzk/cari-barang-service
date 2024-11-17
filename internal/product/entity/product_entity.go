package entity

type Product struct {
	ID          uint   `gorm:"primaryKey;column:id"`
	Name        string `gorm:"column:name;not null"`
	Slug        string `gorm:"column:slug;unique;not null"`
	IsActive    int    `gorm:"column:is_active;default:1"`
	Description string `gorm:"column:description"`
	CreatedAt   uint   `gorm:"column:created_at;autoUpdateTime;type:bigint"`
	UpdatedAt   uint   `gorm:"column:updated_at;autoUpdateTime;type:bigint"`
}

func (Product) TableName() string {
	return "product"
}
