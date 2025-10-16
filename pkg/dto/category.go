package dto

// CategoryCreateRequest represents a category for creation
// @Description Category data required to create a new category
type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required,one of=income expense"` // income | expense
}

// CategoryResponse represents a category returned in responses
// @Description Category data returned by the API
type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
