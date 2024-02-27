package cardtoken

import "time"

type Response struct {
	DateCreated     *time.Time         `json:"date_created"`
	DateLastUpdated *time.Time         `json:"date_last_updated"`
	DateDue         *time.Time         `json:"date_due"`
	Cardholder      CardholderResponse `json:"cardholder"`

	ID                 string `json:"id"`
	FirstSixDigits     string `json:"first_six_digits"`
	LastFourDigits     string `json:"last_four_digits"`
	Status             string `json:"status"`
	LuhnValidation     bool   `json:"luhn_validation"`
	LiveMode           bool   `json:"live_mode"`
	RequireEsc         bool   `json:"require_esc"`
	ExpirationMonth    int    `json:"expiration_month"`
	ExpirationYear     int    `json:"expiration_year"`
	CardNumberLength   int    `json:"card_number_length"`
	SecurityCodeLength int    `json:"security_code_length"`
}

type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification,omitempty"`

	Name string `json:"name"`
}

type IdentificationResponse struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}
