package cardtoken

type Request struct {
	SiteId          string     `json:"site_id"`
	CardNumber      string     `json:"card_number"`
	ExpirationYear  string     `json:"expiration_year"`
	ExpirationMonth string     `json:"expiration_month"`
	SecurityCode    string     `json:"security_code"`
	Cardholder      Cardholder `json:"cardholder"`
}

type Cardholder struct {
	Identification Identification `json:"identification"`
	Name           string         `json:"name"`
}

type Identification struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}
