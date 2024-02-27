package customer

import "time"

// Request represents a request for creating a customer.
type Request struct {
	DateRegistered *time.Time             `json:"date_registered,omitempty"`
	Address        *AddressResponse       `json:"address,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`

	DefaultAddress string `json:"default_address,omitempty"`
	DefaultCard    string `json:"default_card,omitempty"`
	Description    string `json:"description,omitempty"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
}

// AddressRequest represents a request for an address.
type AddressRequest struct {
	City *CityRequest `json:"city,omitempty"`

	ID           string `json:"id,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber int    `json:"street_number,omitempty"`
}

// CityRequest represents a request for a city.
type CityRequest struct {
	Name string `json:"name,omitempty"`
}

// IdentificationRequest represents a request for an identification.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// PhoneRequest represents a request for a phone.
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}
