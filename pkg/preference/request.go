package preference

import (
	"time"
)

// Request contains parameters to create/update a preference.
type Request struct {
	BackURLs            *BackURLsRequest            `json:"back_urls,omitempty"`
	DifferentialPricing *DifferentialPricingRequest `json:"differential_pricing,omitempty"`
	Payer               *PayerRequest               `json:"payer,omitempty"`
	PaymentMethods      *PaymentMethodsRequest      `json:"payment_methods,omitempty"`
	Shipments           *ShipmentsRequest           `json:"shipments,omitempty"`
	DateOfExpiration    *time.Time                  `json:"date_of_expiration,omitempty"`
	ExpirationDateFrom  *time.Time                  `json:"expiration_date_from,omitempty"`
	ExpirationDateTo    *time.Time                  `json:"expiration_date_to,omitempty"`
	Items               []ItemRequest               `json:"items,omitempty"`
	Taxes               []TaxRequest                `json:"taxes,omitempty"`
	Tracks              []TrackRequest              `json:"tracks,omitempty"`
	Amounts             *AmountsRequest             `json:"amounts,omitempty"`
	CounterCurrency     *CounterCurrencyRequest     `json:"counter_currency,omitempty"`

	AdditionalInfo      string         `json:"additional_info,omitempty"`
	AutoReturn          string         `json:"auto_return,omitempty"`
	ExternalReference   string         `json:"external_reference,omitempty"`
	Marketplace         string         `json:"marketplace,omitempty"`
	OperationType       string         `json:"operation_type,omitempty"`
	NotificationURL     string         `json:"notification_url,omitempty"`
	Purpose             string         `json:"purpose,omitempty"`
	StatementDescriptor string         `json:"statement_descriptor,omitempty"`
	MarketplaceFee      float64        `json:"marketplace_fee,omitempty"`
	BinaryMode          bool           `json:"binary_mode,omitempty"`
	Expires             bool           `json:"expires,omitempty"`
	ProcessingModes     []string       `json:"processing_modes,omitempty"`
	Metadata            map[string]any `json:"metadata,omitempty"`
	DeviceID            string         `json:"device_id,omitempty"`
}

// BackURLsRequest contains callback URLs.
type BackURLsRequest struct {
	Success string `json:"success,omitempty"`
	Pending string `json:"pending,omitempty"`
	Failure string `json:"failure,omitempty"`
}

// DifferentialPricingRequest contains information about differential pricing configuration.
type DifferentialPricingRequest struct {
	ID int `json:"id,omitempty"`
}

type PassengerRequest struct {
	FirstName            string `json:"first_name,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	IdentificationType   string `json:"identification_type,omitempty"`
	IdentificationNumber string `json:"identification_number,omitempty"`
}

type RouteRequest struct {
	Departure         string     `json:"departure,omitempty"`
	Destination       string     `json:"destination,omitempty"`
	DepartureDateTime *time.Time `json:"departure_date_time,omitempty"`
	ArrivalDateTime   *time.Time `json:"arrival_date_time,omitempty"`
	Company           string     `json:"company,omitempty"`
}

type CategoryDescriptorRequest struct {
	EventDate *time.Time `json:"event_date,omitempty"`
	Type      string     `json:"type,omitempty"`

	Passenger *PassengerRequest `json:"passenger,omitempty"`
	Route     *RouteRequest     `json:"route,omitempty"`
}

// ItemRequest represents a purchased item.
type ItemRequest struct {
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Type        string  `json:"type,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	CurrencyID  string  `json:"currency_id,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	Warranty    string  `json:"warranty,omitempty"`

	CategoryDescriptor CategoryDescriptorRequest `json:"category_descriptor,omitempty"`
}

// PayerRequest contains payer information in the preference.
type PayerRequest struct {
	DateCreated           *time.Time `json:"date_created,omitempty"`
	Name                  string     `json:"name,omitempty"`
	Surname               string     `json:"surname,omitempty"`
	Email                 string     `json:"email,omitempty"`
	AuthenticationType    string     `json:"authentication_type,omitempty"`
	IsPrimeUser           bool       `json:"is_prime_user,omitempty"`
	IsFirstPurchaseOnline bool       `json:"is_first_purchase_online,omitempty"`
	RegistrationDate      *time.Time `json:"registration_date,omitempty"`
	LastPurchaseDate      *time.Time `json:"last_purchase_date,omitempty"`

	Identification *IdentificationRequest `json:"identification,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`
	Address        *AddressRequest        `json:"address,omitempty"`
}

// PhoneRequest represents a telephone number.
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// IdentificationRequest is a base type that represents identifications, such as customer identification.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// AddressRequest represents an address.
type AddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// PaymentMethodsRequest contains information about payment methods in the preference.
type PaymentMethodsRequest struct {
	ExcludedPaymentMethods []ExcludedPaymentMethodRequest `json:"excluded_payment_methods,omitempty"`
	ExcludedPaymentTypes   []ExcludedPaymentTypeRequest   `json:"excluded_payment_types,omitempty"`

	DefaultPaymentMethodID string `json:"default_payment_method_id,omitempty"`
	Installments           int    `json:"installments,omitempty"`
	DefaultInstallments    int    `json:"default_installments,omitempty"`
}

// ExcludedPaymentMethodRequest contains information about the payment method in the preference.
type ExcludedPaymentMethodRequest struct {
	ID string `json:"id,omitempty"`
}

// ExcludedPaymentTypeRequest contains information about the payment type in the preference.
type ExcludedPaymentTypeRequest struct {
	ID string `json:"id,omitempty"`
}

// ShipmentsRequest contains information about shipments in the preference.
type ShipmentsRequest struct {
	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`
	FreeMethods     []FreeMethodRequest     `json:"free_methods,omitempty"`

	Mode                  string  `json:"mode,omitempty"`
	Dimensions            string  `json:"dimensions,omitempty"`
	DefaultShippingMethod string  `json:"default_shipping_method,omitempty"`
	Cost                  float64 `json:"cost,omitempty"`
	LocalPickup           bool    `json:"local_pickup,omitempty"`
	FreeShipping          bool    `json:"free_shipping,omitempty"`
	ExpressShipment       bool    `json:"express_shipment,omitempty"`
}

// FreeMethodRequest contains information about free shipping methods in the preference.
type FreeMethodRequest struct {
	ID int `json:"id,omitempty"`
}

// ReceiverAddressRequest contains information about the send address in the preference.
type ReceiverAddressRequest struct {
	ZipCode         string `json:"zip_code,omitempty"`
	StreetName      string `json:"street_name,omitempty"`
	StreetNumber    string `json:"street_number,omitempty"`
	CountryName     string `json:"country_name,omitempty"`
	StateName       string `json:"state_name,omitempty"`
	Floor           string `json:"floor,omitempty"`
	Apartment       string `json:"apartment,omitempty"`
	CityName        string `json:"city_name,omitempty"`
	LocalPickup     bool   `json:"local_pickup,omitempty"`
	ExpressShipment bool   `json:"express_shipment,omitempty"`
}

// TaxRequest contains information about taxes in the preference.
type TaxRequest struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// TrackRequest contains information about the tracking to be performed during user interaction in the Checkout flow.
type TrackRequest struct {
	Values *ValuesRequest `json:"values,omitempty"`

	Type string `json:"type,omitempty"`
}

// ValuesRequest contains the values of the tracks to be executed during user interaction in the Checkout flow.
type ValuesRequest struct {
	ConversionID    string `json:"conversion_id,omitempty"`
	ConversionLabel string `json:"conversion_label,omitempty"`
	PixelID         string `json:"pixel_id,omitempty"`
}

// AmountsRequest represents amounts request within Request.
type AmountsRequest struct {
	Collector UserAmountsRequest `json:"collector,omitempty"`
	Payer     UserAmountsRequest `json:"payer,omitempty"`
}

// UserAmountsRequest represents user amounts request within AmountsRequest.
type UserAmountsRequest struct {
	CurrencyID  string  `json:"currency_id,omitempty"`
	Transaction float64 `json:"transaction,omitempty"`
}

// CounterCurrencyRequest represents counter currency request within Request.
type CounterCurrencyRequest struct {
	CurrencyID string `json:"currency_id,omitempty"`
}
