package preference

import "time"

// SearchResponse contains information about a preference for searching.
type SearchResponse struct {
	DateCreated        *time.Time `json:"date_created"`
	ExpirationDateFrom *time.Time `json:"expiration_date_from"`
	ExpirationDateTo   *time.Time `json:"expiration_date_to"`
	LastUpdated        *time.Time `json:"last_updated"`

	ID                string   `json:"id"`
	ClientID          string   `json:"client_id"`
	Marketplace       string   `json:"marketplace"`
	OperationType     string   `json:"operation_type"`
	PayerEmail        string   `json:"payer_email"`
	PayerID           string   `json:"payer_id"`
	PlatformID        string   `json:"platform_id"`
	ShippingMode      string   `json:"shipping_mode"`
	ExternalReference string   `json:"external_reference"`
	ProductID         string   `json:"product_id"`
	Purpose           string   `json:"purpose"`
	SiteID            string   `json:"site_id"`
	CollectorID       int64    `json:"collector_id"`
	SponsorID         int64    `json:"sponsor_id"`
	LiveMode          bool     `json:"live_mode"`
	Expires           bool     `json:"expires"`
	Items             []string `json:"items"`
	ProcessingModes   []string `json:"processing_modes"`
}

// SearchPageResponse is a search page that contains elements.
type SearchResponsePage struct {
	Elements []SearchResponse `json:"elements"`

	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`
}
