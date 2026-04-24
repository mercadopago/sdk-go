package cardtoken

import "time"

// Response represents the card token returned by the MercadoPago Card Tokens API.
// It contains masked card details, cardholder information, validation metadata,
// and the token's expiration timestamp (DateDue).
type Response struct {
	// Cardholder contains the cardholder's name and identification document.
	Cardholder CardholderResponse `json:"cardholder"`
	// DateCreated is the date when the card token was created.
	DateCreated time.Time `json:"date_created"`
	// DateLastUpdated is the date when the card token was last modified.
	DateLastUpdated time.Time `json:"date_last_updated"`
	// DateDue is the expiration date of the card token, after which it can no longer be used.
	DateDue time.Time `json:"date_due"`

	// ID is the unique identifier for the card token, used to reference it in payment or
	// card-saving operations.
	ID string `json:"id"`
	// FirstSixDigits contains the first six digits (BIN) of the card number.
	FirstSixDigits string `json:"first_six_digits"`
	// LastFourDigits contains the last four digits of the card number.
	LastFourDigits string `json:"last_four_digits"`
	// Status is the current status of the card token (e.g., "active").
	Status string `json:"status"`
	// ExpirationMonth is the card's expiration month (1-12).
	ExpirationMonth int `json:"expiration_month"`
	// ExpirationYear is the card's expiration year (four digits).
	ExpirationYear int `json:"expiration_year"`
	// CardNumberLength is the total number of digits in the card number.
	CardNumberLength int `json:"card_number_length"`
	// SecurityCodeLength is the number of digits in the card's security code (CVV/CVC).
	SecurityCodeLength int `json:"security_code_length"`
	// LuhnValidation indicates whether the card number passed Luhn algorithm validation.
	LuhnValidation bool `json:"luhn_validation"`
	// LiveMode indicates whether the token was created in production (true) or sandbox (false).
	LiveMode bool `json:"live_mode"`
	// RequireEsc indicates whether the card requires E-commerce Security Code validation.
	RequireEsc bool `json:"require_esc"`
}

// CardholderResponse represents the cardholder's information within a card token [Response].
type CardholderResponse struct {
	// Identification is the cardholder's identity document.
	Identification IdentificationResponse `json:"identification"`

	// Name is the cardholder's full name as printed on the card.
	Name string `json:"name"`
}

// IdentificationResponse represents an identity document (e.g., CPF, DNI) associated
// with the cardholder in a card token [Response].
type IdentificationResponse struct {
	// Number is the identification document number.
	Number string `json:"number"`
	// Type is the identification document type (e.g., "CPF", "DNI", "CC").
	Type string `json:"type"`
}
