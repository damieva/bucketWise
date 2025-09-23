package domain

type CategoryType string

// El tag bson le dice al driver de MongoDB (como mongo-go-driver) cómo mapear los campos de tus structs Go a los documentos BSON de MongoDB y viceversa.
type Category struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
