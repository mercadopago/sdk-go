package customer

import "time"

// Response represents a customer.
type Response struct {
	Phone           PhoneResponse             `json:"phone"`             // customer's phone
	Identification  IdentificationResponse    `json:"identification"`    // customer's identification
	Address         AddressResponse           `json:"address"`           // customer's address
	DateRegistered  time.Time                 `json:"date_registered"`   // customer's date registered
	DateCreated     time.Time                 `json:"date_created"`      // customer's date created
	DateLastUpdated time.Time                 `json:"date_last_updated"` // customer's date last updated
	Cards           []CardResponse            `json:"cards"`             // customer's cards
	Addresses       []CompleteAddressResponse `json:"addresses"`         // customer's addresses

	ID             string `json:"id"`          // customer's identification
	Email          string `json:"email"`       // customer's email
	FirstName      string `json:"first_name"`  // customer's first name
	LastName       string `json:"last_name"`   // customer's last name
	Description    string `json:"description"` // customer's description
	DefaultCard    string `json:"default_card"`
	DefaultAddress string `json:"default_address"`
	Status         string `json:"status"`
	UserID         int    `json:"user_id"`
	MerchantID     int    `json:"merchant_id"`
	ClientID       int    `json:"client_id"`
	LiveMode       bool   `json:"live_mode"`
}

// PhoneResponse represents a response for a phone.
type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"` // number
}

// AddressResponse represents a response for an address.
type AddressResponse struct {
	ID           string `json:"id"`            // address's identification
	ZipCode      string `json:"zip_code"`      // zip code
	StreetName   string `json:"street_name"`   // street name
	StreetNumber int    `json:"street_number"` // street number
}

// CardResponse represents a response for a card.
type CardResponse struct {
	Cardholder      CardholderResponse    `json:"cardholder"` // card's cardholder data
	Issuer          IssuerResponse        `json:"issuer"`
	PaymentMethod   PaymentMethodResponse `json:"payment_method"`    // card's payment method data
	SecurityCode    SecurityCodeResponse  `json:"security_code"`     // card's security code data
	DateCreated     time.Time             `json:"date_created"`      // card's date created
	DateLastUpdated time.Time             `json:"date_last_updated"` // card's date last updated

	ID              string `json:"id"`          // card's identification
	CustomerID      string `json:"customer_id"` // card's customer identification
	UserID          string `json:"user_id"`
	FirstSixDigits  string `json:"first_six_digits"` // card's first six digits
	LastFourDigits  string `json:"last_four_digits"` // card's last four digits
	ExpirationMonth int    `json:"expiration_month"` // card's expiration month
	ExpirationYear  int    `json:"expiration_year"`  // card's expiration year
}

// CardholderResponse represents a response for a cardholder.
type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification"` // cardholder's identification data

	Name string `json:"name"` // cardholder's name
}

// IdentificationResponse represents a response for an identification.
type IdentificationResponse struct {
	Type   string `json:"type"`   // type (can change depending on the country)
	Number string `json:"number"` // number (its format can change depending on the country)
}

// IssuerResponse represents a response for an issuer.
type IssuerResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// PaymentMethodResponse represents a response for a payment method.
type PaymentMethodResponse struct {
	ID              string `json:"id"`              // payment method identification
	Name            string `json:"name"`            // payment method name
	PaymentTypeID   string `json:"payment_type_id"` // payment method type
	Thumbnail       string `json:"thumbnail"`       // payment method thumbnail
	SecureThumbnail string `json:"secure_thumbnail"`
}

// SecurityCodeResponse represents a response for a security code.
type SecurityCodeResponse struct {
	CardLocation string `json:"card_location"`
	Length       int    `json:"length"`
}

// CompleteAddressResponse represents a response for a complete address.
type CompleteAddressResponse struct {
	City         CityResponse         `json:"city"`         // address's city data
	State        StateResponse        `json:"state"`        // address's state data
	Country      CountryResponse      `json:"country"`      // address's country data
	Neighborhood NeighborhoodResponse `json:"neighborhood"` // address's neighborhood data
	DateCreated  time.Time            `json:"date_created"` // address's date created data

	ID         string `json:"id"`          // address's identification
	StreetName string `json:"street_name"` // address's street name
	ZipCode    string `json:"zip_code"`    // address's zip code
}

// CityResponse represents a response for a city.
type CityResponse struct {
	ID   string `json:"id"`   // city's identification
	Name string `json:"name"` // city's name
}

// StateResponse represents a response for a state.
type StateResponse struct {
	ID   string `json:"id"`   // state's identification
	Name string `json:"name"` // state's name
}

// CountryResponse represents a response for a country.
type CountryResponse struct {
	ID   string `json:"id"`   // country's identification
	Name string `json:"name"` // country's name
}

// NeighborhoodResponse represents a response for a neighborhood.
type NeighborhoodResponse struct {
	Name string `json:"name"` // neighborhood's name
}
