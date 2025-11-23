package domain

type CategoryType string

const (
	IncomeCategory  CategoryType = "income"
	ExpenseCategory CategoryType = "expense"
)

type Category struct {
	ID   ID           `bson:"_id,omitempty"`
	Name string       `bson:"name"`
	Type CategoryType `bson:"type"` // income | expense
}
