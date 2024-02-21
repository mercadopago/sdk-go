package merchantorder

import (
	"time"
)

// Response represents a merchant order resource.
type Response struct {
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
	ID                int64   `json:"id"`
	Cancelled         bool    `json:"cancelled"`
	PaidAmount        float64 `json:"paid_amount"`
	RefundedAmount    float64 `json:"refunded_amount"`
	ShippingCost      float64 `json:"shipping_cost"`
	TotalAmount       float64 `json:"total_amount"`

	DateCreated *time.Time         `json:"date_created"`
	LastUpdated *time.Time         `json:"last_updated"`
	Payer       PayerResponse      `json:"payer"`
	Collector   CollectorResponse  `json:"collector"`
	Items       []ItemResponse     `json:"items"`
	Shipments   []ShipmentResponse `json:"shipments"`
	Payments    []PaymentResponse  `json:"payments"`
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
	Status            string  `json:"status"`
	StatusDetails     string  `json:"status_details"`
	OperationType     string  `json:"operation_type"`
	CurrencyID        string  `json:"currency_id"`
	ID                int64   `json:"id"`
	TransactionAmount float64 `json:"transaction_amount"`
	TotalPaidAmount   float64 `json:"total_paid_amount"`
	ShippingCost      float64 `json:"shipping_cost"`
	AmountRefunded    float64 `json:"amount_refunded"`

	DateApproved *time.Time `json:"date_approved"`
	DateCreated  *time.Time `json:"date_created"`
	LastModified *time.Time `json:"last_modified"`
}

// ItemResponse represents item information.
type ItemResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PictureURL  string  `json:"picture_url"`
	CurrencyID  string  `json:"currency_id"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

// ShipmentResponse represents shipment information.
type ShipmentResponse struct {
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
	ReceiverAddress  ReceiverAddressResponse `json:"receiver_address"`
	ShippingOption   ShippingOptionResponse  `json:"shipping_option"`
}

// ReceiverAddressResponse represents receiver address information.
type ReceiverAddressResponse struct {
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

	City    ReceiverAddressCityResponse    `json:"city"`
	State   ReceiverAddressStateResponse   `json:"state"`
	Country ReceiverAddressCountryResponse `json:"country"`
}

// ShippingOptionResponse represents shipping option information.
type ShippingOptionResponse struct {
	Name             string  `json:"name"`
	CurrencyID       string  `json:"currency_id"`
	ID               int64   `json:"id"`
	ShippingMethodID int64   `json:"shipping_method_id"`
	Cost             float64 `json:"cost"`
	ListCost         float64 `json:"list_cost"`

	Speed             ShippingSpeedResponse             `json:"speed"`
	EstimatedDelivery ShippingEstimatedDeliveryResponse `json:"estimated_delivery"`
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
	TimeFrom string `json:"time_from"`
	TimeTo   string `json:"time_to"`

	Date *time.Time `json:"date"`
}

// ShippingSpeedResponse represents shipping speed information.
type ShippingSpeedResponse struct {
	Handling int64 `json:"handling"`
	Shipping int64 `json:"shipping"`
}
