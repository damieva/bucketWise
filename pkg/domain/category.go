package domain

// Category represents a category in the system
// @Description Category entity used for classification
type Category struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
