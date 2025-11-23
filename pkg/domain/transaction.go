package domain

import (
	"time"
)

type Transaction struct {
	ID           ID           `bson:"_id,omitempty"`
	Amount       float64      `bson:"amount"`
	Date         time.Time    `bson:"date"`
	Description  string       `bson:"description"`
	CategoryID   string       `bson:"category_id"`
	CategoryName string       `bson:"category_name"`
	Type         CategoryType `bson:"type"`
}
