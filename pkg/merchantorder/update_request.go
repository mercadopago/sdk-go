package merchantorder

import "time"

// UpdateRequest represents the body payload sent to the MercadoPago Merchant Orders API
// when updating an existing merchant order via [Client.Update]. All fields use omitempty
// so only the provided values are sent to the API.
//
// Unlike [Request], UpdateRequest supports shipment modifications and uses
// [ItemUpdateRequest] to allow updating item quantities by ID.
type UpdateRequest struct {
	// Collector identifies the seller (collector) who will receive the payment.
	Collector *CollectorRequest `json:"collector,omitempty"`

	// Payer identifies the buyer for this merchant order.
	Payer *PayerRequest `json:"payer,omitempty"`

	// Items lists the items to update, allowing quantity changes by item ID.
	Items []ItemUpdateRequest `json:"items,omitempty"`

	// Shipments lists the shipments to update or add to the merchant order.
	Shipments []ShipmentRequest `json:"shipments,omitempty"`

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

// ShipmentRequest represents shipment information sent when updating a merchant order.
// It includes the full set of shipment details including address, shipping option,
// status, and lifecycle timestamps.
type ShipmentRequest struct {
	// ReceiverAddress is the delivery address for this shipment.
	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`

	// ShippingOption contains the shipping method and cost details.
	ShippingOption *ShippingOptionRequest `json:"shipping_option,omitempty"`

	// DateCreated is the date when the shipment was created.
	DateCreated *time.Time `json:"date_created,omitempty"`

	// LastModified is the date when the shipment was last modified.
	LastModified *time.Time `json:"last_modified,omitempty"`

	// DateFirstPrinted is the date when the shipping label was first printed.
	DateFirstPrinted *time.Time `json:"date_first_printed,omitempty"`

	// ShippingType is the type of shipping (e.g., "custom", "mercado_envios").
	ShippingType string `json:"shipping_type,omitempty"`

	// ShippingMode is the shipping mode (e.g., "me2", "custom").
	ShippingMode string `json:"shipping_mode,omitempty"`

	// PickingType is the picking type for the shipment.
	PickingType string `json:"picking_type,omitempty"`

	// Status is the current shipment status (e.g., "pending", "shipped", "delivered").
	Status string `json:"status,omitempty"`

	// ShippingSubstatus provides additional detail about the current shipping status.
	ShippingSubstatus string `json:"shipping_substatus,omitempty"`

	// ServiceID is the identifier of the shipping service used.
	ServiceID string `json:"service_id,omitempty"`

	// ID is the unique shipment identifier assigned by MercadoPago.
	ID int `json:"id,omitempty"`

	// SenderID is the MercadoPago user ID of the sender.
	SenderID int `json:"sender_id,omitempty"`

	// ReceiverID is the MercadoPago user ID of the receiver.
	ReceiverID int `json:"receiver_id,omitempty"`

	// Items lists the items included in this shipment as key-value pairs.
	Items []map[string]any `json:"items,omitempty"`
}

// ReceiverAddressRequest represents the delivery address for a shipment when
// updating a merchant order, including geographic coordinates and contact details.
type ReceiverAddressRequest struct {
	// City contains city information for the delivery address.
	City *CityRequest `json:"city,omitempty"`

	// State contains state or province information for the delivery address.
	State *StateRequest `json:"state,omitempty"`

	// Country contains country information for the delivery address.
	Country *CountryRequest `json:"country,omitempty"`

	// AddressLine is the full address line.
	AddressLine string `json:"address_line,omitempty"`

	// Apartment is the apartment or unit number.
	Apartment string `json:"apartment,omitempty"`

	// Comment is an additional comment about the delivery address.
	Comment string `json:"comment,omitempty"`

	// Contact is the contact name at the delivery address.
	Contact string `json:"contact,omitempty"`

	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code,omitempty"`

	// StreetName is the street name.
	StreetName string `json:"street_name,omitempty"`

	// StreetNumber is the street number.
	StreetNumber string `json:"street_number,omitempty"`

	// Floor is the floor number, if applicable.
	Floor string `json:"floor,omitempty"`

	// Phone is the contact phone number at the delivery address.
	Phone string `json:"phone,omitempty"`

	// Latitude is the geographic latitude of the delivery address.
	Latitude string `json:"latitude,omitempty"`

	// Longitude is the geographic longitude of the delivery address.
	Longitude string `json:"longitude,omitempty"`

	// ID is the unique address identifier.
	ID int `json:"id,omitempty"`
}

// ShippingOptionRequest represents the shipping option details when updating a
// merchant order, including cost, method, speed, and estimated delivery information.
type ShippingOptionRequest struct {
	// EstimatedDelivery contains the estimated delivery date and time window.
	EstimatedDelivery *EstimatedDeliveryRequest `json:"estimated_delivery,omitempty"`

	// Speed contains handling and shipping time information.
	Speed *SpeedRequest `json:"speed,omitempty"`

	// Name is the name of the shipping option.
	Name string `json:"name,omitempty"`

	// CurrencyID is the ISO 4217 currency code for the shipping cost.
	CurrencyID string `json:"currency_id,omitempty"`

	// Cost is the actual shipping cost charged to the buyer.
	Cost float64 `json:"cost,omitempty"`

	// ListCost is the listed (original) shipping cost before any discounts.
	ListCost float64 `json:"list_cost,omitempty"`

	// ShippingMethodID is the identifier of the shipping method used.
	ShippingMethodID int `json:"shipping_method_id,omitempty"`

	// ID is the unique shipping option identifier.
	ID int `json:"id,omitempty"`
}

// CityRequest represents city information for a shipment delivery address.
type CityRequest struct {
	// ID is the city identifier.
	ID string `json:"id,omitempty"`

	// Name is the city name.
	Name string `json:"name,omitempty"`
}

// StateRequest represents state or province information for a shipment delivery address.
type StateRequest struct {
	// ID is the state identifier.
	ID string `json:"id,omitempty"`

	// Name is the state name.
	Name string `json:"name,omitempty"`
}

// CountryRequest represents country information for a shipment delivery address.
type CountryRequest struct {
	// ID is the ISO 3166-1 country code.
	ID string `json:"id,omitempty"`

	// Name is the country name.
	Name string `json:"name,omitempty"`
}

// EstimatedDeliveryRequest represents the estimated delivery date and time window
// for a shipping option in a merchant order update.
type EstimatedDeliveryRequest struct {
	// Date is the estimated delivery date.
	Date *time.Time `json:"date,omitempty"`

	// TimeFrom is the start of the estimated delivery time window (e.g., "09:00").
	TimeFrom string `json:"time_from,omitempty"`

	// TimeTo is the end of the estimated delivery time window (e.g., "18:00").
	TimeTo string `json:"time_to,omitempty"`
}

// SpeedRequest represents handling and shipping time for a shipping option,
// measured in hours.
type SpeedRequest struct {
	// Handling is the handling time in hours before the package is shipped.
	Handling int `json:"handling,omitempty"`

	// Shipping is the shipping transit time in hours.
	Shipping int `json:"shipping,omitempty"`
}

// ItemUpdateRequest represents an item quantity update within a merchant order
// [UpdateRequest]. It allows changing the quantity of an existing item by its ID.
type ItemUpdateRequest struct {
	// ID is the item identifier to update.
	ID string `json:"id,omitempty"`

	// Quantity is the new number of units for this item.
	Quantity int `json:"quantity,omitempty"`
}
