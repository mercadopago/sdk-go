package merchantorder

// Request represents the body payload sent to the MercadoPago Merchant Orders API
// when creating a new merchant order via [Client.Create]. All fields use omitempty
// so only the provided values are sent to the API.
type Request struct {
	// Collector identifies the seller (collector) who will receive the payment.
	Collector *CollectorRequest `json:"collector,omitempty"`

	// Payer identifies the buyer for this merchant order.
	Payer *PayerRequest `json:"payer,omitempty"`

	// Items lists the products or services included in this merchant order.
	Items []ItemRequest `json:"items,omitempty"`

	// PreferenceID is the checkout preference identifier that originated this order.
	PreferenceID string `json:"preference_id,omitempty"`

	// ApplicationID is the application ID that created the merchant order.
	ApplicationID string `json:"application_id,omitempty"`

	// SiteID is the MercadoPago site identifier (e.g., "MLB" for Brazil, "MLA" for Argentina).
	SiteID string `json:"site_id,omitempty"`

	// NotificationURL is the URL where MercadoPago sends webhook notifications about order status changes.
	NotificationURL string `json:"notification_url,omitempty"`

	// AdditionalInfo is free-text information attached to the order for the integrator's reference.
	AdditionalInfo string `json:"additional_info,omitempty"`

	// ExternalReference is an external identifier for reconciliation with the integrator's system.
	ExternalReference string `json:"external_reference,omitempty"`

	// Marketplace identifies the marketplace that originated the order.
	Marketplace string `json:"marketplace,omitempty"`

	// Version is the version number of the merchant order resource.
	Version int `json:"version,omitempty"`

	// SponsorID is the MercadoPago user ID of the sponsor in marketplace scenarios.
	SponsorID int `json:"sponsor_id,omitempty"`
}

// PayerRequest represents the buyer's information when creating a merchant order.
type PayerRequest struct {
	// Nickname is the buyer's MercadoPago username.
	Nickname string `json:"nickname,omitempty"`

	// ID is the buyer's MercadoPago user identifier.
	ID int `json:"id,omitempty"`
}

// ItemRequest represents a product or service included in a merchant order creation request.
type ItemRequest struct {
	// ID is the item identifier in the integrator's system.
	ID string `json:"id,omitempty"`

	// Title is the item name displayed to the buyer.
	Title string `json:"title,omitempty"`

	// Description is a detailed description of the item.
	Description string `json:"description,omitempty"`

	// PictureURL is the URL of the item image.
	PictureURL string `json:"picture_url,omitempty"`

	// CategoryID is the MercadoPago category identifier for the item.
	CategoryID string `json:"category_id,omitempty"`

	// CurrencyID is the ISO 4217 currency code for the item price (e.g., "BRL", "ARS").
	CurrencyID string `json:"currency_id,omitempty"`

	// UnitPrice is the price per unit of the item.
	UnitPrice float64 `json:"unit_price,omitempty"`

	// Quantity is the number of units of this item.
	Quantity int `json:"quantity,omitempty"`
}

// CollectorRequest identifies the seller (collector) who will receive the payment
// for a merchant order.
type CollectorRequest struct {
	// ID is the seller's MercadoPago user identifier.
	ID int64 `json:"id,omitempty"`
}
