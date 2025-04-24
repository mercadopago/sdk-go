package preference

import (
	"time"
)

// Response represents the response to a preference endpoint.
type Response struct {
	Payer               PayerResponse               `json:"payer"`
	PaymentMethods      PaymentMethodsResponse      `json:"payment_methods"`
	BackURLs            BackURLsResponse            `json:"back_urls"`
	Shipments           ShipmentsResponse           `json:"shipments"`
	DifferentialPricing DifferentialPricingResponse `json:"differential_pricing"`
	DateOfExpiration    time.Time                   `json:"date_of_expiration"`
	ExpirationDateFrom  time.Time                   `json:"expiration_date_from"`
	ExpirationDateTo    time.Time                   `json:"expiration_date_to"`
	LastUpdated         time.Time                   `json:"last_updated"`
	DateCreated         time.Time                   `json:"date_created"`
	Taxes               []TaxResponse               `json:"taxes"`
	Tracks              []TrackResponse             `json:"tracks"`
	Items               []ItemResponse              `json:"items"`

	ID                  string         `json:"id"`
	ClientID            string         `json:"client_id"`
	NotificationURL     string         `json:"notification_url"`
	StatementDescriptor string         `json:"statement_descriptor"`
	Marketplace         string         `json:"marketplace"`
	ExternalReference   string         `json:"external_reference"`
	AdditionalInfo      string         `json:"additional_info"`
	AutoReturn          string         `json:"auto_return"`
	OperationType       string         `json:"operation_type"`
	InitPoint           string         `json:"init_point"`
	SandboxInitPoint    string         `json:"sandbox_init_point"`
	SiteID              string         `json:"site_id"`
	MarketplaceFee      float64        `json:"marketplace_fee"`
	CollectorID         int64          `json:"collector_id"`
	Expires             bool           `json:"expires"`
	BinaryMode          bool           `json:"binary_mode"`
	ProcessingModes     []string       `json:"processing_modes"`
	Metadata            map[string]any `json:"metadata"`
}

// ItemResponse represents an item.
type ItemResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CurrencyID  string  `json:"currency_id"`
	PictureURL  string  `json:"picture_url"`
	CategoryID  string  `json:"category_id"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	Warranty    string  `json:"warranty"`

	CategoryDescriptor CategoryDescriptorResponse `json:"category_descriptor"`
}

type CategoryDescriptorResponse struct {
	Passenger PassengerRequest `json:"passenger"`
}

// PayerResponse contains payer information in the preference.
type PayerResponse struct {
	DateCreated           time.Time `json:"date_created"`
	LastPurchase          time.Time `json:"last_purchase"`
	Name                  string    `json:"name"`
	Surname               string    `json:"surname"`
	Email                 string    `json:"email"`
	AuthenticationType    string    `json:"authentication_type"`
	IsPrimerUser          bool      `json:"is_primer_user"`
	IsFirstPurchaseOnLine bool      `json:"is_first_purchase_on_line"`
	RegistrationDate      time.Time `json:"registration_date"`
	LastPurchaseDate      time.Time `json:"last_purchase_date"`

	Identification IdentificationResponse `json:"identification"`
	Phone          PhoneResponse          `json:"phone"`
	Address        AddressResponse        `json:"address"`
}

// PaymentMethodsResponse contains information about payment methods in the preference.
type PaymentMethodsResponse struct {
	ExcludedPaymentMethods []ExcludedPaymentMethodResponse `json:"excluded_payment_methods"`
	ExcludedPaymentTypes   []ExcludedPaymentTypeResponse   `json:"excluded_payment_types"`

	DefaultPaymentMethodID string `json:"default_payment_method_id"`
	Installments           int    `json:"installments"`
	DefaultInstallments    int    `json:"default_installments"`
}

// ExcludedPaymentMethodResponse contains information about the payment method in the preference.
type ExcludedPaymentMethodResponse struct {
	ID string `json:"id"`
}

// ExcludedPaymentTypeResponse contains information about the type of payment in the preference.
type ExcludedPaymentTypeResponse struct {
	ID string `json:"id"`
}

// BackURLsResponse contains preference back URLs.
type BackURLsResponse struct {
	Success string `json:"success"`
	Pending string `json:"pending"`
	Failure string `json:"failure"`
}

// ShipmentsResponse contains preference shipping information.
type ShipmentsResponse struct {
	ReceiverAddress ReceiverAddressResponse `json:"receiver_address"`
	FreeMethods     []FreeMethodResponse    `json:"free_methods"`

	Mode                  string  `json:"mode"`
	Dimensions            string  `json:"dimensions"`
	DefaultShippingMethod string  `json:"default_shipping_method"`
	Cost                  float64 `json:"cost"`
	LocalPickup           bool    `json:"local_pickup"`
	FreeShipping          bool    `json:"free_shipping"`
	ExpressShipment       bool    `json:"express_shipment"`
}

// FreeMethodResponse contains information about free shipping methods.
type FreeMethodResponse struct {
	ID int `json:"id"`
}

// ReceiverAddressResponse represents a sending address.
type ReceiverAddressResponse struct {
	Address AddressResponse `json:"address"`

	CountryName string `json:"country_name"`
	StateName   string `json:"state_name"`
	Floor       string `json:"floor"`
	Apartment   string `json:"apartment"`
	CityName    string `json:"city_name"`
}

// DifferentialPricingResponse contains information about the differential pricing configuration in the preference.
type DifferentialPricingResponse struct {
	ID int `json:"id"`
}

// TaxResponse contains information about taxes in the preference.
type TaxResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// TrackResponse represents a trace to be executed during user interaction in the Checkout flow.
type TrackResponse struct {
	Values ValuesResponse `json:"values"`

	Type string `json:"type"`
}

// ValuesResponse contains the values of the tracks to be executed during user interaction in the Checkout flow.
type ValuesResponse struct {
	ConversionID    string `json:"conversion_id"`
	ConversionLabel string `json:"conversion_label"`
	PixelID         string `json:"pixel_id"`
}

// PhoneResponse represents a telephone number.
type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

// IdentificationResponse is a base type that represents identifications, such as customer identification.
type IdentificationResponse struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// AddressResponse represents an address.
type AddressResponse struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}
