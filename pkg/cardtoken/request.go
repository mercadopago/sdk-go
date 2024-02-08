package cardtoken

type Request struct {
	SiteID          string      `json:"site_id"`
	CardNumber      string      `json:"card_number"`
	ExpirationYear  string      `json:"expiration_year"`
	ExpirationMonth string      `json:"expiration_month"`
	SecurityCode    string      `json:"security_code"`
	Cardholder      *Cardholder `json:"cardholder,omitempty"`
}

type Cardholder struct {
	Identification *Identification `json:"identification,omitempty"`
	Name           string          `json:"name"`
}

type Identification struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}
