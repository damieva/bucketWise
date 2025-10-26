package dto

// CategoryCreateRequest represents a category for creation
// @Description Category data required to create a new category
type CategoryCreateRequest struct {
	Name string `json:"name" example:"fixed costs" binding:"required"`
	Type string `json:"type" example:"expense" binding:"required"`
}

// CategoryResponse represents a category returned in responses
// @Description Category data returned by the API
type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
