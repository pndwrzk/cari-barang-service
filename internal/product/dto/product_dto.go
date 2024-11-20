package dto

type RequestBodyProduct struct {
	Name        string `json:"name"`
	Description string `gorm:"column:description"`
	CategoryID  []int  `json:"category_id"`
}

type ResponseGetProduct struct {
	Name        string `gorm:"column:name;not null"`
	Slug        string `gorm:"column:slug;unique;not null"`
	IsActive    int    `gorm:"column:is_active;default:1"`
	Description string `gorm:"column:description"`
}

type RequestUpdateStatusProduct struct {
	Status int `json:"status"`
}
