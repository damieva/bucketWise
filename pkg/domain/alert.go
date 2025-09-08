package domain

import "time"

type Alert struct {
	ID           string    `json:"id"`
	Message      string    `json:"message"`
	Status       string    `json:"status"`
	BudgetID     string    `json:"budgetID"`
	CreationTime time.Time `json:"creationTime"`
}
