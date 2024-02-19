package merchantorder

import "time"

// UpdateRequest represents merchant order update.
type UpdateRequest struct {
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
	Shipments []ShipmentRequest `json:"shipments"`
}

// ShipmentRequest represents shipment information.
type ShipmentRequest struct {
	ShippingType      string                   `json:"shipping_type"`
	ShippingMode      string                   `json:"shipping_mode"`
	PickingType       string                   `json:"picking_type"`
	Status            string                   `json:"status"`
	ShippingSubstatus string                   `json:"shipping_substatus"`
	ServiceID         string                   `json:"service_id"`
	ID                int64                    `json:"id"`
	SenderID          int64                    `json:"sender_id"`
	ReceiverID        int64                    `json:"receiver_id"`
	Items             []map[string]interface{} `json:"items"`

	DateCreated      *time.Time              `json:"date_created"`
	LastModified     *time.Time              `json:"last_modified"`
	DateFirstPrinted *time.Time              `json:"date_first_printed"`
	ReceiverAddress  *ReceiverAddressRequest `json:"receiver_address"`
	ShippingOption   *ShippingOptionRequest  `json:"shipping_option"`
}

// ReceiverAddressRequest represents receiver address information.
type ReceiverAddressRequest struct {
	AddressLine  string `json:"address_line"`
	Apartment    string `json:"apartment"`
	Comment      string `json:"comment"`
	Contact      string `json:"contact"`
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Floor        string `json:"floor"`
	Phone        string `json:"phone"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	ID           int64  `json:"id"`

	City    *ReceiverAddressCityRequest    `json:"city"`
	State   *ReceiverAddressStateRequest   `json:"state"`
	Country *ReceiverAddressCountryRequest `json:"country"`
}

// ShippingOptionRequest represents shipping option information.
type ShippingOptionRequest struct {
	Name             string  `json:"name"`
	CurrencyID       string  `json:"currency_id"`
	ShippingMethodID int64   `json:"shipping_method_id"`
	ID               int64   `json:"id"`
	Cost             float64 `json:"cost"`
	ListCost         float64 `json:"list_cost"`

	EstimatedDelivery *ShippingEstimatedDeliveryRequest `json:"estimated_delivery"`
	Speed             *ShippingSpeedRequest             `json:"speed"`
}

// ReceiverAddressCityRequest represents city information.
type ReceiverAddressCityRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ReceiverAddressStateRequest represents state information.
type ReceiverAddressStateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ReceiverAddressCountryRequest represents country information.
type ReceiverAddressCountryRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ShippingEstimatedDeliveryRequest represents estimated delivery information.
type ShippingEstimatedDeliveryRequest struct {
	TimeFrom string `json:"time_from"`
	TimeTo   string `json:"time_to"`

	Date *time.Time `json:"date"`
}

// ShippingSpeedRequest represents shipping speed information.
type ShippingSpeedRequest struct {
	Handling int64 `json:"handling"`
	Shipping int64 `json:"shipping"`
}
