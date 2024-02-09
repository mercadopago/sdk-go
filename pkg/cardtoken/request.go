package cardtoken

type Request struct {
	SiteID          string      `json:"site_id,omitempty"`
	CardNumber      string      `json:"card_number,omitempty"`
	ExpirationYear  string      `json:"expiration_year,omitempty"`
	ExpirationMonth string      `json:"expiration_month,omitempty"`
	SecurityCode    string      `json:"security_code,omitempty"`
	Cardholder      *Cardholder `json:"cardholder,omitempty,omitempty"`
}

type Cardholder struct {
	Identification *Identification `json:"identification,omitempty"`
	Name           string          `json:"name,omitempty"`
}

type Identification struct {
	Number string `json:"number,omitempty"`
	Type   string `json:"type,omitempty"`
}
