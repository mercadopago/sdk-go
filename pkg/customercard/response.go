package customercard

import "time"

// Response represents a customer card.
type Response struct {
	Issuer          IssuerResponse         `json:"issuer"`
	Cardholder      CardholderResponse     `json:"cardholder"`
	AdditionalInfo  AdditionalInfoResponse `json:"additional_info"`
	PaymentMethod   PaymentMethodResponse  `json:"payment_method"`
	SecurityCode    SecurityCodeResponse   `json:"security_code"`
	DateCreated     time.Time              `json:"date_created"`
	DateLastUpdated time.Time              `json:"date_last_updated"`

	ID              string `json:"id"`
	CustomerID      string `json:"customer_id"`
	UserID          string `json:"user_id"`
	CardNumberID    string `json:"card_number_id"`
	FirstSixDigits  string `json:"first_six_digits"`
	LastFourDigits  string `json:"last_four_digits"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
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
	Identification IdentificationResponse `json:"identification"`

	Name string `json:"name"`
}

// IdentificationResponse represents the cardholder's document.
type IdentificationResponse struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

// IssuerResponse represents the card issuer code.
type IssuerResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// PaymentMethodResponse represents the card's payment method.
type PaymentMethodResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeID   string `json:"payment_type_id"`
	Thumbnail       string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}

// SecurityCode represents the card's security code.
type SecurityCodeResponse struct {
	CardLocation string `json:"card_location"`
	Length       int    `json:"length"`
}
