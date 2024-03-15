package merchantorder

import (
	"time"
)

// Response represents a merchant order resource.
type Response struct {
	Payer       PayerResponse      `json:"payer"`
	Collector   CollectorResponse  `json:"collector"`
	DateCreated time.Time          `json:"date_created"`
	LastUpdated time.Time          `json:"last_updated"`
	Items       []ItemResponse     `json:"items"`
	Shipments   []ShipmentResponse `json:"shipments"`
	Payments    []PaymentResponse  `json:"payments"`

	PreferenceID      string  `json:"preference_id"`
	ApplicationID     string  `json:"application_id"`
	Status            string  `json:"status"`
	SiteID            string  `json:"site_id"`
	NotificationURL   string  `json:"notification_url"`
	AdditionalInfo    string  `json:"additional_info"`
	ExternalReference string  `json:"external_reference"`
	Marketplace       string  `json:"marketplace"`
	SponsorID         string  `json:"sponsor_id"`
	OrderStatus       string  `json:"order_status"`
	PaidAmount        float64 `json:"paid_amount"`
	RefundedAmount    float64 `json:"refunded_amount"`
	ShippingCost      float64 `json:"shipping_cost"`
	TotalAmount       float64 `json:"total_amount"`
	ID                int64   `json:"id"`
	Cancelled         bool    `json:"cancelled"`
}

// PayerResponse represents buyer information.
type PayerResponse struct {
	Nickname string `json:"nickname"`
	ID       int64  `json:"id"`
}

// CollectorResponse represents collector information.
type CollectorResponse struct {
	Nickname string `json:"nickname"`
	ID       int64  `json:"id"`
}

// PaymentResponse represents payment information.
type PaymentResponse struct {
	DateApproved time.Time `json:"date_approved"`
	DateCreated  time.Time `json:"date_created"`
	LastModified time.Time `json:"last_modified"`

	Status            string  `json:"status"`
	StatusDetails     string  `json:"status_details"`
	OperationType     string  `json:"operation_type"`
	CurrencyID        string  `json:"currency_id"`
	TransactionAmount float64 `json:"transaction_amount"`
	TotalPaidAmount   float64 `json:"total_paid_amount"`
	ShippingCost      float64 `json:"shipping_cost"`
	AmountRefunded    float64 `json:"amount_refunded"`
	ID                int64   `json:"id"`
}

// ItemResponse represents item information.
type ItemResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PictureURL  string  `json:"picture_url"`
	CurrencyID  string  `json:"currency_id"`
	CategoryID  string  `json:"category_id"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
}

// ShipmentResponse represents shipment information.
type ShipmentResponse struct {
	ReceiverAddress  ReceiverAddressResponse `json:"receiver_address"`
	ShippingOption   ShippingOptionResponse  `json:"shipping_option"`
	DateCreated      time.Time               `json:"date_created"`
	LastModified     time.Time               `json:"last_modified"`
	DateFirstPrinted time.Time               `json:"date_first_printed"`

	ShippingType      string           `json:"shipping_type"`
	ShippingMode      string           `json:"shipping_mode"`
	PickingType       string           `json:"picking_type"`
	Status            string           `json:"status"`
	ShippingSubstatus string           `json:"shipping_substatus"`
	ServiceID         string           `json:"service_id"`
	ID                int64            `json:"id"`
	SenderID          int64            `json:"sender_id"`
	ReceiverID        int64            `json:"receiver_id"`
	Items             []map[string]any `json:"items"`
}

// ReceiverAddressResponse represents receiver address information.
type ReceiverAddressResponse struct {
	City    ReceiverAddressCityResponse    `json:"city"`
	State   ReceiverAddressStateResponse   `json:"state"`
	Country ReceiverAddressCountryResponse `json:"country"`

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
}

// ShippingOptionResponse represents shipping option information.
type ShippingOptionResponse struct {
	Speed             ShippingSpeedResponse             `json:"speed"`
	EstimatedDelivery ShippingEstimatedDeliveryResponse `json:"estimated_delivery"`

	Name             string  `json:"name"`
	CurrencyID       string  `json:"currency_id"`
	Cost             float64 `json:"cost"`
	ListCost         float64 `json:"list_cost"`
	ID               int64   `json:"id"`
	ShippingMethodID int64   `json:"shipping_method_id"`
}

// ReceiverAddressCityResponse represents city information.
type ReceiverAddressCityResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ReceiverAddressStateResponse represents state information.
type ReceiverAddressStateResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ReceiverAddressCountryResponse represents country information.
type ReceiverAddressCountryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ShippingEstimatedDeliveryResponse represents estimated delivery information.
type ShippingEstimatedDeliveryResponse struct {
	Date time.Time `json:"date"`

	TimeFrom string `json:"time_from"`
	TimeTo   string `json:"time_to"`
}

// ShippingSpeedResponse represents shipping speed information.
type ShippingSpeedResponse struct {
	Handling int64 `json:"handling"`
	Shipping int64 `json:"shipping"`
}
