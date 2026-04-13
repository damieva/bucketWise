package dto

import "time"

// TransactionCreateRequest represents a transaction for creation
// @Description Transaction data required to create a new transaction
type TransactionCreateRequest struct {
	Amount      float64   `json:"amount" example:"49.99" binding:"required"`
	Date        time.Time `json:"date" example:"2025-10-26T19:11:58+01:00" binding:"required"`
	Description string    `json:"description" example:"Netflix subscription" binding:"required"`
}

// TransactionResponse represents a transaction returned in responses
// @Description Transaction data returned by the API
type TransactionResponse struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	CategoryID   string    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Type         string    `json:"type"` // income | expense
}

// TransactionsDeleteRequest represents the request body used to delete multiple transactions.
type TransactionsDeleteRequest struct {
	IDs []string `json:"ids"`
}
