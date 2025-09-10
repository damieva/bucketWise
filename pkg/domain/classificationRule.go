package domain

type RuleType string

const (
	Merchant    RuleType = "Merchant"
	Description RuleType = "Description"
)

type ClassificationRule struct {
	ID         string   `json:"id"`
	Keyword    string   `json:"keyword"`
	CategoryID string   `json:"categoryID"`
	Rule       RuleType `json:"rule"`
}
