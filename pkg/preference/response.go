package preference

import (
	"time"
)

type Response struct {
	ID                  string                        `json:"id"`
	Items               []PreferenceItem              `json:"items"`
	Payer               PreferencePayer               `json:"payer"`
	ClientID            string                        `json:"client_id"`
	PaymentMethods      PreferencePaymentMethods      `json:"payment_methods"`
	BackUrls            PreferenceBackUrls            `json:"back_urls"`
	Shipments           PreferenceShipments           `json:"shipments"`
	NotificationURL     string                        `json:"notification_url"`
	StatementDescriptor string                        `json:"statement_descriptor"`
	ExternalReference   string                        `json:"external_reference"`
	Expires             bool                          `json:"expires"`
	DateOfExpiration    time.Time                     `json:"date_of_expiration"`
	ExpirationDateFrom  time.Time                     `json:"expiration_date_from"`
	ExpirationDateTo    time.Time                     `json:"expiration_date_to"`
	CollectorID         int64                         `json:"collector_id"`
	Marketplace         string                        `json:"marketplace"`
	MarketplaceFee      float64                       `json:"marketplace_fee"`
	AdditionalInfo      string                        `json:"additional_info"`
	AutoReturn          string                        `json:"auto_return"`
	OperationType       string                        `json:"operation_type"`
	DifferentialPricing PreferenceDifferentialPricing `json:"differential_pricing"`
	ProcessingModes     []string                      `json:"processing_modes"`
	BinaryMode          bool                          `json:"binary_mode"`
	Taxes               []PreferenceTax               `json:"taxes"`
	Tracks              []PreferenceTrack             `json:"tracks"`
	Metadata            map[string]interface{}        `json:"metadata"`
	InitPoint           string                        `json:"init_point"`
	SandboxInitPoint    string                        `json:"sandbox_init_point"`
	DateCreated         time.Time                     `json:"date_created"`
	SiteID              string                        `json:"site_id"`
	LastUpdated         time.Time                     `json:"last_updated"`
}

// PreferenceItem represents an item.
type PreferenceItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PictureURL  string  `json:"picture_url"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	CurrencyID  string  `json:"currency_id"`
}

// PreferencePayer contains payer information in the preference.
type PreferencePayer struct {
	Name           string         `json:"name"`
	Surname        string         `json:"surname"`
	Email          string         `json:"email"`
	Phone          Phone          `json:"phone"`
	Identification Identification `json:"identification"`
	Address        Address        `json:"address"`
	DateCreated    time.Time      `json:"date_created"`
	LastPurchase   time.Time      `json:"last_purchase"`
}

// PreferencePaymentMethods contains information about payment methods in the preference.
type PreferencePaymentMethods struct {
	ExcludedPaymentMethods []PreferencePaymentMethod `json:"excluded_payment_methods"`
	ExcludedPaymentTypes   []PreferencePaymentType   `json:"excluded_payment_types"`
	DefaultPaymentMethodID string                    `json:"default_payment_method_id"`
	Installments           int                       `json:"installments"`
	DefaultInstallments    int                       `json:"default_installments"`
}

// PreferencePaymentMethod contains information about the payment method in the preference.
type PreferencePaymentMethod struct {
	ID string `json:"id"`
}

// PreferencePaymentType contains information about the type of payment in the preference.
type PreferencePaymentType struct {
	ID string `json:"id"`
}

// PreferenceBackUrls contains preference back URLs.
type PreferenceBackUrls struct {
	Success string `json:"success"`
	Pending string `json:"pending"`
	Failure string `json:"failure"`
}

// PreferenceShipments contains preference shipping information.
type PreferenceShipments struct {
	Mode                  string                    `json:"mode"`
	LocalPickup           bool                      `json:"local_pickup"`
	Dimensions            string                    `json:"dimensions"`
	DefaultShippingMethod string                    `json:"default_shipping_method"`
	FreeMethods           []PreferenceFreeMethod    `json:"free_methods"`
	Cost                  float64                   `json:"cost"`
	FreeShipping          bool                      `json:"free_shipping"`
	ReceiverAddress       PreferenceReceiverAddress `json:"receiver_address"`
	ExpressShipment       bool                      `json:"express_shipment"`
}

// PreferenceFreeMethod contains information about free shipping methods.
type PreferenceFreeMethod struct {
	ID int64 `json:"id"`
}

// PreferenceReceiverAddress represents a sending address.
type PreferenceReceiverAddress struct {
	Address
	CountryName string `json:"country_name"`
	StateName   string `json:"state_name"`
	Floor       string `json:"floor"`
	Apartment   string `json:"apartment"`
	CityName    string `json:"city_name"`
}

// PreferenceDifferentialPricing contains information about the differential pricing configuration in the preference.
type PreferenceDifferentialPricing struct {
	ID int64 `json:"id"`
}

// PreferenceTax contains information about taxes in the preference.
type PreferenceTax struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// PreferenceTrack represents a trace to be executed during user interaction in the Checkout flow.
type PreferenceTrack struct {
	Type   string                        `json:"type"`
	Values PreferenceTrackValuesResponse `json:"values"`
}

// PreferenceTrackValuesRequest contains the values ​​of the tracks to be executed during user interaction in the Checkout flow.
type PreferenceTrackValuesResponse struct {
	ConversionID    string `json:"conversion_id"`
	ConversionLabel string `json:"conversion_label"`
	PixelID         string `json:"pixel_id"`
}

// Phone represents a telephone number.
type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

// Identification is a base type that represents identifications, such as customer identification.
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// Address represents an address.
type Address struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}
