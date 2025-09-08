package domain

import "time"

type CategoryType string

const (
	FixedExpense    CategoryType = "income"
	VariableExpense CategoryType = "expense"
	Income          CategoryType = "income"
)

type Category struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         CategoryType `json:"type"`
	CreationTime time.Time    `json:"creationTime"`
}
