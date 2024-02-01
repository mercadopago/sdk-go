package preference

import "time"

// PreferenceSearch contains information about a preference for searching.
type SearchResponse struct {
	ID                  string        `json:"id"`
	ClientID            string        `json:"client_id"`
	CollectorID         int64         `json:"collector_id"`
	DateCreated         time.Time     `json:"date_created"`
	ExpirationDateFrom  time.Time     `json:"expiration_date_from"`
	ExpirationDateTo    time.Time     `json:"expiration_date_to"`
	Expires             bool          `json:"expires"`
	ExternalReference   string        `json:"external_reference"`
	Items               []string      `json:"items"`
	LastUpdated         time.Time     `json:"last_updated"`
	LiveMode            bool          `json:"live_mode"`
	Marketplace         string        `json:"marketplace"`
	OperationType       string        `json:"operation_type"`
	PayerEmail          string        `json:"payer_email"`
	PayerID             string        `json:"payer_id"`
	PlatformID          string        `json:"platform_id"`
	ProcessingModes     []string      `json:"processing_modes"`
	ProductID           string        `json:"product_id"`
	Purpose             string        `json:"purpose"`
	SiteID              string        `json:"site_id"`
	SponsorID           int64         `json:"sponsor_id"`
	ShippingMode        string        `json:"shipping_mode"`
}

// MPElementsResourcesPage is a search page that contains elements.
type SearchResponsePage struct {
	Total      int         `json:"total"`
	NextOffset int         `json:"next_offset"`
	Elements   []SearchResponse `json:"elements"`
}