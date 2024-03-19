package cardtoken

import "time"

// Response contains every field returned by Card Tokens API on card token creation.
type Response struct {
	Cardholder      CardholderResponse `json:"cardholder"`
	DateCreated     time.Time          `json:"date_created"`
	DateLastUpdated time.Time          `json:"date_last_updated"`
	DateDue         time.Time          `json:"date_due"`

	ID                 string `json:"id"`
	FirstSixDigits     string `json:"first_six_digits"`
	LastFourDigits     string `json:"last_four_digits"`
	Status             string `json:"status"`
	ExpirationMonth    int    `json:"expiration_month"`
	ExpirationYear     int    `json:"expiration_year"`
	CardNumberLength   int    `json:"card_number_length"`
	SecurityCodeLength int    `json:"security_code_length"`
	LuhnValidation     bool   `json:"luhn_validation"`
	LiveMode           bool   `json:"live_mode"`
	RequireEsc         bool   `json:"require_esc"`
}

// CardholderResponse contains cardholder data used at generate card token.
type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification,omitempty"`

	Name string `json:"name"`
}

// IdentificationResponse contains cardholder identification data used at generate card token.
type IdentificationResponse struct {
	Number string `json:"number"` // idenfication number, its format can change depending on the country
	Type   string `json:"type"`   // idenfication type, its value can change depending on the country
}
