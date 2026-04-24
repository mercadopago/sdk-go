package preference

import "time"

// SearchResponse represents a summarized checkout preference returned as part of a
// search result from [Client.Search]. It contains a subset of the full preference
// fields, optimized for listing and filtering purposes.
type SearchResponse struct {
	// DateCreated is the date when the preference was created.
	DateCreated time.Time `json:"date_created"`

	// ExpirationDateFrom is the date from which the preference becomes active.
	ExpirationDateFrom time.Time `json:"expiration_date_from"`

	// ExpirationDateTo is the date until which the preference remains active.
	ExpirationDateTo time.Time `json:"expiration_date_to"`

	// LastUpdated is the date when the preference was last modified.
	LastUpdated time.Time `json:"last_updated"`

	// ID is the unique preference identifier assigned by MercadoPago.
	ID string `json:"id"`

	// ClientID is the application ID that created the preference.
	ClientID string `json:"client_id"`

	// Marketplace identifies the marketplace that originated the preference.
	Marketplace string `json:"marketplace"`

	// OperationType is the type of operation (e.g., "regular_payment").
	OperationType string `json:"operation_type"`

	// PayerEmail is the buyer's email address.
	PayerEmail string `json:"payer_email"`

	// PayerID is the buyer's MercadoPago user identifier.
	PayerID string `json:"payer_id"`

	// PlatformID is the platform identifier associated with the preference.
	PlatformID string `json:"platform_id"`

	// ShippingMode is the shipping mode configured for the preference.
	ShippingMode string `json:"shipping_mode"`

	// ExternalReference is the integrator-provided external identifier.
	ExternalReference string `json:"external_reference"`

	// ProductID is the product identifier associated with the preference.
	ProductID string `json:"product_id"`

	// Purpose defines the checkout purpose (e.g., "wallet_purchase").
	Purpose string `json:"purpose"`

	// SiteID is the MercadoPago site identifier (e.g., "MLB", "MLA").
	SiteID string `json:"site_id"`

	// CollectorID is the MercadoPago user ID of the seller who receives payments.
	CollectorID int64 `json:"collector_id"`

	// SponsorID is the sponsor's MercadoPago user ID in marketplace scenarios.
	SponsorID int `json:"sponsor_id"`

	// LiveMode indicates whether the preference operates in production (true) or sandbox (false).
	LiveMode bool `json:"live_mode"`

	// Expires indicates whether the preference has an active expiration window.
	Expires bool `json:"expires"`

	// Items lists the item identifiers included in the preference.
	Items []string `json:"items"`

	// ProcessingModes lists the processing modes configured for the preference.
	ProcessingModes []string `json:"processing_modes"`
}

// PagingResponse represents a paginated list of search results returned by [Client.Search].
// It contains the matched preferences along with pagination metadata.
type PagingResponse struct {
	// Elements is the list of preferences matching the search criteria.
	Elements []SearchResponse `json:"elements"`

	// Total is the total number of preferences matching the search criteria across all pages.
	Total int `json:"total"`

	// NextOffset is the offset value to use for retrieving the next page of results.
	NextOffset int `json:"next_offset"`
}
