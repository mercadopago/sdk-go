package customer

import "time"

// Response represents the full customer profile returned by the MercadoPago Customers API.
// It contains personal information, saved cards, registered addresses, and metadata
// such as creation timestamps and live-mode status.
type Response struct {
	// Phone is the customer's phone contact information.
	Phone PhoneResponse `json:"phone"`
	// Identification is the customer's identity document.
	Identification IdentificationResponse `json:"identification"`
	// Address is the customer's primary address.
	Address AddressResponse `json:"address"`
	// DateRegistered is the date when the customer was registered in the merchant's system.
	DateRegistered time.Time `json:"date_registered"`
	// DateCreated is the date when the customer record was created in MercadoPago.
	DateCreated time.Time `json:"date_created"`
	// DateLastUpdated is the date when the customer record was last modified.
	DateLastUpdated time.Time `json:"date_last_updated"`
	// Cards is the list of saved payment cards associated with the customer.
	Cards []CardResponse `json:"cards"`
	// Addresses is the list of all addresses registered for the customer.
	Addresses []CompleteAddressResponse `json:"addresses"`

	// ID is the unique MercadoPago identifier for the customer.
	ID string `json:"id"`
	// Email is the customer's email address.
	Email string `json:"email"`
	// FirstName is the customer's first name.
	FirstName string `json:"first_name"`
	// LastName is the customer's last name.
	LastName string `json:"last_name"`
	// Description is a free-text reference or label for the customer.
	Description string `json:"description"`
	// DefaultCard is the ID of the customer's default saved card.
	DefaultCard string `json:"default_card"`
	// DefaultAddress is the ID of the customer's default address.
	DefaultAddress string `json:"default_address"`
	// Status is the current status of the customer record (e.g., "active").
	Status string `json:"status"`
	// UserID is the MercadoPago user ID linked to the customer.
	UserID int64 `json:"user_id"`
	// MerchantID is the merchant ID that owns this customer record.
	MerchantID int `json:"merchant_id"`
	// ClientID is the application client ID that created the customer.
	ClientID int `json:"client_id"`
	// LiveMode indicates whether the customer was created in production (true) or sandbox (false).
	LiveMode bool `json:"live_mode"`
}

// PhoneResponse represents the customer's phone contact information returned by the API.
type PhoneResponse struct {
	// AreaCode is the phone area code.
	AreaCode string `json:"area_code"`
	// Number is the phone number without the area code.
	Number string `json:"number"`
}

// AddressResponse represents the customer's primary address returned by the API.
type AddressResponse struct {
	// ID is the address identifier.
	ID string `json:"id"`
	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code"`
	// StreetName is the street name.
	StreetName string `json:"street_name"`
	// StreetNumber is the street number.
	StreetNumber int `json:"street_number"`
}

// CardResponse represents a saved payment card associated with a customer, including
// masked card details, cardholder information, and the payment method.
type CardResponse struct {
	// Cardholder contains the cardholder's name and identification.
	Cardholder CardholderResponse `json:"cardholder"`
	// Issuer contains information about the card issuer (e.g., bank).
	Issuer IssuerResponse `json:"issuer"`
	// PaymentMethod describes the payment method associated with this card (e.g., Visa, Mastercard).
	PaymentMethod PaymentMethodResponse `json:"payment_method"`
	// SecurityCode contains metadata about the card's security code.
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
	// FirstSixDigits contains the first six digits (BIN) of the card number.
	FirstSixDigits string `json:"first_six_digits"`
	// LastFourDigits contains the last four digits of the card number.
	LastFourDigits string `json:"last_four_digits"`
	// ExpirationMonth is the card's expiration month (1-12).
	ExpirationMonth int `json:"expiration_month"`
	// ExpirationYear is the card's expiration year (four digits).
	ExpirationYear int `json:"expiration_year"`
}

// CardholderResponse represents the cardholder's name and identification document
// as returned by the API.
type CardholderResponse struct {
	// Identification is the cardholder's identity document.
	Identification IdentificationResponse `json:"identification"`

	// Name is the cardholder's full name as printed on the card.
	Name string `json:"name"`
}

// IdentificationResponse represents an identity document (e.g., CPF, DNI) returned
// by the API, typically associated with a customer or cardholder.
type IdentificationResponse struct {
	// Type is the identification document type (e.g., "CPF", "DNI").
	Type string `json:"type"`
	// Number is the identification document number.
	Number string `json:"number"`
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

// CompleteAddressResponse represents a fully detailed address registered for a customer,
// including city, state, country, and neighborhood information.
type CompleteAddressResponse struct {
	// City contains the city information.
	City CityResponse `json:"city"`
	// State contains the state or province information.
	State StateResponse `json:"state"`
	// Country contains the country information.
	Country CountryResponse `json:"country"`
	// Neighborhood contains the neighborhood information.
	Neighborhood NeighborhoodResponse `json:"neighborhood"`
	// DateCreated is the date when this address was registered.
	DateCreated time.Time `json:"date_created"`

	// ID is the address identifier.
	ID string `json:"id"`
	// StreetName is the street name.
	StreetName string `json:"street_name"`
	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code"`
}

// CityResponse represents city information within a [CompleteAddressResponse].
type CityResponse struct {
	// ID is the city identifier.
	ID string `json:"id"`
	// Name is the city name.
	Name string `json:"name"`
}

// StateResponse represents state or province information within a [CompleteAddressResponse].
type StateResponse struct {
	// ID is the state identifier.
	ID string `json:"id"`
	// Name is the state name.
	Name string `json:"name"`
}

// CountryResponse represents country information within a [CompleteAddressResponse].
type CountryResponse struct {
	// ID is the country identifier (e.g., "AR", "BR", "MX").
	ID string `json:"id"`
	// Name is the country name.
	Name string `json:"name"`
}

// NeighborhoodResponse represents neighborhood information within a [CompleteAddressResponse].
type NeighborhoodResponse struct {
	// Name is the neighborhood name.
	Name string `json:"name"`
}
