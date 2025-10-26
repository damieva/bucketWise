package dto

import "time"

// TransactionCreateRequest represents a transaction for creation
// @Description Transaction data required to create a new transaction
type TransactionCreateRequest struct {
	Amount      float64   `json:"amount" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

// TransactionResponse represents a transaction returned in responses
// @Description Transaction data returned by the API
type TransactionResponse struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	CategoryID   string    `json:"categoryID"`
	CategoryName string    `json:"categoryName"`
	Type         string    `json:"type"` // income | expense
}
