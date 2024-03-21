package identificationtype

type Response struct {
	ID        string `json:"id"`         // identification type's id
	Name      string `json:"name"`       // identification type's name
	Type      string `json:"type"`       // identification type's type
	MinLength int    `json:"min_length"` // identification type's min_length
	MaxLength int    `json:"max_length"` // identification type's max_length
}
