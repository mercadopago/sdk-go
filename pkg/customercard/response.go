package customercard

import "time"

// Response represents a customer card.
type Response struct {
	Issuer          IssuerResponse         `json:"issuer"`            // card's issuer data
	Cardholder      CardholderResponse     `json:"cardholder"`        // card's cardholder data
	AdditionalInfo  AdditionalInfoResponse `json:"additional_info"`   // card's additional info data
	PaymentMethod   PaymentMethodResponse  `json:"payment_method"`    // card's payment method data
	SecurityCode    SecurityCodeResponse   `json:"security_code"`     // card's security code data
	DateCreated     time.Time              `json:"date_created"`      // card's date created
	DateLastUpdated time.Time              `json:"date_last_updated"` // card's date last updated

	ID              string `json:"id"`          // card's id
	CustomerID      string `json:"customer_id"` // customer's id
	UserID          string `json:"user_id"`
	CardNumberID    string `json:"card_number_id"`
	FirstSixDigits  string `json:"first_six_digits"` // card's first six digits
	LastFourDigits  string `json:"last_four_digits"` // card's last four digits
	ExpirationMonth int    `json:"expiration_month"` // card's expiration month
	ExpirationYear  int    `json:"expiration_year"`  // card's expiration year
	LiveMode        bool   `json:"live_mode"`
}

// AdditionalInfoResponse represents additional customer card information.
type AdditionalInfoResponse struct {
	RequestPublic        string `json:"request_public"`
	ApiClientApplication string `json:"api_client_application"`
	ApiClientScope       string `json:"api_client_scope"`
}

// CardholderResponse represents information about the cardholder.
type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification"` // cardholder`s identification

	Name string `json:"name"` // cardholder's name
}

// IdentificationResponse represents the cardholder's document.
type IdentificationResponse struct {
	Type   string `json:"type"`   // type (can change depending on the country)
	Number string `json:"number"` // number (its format can change depending on the country)
}

// IssuerResponse represents the card issuer code.
type IssuerResponse struct {
	Name string `json:"name"` // issuer's name
	ID   int    `json:"id"`   // issuer's id
}

// PaymentMethodResponse represents the card's payment method.
type PaymentMethodResponse struct {
	ID              string `json:"id"`              // payment method's id
	Name            string `json:"name"`            // payment method's name
	PaymentTypeID   string `json:"payment_type_id"` // payment method's type
	Thumbnail       string `json:"thumbnail"`       // payment method's thumbnail
	SecureThumbnail string `json:"secure_thumbnail"`
}

// SecurityCode represents the card's security code.
type SecurityCodeResponse struct {
	CardLocation string `json:"card_location"`
	Length       int    `json:"length"`
}
