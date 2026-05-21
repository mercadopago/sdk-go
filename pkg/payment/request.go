package payment

import (
	"time"
)

// Request represents the body sent to the MercadoPago Payments API when creating a new payment.
// It maps to the JSON payload described at
// https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/create-payment/post
//
// Fields are serialized as JSON with omitempty so only populated values are sent.
type Request struct {
	AdditionalInfo     *AdditionalInfoRequest     `json:"additional_info,omitempty"`
	MerchantServices   *MerchantServicesRequest   `json:"merchant_services,omitempty"`
	Order              *OrderRequest              `json:"order,omitempty"`
	Payer              *PayerRequest              `json:"payer,omitempty"`
	ForwardData        *ForwardDataRequest        `json:"forward_data,omitempty"`
	TransactionDetails *TransactionDetailsRequest `json:"transaction_details,omitempty"`
	PointOfInteraction *PointOfInteractionRequest `json:"point_of_interaction,omitempty"`
	PaymentMethod      *PaymentMethodRequest      `json:"payment_method,omitempty"`
	DateOfExpiration   *time.Time                 `json:"date_of_expiration,omitempty"`
	Taxes              []TaxRequest               `json:"taxes,omitempty"`
	Amounts            *AmountsRequest            `json:"amounts,omitempty"`
	CounterCurrency    *CounterCurrencyRequest    `json:"counter_currency,omitempty"`

	CallbackURL           string         `json:"callback_url,omitempty"`
	CouponCode            string         `json:"coupon_code,omitempty"`
	Description           string         `json:"description,omitempty"`
	ExternalReference     string         `json:"external_reference,omitempty"`
	IssuerID              string         `json:"issuer_id,omitempty"`
	MerchantAccountID     string         `json:"merchant_account_id,omitempty"`
	NotificationURL       string         `json:"notification_url,omitempty"`
	PaymentMethodID       string         `json:"payment_method_id,omitempty"`
	ProcessingMode        string         `json:"processing_mode,omitempty"`
	Token                 string         `json:"token,omitempty"`
	PaymentMethodOptionID string         `json:"payment_method_option_id,omitempty"`
	StatementDescriptor   string         `json:"statement_descriptor,omitempty"`
	ThreeDSecureMode      string         `json:"three_d_secure_mode,omitempty"`
	ApplicationFee        float64        `json:"application_fee,omitempty"`
	CouponAmount          float64        `json:"coupon_amount,omitempty"`
	NetAmount             float64        `json:"net_amount,omitempty"`
	TransactionAmount     float64        `json:"transaction_amount,omitempty"`
	Installments          int            `json:"installments,omitempty"`
	CampaignID            int            `json:"campaign_id,omitempty"`
	DifferentialPricingID int            `json:"differential_pricing_id,omitempty"`
	SponsorID             int            `json:"sponsor_id,omitempty"`
	BinaryMode            bool           `json:"binary_mode,omitempty"`
	Capture               bool           `json:"capture,omitempty"`
	Metadata              map[string]any `json:"metadata,omitempty"`
	DeviceID              string         `json:"device_id,omitempty"`
	BackURLs              []string       `json:"back_urls,omitempty"`
}

// AdditionalInfoRequest contains supplementary data about the payment such as payer details,
// shipping information, item details, and IP address. It is embedded in [Request].
type AdditionalInfoRequest struct {
	Payer     *AdditionalInfoPayerRequest `json:"payer,omitempty"`
	Shipments *ShipmentsRequest           `json:"shipments,omitempty"`
	Barcode   *BarcodeRequest             `json:"barcode,omitempty"`
	Items     []ItemRequest               `json:"items,omitempty"`

	IPAddress string `json:"ip_address,omitempty"`
}

// AdditionalInfoPayerRequest contains additional payer information such as name, phone, address,
// registration date, and purchase history. It is used within [AdditionalInfoRequest].
type AdditionalInfoPayerRequest struct {
	Phone            *AdditionalInfoPayerPhoneRequest   `json:"phone,omitempty"`
	Address          *AdditionalInfoPayerAddressRequest `json:"address,omitempty"`
	RegistrationDate *time.Time                         `json:"registration_date,omitempty"`
	LastPurchase     *time.Time                         `json:"last_purchase,omitempty"`

	FirstName             string `json:"first_name,omitempty"`
	LastName              string `json:"last_name,omitempty"`
	AuthenticationType    string `json:"authentication_type,omitempty"`
	IsPrimeUser           bool   `json:"is_prime_user,omitempty"`
	IsFirstPurchaseOnline bool   `json:"is_first_purchase_online,omitempty"`
}

// AdditionalInfoPayerPhoneRequest contains the payer's phone number split into area code and number.
// It is used within [AdditionalInfoPayerRequest].
type AdditionalInfoPayerPhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// AdditionalInfoPayerAddressRequest contains the payer's address details for fraud prevention
// and risk analysis. It is used within [AdditionalInfoPayerRequest].
type AdditionalInfoPayerAddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// ShipmentsRequest contains shipping details including the receiver address and delivery options.
// It is used within [AdditionalInfoRequest] to provide logistics context for the payment.
type ShipmentsRequest struct {
	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`

	LocalPickup     bool `json:"local_pickup,omitempty"`
	ExpressShipment bool `json:"express_shipment,omitempty"`
}

// ReceiverAddressRequest contains the destination address for the shipment.
// It is used within [ShipmentsRequest].
type ReceiverAddressRequest struct {
	StateName       string `json:"state_name,omitempty"`
	CityName        string `json:"city_name,omitempty"`
	Floor           string `json:"floor,omitempty"`
	Apartment       string `json:"apartment,omitempty"`
	ZipCode         string `json:"zip_code,omitempty"`
	StreetName      string `json:"street_name,omitempty"`
	StreetNumber    string `json:"street_number,omitempty"`
	LocalPickup     bool   `json:"local_pickup,omitempty"`
	ExpressShipment bool   `json:"express_shipment,omitempty"`
}

// BarcodeRequest contains barcode information for ticket-based or offline payment methods.
// It is used within [AdditionalInfoRequest].
type BarcodeRequest struct {
	Type    string  `json:"type,omitempty"`
	Content string  `json:"content,omitempty"`
	Width   float64 `json:"width,omitempty"`
	Height  float64 `json:"height,omitempty"`
}

// ItemRequest describes a purchased item or service associated with the payment.
// Multiple items can be included in [AdditionalInfoRequest] to improve fraud analysis.
type ItemRequest struct {
	EventDate          string  `json:"event_date,omitempty"`
	ID                 string  `json:"id,omitempty"`
	Type               string  `json:"type,omitempty"`
	Title              string  `json:"title,omitempty"`
	Description        string  `json:"description,omitempty"`
	PictureURL         string  `json:"picture_url,omitempty"`
	CategoryID         string  `json:"category_id,omitempty"`
	CurrencyIdentifier string  `json:"currency_identifier,omitempty"`
	UnitPrice          float64 `json:"unit_price,omitempty"`
	Quantity           int     `json:"quantity,omitempty"`
	Warranty           bool    `json:"warranty,omitempty"`

	CategoryDescriptor CategoryDescriptorRequest `json:"category_descriptor,omitempty"`
}

// CategoryDescriptorRequest provides category-specific metadata for an item, such as
// passenger and route information for travel-related payments. It is used within [ItemRequest].
type CategoryDescriptorRequest struct {
	EventDate *time.Time `json:"event_date,omitempty"`
	Type      string     `json:"type,omitempty"`

	Passenger *PassengerRequest `json:"passenger,omitempty"`
	Route     *RouteRequest     `json:"route,omitempty"`
}

// PassengerRequest contains passenger identification data for travel-related payments.
// It is used within [CategoryDescriptorRequest].
type PassengerRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"`

	FirstName            string `json:"first_name,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	IdentificationType   string `json:"identification_type,omitempty"`
	IdentificationNumber string `json:"identification_number,omitempty"`
}

// IdentificationRequest contains a personal identification document type and number.
// It is used within [PassengerRequest] and [PayerRequest] to identify individuals.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// RouteRequest describes a travel route with departure and destination details.
// It is used within [CategoryDescriptorRequest] for transportation-related items.
type RouteRequest struct {
	Departure         string `json:"departure,omitempty"`
	Destination       string `json:"destination,omitempty"`
	DepartureDateTime string `json:"departure_date_time,omitempty"`
	ArrivalDateTime   string `json:"arrival_date_time,omitempty"`
	Company           string `json:"company,omitempty"`
}

// MerchantServicesRequest configures optional merchant-level services such as
// fraud scoring and manual fraud review. It is used within [Request].
type MerchantServicesRequest struct {
	FraudScoring      bool `json:"fraud_scoring,omitempty"`
	FraudManualReview bool `json:"fraud_manual_review,omitempty"`
}

// OrderRequest associates the payment with an existing MercadoPago order.
// It is used within [Request].
type OrderRequest struct {
	Type string `json:"type,omitempty"`
	ID   int    `json:"id,omitempty"`
}

// PayerRequest identifies the person or entity making the payment including
// contact information and identification documents. It is used within [Request].
type PayerRequest struct {
	Type                  string `json:"type,omitempty"`
	ID                    string `json:"id,omitempty"`
	Email                 string `json:"email,omitempty"`
	FirstName             string `json:"first_name,omitempty"`
	LastName              string `json:"last_name,omitempty"`
	EntityType            string `json:"entity_type,omitempty"`
	AuthenticationType    string `json:"authentication_type,omitempty"`
	IsPrimeUser           bool   `json:"is_prime_user,omitempty"`
	IsFirstPurchaseOnline bool   `json:"is_first_purchase_online,omitempty"`
	RegistrationDate      string `json:"registration_date,omitempty"`
	LastPurchase          string `json:"last_purchase,omitempty"`

	Identification *IdentificationRequest `json:"identification,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`
	Address        *AddressRequest        `json:"address,omitempty"`
}

// ForwardDataRequest contains data forwarded to acquirers or processors under special conditions,
// such as payment facilitator (sub-merchant) details or network transaction references.
// It is used within [Request].
type ForwardDataRequest struct {
	SubMerchant            *SubMerchantRequest            `json:"sub_merchant,omitempty"`
	NetworkTransactionData *NetworkTransactionDataRequest `json:"network_transaction_data,omitempty"`
}

// NetworkTransactionDataRequest contains network-level transaction identifiers used
// for recurring or credential-on-file transactions. It is used within [ForwardDataRequest].
type NetworkTransactionDataRequest struct {
	NetworkTransactionID string `json:"network_transaction_id,omitempty"`
}

// SubMerchantRequest contains sub-merchant information for payment facilitator flows.
// Acquirers require this data when the collector operates as a marketplace or aggregator.
// It is used within [ForwardDataRequest].
type SubMerchantRequest struct {
	SubMerchantID     string `json:"sub_merchant_id,omitempty"`
	MCC               string `json:"mcc,omitempty"`
	Country           string `json:"country,omitempty"`
	ZIP               string `json:"zip,omitempty"`
	DocumentNumber    string `json:"document_number,omitempty"`
	City              string `json:"city,omitempty"`
	AddressStreet     string `json:"address_street,omitempty"`
	LegalName         string `json:"legal_name,omitempty"`
	RegionCodeISO     string `json:"region_code_iso,omitempty"`
	RegionCode        string `json:"region_code,omitempty"`
	DocumentType      string `json:"document_type,omitempty"`
	Phone             string `json:"phone,omitempty"`
	URL               string `json:"url,omitempty"`
	AddressDoorNumber int    `json:"address_door_number,omitempty"`
}

// AddressRequest contains the payer's billing address. It is used within [PayerRequest].
type AddressRequest struct {
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	FederalUnit  string `json:"federal_unit,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// PhoneRequest contains the payer's phone number split into area code and number.
// It is used within [PayerRequest].
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// TransactionDetailsRequest contains additional transaction-level details such as the
// financial institution involved. It is used within [Request].
type TransactionDetailsRequest struct {
	FinancialInstitution string `json:"financial_institution,omitempty"`
}

// PointOfInteractionRequest describes the context in which the payment interaction originates,
// such as a physical POS device or an online checkout. It is used within [Request].
type PointOfInteractionRequest struct {
	TransactionData *TransactionDataRequest `json:"transaction_data,omitempty"`

	LinkedTo string `json:"linked_to,omitempty"`
	Type     string `json:"type,omitempty"`
	SubType  string `json:"sub_type,omitempty"`
}

// TransactionDataRequest contains transaction-specific data for the point of interaction,
// including subscription and invoice details. It is used within [PointOfInteractionRequest].
type TransactionDataRequest struct {
	SubscriptionSequence *SubscriptionSequenceRequest `json:"subscription_sequence,omitempty"`
	InvoicePeriod        *InvoicePeriodRequest        `json:"invoice_period,omitempty"`
	PaymentReference     *PaymentReferenceRequest     `json:"payment_reference,omitempty"`

	SubscriptionID string `json:"subscription_id,omitempty"`
	BillingDate    string `json:"billing_date,omitempty"`
	FirstTimeUse   bool   `json:"first_time_use,omitempty"`
}

// SubscriptionSequenceRequest tracks the position of a payment within a recurring subscription.
// It is used within [TransactionDataRequest].
type SubscriptionSequenceRequest struct {
	Number int `json:"number,omitempty"`
	Total  int `json:"total,omitempty"`
}

// InvoicePeriodRequest defines the billing period type and length for subscription-based payments.
// It is used within [TransactionDataRequest].
type InvoicePeriodRequest struct {
	Type   string `json:"type,omitempty"`
	Period int    `json:"period,omitempty"`
}

// PaymentReferenceRequest contains a reference identifier linking to a related payment.
// It is used within [TransactionDataRequest].
type PaymentReferenceRequest struct {
	ID string `json:"id,omitempty"`
}

// PaymentMethodRequest specifies the payment method type and associated data such as
// authentication and rules. It is used within [Request].
type PaymentMethodRequest struct {
	Data *DataRequest `json:"data,omitempty"`

	Type string `json:"type,omitempty"`
}

// DataRequest contains authentication credentials and payment rules for a specific payment method.
// It is used within [PaymentMethodRequest].
type DataRequest struct {
	Authentication *AuthenticationRequest `json:"authentication,omitempty"`
	Rules          *RulesRequest          `json:"rules,omitempty"`
}

// RulesRequest defines fine, interest, and discount rules applied to a payment method.
// These are typically used for boleto-style payment methods. It is used within [DataRequest].
type RulesRequest struct {
	Fine      *FeeRequest       `json:"fine,omitempty"`
	Interest  *FeeRequest       `json:"interest,omitempty"`
	Discounts []DiscountRequest `json:"discounts,omitempty"`
}

// AuthenticationRequest provides 3DS (Three-Domain Secure) authentication data for card payments.
// It is used within [DataRequest] to supply cryptograms, transaction IDs, and ECI values.
type AuthenticationRequest struct {
	Type                 string `json:"type,omitempty"`
	Cryptogram           string `json:"cryptogram,omitempty"`
	ThreeDSServerTransID string `json:"three_ds_server_trans_id,omitempty"`
	ECI                  string `json:"eci,omitempty"`
	DSTransID            string `json:"ds_trans_id,omitempty"`
	ACSTransID           string `json:"acs_trans_id,omitempty"`
	ThreeDSVersion       string `json:"three_ds_version,omitempty"`
	AuthenticationStatus string `json:"authentication_status,omitempty"`
}

// FeeRequest defines a fine or interest fee to be applied to the payment.
// It is used within [RulesRequest].
type FeeRequest struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// DiscountRequest defines an early-payment discount with an optional deadline.
// It is used within [RulesRequest].
type DiscountRequest struct {
	LimitDate *time.Time `json:"limit_date,omitempty"`

	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// TaxRequest describes a tax applied to the payment, either as a fixed value or a percentage.
// It is used within [Request].
type TaxRequest struct {
	Type       string  `json:"type,omitempty"`
	Value      float64 `json:"value,omitempty"`
	Percentage bool    `json:"percentage,omitempty"`
}

// AmountsRequest specifies currency-specific transaction amounts for both the collector and
// the payer, enabling cross-border payment scenarios. It is used within [Request].
type AmountsRequest struct {
	Collector UserAmountsRequest `json:"collector,omitempty"`
	Payer     UserAmountsRequest `json:"payer,omitempty"`
}

// UserAmountsRequest defines a transaction amount in a specific currency for one party
// (collector or payer). It is used within [AmountsRequest].
type UserAmountsRequest struct {
	CurrencyID  string  `json:"currency_id,omitempty"`
	Transaction float64 `json:"transaction,omitempty"`
}

// CounterCurrencyRequest specifies the counter currency for cross-currency payments.
// It is used within [Request].
type CounterCurrencyRequest struct {
	CurrencyID string `json:"currency_id,omitempty"`
}
