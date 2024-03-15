package merchantorder

import "time"

// UpdateRequest represents merchant order update.
type UpdateRequest struct {
	Collector *CollectorRequest   `json:"collector,omitempty"`
	Payer     *PayerRequest       `json:"payer,omitempty"`
	Items     []ItemUpdateRequest `json:"items,omitempty"`
	Shipments []ShipmentRequest   `json:"shipments,omitempty"`

	PreferenceID      string `json:"preference_id,omitempty"`
	ApplicationID     string `json:"application_id,omitempty"`
	SiteID            string `json:"site_id,omitempty"`
	NotificationURL   string `json:"notification_url,omitempty"`
	AdditionalInfo    string `json:"additional_info,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Marketplace       string `json:"marketplace,omitempty"`
	Version           int64  `json:"version,omitempty"`
	SponsorID         int64  `json:"sponsor_id,omitempty"`
}

// ShipmentRequest represents shipment information.
type ShipmentRequest struct {
	ReceiverAddress  *ReceiverAddressRequest `json:"receiver_address,omitempty"`
	ShippingOption   *ShippingOptionRequest  `json:"shipping_option,omitempty"`
	DateCreated      *time.Time              `json:"date_created,omitempty"`
	LastModified     *time.Time              `json:"last_modified,omitempty"`
	DateFirstPrinted *time.Time              `json:"date_first_printed,omitempty"`

	ShippingType      string           `json:"shipping_type,omitempty"`
	ShippingMode      string           `json:"shipping_mode,omitempty"`
	PickingType       string           `json:"picking_type,omitempty"`
	Status            string           `json:"status,omitempty"`
	ShippingSubstatus string           `json:"shipping_substatus,omitempty"`
	ServiceID         string           `json:"service_id,omitempty"`
	ID                int64            `json:"id,omitempty"`
	SenderID          int64            `json:"sender_id,omitempty"`
	ReceiverID        int64            `json:"receiver_id,omitempty"`
	Items             []map[string]any `json:"items,omitempty"`
}

// ReceiverAddressRequest represents receiver address information.
type ReceiverAddressRequest struct {
	City    *ReceiverAddressCityRequest    `json:"city,omitempty"`
	State   *ReceiverAddressStateRequest   `json:"state,omitempty"`
	Country *ReceiverAddressCountryRequest `json:"country,omitempty"`

	AddressLine  string `json:"address_line,omitempty"`
	Apartment    string `json:"apartment,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Contact      string `json:"contact,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	Floor        string `json:"floor,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	ID           int64  `json:"id,omitempty"`
}

// ShippingOptionRequest represents shipping option information.
type ShippingOptionRequest struct {
	EstimatedDelivery *ShippingEstimatedDeliveryRequest `json:"estimated_delivery,omitempty"`
	Speed             *ShippingSpeedRequest             `json:"speed,omitempty"`

	Name             string  `json:"name,omitempty"`
	CurrencyID       string  `json:"currency_id,omitempty"`
	Cost             float64 `json:"cost,omitempty"`
	ListCost         float64 `json:"list_cost,omitempty"`
	ShippingMethodID int64   `json:"shipping_method_id,omitempty"`
	ID               int64   `json:"id,omitempty"`
}

// ReceiverAddressCityRequest represents city information.
type ReceiverAddressCityRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ReceiverAddressStateRequest represents state information.
type ReceiverAddressStateRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ReceiverAddressCountryRequest represents country information.
type ReceiverAddressCountryRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ShippingEstimatedDeliveryRequest represents estimated delivery information.
type ShippingEstimatedDeliveryRequest struct {
	Date *time.Time `json:"date,omitempty"`

	TimeFrom string `json:"time_from,omitempty"`
	TimeTo   string `json:"time_to,omitempty"`
}

// ShippingSpeedRequest represents shipping speed information.
type ShippingSpeedRequest struct {
	Handling int64 `json:"handling,omitempty"`
	Shipping int64 `json:"shipping,omitempty"`
}

// ItemUpdateRequest represents item information.
type ItemUpdateRequest struct {
	ID       string `json:"id,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
