package domain

type CategoryType string

const (
	IncomeCategory  CategoryType = "income"
	ExpenseCategory CategoryType = "expense"
)

type Category struct {
	ID   string       `bson:"_id,omitempty"`
	Name string       `bson:"name"`
	Type CategoryType `bson:"type"` // income | expense
}
