package domain

type BudgetPeriod string

const (
	Weekly   BudgetPeriod = "weekly"
	Monthtly BudgetPeriod = "monthtly"
)

type Budget struct {
	ID             string       `json:"id"`
	AssignedAmount float64      `json:"assignedAmount"`
	SpentAmount    float64      `json:"spentAmount"`
	Period         BudgetPeriod `json:"period"`
	CategoryID     string       `json:"categoryID"`
}
