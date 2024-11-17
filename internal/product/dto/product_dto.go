package dto

type RequestBodyProduct struct {
	Name        string `json:"name"`
	Description string `gorm:"column:description"`
	CategoryID  []int  `json:"category_id"`
}
