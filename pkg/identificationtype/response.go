package identificationtype

// Response represents an identification document type returned by the MercadoPago
// Identification Types API. Each entry describes a kind of identity document accepted
// in the country associated with the authenticated credentials (e.g., CPF in Brazil,
// DNI in Argentina).
type Response struct {
	// ID is the unique identifier for the document type (e.g., "CPF", "DNI", "CC").
	ID string `json:"id"`
	// Name is the human-readable name of the document type (e.g., "CPF", "DNI").
	Name string `json:"name"`
	// Type is the category of the identification (e.g., "number").
	Type string `json:"type"`
	// MinLength is the minimum allowed number of characters for this document type.
	MinLength int `json:"min_length"`
	// MaxLength is the maximum allowed number of characters for this document type.
	MaxLength int `json:"max_length"`
}
