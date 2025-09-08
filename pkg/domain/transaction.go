package domain

import (
	"time"
)

type Transaction struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	Description  string    `json:"description"`
	CategoryID   string    `json:"categoryID"`
	CreationTime time.Time `json:"creationTime"`
}
