package entity

type Category struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name;not null"`
	Slug      string    `gorm:"column:slug;unique;not null"`
	IsActive  int       `gorm:"column:is_active;default:1"`
	ParentID  *uint     `gorm:"column:parent_id;index"`
	CreatedAt uint      `gorm:"column:created_at;autoUpdateTime;type:bigint"`
	UpdatedAt uint      `gorm:"column:updated_at;autoUpdateTime;type:bigint"`
	Parent    *Category `gorm:"foreignKey:ParentID;references:ID"`
}

func (Category) TableName() string {
	return "category"
}
