package merchantorder

import (
	"time"
)

// Response represents the full merchant order resource returned by the MercadoPago
// Merchant Orders API. It is returned by [Client.Get], [Client.Create], and [Client.Update].
//
// A merchant order aggregates items, payments, and shipments into a single commercial
// transaction, tracking the overall financial status including paid, refunded, and total amounts.
type Response struct {
	// Payer contains information about the buyer.
	Payer PayerResponse `json:"payer"`

	// Collector contains information about the seller who receives the payment.
	Collector CollectorResponse `json:"collector"`

	// DateCreated is the date when the merchant order was created.
	DateCreated time.Time `json:"date_created"`

	// LastUpdated is the date when the merchant order was last modified.
	LastUpdated time.Time `json:"last_updated"`

	// Items lists the products or services included in this merchant order.
	Items []ItemResponse `json:"items"`

	// Shipments lists the shipments associated with this merchant order.
	Shipments []ShipmentResponse `json:"shipments"`

	// Payments lists the payments associated with this merchant order.
	Payments []PaymentResponse `json:"payments"`

	// PreferenceID is the checkout preference identifier that originated this order.
	PreferenceID string `json:"preference_id"`

	// ApplicationID is the application ID that created the merchant order.
	ApplicationID string `json:"application_id"`

	// Status is the merchant order status (e.g., "opened", "closed").
	Status string `json:"status"`

	// SiteID is the MercadoPago site identifier (e.g., "MLB" for Brazil, "MLA" for Argentina).
	SiteID string `json:"site_id"`

	// NotificationURL is the webhook URL for order status notifications.
	NotificationURL string `json:"notification_url"`

	// AdditionalInfo is free-text information attached to the order.
	AdditionalInfo string `json:"additional_info"`

	// ExternalReference is the integrator-provided external identifier for reconciliation.
	ExternalReference string `json:"external_reference"`

	// Marketplace identifies the marketplace that originated the order.
	Marketplace string `json:"marketplace"`

	// SponsorID is the sponsor's identifier in marketplace scenarios.
	SponsorID string `json:"sponsor_id"`

	// OrderStatus is the payment-aggregated status of the order (e.g., "paid", "partially_paid", "unpaid").
	OrderStatus string `json:"order_status"`

	// PaidAmount is the total amount already paid across all associated payments.
	PaidAmount float64 `json:"paid_amount"`

	// RefundedAmount is the total amount refunded across all associated payments.
	RefundedAmount float64 `json:"refunded_amount"`

	// ShippingCost is the total shipping cost for the order.
	ShippingCost float64 `json:"shipping_cost"`

	// TotalAmount is the total amount of the order, including items and shipping.
	TotalAmount float64 `json:"total_amount"`

	// ID is the unique merchant order identifier assigned by MercadoPago.
	ID int `json:"id"`

	// Cancelled indicates whether the merchant order has been cancelled.
	Cancelled bool `json:"cancelled"`
}

// PayerResponse represents the buyer's information returned within a merchant order [Response].
type PayerResponse struct {
	// Nickname is the buyer's MercadoPago username.
	Nickname string `json:"nickname"`

	// ID is the buyer's MercadoPago user identifier.
	ID int `json:"id"`
}

// CollectorResponse represents the seller's (collector's) information returned within
// a merchant order [Response].
type CollectorResponse struct {
	// Nickname is the seller's MercadoPago username.
	Nickname string `json:"nickname"`

	// ID is the seller's MercadoPago user identifier.
	ID int `json:"id"`
}

// PaymentResponse represents a payment associated with a merchant order, as returned
// within a [Response]. Each merchant order can have multiple payments.
type PaymentResponse struct {
	// DateApproved is the date when the payment was approved.
	DateApproved time.Time `json:"date_approved"`

	// DateCreated is the date when the payment was created.
	DateCreated time.Time `json:"date_created"`

	// LastModified is the date when the payment was last modified.
	LastModified time.Time `json:"last_modified"`

	// Status is the payment status (e.g., "approved", "pending", "rejected").
	Status string `json:"status"`

	// OperationType is the type of payment operation (e.g., "regular_payment").
	OperationType string `json:"operation_type"`

	// CurrencyID is the ISO 4217 currency code for the payment (e.g., "BRL", "ARS").
	CurrencyID string `json:"currency_id"`

	// TransactionAmount is the total transaction amount of the payment.
	TransactionAmount float64 `json:"transaction_amount"`

	// TotalPaidAmount is the amount actually paid, including fees.
	TotalPaidAmount float64 `json:"total_paid_amount"`

	// ShippingCost is the shipping cost included in this payment.
	ShippingCost float64 `json:"shipping_cost"`

	// AmountRefunded is the amount refunded from this payment.
	AmountRefunded float64 `json:"amount_refunded"`

	// ID is the unique payment identifier assigned by MercadoPago.
	ID int `json:"id"`
}

// ItemResponse represents a product or service item returned within a merchant order [Response].
type ItemResponse struct {
	// ID is the item identifier.
	ID string `json:"id"`

	// Title is the item name displayed to the buyer.
	Title string `json:"title"`

	// Description is a detailed description of the item.
	Description string `json:"description"`

	// PictureURL is the URL of the item image.
	PictureURL string `json:"picture_url"`

	// CurrencyID is the ISO 4217 currency code for the item price.
	CurrencyID string `json:"currency_id"`

	// CategoryID is the MercadoPago category identifier for the item.
	CategoryID string `json:"category_id"`

	// UnitPrice is the price per unit of the item.
	UnitPrice float64 `json:"unit_price"`

	// Quantity is the number of units of this item.
	Quantity int `json:"quantity"`
}

// ShipmentResponse represents a shipment associated with a merchant order, as returned
// within a [Response]. It includes the full lifecycle of a shipment including status,
// tracking, and delivery details.
type ShipmentResponse struct {
	// ReceiverAddress is the delivery address for this shipment.
	ReceiverAddress ReceiverAddressResponse `json:"receiver_address"`

	// ShippingOption contains the selected shipping method and cost details.
	ShippingOption ShippingOptionResponse `json:"shipping_option"`

	// DateCreated is the date when the shipment was created.
	DateCreated time.Time `json:"date_created"`

	// LastModified is the date when the shipment was last modified.
	LastModified time.Time `json:"last_modified"`

	// DateFirstPrinted is the date when the shipping label was first printed.
	DateFirstPrinted time.Time `json:"date_first_printed"`

	// ShippingType is the type of shipping (e.g., "custom", "mercado_envios").
	ShippingType string `json:"shipping_type"`

	// ShippingMode is the shipping mode (e.g., "me2", "custom").
	ShippingMode string `json:"shipping_mode"`

	// PickingType is the picking type for the shipment.
	PickingType string `json:"picking_type"`

	// Status is the current shipment status (e.g., "pending", "shipped", "delivered").
	Status string `json:"status"`

	// ShippingSubstatus provides additional detail about the current shipping status.
	ShippingSubstatus string `json:"shipping_substatus"`

	// ServiceID is the identifier of the shipping service used.
	ServiceID string `json:"service_id"`

	// ID is the unique shipment identifier assigned by MercadoPago.
	ID int `json:"id"`

	// SenderID is the MercadoPago user ID of the sender.
	SenderID int `json:"sender_id"`

	// ReceiverID is the MercadoPago user ID of the receiver.
	ReceiverID int `json:"receiver_id"`

	// Items lists the items included in this shipment as key-value pairs.
	Items []map[string]any `json:"items"`
}

// ReceiverAddressResponse represents the delivery address for a shipment within a
// merchant order [Response], including geographic coordinates and contact details.
type ReceiverAddressResponse struct {
	// City contains city information for the delivery address.
	City CityResponse `json:"city"`

	// State contains state or province information for the delivery address.
	State StateResponse `json:"state"`

	// Country contains country information for the delivery address.
	Country CountryResponse `json:"country"`

	// AddressLine is the full address line.
	AddressLine string `json:"address_line"`

	// Apartment is the apartment or unit number.
	Apartment string `json:"apartment"`

	// Comment is an additional comment about the delivery address.
	Comment string `json:"comment"`

	// Contact is the contact name at the delivery address.
	Contact string `json:"contact"`

	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code"`

	// StreetName is the street name.
	StreetName string `json:"street_name"`

	// StreetNumber is the street number.
	StreetNumber string `json:"street_number"`

	// Floor is the floor number, if applicable.
	Floor string `json:"floor"`

	// Phone is the contact phone number at the delivery address.
	Phone string `json:"phone"`

	// Latitude is the geographic latitude of the delivery address.
	Latitude string `json:"latitude"`

	// Longitude is the geographic longitude of the delivery address.
	Longitude string `json:"longitude"`

	// ID is the unique address identifier.
	ID int `json:"id"`
}

// ShippingOptionResponse represents the selected shipping option for a shipment,
// including cost, speed, and estimated delivery details.
type ShippingOptionResponse struct {
	// Speed contains handling and shipping time information.
	Speed SpeedResponse `json:"speed"`

	// EstimatedDelivery contains the estimated delivery date and time window.
	EstimatedDelivery EstimatedDeliveryResponse `json:"estimated_delivery"`

	// Name is the name of the shipping option.
	Name string `json:"name"`

	// CurrencyID is the ISO 4217 currency code for the shipping cost.
	CurrencyID string `json:"currency_id"`

	// Cost is the actual shipping cost charged to the buyer.
	Cost float64 `json:"cost"`

	// ListCost is the listed (original) shipping cost before any discounts.
	ListCost float64 `json:"list_cost"`

	// ID is the unique shipping option identifier.
	ID int `json:"id"`

	// ShippingMethodID is the identifier of the shipping method used.
	ShippingMethodID int `json:"shipping_method_id"`
}

// CityResponse represents a city returned within a shipment address in the MercadoPago API.
type CityResponse struct {
	// ID is the city identifier.
	ID string `json:"id"`

	// Name is the city name.
	Name string `json:"name"`
}

// StateResponse represents a state or province returned within a shipment address
// in the MercadoPago API.
type StateResponse struct {
	// ID is the state identifier.
	ID string `json:"id"`

	// Name is the state name.
	Name string `json:"name"`
}

// CountryResponse represents a country returned within a shipment address in the MercadoPago API.
type CountryResponse struct {
	// ID is the ISO 3166-1 country code.
	ID string `json:"id"`

	// Name is the country name.
	Name string `json:"name"`
}

// EstimatedDeliveryResponse represents the estimated delivery date and time window
// for a shipment within a merchant order.
type EstimatedDeliveryResponse struct {
	// Date is the estimated delivery date.
	Date time.Time `json:"date"`

	// TimeFrom is the start of the estimated delivery time window (e.g., "09:00").
	TimeFrom string `json:"time_from"`

	// TimeTo is the end of the estimated delivery time window (e.g., "18:00").
	TimeTo string `json:"time_to"`
}

// SpeedResponse represents the handling and shipping time for a shipping option,
// measured in hours.
type SpeedResponse struct {
	// Handling is the handling time in hours before the package is shipped.
	Handling int `json:"handling"`

	// Shipping is the shipping transit time in hours.
	Shipping int `json:"shipping"`
}
