package merchantorder

// Request represents merchant order request.
type Request struct {
	PreferenceID      string `json:"preference_id,omitempty"`
	ApplicationID     string `json:"application_id,omitempty"`
	SiteID            string `json:"site_id,omitempty"`
	NotificationURL   string `json:"notification_url,omitempty"`
	AdditionalInfo    string `json:"additional_info,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Marketplace       string `json:"marketplace,omitempty"`
	Version           int64  `json:"version,omitempty"`
	SponsorID         int64  `json:"sponsor_id,omitempty"`

	Collector *CollectorRequest `json:"collector,omitempty"`
	Payer     *PayerRequest     `json:"payer,omitempty"`
	Items     []ItemRequest     `json:"items,omitempty"`
}

// PayerRequest represents payer information.
type PayerRequest struct {
	Nickname string `json:"nickname,omitempty"`
	ID       int64  `json:"id,omitempty"`
}

// ItemRequest represents item information.
type ItemRequest struct {
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	CurrencyID  string  `json:"currency_id,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
}

// CollectorRequest represents seller information.
type CollectorRequest struct {
	ID int64 `json:"id,omitempty"`
}
