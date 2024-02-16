package customer

// UpdateRequest represents a request for updating a customer.
type UpdateRequest struct {
	DateRegistered string `json:"date_registered,omitempty"`
	DefaultAddress string `json:"default_address,omitempty"`
	DefaultCard    string `json:"default_card,omitempty"`
	Description    string `json:"description,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`

	Address        AddressRequest        `json:"address,omitempty"`
	Identification IdentificationRequest `json:"identification,omitempty"`
	Phone          PhoneRequest          `json:"phone,omitempty"`
}
