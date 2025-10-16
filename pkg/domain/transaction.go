package domain

import (
	"time"
)

type Transaction struct {
	ID          string    `bson:"_id,omitempty"`
	Amount      float64   `bson:"amount"`
	Date        time.Time `bson:"date"`
	Description string    `bson:"description"`
	CategoryID  string    `bson:"categoryID"`
}
