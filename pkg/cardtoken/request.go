package cardtoken

// Request contains every field accepted by Card Tokens API.
type Request struct {
	Cardholder *CardholderRequest `json:"cardholder,omitempty"`

	SiteID          string `json:"site_id,omitempty"`
	CardNumber      string `json:"card_number,omitempty"`
	ExpirationYear  string `json:"expiration_year,omitempty"`
	ExpirationMonth string `json:"expiration_month,omitempty"`
	SecurityCode    string `json:"security_code,omitempty"`
}

type CardholderRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"`

	Name string `json:"name,omitempty"`
}

type IdentificationRequest struct {
	Number string `json:"number,omitempty"`
	Type   string `json:"type,omitempty"`
}
