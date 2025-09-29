package dto

// CategoryCreateRequest representa los datos necesarios para crear una categoría
type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

// CategoryResponse representa la categoría devuelta al cliente
type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
