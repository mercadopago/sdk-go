package identificationtype

// Response returns a list of available document types by country.
type Response struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}
