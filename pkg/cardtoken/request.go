package cardtoken

// Request contains every field accepted by Card Tokens API.
type Request struct {
	SiteID          string `json:"site_id,omitempty"`          // represents the country: MLA (Argentina), MLB (Brazil), MLC (Chile), MLM (Mexico), MLU (Uruguay), MCO (Colombia), MPE (Peru)
	CardNumber      string `json:"card_number,omitempty"`      // card number used to generate card token
	ExpirationYear  string `json:"expiration_year,omitempty"`  // expiration year used to generate card token
	ExpirationMonth string `json:"expiration_month,omitempty"` // expiration month used to generate card token
	SecurityCode    string `json:"security_code,omitempty"`    // security code used to generate card token

	Cardholder *CardholderRequest `json:"cardholder,omitempty"` // contains cardholder data used to generate card token
}

// CardholderRequest contains cardholder data used to generate card token.
type CardholderRequest struct {
	Name string `json:"name,omitempty"` // cardholder name

	Identification *IdentificationRequest `json:"identification,omitempty"` // cardholder identification
}

// IdentificationRequest contains cardholder identification data used to generate card token.
type IdentificationRequest struct {
	Number string `json:"number,omitempty"` // idenfication number, its format can change depending on the country
	Type   string `json:"type,omitempty"`   // idenfication type, its value can change depending on the country
}
