package cardtoken

import "time"

// Response contains every field returned by Card Tokens API on card token creation.
type Response struct {
	Cardholder      CardholderResponse `json:"cardholder"`        // token's cardholder data
	DateCreated     time.Time          `json:"date_created"`      // token's creation time
	DateLastUpdated time.Time          `json:"date_last_updated"` // token's last update time
	DateDue         time.Time          `json:"date_due"`          // token's due date

	ID                 string `json:"id"`                   // generated token and that must be sent on payment creation
	FirstSixDigits     string `json:"first_six_digits"`     // card's first six digits
	LastFourDigits     string `json:"last_four_digits"`     // card's last four digits
	Status             string `json:"status"`               // says if the generated token is active or not
	ExpirationMonth    int    `json:"expiration_month"`     // card's  expiration month
	ExpirationYear     int    `json:"expiration_year"`      // card's  expiration year
	CardNumberLength   int    `json:"card_number_length"`   // card's  number length
	SecurityCodeLength int    `json:"security_code_length"` // security code length
	LuhnValidation     bool   `json:"luhn_validation"`      // it is true for valid card number and false for invalid
	LiveMode           bool   `json:"live_mode"`
	RequireEsc         bool   `json:"require_esc"`
}

// CardholderResponse contains cardholder data used at generate card token.
type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification,omitempty"` // cardholder's identification data

	Name string `json:"name"` // cardholder's name
}

// IdentificationResponse contains cardholder identification data used at generate card token.
type IdentificationResponse struct {
	Number string `json:"number"` // idenfication number, its format can change depending on the country
	Type   string `json:"type"`   // idenfication type, its value can change depending on the country
}
