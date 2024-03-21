package customer

import "time"

// Request represents a request for creating a customer.
type Request struct {
	Address        *AddressRequest        `json:"address,omitempty"`         // customer's address
	Identification *IdentificationRequest `json:"identification,omitempty"`  // customer's identification
	Phone          *PhoneRequest          `json:"phone,omitempty"`           // customer's phone
	DateRegistered *time.Time             `json:"date_registered,omitempty"` // customer's registered date

	DefaultAddress string `json:"default_address,omitempty"`
	DefaultCard    string `json:"default_card,omitempty"`
	Description    string `json:"description,omitempty"` // customer's description
	Email          string `json:"email,omitempty"`       // customer's email
	FirstName      string `json:"first_name,omitempty"`  // customer's first name
	LastName       string `json:"last_name,omitempty"`   // customer's last name
}

// AddressRequest represents a request for an address.
type AddressRequest struct {
	City *CityRequest `json:"city,omitempty"` // address's city

	ID           string `json:"id,omitempty"`            // address's identification
	ZipCode      string `json:"zip_code,omitempty"`      // zip code
	StreetName   string `json:"street_name,omitempty"`   // street name
	StreetNumber int    `json:"street_number,omitempty"` // street number
}

// CityRequest represents a request for a city.
type CityRequest struct {
	Name string `json:"name,omitempty"` // city's name
}

// IdentificationRequest represents a request for an identification.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`   // type (can change depending on the country)
	Number string `json:"number,omitempty"` // number (its format can change depending on the country)
}

// PhoneRequest represents a request for a phone.
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"` // number
}
