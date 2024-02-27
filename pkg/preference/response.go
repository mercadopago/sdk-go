package preference

import (
	"time"
)

type Response struct {
	DateOfExpiration    *time.Time                            `json:"date_of_expiration"`
	ExpirationDateFrom  *time.Time                            `json:"expiration_date_from"`
	ExpirationDateTo    *time.Time                            `json:"expiration_date_to"`
	LastUpdated         *time.Time                            `json:"last_updated"`
	DateCreated         *time.Time                            `json:"date_created"`
	Payer               PreferencePayerResponse               `json:"payer"`
	PaymentMethods      PreferencePaymentMethodsResponse      `json:"payment_methods"`
	BackURLS            PreferenceBackUrlsResponse            `json:"back_urls"`
	Shipments           PreferenceShipmentsResponse           `json:"shipments"`
	DifferentialPricing PreferenceDifferentialPricingResponse `json:"differential_pricing"`
	Taxes               []PreferenceTaxResponse               `json:"taxes"`
	Tracks              []PreferenceTrackResponse             `json:"tracks"`
	Items               []PreferenceItemResponse              `json:"items"`

	ID                  string                 `json:"id"`
	ClientID            string                 `json:"client_id"`
	NotificationURL     string                 `json:"notification_url"`
	StatementDescriptor string                 `json:"statement_descriptor"`
	Marketplace         string                 `json:"marketplace"`
	ExternalReference   string                 `json:"external_reference"`
	AdditionalInfo      string                 `json:"additional_info"`
	AutoReturn          string                 `json:"auto_return"`
	OperationType       string                 `json:"operation_type"`
	InitPoint           string                 `json:"init_point"`
	SandboxInitPoint    string                 `json:"sandbox_init_point"`
	SiteID              string                 `json:"site_id"`
	CollectorID         int64                  `json:"collector_id"`
	Expires             bool                   `json:"expires"`
	BinaryMode          bool                   `json:"binary_mode"`
	MarketplaceFee      float64                `json:"marketplace_fee"`
	ProcessingModes     []string               `json:"processing_modes"`
	Metadata            map[string]interface{} `json:"metadata"`
}

// PreferenceItemResponse represents an item.
type PreferenceItemResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CurrencyID  string  `json:"currency_id"`
	PictureURL  string  `json:"picture_url"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

// PreferencePayerResponse contains payer information in the preference.
type PreferencePayerResponse struct {
	Phone          PhoneResponse          `json:"phone"`
	Identification IdentificationResponse `json:"identification"`
	Address        AddressResponse        `json:"address"`
	DateCreated    *time.Time             `json:"date_created"`
	LastPurchase   *time.Time             `json:"last_purchase"`

	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// PreferencePaymentMethodsResponse contains information about payment methods in the preference.
type PreferencePaymentMethodsResponse struct {
	ExcludedPaymentMethods []PreferencePaymentMethodResponse `json:"excluded_payment_methods"`
	ExcludedPaymentTypes   []PreferencePaymentTypeResponse   `json:"excluded_payment_types"`

	DefaultPaymentMethodID string `json:"default_payment_method_id"`
	Installments           int    `json:"installments"`
	DefaultInstallments    int    `json:"default_installments"`
}

// PreferencePaymentMethodResponse contains information about the payment method in the preference.
type PreferencePaymentMethodResponse struct {
	ID string `json:"id"`
}

// PreferencePaymentTypeResponse contains information about the type of payment in the preference.
type PreferencePaymentTypeResponse struct {
	ID string `json:"id"`
}

// PreferenceBackUrlsResponse contains preference back URLs.
type PreferenceBackUrlsResponse struct {
	Success string `json:"success"`
	Pending string `json:"pending"`
	Failure string `json:"failure"`
}

// PreferenceShipmentsResponse contains preference shipping information.
type PreferenceShipmentsResponse struct {
	ReceiverAddress PreferenceReceiverAddressResponse `json:"receiver_address"`
	FreeMethods     []PreferenceFreeMethodResponse    `json:"free_methods"`

	Mode                  string  `json:"mode"`
	Dimensions            string  `json:"dimensions"`
	DefaultShippingMethod string  `json:"default_shipping_method"`
	Cost                  float64 `json:"cost"`
	LocalPickup           bool    `json:"local_pickup"`
	FreeShipping          bool    `json:"free_shipping"`
	ExpressShipment       bool    `json:"express_shipment"`
}

// PreferenceFreeMethodResponse contains information about free shipping methods.
type PreferenceFreeMethodResponse struct {
	ID int64 `json:"id"`
}

// PreferenceReceiverAddressResponse represents a sending address.
type PreferenceReceiverAddressResponse struct {
	AddressResponse

	CountryName string `json:"country_name"`
	StateName   string `json:"state_name"`
	Floor       string `json:"floor"`
	Apartment   string `json:"apartment"`
	CityName    string `json:"city_name"`
}

// PreferenceDifferentialPricingResponse contains information about the differential pricing configuration in the preference.
type PreferenceDifferentialPricingResponse struct {
	ID int64 `json:"id"`
}

// PreferenceTaxResponse contains information about taxes in the preference.
type PreferenceTaxResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// PreferenceTrackResponse represents a trace to be executed during user interaction in the Checkout flow.
type PreferenceTrackResponse struct {
	Values PreferenceTrackValuesResponse `json:"values"`

	Type string `json:"type"`
}

// PreferenceTrackValuesResponse contains the values ​​of the tracks to be executed during user interaction in the Checkout flow.
type PreferenceTrackValuesResponse struct {
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
