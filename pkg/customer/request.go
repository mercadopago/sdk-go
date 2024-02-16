package customer

import "time"

// Request represents a request for creating a customer.
type Request struct {
	DefaultAddress string `json:"default_address,omitempty"`
	DefaultCard    string `json:"default_card,omitempty"`
	Description    string `json:"description,omitempty"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`

	DateRegistered *time.Time            `json:"date_registered,omitempty"`
	Address        AddressResponse       `json:"address,omitempty"`
	Identification IdentificationRequest `json:"identification,omitempty"`
	Phone          PhoneRequest          `json:"phone,omitempty"`
}

type AddressRequest struct {
	ID           string `json:"id,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber int    `json:"street_number,omitempty"`

	City CityRequest `json:"city,omitempty"`
}

type CityRequest struct {
	Name string `json:"name,omitempty"`
}

type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}
