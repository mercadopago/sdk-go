package cardtoken

type Request struct {
	SiteId          string `json:"site_id"`
	CardNumber      string `json:"card_number"`
	ExpirationYear  string `json:"expiration_year"`
	ExpirationMonth string `json:"expiration_month"`
	SecurityCode    string `json:"security_code"`
	Cardholder      struct {
		Identification struct {
			Type   string `json:"type"`
			Number string `json:"number"`
		} `json:"identification"`
		Name string `json:"name"`
	} `json:"cardholder"`
}
