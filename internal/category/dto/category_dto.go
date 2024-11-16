package dto

type ResponseGetCategory struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentID *uint  `json:"parent_id"`
	IsActive int    `json:"is_active"`
}

type RequestBodyCategory struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}

type RequestUpdateCategory struct {
	Name     string `json:"name"`
	Slug     string `json:"slugs"`
	IsActive int    `json:"is_active"`
	ParentID *uint  `json:"parent_id"`
}
