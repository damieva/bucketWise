package domain

// ID es un tipo semántico que representa un identificador único de entidad.
// Se define como string para poder usarse tanto en MongoDB (ObjectID.Hex())
// como en Postgres (ids numéricos convertidos a string).
type ID string

// NewID crea un nuevo ID a partir de un string.
// Puedes extenderlo después para generar UUIDs, por ejemplo.
func NewID(value string) ID {
	return ID(value)
}

// String implementa la interfaz fmt.Stringer.
// Go llama automáticamente a este metodo cuando usamos fmt.Print, fmt.Println, fmt.Sprintf, etc.
// Esto nos permite controlar cómo se imprime el ID y aplicar lógica adicional si queremos.
// Sin este metodo, Go imprimiría el valor crudo del tipo ID.
func (id ID) String() string {
	return string(id)
}
