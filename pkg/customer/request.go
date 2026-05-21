package customer

import "time"

// Request represents the body payload for creating or updating a customer via the
// MercadoPago Customers API. All fields are optional (omitempty) so that partial
// updates can be performed without overwriting existing data.
type Request struct {
	// Address is the customer's primary address information.
	Address *AddressRequest `json:"address,omitempty"`
	// Identification is the customer's identity document (e.g., CPF, DNI).
	Identification *IdentificationRequest `json:"identification,omitempty"`
	// Phone is the customer's phone contact information.
	Phone *PhoneRequest `json:"phone,omitempty"`
	// DateRegistered is the date when the customer was registered in the merchant's system.
	DateRegistered *time.Time `json:"date_registered,omitempty"`

	// DefaultAddress is the ID of the customer's default address.
	DefaultAddress string `json:"default_address,omitempty"`
	// DefaultCard is the ID of the customer's default saved card.
	DefaultCard string `json:"default_card,omitempty"`
	// Description is a free-text reference or label for the customer.
	Description string `json:"description,omitempty"`
	// Email is the customer's email address, which must be unique per merchant.
	Email string `json:"email,omitempty"`
	// FirstName is the customer's first name.
	FirstName string `json:"first_name,omitempty"`
	// LastName is the customer's last name.
	LastName string `json:"last_name,omitempty"`
}

// AddressRequest represents the address information included in a customer [Request].
// It maps to the address object in the MercadoPago Customers API.
type AddressRequest struct {
	// City contains the city details for this address.
	City *CityRequest `json:"city,omitempty"`

	// ID is the address identifier.
	ID string `json:"id,omitempty"`
	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code,omitempty"`
	// StreetName is the street name.
	StreetName string `json:"street_name,omitempty"`
	// StreetNumber is the street number.
	StreetNumber int `json:"street_number,omitempty"`
}

// CityRequest represents city information within an [AddressRequest].
type CityRequest struct {
	// Name is the city name.
	Name string `json:"name,omitempty"`
}

// IdentificationRequest represents the customer's identity document (e.g., CPF, DNI, CC)
// within a customer [Request].
type IdentificationRequest struct {
	// Type is the identification document type (e.g., "CPF", "DNI", "CC").
	Type string `json:"type,omitempty"`
	// Number is the identification document number.
	Number string `json:"number,omitempty"`
}

// PhoneRequest represents the customer's phone contact information within a customer [Request].
type PhoneRequest struct {
	// AreaCode is the phone area code (e.g., "55", "11").
	AreaCode string `json:"area_code,omitempty"`
	// Number is the phone number without the area code.
	Number string `json:"number,omitempty"`
}
