package customercard

import "time"

// Response represents a customer card.
type Response struct {
	ID              string                 `json:"id"`
	CustomerID      string                 `json:"customer_id"`
	UserID          string                 `json:"user_id"`
	CardNumberID    string                 `json:"card_number_id"`
	FirstSixDigits  string                 `json:"first_six_digits"`
	LastFourDigits  string                 `json:"last_four_digits"`
	ExpirationMonth int                    `json:"expiration_month"`
	ExpirationYear  int                    `json:"expiration_year"`
	LiveMode        bool                   `json:"live_mode"`
	DateCreated     *time.Time             `json:"date_created"`
	DateLastUpdated *time.Time             `json:"date_last_updated"`
	Issuer          IssuerResponse         `json:"issuer"`
	Cardholder      CardholderResponse     `json:"cardholder"`
	AdditionalInfo  AdditionalInfoResponse `json:"additional_info"`
	PaymentMethod   PaymentMethodResponse  `json:"payment_method"`
	SecurityCode    SecurityCode           `json:"security_code"`
}

// AdditionalInfoResponse represents additional customer card information.
type AdditionalInfoResponse struct {
	RequestPublic        string `json:"request_public"`
	ApiClientApplication string `json:"api_client_application"`
	ApiClientScope       string `json:"api_client_scope"`
}

// CardholderResponse represents information about the cardholder.
type CardholderResponse struct {
	Name                   string `json:"name"`
	IdentificationResponse `json:"identification"`
}

// IdentificationResponse represents the cardholder's document.
type IdentificationResponse struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

// IssuerResponse represents the card issuer code
type IssuerResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PaymentMethodResponse represents the card's payment method
type PaymentMethodResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeID   string `json:"payment_type_id"`
	Thumbnail       string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}

// SecurityCode represents the card's security code
type SecurityCode struct {
	Length       int    `json:"length"`
	CardLocation string `json:"card_location"`
}
