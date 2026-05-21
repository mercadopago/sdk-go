package preference

import (
	"time"
)

// Response represents the full checkout preference resource returned by the MercadoPago
// Checkout Preferences API. It is returned by [Client.Create], [Client.Get], and [Client.Update].
//
// The InitPoint and SandboxInitPoint fields contain the URLs that redirect buyers to
// the MercadoPago checkout flow in production and sandbox environments, respectively.
type Response struct {
	// Payer contains information about the buyer associated with this preference.
	Payer PayerResponse `json:"payer"`

	// PaymentMethods describes the payment methods configuration for this preference.
	PaymentMethods PaymentMethodsResponse `json:"payment_methods"`

	// BackURLs contains the redirect URLs for different payment outcomes.
	BackURLs BackURLsResponse `json:"back_urls"`

	// Shipments contains the shipping configuration for this preference.
	Shipments ShipmentsResponse `json:"shipments"`

	// DifferentialPricing contains the differential pricing configuration, if set.
	DifferentialPricing DifferentialPricingResponse `json:"differential_pricing"`

	// DateOfExpiration is the absolute expiration date for the preference.
	DateOfExpiration time.Time `json:"date_of_expiration"`

	// ExpirationDateFrom is the date from which the preference becomes active.
	ExpirationDateFrom time.Time `json:"expiration_date_from"`

	// ExpirationDateTo is the date until which the preference remains active.
	ExpirationDateTo time.Time `json:"expiration_date_to"`

	// LastUpdated is the date when the preference was last modified.
	LastUpdated time.Time `json:"last_updated"`

	// DateCreated is the date when the preference was created.
	DateCreated time.Time `json:"date_created"`

	// Taxes lists the taxes applied to this preference.
	Taxes []TaxResponse `json:"taxes"`

	// Tracks lists the conversion tracking configurations for the checkout flow.
	Tracks []TrackResponse `json:"tracks"`

	// Items lists the products or services included in this preference.
	Items []ItemResponse `json:"items"`

	// Amounts contains custom transaction amounts for cross-currency scenarios, if configured.
	Amounts *AmountsResponse `json:"amounts,omitempty"`

	// CounterCurrency contains the alternative currency configuration, if set.
	CounterCurrency *CounterCurrencyResponse `json:"counter_currency,omitempty"`

	// ID is the unique preference identifier assigned by MercadoPago.
	ID string `json:"id"`

	// ClientID is the application ID that created the preference.
	ClientID string `json:"client_id"`

	// NotificationURL is the webhook URL for payment status notifications.
	NotificationURL string `json:"notification_url"`

	// StatementDescriptor is the description shown on the buyer's credit card statement.
	StatementDescriptor string `json:"statement_descriptor"`

	// Marketplace identifies the marketplace that originated the preference.
	Marketplace string `json:"marketplace"`

	// ExternalReference is the integrator-provided external identifier for reconciliation.
	ExternalReference string `json:"external_reference"`

	// AdditionalInfo is free-text information attached to the preference.
	AdditionalInfo string `json:"additional_info"`

	// AutoReturn indicates the auto-return behavior after payment (e.g., "approved", "all").
	AutoReturn string `json:"auto_return"`

	// OperationType is the type of operation (e.g., "regular_payment").
	OperationType string `json:"operation_type"`

	// InitPoint is the production URL that redirects the buyer to the MercadoPago checkout flow.
	InitPoint string `json:"init_point"`

	// SandboxInitPoint is the sandbox URL for testing the checkout flow without real transactions.
	SandboxInitPoint string `json:"sandbox_init_point"`

	// SiteID is the MercadoPago site identifier (e.g., "MLB" for Brazil, "MLA" for Argentina).
	SiteID string `json:"site_id"`

	// MarketplaceFee is the marketplace fee amount in split-payment scenarios.
	MarketplaceFee float64 `json:"marketplace_fee"`

	// CollectorID is the MercadoPago user ID of the seller (collector) who receives the payment.
	CollectorID int64 `json:"collector_id"`

	// Expires indicates whether the preference has an active expiration window.
	Expires bool `json:"expires"`

	// BinaryMode indicates whether payments are limited to approved or rejected (no pending status).
	BinaryMode bool `json:"binary_mode"`

	// ProcessingModes lists the processing modes configured for this preference.
	ProcessingModes []string `json:"processing_modes"`

	// Metadata contains arbitrary key-value pairs attached to the preference by the integrator.
	Metadata map[string]any `json:"metadata"`
}

// ItemResponse represents a product or service item returned within a preference [Response].
type ItemResponse struct {
	// ID is the item identifier.
	ID string `json:"id"`

	// Title is the item name displayed to the buyer.
	Title string `json:"title"`

	// Description is a detailed description of the item.
	Description string `json:"description"`

	// CurrencyID is the ISO 4217 currency code for the item price.
	CurrencyID string `json:"currency_id"`

	// PictureURL is the URL of the item image.
	PictureURL string `json:"picture_url"`

	// CategoryID is the MercadoPago category identifier for the item.
	CategoryID string `json:"category_id"`

	// UnitPrice is the price per unit of the item.
	UnitPrice float64 `json:"unit_price"`

	// Quantity is the number of units purchased.
	Quantity int `json:"quantity"`

	// Warranty is a description of the item warranty.
	Warranty string `json:"warranty"`

	// CategoryDescriptor provides additional categorization metadata for the item.
	CategoryDescriptor CategoryDescriptorResponse `json:"category_descriptor"`
}

// CategoryDescriptorResponse contains additional categorization metadata for an item,
// such as passenger information for travel-related products.
type CategoryDescriptorResponse struct {
	// Passenger contains passenger details for travel-related items.
	Passenger PassengerResponse `json:"passenger"`
}

// PassengerResponse represents passenger information associated with a travel-related
// item in the preference response.
type PassengerResponse struct {
	// FirstName is the passenger's first name.
	FirstName string `json:"first_name"`

	// LastName is the passenger's last name.
	LastName string `json:"last_name"`

	// IdentificationType is the type of the passenger's identification document.
	IdentificationType string `json:"identification_type"`

	// IdentificationNumber is the passenger's identification document number.
	IdentificationNumber string `json:"identification_number"`
}

// PayerResponse represents the buyer's information returned within a preference [Response].
// It includes personal details, identification, contact information, and purchase history.
type PayerResponse struct {
	// DateCreated is the date when the payer account was created.
	DateCreated time.Time `json:"date_created"`

	// LastPurchase is the date of the buyer's last purchase.
	LastPurchase time.Time `json:"last_purchase"`

	// Name is the buyer's first name.
	Name string `json:"name"`

	// Surname is the buyer's last name.
	Surname string `json:"surname"`

	// Email is the buyer's email address.
	Email string `json:"email"`

	// AuthenticationType indicates the authentication method used by the buyer.
	AuthenticationType string `json:"authentication_type"`

	// IsPrimerUser indicates whether the buyer is a premium or loyalty program member.
	IsPrimerUser bool `json:"is_prime_user"`

	// IsFirstPurchaseOnLine indicates whether this is the buyer's first online purchase.
	IsFirstPurchaseOnLine bool `json:"is_first_purchase_online"`

	// RegistrationDate is the date when the buyer registered in the integrator's platform.
	RegistrationDate time.Time `json:"registration_date"`

	// LastPurchaseDate is the date of the buyer's most recent purchase.
	LastPurchaseDate time.Time `json:"last_purchase_date"`

	// Identification contains the buyer's identity document details.
	Identification IdentificationResponse `json:"identification"`

	// Phone contains the buyer's phone number.
	Phone PhoneResponse `json:"phone"`

	// Address contains the buyer's address information.
	Address AddressResponse `json:"address"`
}

// PaymentMethodsResponse describes the payment methods configuration returned within
// a preference [Response], including excluded methods, excluded types, and defaults.
type PaymentMethodsResponse struct {
	// ExcludedPaymentMethods lists the payment methods excluded from checkout.
	ExcludedPaymentMethods []ExcludedPaymentMethodResponse `json:"excluded_payment_methods"`

	// ExcludedPaymentTypes lists the payment types excluded from checkout.
	ExcludedPaymentTypes []ExcludedPaymentTypeResponse `json:"excluded_payment_types"`

	// DefaultPaymentMethodID is the pre-selected payment method for the checkout.
	DefaultPaymentMethodID string `json:"default_payment_method_id"`

	// Installments is the maximum number of installments allowed.
	Installments int `json:"installments"`

	// DefaultInstallments is the pre-selected number of installments.
	DefaultInstallments int `json:"default_installments"`
}

// ExcludedPaymentMethodResponse identifies a payment method excluded from a checkout preference.
type ExcludedPaymentMethodResponse struct {
	// ID is the MercadoPago payment method identifier.
	ID string `json:"id"`
}

// ExcludedPaymentTypeResponse identifies a payment type excluded from a checkout preference.
type ExcludedPaymentTypeResponse struct {
	// ID is the MercadoPago payment type identifier.
	ID string `json:"id"`
}

// BackURLsResponse represents the callback URLs that the buyer is redirected to
// after a checkout session ends, as returned in a preference [Response].
type BackURLsResponse struct {
	// Success is the URL for approved payments.
	Success string `json:"success"`

	// Pending is the URL for payments still being processed.
	Pending string `json:"pending"`

	// Failure is the URL for rejected or failed payments.
	Failure string `json:"failure"`
}

// ShipmentsResponse contains the shipping configuration returned within a preference [Response],
// including the delivery address, shipping methods, cost, and mode.
type ShipmentsResponse struct {
	// ReceiverAddress is the delivery address for the shipment.
	ReceiverAddress ReceiverAddressResponse `json:"receiver_address"`

	// FreeMethods lists the shipping methods offered for free to the buyer.
	FreeMethods []FreeMethodResponse `json:"free_methods"`

	// Mode is the shipping mode (e.g., "custom", "me2", "not_specified").
	Mode string `json:"mode"`

	// Dimensions specifies the package dimensions.
	Dimensions string `json:"dimensions"`

	// DefaultShippingMethod is the pre-selected shipping method.
	DefaultShippingMethod string `json:"default_shipping_method"`

	// Cost is the shipping cost.
	Cost float64 `json:"cost"`

	// LocalPickup indicates whether local pickup is available.
	LocalPickup bool `json:"local_pickup"`

	// FreeShipping indicates whether shipping is free.
	FreeShipping bool `json:"free_shipping"`

	// ExpressShipment indicates whether express shipping is available.
	ExpressShipment bool `json:"express_shipment"`
}

// FreeMethodResponse identifies a shipping method offered for free within the preference.
type FreeMethodResponse struct {
	// ID is the MercadoPago shipping method identifier.
	ID int `json:"id"`
}

// ReceiverAddressResponse represents the shipment destination address returned within
// a preference [Response]. It includes the base address and geographic details.
type ReceiverAddressResponse struct {
	// Address contains the base street address information.
	Address AddressResponse `json:"address"`

	// CountryName is the country name of the delivery address.
	CountryName string `json:"country_name"`

	// StateName is the state or province name of the delivery address.
	StateName string `json:"state_name"`

	// Floor is the floor number, if applicable.
	Floor string `json:"floor"`

	// Apartment is the apartment or unit number, if applicable.
	Apartment string `json:"apartment"`

	// CityName is the city name of the delivery address.
	CityName string `json:"city_name"`
}

// DifferentialPricingResponse represents the differential pricing configuration
// returned within a preference [Response].
type DifferentialPricingResponse struct {
	// ID is the unique identifier of the differential pricing configuration.
	ID int `json:"id"`
}

// TaxResponse represents a tax applied to the checkout preference,
// as returned in a preference [Response].
type TaxResponse struct {
	// Type is the tax type (e.g., "IVA", "ISR").
	Type string `json:"type"`

	// Value is the tax amount.
	Value float64 `json:"value"`
}

// TrackResponse represents a conversion tracking configuration returned within
// a preference [Response], used for analytics integration during the checkout flow.
type TrackResponse struct {
	// Values contains the tracking identifiers for the analytics platform.
	Values ValuesResponse `json:"values"`

	// Type is the tracking type (e.g., "google_ad", "facebook_ad").
	Type string `json:"type"`
}

// ValuesResponse contains the analytics tracking identifiers returned within a [TrackResponse].
type ValuesResponse struct {
	// ConversionID is the Google Ads conversion ID.
	ConversionID string `json:"conversion_id"`

	// ConversionLabel is the Google Ads conversion label.
	ConversionLabel string `json:"conversion_label"`

	// PixelID is the Facebook Pixel identifier.
	PixelID string `json:"pixel_id"`
}

// PhoneResponse represents a phone number returned within a preference [Response].
type PhoneResponse struct {
	// AreaCode is the phone number area code.
	AreaCode string `json:"area_code"`

	// Number is the phone number without the area code.
	Number string `json:"number"`
}

// IdentificationResponse represents an identity document returned within a preference [Response],
// such as the buyer's CPF (Brazil), DNI (Argentina), or similar national identification.
type IdentificationResponse struct {
	// Type is the identification document type (e.g., "CPF", "DNI", "CC").
	Type string `json:"type"`

	// Number is the identification document number.
	Number string `json:"number"`
}

// AddressResponse represents a postal address returned within a preference [Response].
type AddressResponse struct {
	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code"`

	// StreetName is the street name.
	StreetName string `json:"street_name"`

	// StreetNumber is the street number.
	StreetNumber string `json:"street_number"`
}

// AmountsResponse represents the custom transaction amounts for collector and payer
// in cross-currency scenarios, as returned in a preference [Response].
type AmountsResponse struct {
	// Collector contains the transaction amount and currency for the seller.
	Collector UserAmountsResponse `json:"collector,omitempty"`

	// Payer contains the transaction amount and currency for the buyer.
	Payer UserAmountsResponse `json:"payer,omitempty"`
}

// UserAmountsResponse represents the transaction amount and currency for a specific
// party (collector or payer) in a cross-currency preference [Response].
type UserAmountsResponse struct {
	// CurrencyID is the ISO 4217 currency code.
	CurrencyID string `json:"currency_id,omitempty"`

	// Transaction is the transaction amount in the specified currency.
	Transaction float64 `json:"transaction,omitempty"`
}

// CounterCurrencyResponse represents the alternative currency configuration
// returned in a preference [Response] for cross-currency payment scenarios.
type CounterCurrencyResponse struct {
	// CurrencyID is the ISO 4217 currency code for the alternative currency.
	CurrencyID string `json:"currency_id,omitempty"`
}
