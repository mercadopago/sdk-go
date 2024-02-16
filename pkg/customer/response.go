package customer

import "time"

// Response represents a customer.
type Response struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Description    string `json:"description"`
	DefaultCard    string `json:"default_card"`
	DefaultAddress string `json:"default_address"`
	Status         string `json:"status"`
	UserID         int    `json:"user_id"`
	MerchantID     int    `json:"merchant_id"`
	ClientID       int64  `json:"client_id"`
	LiveMode       bool   `json:"live_mode"`

	DateRegistered  *time.Time                `json:"date_registered"`
	DateCreated     *time.Time                `json:"date_created"`
	DateLastUpdated *time.Time                `json:"date_last_updated"`
	Phone           PhoneResponse             `json:"phone"`
	Identification  IdentificationResponse    `json:"identification"`
	Address         AddressResponse           `json:"address"`
	Cards           []CardResponse            `json:"cards"`
	Addresses       []CompleteAddressResponse `json:"addresses"`
}

// PhoneResponse represents a response for a phone.
type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

// AddressResponse represents a response for an address.
type AddressResponse struct {
	ID           string `json:"id"`
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber int64  `json:"street_number"`
}

// CardResponse represents a response for a card.
type CardResponse struct {
	CustomerID      string `json:"customer_id"`
	FirstSixDigits  string `json:"first_six_digits"`
	ID              string `json:"id"`
	LastFourDigits  string `json:"last_four_digits"`
	UserId          string `json:"user_id"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`

	DateCreated     *time.Time            `json:"date_created"`
	DateLastUpdated *time.Time            `json:"date_last_updated"`
	Cardholder      CardholderResponse    `json:"cardholder"`
	Issuer          IssuerResponse        `json:"issuer"`
	PaymentMethod   PaymentMethodResponse `json:"payment_method"`
	SecurityCode    SecurityCodeResponse  `json:"security_code"`
}

// CardholderResponse represents a response for a cardholder.
type CardholderResponse struct {
	Name string `json:"name"`

	Identification IdentificationResponse `json:"identification"`
}

// IdentificationResponse represents a response for an identification.
type IdentificationResponse struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// IssuerResponse represents a response for an issuer.
type IssuerResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PaymentMethodResponse represents a response for a payment method.
type PaymentMethodResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeId   string `json:"payment_type_id"`
	Thumbnail       string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}

// SecurityCodeResponse represents a response for a security code.
type SecurityCodeResponse struct {
	Length       int    `json:"length"`
	CardLocation string `json:"card_location"`
}

// CompleteAddressResponse represents a response for a complete address.
type CompleteAddressResponse struct {
	ID         string `json:"id"`
	StreetName string `json:"street_name"`
	ZipCode    string `json:"zip_code"`

	DateCreated  *time.Time           `json:"date_created"`
	City         CityResponse         `json:"city"`
	State        StateResponse        `json:"state"`
	Country      CountryResponse      `json:"country"`
	Neighborhood NeighborhoodResponse `json:"neighborhood"`
}

// CityResponse represents a response for a city.
type CityResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// StateResponse represents a response for a state.
type StateResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CountryResponse represents a response for a country.
type CountryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NeighborhoodResponse represents a response for a neighborhood.
type NeighborhoodResponse struct {
	Name string `json:"name"`
}
