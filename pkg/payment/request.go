package payment

import (
	"time"
)

// Request represents a request for creating or updating a payment.
type Request struct {
	AdditionalInfo     *AdditionalInfoRequest     `json:"additional_info,omitempty"`
	MerchantServices   *MerchantServicesRequest   `json:"merchant_services,omitempty"`
	Order              *OrderRequest              `json:"order,omitempty"`
	Payer              *PayerRequest              `json:"payer,omitempty"`
	TransactionDetails *TransactionDetailsRequest `json:"transaction_details,omitempty"`
	PointOfInteraction *PointOfInteractionRequest `json:"point_of_interaction,omitempty"`
	PaymentMethod      *PaymentMethodRequest      `json:"payment_method,omitempty"`
	DateOfExpiration   *time.Time                 `json:"date_of_expiration,omitempty"`
	Taxes              []TaxRequest               `json:"taxes,omitempty"`

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
	CampaignID            int64          `json:"campaign_id,omitempty"`
	DifferentialPricingID int64          `json:"differential_pricing_id,omitempty"`
	SponsorID             int64          `json:"sponsor_id,omitempty"`
	BinaryMode            bool           `json:"binary_mode,omitempty"`
	Capture               bool           `json:"capture,omitempty"`
	Metadata              map[string]any `json:"metadata,omitempty"`
}

// AdditionalInfoRequest represents additional information request within Request.
type AdditionalInfoRequest struct {
	Payer     *AdditionalInfoPayerRequest   `json:"payer,omitempty"`
	Shipments *ShipmentsRequest             `json:"shipments,omitempty"`
	Barcode   *AdditionalInfoBarcodeRequest `json:"barcode,omitempty"`
	Items     []ItemRequest                 `json:"items,omitempty"`

	IPAddress string `json:"ip_address,omitempty"`
}

// AdditionalInfoPayerRequest represents payer information request within AdditionalInfoPayerRequest.
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

// AdditionalInfoPayerPhoneRequest represents phone request within AdditionalInfoPayerRequest.
type AdditionalInfoPayerPhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// AdditionalInfoPayerAddressRequest represents address request within AdditionalInfoPayerRequest.
type AdditionalInfoPayerAddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// ShipmentsRequest represents shipments request within AdditionalInfoRequest.
type ShipmentsRequest struct {
	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`

	LocalPickup     bool `json:"local_pickup,omitempty"`
	ExpressShipment bool `json:"express_shipment,omitempty"`
}

// ReceiverAddressRequest represents receiver address request within ShipmentsRequest.
type ReceiverAddressRequest struct {
	StateName    string `json:"state_name,omitempty"`
	CityName     string `json:"city_name,omitempty"`
	Floor        string `json:"floor,omitempty"`
	Apartment    string `json:"apartment,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// AdditionalInfoBarcodeRequest represents barcode request within AdditionalInfoRequest.
type AdditionalInfoBarcodeRequest struct {
	Type    string  `json:"type,omitempty"`
	Content string  `json:"content,omitempty"`
	Width   float64 `json:"width,omitempty"`
	Height  float64 `json:"height,omitempty"`
}

// ItemRequest represents an item request within AdditionalInfoRequest.
type ItemRequest struct {
	CategoryDescriptor *CategoryDescriptorRequest `json:"category_descriptor,omitempty"`
	EventDate          *time.Time                 `json:"event_date,omitempty"`

	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
	Quantity    int64   `json:"quantity,omitempty"`
	Warranty    bool    `json:"warranty,omitempty"`
}

// CategoryDescriptorRequest represents category descriptor request within ItemRequest.
type CategoryDescriptorRequest struct {
	Passenger *PassengerRequest `json:"passenger,omitempty"`
	Route     *RouteRequest     `json:"route,omitempty"`
}

// PassengerRequest represents passenger request within CategoryDescriptorRequest.
type PassengerRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"`

	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// IdentificationRequest represents identification request within PaymentPassengerRequest.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// RouteRequest represents route request within CategoryDescriptorRequest.
type RouteRequest struct {
	DepartureDateTime *time.Time `json:"departure_date_time,omitempty"`
	ArrivalDateTime   *time.Time `json:"arrival_date_time,omitempty"`

	Departure   string `json:"departure,omitempty"`
	Destination string `json:"destination,omitempty"`
	Company     string `json:"company,omitempty"`
}

// MerchantServicesRequest represents merchant services request within Request.
type MerchantServicesRequest struct {
	FraudScoring      bool `json:"fraud_scoring,omitempty"`
	FraudManualReview bool `json:"fraud_manual_review,omitempty"`
}

// OrderRequest represents order request within Request.
type OrderRequest struct {
	Type string `json:"type,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

// PayerRequest represents payer request within Request.
type PayerRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"`
	Address        *PayerAddressRequest   `json:"address,omitempty"`
	Phone          *PayerPhoneRequest     `json:"phone,omitempty"`

	Type       string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	Email      string `json:"email,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}

// PayerAddressRequest represents payer address request within PayerRequest.
type PayerAddressRequest struct {
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	FederalUnit  string `json:"federal_unit,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// PayerPhoneRequest represents payer phone request within PayerRequest.
type PayerPhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// TransactionDetailsRequest represents transaction details request within Request.
type TransactionDetailsRequest struct {
	FinancialInstitution string `json:"financial_institution,omitempty"`
}

// PointOfInteractionRequest represents point of interaction request within Request.
type PointOfInteractionRequest struct {
	TransactionData *TransactionDataRequest `json:"transaction_data,omitempty"`

	LinkedTo string `json:"linked_to,omitempty"`
	Type     string `json:"type,omitempty"`
}

type TransactionDataRequest struct {
	SubscriptionSequence *SubscriptionSequenceRequest `json:"subscription_sequence,omitempty"`
	InvoicePeriod        *InvoicePeriodRequest        `json:"invoice_period,omitempty"`
	PaymentReference     *PaymentReferenceRequest     `json:"payment_reference,omitempty"`

	SubscriptionID string `json:"subscription_id,omitempty"`
	BillingDate    string `json:"billing_date,omitempty"`
	FirstTimeUse   bool   `json:"first_time_use,omitempty"`
}

type SubscriptionSequenceRequest struct {
	Number int64 `json:"number,omitempty"`
	Total  int64 `json:"total,omitempty"`
}

type InvoicePeriodRequest struct {
	Type   string `json:"type,omitempty"`
	Period int64  `json:"period,omitempty"`
}

type PaymentReferenceRequest struct {
	ID string `json:"id,omitempty"`
}

// PaymentMethodRequest represents payment method request within Request.
type PaymentMethodRequest struct {
	Data *DataRequest `json:"data,omitempty"`

	Type string `json:"type,omitempty"`
}

// DataRequest represents payment data request within PaymentMethodRequest.
type DataRequest struct {
	Authentication *AuthenticationRequest `json:"authentication,omitempty"`
	Rules          *RulesRequest          `json:"rules,omitempty"`
}

// RulesRequest represents payment rules request within DataRequest.
type RulesRequest struct {
	Fine      *FeeRequest       `json:"fine,omitempty"`
	Interest  *FeeRequest       `json:"interest,omitempty"`
	Discounts []DiscountRequest `json:"discounts,omitempty"`
}

// AuthenticationRequest represents payment authentication request within DataRequest.
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

// FeeRequest represents fee request within RulesRequest.
type FeeRequest struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// DiscountRequest represents discount request within RulesRequest.
type DiscountRequest struct {
	LimitDate *time.Time `json:"limit_date,omitempty"`

	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// TaxRequest represents tax request within Request.
type TaxRequest struct {
	Type       string  `json:"type,omitempty"`
	Value      float64 `json:"value,omitempty"`
	Percentage bool    `json:"percentage,omitempty"`
}
