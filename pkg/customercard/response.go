package customercard

import "time"

// Response represents a saved customer card returned by the MercadoPago Customer Cards API.
// It contains masked card details, cardholder information, issuer data, payment method
// metadata, and timestamps.
type Response struct {
	// Issuer contains information about the card-issuing institution.
	Issuer IssuerResponse `json:"issuer"`
	// Cardholder contains the cardholder's name and identification document.
	Cardholder CardholderResponse `json:"cardholder"`
	// AdditionalInfo contains supplementary metadata about the card, such as the API client
	// application and scope that created it.
	AdditionalInfo AdditionalInfoResponse `json:"additional_info"`
	// PaymentMethod describes the payment method associated with this card.
	PaymentMethod PaymentMethodResponse `json:"payment_method"`
	// SecurityCode contains metadata about the card's security code (CVV/CVC).
	SecurityCode SecurityCodeResponse `json:"security_code"`
	// DateCreated is the date when the card was saved.
	DateCreated time.Time `json:"date_created"`
	// DateLastUpdated is the date when the card record was last modified.
	DateLastUpdated time.Time `json:"date_last_updated"`

	// ID is the unique identifier for the saved card.
	ID string `json:"id"`
	// CustomerID is the ID of the customer that owns this card.
	CustomerID string `json:"customer_id"`
	// UserID is the MercadoPago user ID linked to this card.
	UserID string `json:"user_id"`
	// CardNumberID is an internal identifier for the card number.
	CardNumberID string `json:"card_number_id"`
	// FirstSixDigits contains the first six digits (BIN) of the card number.
	FirstSixDigits string `json:"first_six_digits"`
	// LastFourDigits contains the last four digits of the card number.
	LastFourDigits string `json:"last_four_digits"`
	// ExpirationMonth is the card's expiration month (1-12).
	ExpirationMonth int `json:"expiration_month"`
	// ExpirationYear is the card's expiration year (four digits).
	ExpirationYear int `json:"expiration_year"`
	// LiveMode indicates whether the card was created in production (true) or sandbox (false).
	LiveMode bool `json:"live_mode"`
}

// AdditionalInfoResponse represents supplementary metadata about a saved customer card,
// such as the originating API client and access scope.
type AdditionalInfoResponse struct {
	// RequestPublic is the public key used in the request that created this card.
	RequestPublic string `json:"request_public"`
	// APIClientApplication is the name of the API client application that created this card.
	APIClientApplication string `json:"api_client_application"`
	// APIClientScope is the access scope of the API client that created this card.
	APIClientScope string `json:"api_client_scope"`
}

// CardholderResponse represents the cardholder's name and identification document
// as returned by the API.
type CardholderResponse struct {
	// Identification is the cardholder's identity document.
	Identification IdentificationResponse `json:"identification"`

	// Name is the cardholder's full name as printed on the card.
	Name string `json:"name"`
}

// IdentificationResponse represents the cardholder's identity document (e.g., CPF, DNI)
// as returned by the API.
type IdentificationResponse struct {
	// Number is the identification document number.
	Number string `json:"number"`
	// Type is the identification document type (e.g., "CPF", "DNI").
	Type string `json:"type"`
}

// IssuerResponse represents the card-issuing institution (e.g., a bank) returned by the API.
type IssuerResponse struct {
	// Name is the issuer's name.
	Name string `json:"name"`
	// ID is the issuer's numeric identifier in MercadoPago.
	ID int `json:"id"`
}

// PaymentMethodResponse represents the payment method associated with a saved card,
// including display metadata such as thumbnails.
type PaymentMethodResponse struct {
	// ID is the payment method identifier (e.g., "visa", "master").
	ID string `json:"id"`
	// Name is the human-readable payment method name.
	Name string `json:"name"`
	// PaymentTypeID is the payment type category (e.g., "credit_card", "debit_card").
	PaymentTypeID string `json:"payment_type_id"`
	// Thumbnail is the URL of the payment method icon.
	Thumbnail string `json:"thumbnail"`
	// SecureThumbnail is the HTTPS URL of the payment method icon.
	SecureThumbnail string `json:"secure_thumbnail"`
}

// SecurityCodeResponse represents metadata about a card's security code (CVV/CVC)
// as returned by the API.
type SecurityCodeResponse struct {
	// CardLocation indicates where the security code is printed on the card (e.g., "back", "front").
	CardLocation string `json:"card_location"`
	// Length is the number of digits in the security code.
	Length int `json:"length"`
}
