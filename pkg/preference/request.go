package preference

import (
	"time"
)

// Request represents the body payload sent to the MercadoPago Checkout Preferences API
// when creating or updating a preference. All fields are optional (omitempty) so that
// partial updates can be performed by setting only the fields that need to change.
//
// Key fields include Items (the products or services being sold), Payer (buyer details),
// PaymentMethods (allowed or excluded methods), BackURLs (redirect URLs after payment),
// and Shipments (delivery configuration).
type Request struct {
	// BackURLs defines the redirect URLs the buyer is sent to after a successful, pending, or failed payment.
	BackURLs *BackURLsRequest `json:"back_urls,omitempty"`

	// DifferentialPricing sets a differential pricing configuration to apply specific fees per buyer segment.
	DifferentialPricing *DifferentialPricingRequest `json:"differential_pricing,omitempty"`

	// Payer contains information about the buyer, such as name, email, identification, and address.
	Payer *PayerRequest `json:"payer,omitempty"`

	// PaymentMethods defines which payment methods are allowed, excluded, or set as default for this preference.
	PaymentMethods *PaymentMethodsRequest `json:"payment_methods,omitempty"`

	// Shipments configures shipping options, including delivery address, shipping cost, and available methods.
	Shipments *ShipmentsRequest `json:"shipments,omitempty"`

	// DateOfExpiration is the absolute expiration date for the preference, after which it can no longer be used.
	DateOfExpiration *time.Time `json:"date_of_expiration,omitempty"`

	// ExpirationDateFrom is the date from which the preference becomes active.
	ExpirationDateFrom *time.Time `json:"expiration_date_from,omitempty"`

	// ExpirationDateTo is the date until which the preference remains active.
	ExpirationDateTo *time.Time `json:"expiration_date_to,omitempty"`

	// Items is the list of products or services included in this checkout preference. At least one item is required.
	Items []ItemRequest `json:"items,omitempty"`

	// Taxes lists the applicable taxes for this preference.
	Taxes []TaxRequest `json:"taxes,omitempty"`

	// Tracks defines the conversion tracking configurations (e.g., Google Ads, Facebook Pixel) for the checkout flow.
	Tracks []TrackRequest `json:"tracks,omitempty"`

	// Amounts specifies custom transaction amounts for collector and payer, used in cross-currency scenarios.
	Amounts *AmountsRequest `json:"amounts,omitempty"`

	// CounterCurrency specifies the alternative currency for cross-currency payment scenarios.
	CounterCurrency *CounterCurrencyRequest `json:"counter_currency,omitempty"`

	// AdditionalInfo is free-text information attached to the preference for the integrator's reference.
	AdditionalInfo string `json:"additional_info,omitempty"`

	// AutoReturn controls automatic redirection behavior after payment. Accepted values are "approved", "all", or empty.
	AutoReturn string `json:"auto_return,omitempty"`

	// ExternalReference is an external identifier that helps the integrator reconcile the preference with their own system.
	ExternalReference string `json:"external_reference,omitempty"`

	// Marketplace identifies the marketplace that originated the preference, used in split-payment scenarios.
	Marketplace string `json:"marketplace,omitempty"`

	// OperationType indicates the type of operation (e.g., "regular_payment").
	OperationType string `json:"operation_type,omitempty"`

	// NotificationURL is the URL where MercadoPago sends webhook notifications about payment status changes.
	NotificationURL string `json:"notification_url,omitempty"`

	// Purpose defines the checkout purpose. Use "wallet_purchase" to restrict payment to MercadoPago wallet users only.
	Purpose string `json:"purpose,omitempty"`

	// StatementDescriptor is the description shown on the buyer's credit card statement.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// MarketplaceFee is the fee amount charged by the marketplace in split-payment scenarios.
	MarketplaceFee float64 `json:"marketplace_fee,omitempty"`

	// BinaryMode, when true, disables the "pending" payment status so payments are either approved or rejected immediately.
	BinaryMode bool `json:"binary_mode,omitempty"`

	// Expires indicates whether the preference has an expiration window defined by ExpirationDateFrom and ExpirationDateTo.
	Expires bool `json:"expires,omitempty"`

	// ProcessingModes lists the processing modes for the payment (e.g., "aggregator", "gateway").
	ProcessingModes []string `json:"processing_modes,omitempty"`

	// Metadata is a map of arbitrary key-value pairs that the integrator can attach to the preference for custom data.
	Metadata map[string]any `json:"metadata,omitempty"`

	// DeviceID is the unique identifier of the device from which the preference is created, used for fraud prevention.
	DeviceID string `json:"device_id,omitempty"`
}

// BackURLsRequest represents the callback URLs that the buyer will be redirected to
// after a checkout session ends. Each URL corresponds to a different payment outcome.
type BackURLsRequest struct {
	// Success is the URL the buyer is redirected to when the payment is approved.
	Success string `json:"success,omitempty"`

	// Pending is the URL the buyer is redirected to when the payment is still being processed.
	Pending string `json:"pending,omitempty"`

	// Failure is the URL the buyer is redirected to when the payment is rejected or fails.
	Failure string `json:"failure,omitempty"`
}

// DifferentialPricingRequest identifies a differential pricing configuration in the
// MercadoPago API, allowing different fees to be applied per buyer segment.
type DifferentialPricingRequest struct {
	// ID is the unique identifier of the differential pricing configuration.
	ID int `json:"id,omitempty"`
}

// PassengerRequest represents passenger information for travel-related items
// in the checkout preference, such as airline tickets or bus travel.
type PassengerRequest struct {
	// FirstName is the passenger's first name.
	FirstName string `json:"first_name,omitempty"`

	// LastName is the passenger's last name.
	LastName string `json:"last_name,omitempty"`

	// IdentificationType is the type of the passenger's identification document (e.g., "CPF", "DNI").
	IdentificationType string `json:"identification_type,omitempty"`

	// IdentificationNumber is the passenger's identification document number.
	IdentificationNumber string `json:"identification_number,omitempty"`

	// Identification contains the passenger's identity document details as a nested struct.
	Identification *IdentificationRequest `json:"identification,omitempty"`
}

// RouteRequest represents travel route details for travel-related items in the checkout preference.
type RouteRequest struct {
	// Departure is the origin location or airport code.
	Departure string `json:"departure,omitempty"`

	// Destination is the destination location or airport code.
	Destination string `json:"destination,omitempty"`

	// DepartureDateTime is the scheduled departure date and time.
	DepartureDateTime *time.Time `json:"departure_date_time,omitempty"`

	// ArrivalDateTime is the scheduled arrival date and time.
	ArrivalDateTime *time.Time `json:"arrival_date_time,omitempty"`

	// Company is the name of the transport company (e.g., airline or bus operator).
	Company string `json:"company,omitempty"`
}

// CategoryDescriptorRequest provides additional categorization metadata for an item,
// such as event dates and travel details. This is used for items that represent
// services like events, flights, or bus trips.
type CategoryDescriptorRequest struct {
	// EventDate is the date of the event associated with the item.
	EventDate *time.Time `json:"event_date,omitempty"`

	// Type describes the category type of the item (e.g., "event", "travel").
	Type string `json:"type,omitempty"`

	// Passenger contains passenger details for travel-related items.
	Passenger *PassengerRequest `json:"passenger,omitempty"`

	// Route contains route details for travel-related items.
	Route *RouteRequest `json:"route,omitempty"`
}

// ItemRequest represents a product or service being sold through the checkout preference.
// At least one item is required when creating a preference. Each item includes pricing,
// quantity, and descriptive information shown to the buyer during checkout.
type ItemRequest struct {
	// ID is the item identifier in the integrator's system.
	ID string `json:"id,omitempty"`

	// Title is the item name displayed to the buyer.
	Title string `json:"title,omitempty"`

	// Type is the item type (e.g., "product", "service").
	Type string `json:"type,omitempty"`

	// Description is a detailed description of the item.
	Description string `json:"description,omitempty"`

	// PictureURL is the URL of the item image displayed during checkout.
	PictureURL string `json:"picture_url,omitempty"`

	// CategoryID is the MercadoPago category identifier for the item.
	CategoryID string `json:"category_id,omitempty"`

	// CurrencyID is the ISO 4217 currency code for the item price (e.g., "BRL", "ARS", "MXN").
	CurrencyID string `json:"currency_id,omitempty"`

	// UnitPrice is the price per unit of the item.
	UnitPrice float64 `json:"unit_price,omitempty"`

	// Quantity is the number of units of this item being purchased.
	Quantity int `json:"quantity,omitempty"`

	// Warranty indicates whether the item includes a warranty.
	Warranty *bool `json:"warranty,omitempty"`

	// CategoryDescriptor provides additional categorization metadata, such as travel or event details.
	CategoryDescriptor CategoryDescriptorRequest `json:"category_descriptor,omitempty"`
}

// PayerRequest represents the buyer's information included in the checkout preference.
// Pre-filling payer data can improve conversion rates by reducing the amount of
// information the buyer must enter during checkout.
type PayerRequest struct {
	// DateCreated is the date when the payer account was created in the integrator's system.
	DateCreated *time.Time `json:"date_created,omitempty"`

	// Name is the buyer's first name.
	Name string `json:"name,omitempty"`

	// Surname is the buyer's last name.
	Surname string `json:"surname,omitempty"`

	// Email is the buyer's email address, used for notifications and identification.
	Email string `json:"email,omitempty"`

	// AuthenticationType indicates the authentication method used by the buyer (e.g., "gmail", "facebook", "native").
	AuthenticationType string `json:"authentication_type,omitempty"`

	// IsPrimeUser indicates whether the buyer is a premium or loyalty program member.
	IsPrimeUser bool `json:"is_prime_user,omitempty"`

	// IsFirstPurchaseOnline indicates whether this is the buyer's first online purchase.
	IsFirstPurchaseOnline bool `json:"is_first_purchase_online,omitempty"`

	// RegistrationDate is the date when the buyer registered in the integrator's platform.
	RegistrationDate *time.Time `json:"registration_date,omitempty"`

	// LastPurchase is the date of the buyer's most recent purchase.
	LastPurchase *time.Time `json:"last_purchase,omitempty"`

	// Identification contains the buyer's identity document details.
	Identification *IdentificationRequest `json:"identification,omitempty"`

	// Phone contains the buyer's phone number.
	Phone *PhoneRequest `json:"phone,omitempty"`

	// Address contains the buyer's address information.
	Address *AddressRequest `json:"address,omitempty"`
}

// PhoneRequest represents a phone number with area code, used to provide
// the buyer's contact information in the checkout preference.
type PhoneRequest struct {
	// AreaCode is the phone number area code.
	AreaCode string `json:"area_code,omitempty"`

	// Number is the phone number without the area code.
	Number string `json:"number,omitempty"`
}

// IdentificationRequest represents an identity document, used to provide the
// buyer's identification (e.g., CPF in Brazil, DNI in Argentina) in the checkout preference.
type IdentificationRequest struct {
	// Type is the identification document type (e.g., "CPF", "DNI", "CC").
	Type string `json:"type,omitempty"`

	// Number is the identification document number.
	Number string `json:"number,omitempty"`
}

// AddressRequest represents a postal address, used to specify the buyer's
// address in the checkout preference.
type AddressRequest struct {
	// ZipCode is the postal or ZIP code.
	ZipCode string `json:"zip_code,omitempty"`

	// StreetName is the street name.
	StreetName string `json:"street_name,omitempty"`

	// StreetNumber is the street number.
	StreetNumber string `json:"street_number,omitempty"`
}

// PaymentMethodsRequest configures which payment methods are available, excluded, or set
// as default for the checkout preference. This allows integrators to control the payment
// experience, such as limiting installments or hiding specific methods.
type PaymentMethodsRequest struct {
	// ExcludedPaymentMethods lists the payment methods that should not be offered during checkout.
	ExcludedPaymentMethods []ExcludedPaymentMethodRequest `json:"excluded_payment_methods,omitempty"`

	// ExcludedPaymentTypes lists the payment types (e.g., "credit_card", "ticket") to exclude from checkout.
	ExcludedPaymentTypes []ExcludedPaymentTypeRequest `json:"excluded_payment_types,omitempty"`

	// DefaultPaymentMethodID is the payment method pre-selected when the checkout loads.
	DefaultPaymentMethodID string `json:"default_payment_method_id,omitempty"`

	// Installments is the maximum number of installments allowed for this preference.
	Installments int `json:"installments,omitempty"`

	// DefaultInstallments is the number of installments pre-selected when the checkout loads.
	DefaultInstallments int `json:"default_installments,omitempty"`
}

// ExcludedPaymentMethodRequest identifies a specific payment method to be excluded
// from the checkout preference by its MercadoPago payment method ID.
type ExcludedPaymentMethodRequest struct {
	// ID is the MercadoPago payment method identifier to exclude.
	ID string `json:"id,omitempty"`
}

// ExcludedPaymentTypeRequest identifies a payment type to be excluded
// from the checkout preference (e.g., "credit_card", "debit_card", "ticket").
type ExcludedPaymentTypeRequest struct {
	// ID is the MercadoPago payment type identifier to exclude.
	ID string `json:"id,omitempty"`
}

// ShipmentsRequest configures shipping options for the checkout preference,
// including delivery address, available shipping methods, cost, and whether
// free shipping or local pickup is available.
type ShipmentsRequest struct {
	// ReceiverAddress is the address where the shipment will be delivered.
	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`

	// FreeMethods lists shipping methods that are offered for free to the buyer.
	FreeMethods []FreeMethodRequest `json:"free_methods,omitempty"`

	// Mode is the shipping mode (e.g., "custom", "me2", "not_specified").
	Mode string `json:"mode,omitempty"`

	// Dimensions specifies the package dimensions in the format "height x width x length, weight".
	Dimensions string `json:"dimensions,omitempty"`

	// DefaultShippingMethod is the shipping method pre-selected when the checkout loads.
	DefaultShippingMethod string `json:"default_shipping_method,omitempty"`

	// Cost is the fixed shipping cost to charge the buyer.
	Cost float64 `json:"cost,omitempty"`

	// LocalPickup indicates whether local pickup at a physical location is available.
	LocalPickup bool `json:"local_pickup,omitempty"`

	// FreeShipping indicates whether shipping is free for all available methods.
	FreeShipping bool `json:"free_shipping,omitempty"`

	// ExpressShipment indicates whether express (expedited) shipping is available.
	ExpressShipment bool `json:"express_shipment,omitempty"`
}

// FreeMethodRequest identifies a shipping method that should be offered for free
// to the buyer, referenced by its MercadoPago shipping method ID.
type FreeMethodRequest struct {
	// ID is the MercadoPago shipping method identifier.
	ID int `json:"id,omitempty"`
}

// ReceiverAddressRequest represents the destination address for a shipment in
// the checkout preference. It includes full postal information and delivery options.
type ReceiverAddressRequest struct {
	// ZipCode is the postal or ZIP code of the delivery address.
	ZipCode string `json:"zip_code,omitempty"`

	// StreetName is the street name of the delivery address.
	StreetName string `json:"street_name,omitempty"`

	// StreetNumber is the street number of the delivery address.
	StreetNumber string `json:"street_number,omitempty"`

	// CountryName is the country name of the delivery address.
	CountryName string `json:"country_name,omitempty"`

	// StateName is the state or province name of the delivery address.
	StateName string `json:"state_name,omitempty"`

	// Floor is the floor number, if applicable.
	Floor string `json:"floor,omitempty"`

	// Apartment is the apartment or unit number, if applicable.
	Apartment string `json:"apartment,omitempty"`

	// CityName is the city name of the delivery address.
	CityName string `json:"city_name,omitempty"`

	// LocalPickup indicates whether the buyer can pick up the order locally instead of having it shipped.
	LocalPickup bool `json:"local_pickup,omitempty"`

	// ExpressShipment indicates whether express delivery is available to this address.
	ExpressShipment bool `json:"express_shipment,omitempty"`
}

// TaxRequest represents a tax to be applied to the checkout preference.
// It allows integrators to include taxes that are added on top of the item prices.
type TaxRequest struct {
	// Type is the tax type (e.g., "IVA", "ISR").
	Type string `json:"type,omitempty"`

	// Value is the tax amount.
	Value float64 `json:"value,omitempty"`
}

// TrackRequest configures conversion tracking for analytics platforms during the
// checkout flow. This enables integrators to measure checkout conversion using
// services such as Google Ads or Facebook Pixel.
type TrackRequest struct {
	// Values contains the tracking identifiers for the configured analytics platform.
	Values *ValuesRequest `json:"values,omitempty"`

	// Type is the tracking type (e.g., "google_ad", "facebook_ad").
	Type string `json:"type,omitempty"`
}

// ValuesRequest contains the tracking identifiers used by analytics platforms
// to attribute conversions that occur during the checkout flow.
type ValuesRequest struct {
	// ConversionID is the Google Ads conversion ID.
	ConversionID string `json:"conversion_id,omitempty"`

	// ConversionLabel is the Google Ads conversion label.
	ConversionLabel string `json:"conversion_label,omitempty"`

	// PixelID is the Facebook Pixel identifier.
	PixelID string `json:"pixel_id,omitempty"`
}

// AmountsRequest specifies custom transaction amounts for the collector (seller)
// and payer (buyer) in cross-currency payment scenarios within a [Request].
type AmountsRequest struct {
	// Collector defines the transaction amount and currency for the seller.
	Collector UserAmountsRequest `json:"collector,omitempty"`

	// Payer defines the transaction amount and currency for the buyer.
	Payer UserAmountsRequest `json:"payer,omitempty"`
}

// UserAmountsRequest represents the transaction amount and currency for a specific
// party (collector or payer) in cross-currency scenarios within [AmountsRequest].
type UserAmountsRequest struct {
	// CurrencyID is the ISO 4217 currency code (e.g., "BRL", "USD").
	CurrencyID string `json:"currency_id,omitempty"`

	// Transaction is the transaction amount in the specified currency.
	Transaction float64 `json:"transaction,omitempty"`
}

// CounterCurrencyRequest specifies the alternative currency for cross-currency
// payment scenarios within a [Request]. It is used when the buyer pays in a
// different currency than the seller receives.
type CounterCurrencyRequest struct {
	// CurrencyID is the ISO 4217 currency code for the alternative currency.
	CurrencyID string `json:"currency_id,omitempty"`
}
