package cardtoken

type Request struct {
	SiteID          string `json:"site_id,omitempty"`
	CardNumber      string `json:"card_number,omitempty"`
	ExpirationYear  string `json:"expiration_year,omitempty"`
	ExpirationMonth string `json:"expiration_month,omitempty"`
	SecurityCode    string `json:"security_code,omitempty"`

	Cardholder Cardholder `json:"cardholder,omitempty"`
}

type Cardholder struct {
	Name string `json:"name,omitempty"`

	Identification *Identification `json:"identification,omitempty"`
}

type Identification struct {
	Number string `json:"number,omitempty"`
	Type   string `json:"type,omitempty"`
}
