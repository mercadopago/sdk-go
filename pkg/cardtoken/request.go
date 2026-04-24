package cardtoken

// Request represents the body payload for creating a card token via the MercadoPago
// Card Tokens API. It contains the raw card details needed to generate a single-use,
// short-lived token for secure payment processing.
type Request struct {
	// Cardholder contains the name and identification of the card owner.
	Cardholder *CardholderRequest `json:"cardholder,omitempty"`

	// SiteID is the MercadoPago site identifier (e.g., "MLB", "MLA", "MLM")
	// indicating the country where the token is being created.
	SiteID string `json:"site_id,omitempty"`
	// CardNumber is the full card number (PAN). This value is never stored by
	// MercadoPago; it is used only to generate the token.
	CardNumber string `json:"card_number,omitempty"`
	// ExpirationYear is the card's expiration year (e.g., "2025").
	ExpirationYear string `json:"expiration_year,omitempty"`
	// ExpirationMonth is the card's expiration month (e.g., "12").
	ExpirationMonth string `json:"expiration_month,omitempty"`
	// SecurityCode is the card's security code (CVV/CVC).
	SecurityCode string `json:"security_code,omitempty"`
}

// CardholderRequest represents the cardholder's information within a card token [Request].
type CardholderRequest struct {
	// Identification is the cardholder's identity document (e.g., CPF, DNI).
	Identification *IdentificationRequest `json:"identification,omitempty"`

	// Name is the cardholder's full name as printed on the card.
	Name string `json:"name,omitempty"`
}

// IdentificationRequest represents an identity document (e.g., CPF, DNI) associated
// with the cardholder in a card token [Request].
type IdentificationRequest struct {
	// Number is the identification document number.
	Number string `json:"number,omitempty"`
	// Type is the identification document type (e.g., "CPF", "DNI", "CC").
	Type string `json:"type,omitempty"`
}
