package merchantorder

// Request represents merchant order request.
type Request struct {
	PreferenceID      string `json:"preference_id"`
	ApplicationID     string `json:"application_id"`
	SiteID            string `json:"site_id"`
	SponsorID         string `json:"sponsor_id"`
	NotificationURL   string `json:"notification_url"`
	AdditionalInfo    string `json:"additional_info"`
	ExternalReference string `json:"external_reference"`
	Marketplace       string `json:"marketplace"`
	Version           int64  `json:"version"`

	Collector *CollectorRequest `json:"collector"`
	Payer     *PayerRequest     `json:"payer"`
	Items     []ItemRequest     `json:"items"`
}

// PayerRequest represents payer information.
type PayerRequest struct {
	Nickname string `json:"nickname"`
	ID       int64  `json:"id"`
}

// ItemRequest represents item information.
type ItemRequest struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PictureURL  string  `json:"picture_url"`
	CategoryID  string  `json:"category_id"`
	CurrencyID  string  `json:"currency_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

// CollectorRequest represents seller information.
type CollectorRequest struct {
	ID int64 `json:"id"`
}
