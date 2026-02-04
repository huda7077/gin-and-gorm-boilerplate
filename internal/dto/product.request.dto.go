package dto

// Create Product
type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required,min=3,max=255"`
	Description string  `json:"description" binding:"max=1000"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
}

// Update Product
type ProductUpdateRequest struct {
	Name        string  `json:"name" binding:"omitempty,min=3,max=255"`
	Description string  `json:"description" binding:"omitempty,max=1000"`
	Price       float64 `json:"price" binding:"omitempty,gt=0"`
	Stock       int     `json:"stock" binding:"omitempty,gte=0"`
}
