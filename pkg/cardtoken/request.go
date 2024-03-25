package cardtoken

// Request contains parameters to create a cardtoken.
type Request struct {
	Cardholder *CardholderRequest `json:"cardholder,omitempty"`

	SiteID          string `json:"site_id,omitempty"`
	CardNumber      string `json:"card_number,omitempty"`
	ExpirationYear  string `json:"expiration_year,omitempty"`
	ExpirationMonth string `json:"expiration_month,omitempty"`
	SecurityCode    string `json:"security_code,omitempty"`
}

// CardholderRequest contains cardholder information in the cardtoken.
type CardholderRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"`

	Name string `json:"name,omitempty"`
}

// IdentificationRequest is a base type that represents identifications, such as payer identification.
type IdentificationRequest struct {
	Number string `json:"number,omitempty"`
	Type   string `json:"type,omitempty"`
}
