package payment

import (
	"time"
)

// Request represents a request for creating or updating a payment.
type Request struct {
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
	Installments          int            `json:"installments,omitempty"`
	CampaignID            int64          `json:"campaign_id,omitempty"`
	DifferentialPricingID int64          `json:"differential_pricing_id,omitempty"`
	SponsorID             int64          `json:"sponsor_id,omitempty"`
	BinaryMode            bool           `json:"binary_mode,omitempty"`
	Capture               bool           `json:"capture,omitempty"`
	ApplicationFee        float64        `json:"application_fee,omitempty"`
	CouponAmount          float64        `json:"coupon_amount,omitempty"`
	NetAmount             float64        `json:"net_amount,omitempty"`
	TransactionAmount     float64        `json:"transaction_amount,omitempty"`
	Metadata              map[string]any `json:"metadata,omitempty"`

	DateOfExpiration   *time.Time                 `json:"date_of_expiration,omitempty"`
	AdditionalInfo     *AdditionalInfoRequest     `json:"additional_info,omitempty"`
	MerchantServices   *MerchantServicesRequest   `json:"merchant_services,omitempty"`
	Order              *OrderRequest              `json:"order,omitempty"`
	Payer              *PayerRequest              `json:"payer,omitempty"`
	TransactionDetails *TransactionDetailsRequest `json:"transaction_details,omitempty"`
	PointOfInteraction *PointOfInteractionRequest `json:"point_of_interaction,omitempty"`
	PaymentMethod      *PaymentMethodRequest      `json:"payment_method,omitempty"`
	Taxes              []TaxRequest               `json:"taxes,omitempty"`
}

// AdditionalInfoRequest represents additional information request within Request.
type AdditionalInfoRequest struct {
	IPAddress string `json:"ip_address,omitempty"`

	Payer     *AdditionalInfoPayerRequest   `json:"payer,omitempty"`
	Shipments *ShipmentsRequest             `json:"shipments,omitempty"`
	Barcode   *AdditionalInfoBarcodeRequest `json:"barcode,omitempty"`
	Items     []ItemRequest                 `json:"items,omitempty"`
}

// AdditionalInfoPayerRequest represents payer information request within AdditionalInfoPayerRequest.
type AdditionalInfoPayerRequest struct {
	FirstName             string `json:"first_name,omitempty"`
	LastName              string `json:"last_name,omitempty"`
	AuthenticationType    string `json:"authentication_type,omitempty"`
	IsPrimeUser           bool   `json:"is_prime_user,omitempty"`
	IsFirstPurchaseOnline bool   `json:"is_first_purchase_online,omitempty"`

	RegistrationDate *time.Time      `json:"registration_date,omitempty"`
	LastPurchase     *time.Time      `json:"last_purchase,omitempty"`
	Phone            *PhoneRequest   `json:"phone,omitempty"`
	Address          *AddressRequest `json:"address,omitempty"`
}

// PhoneRequest represents phone request within AdditionalInfoPayerRequest.
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// AddressRequest represents address request within AdditionalInfoPayerRequest.
type AddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// ShipmentsRequest represents shipments request within AdditionalInfoRequest.
type ShipmentsRequest struct {
	LocalPickup     bool `json:"local_pickup,omitempty"`
	ExpressShipment bool `json:"express_shipment,omitempty"`

	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`
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
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
	Warranty    bool    `json:"warranty,omitempty"`

	EventDate          *time.Time                 `json:"event_date,omitempty"`
	CategoryDescriptor *CategoryDescriptorRequest `json:"category_descriptor,omitempty"`
}

// CategoryDescriptorRequest represents category descriptor request within ItemRequest.
type CategoryDescriptorRequest struct {
	Passenger *PassengerRequest `json:"passenger,omitempty"`
	Route     *RouteRequest     `json:"route,omitempty"`
}

// PassengerRequest represents passenger request within CategoryDescriptorRequest.
type PassengerRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`

	Identification *IdentificationRequest `json:"identification,omitempty"`
}

// IdentificationRequest represents identification request within PaymentPassengerRequest.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// RouteRequest represents route request within CategoryDescriptorRequest.
type RouteRequest struct {
	Departure   string `json:"departure,omitempty"`
	Destination string `json:"destination,omitempty"`
	Company     string `json:"company,omitempty"`

	DepartureDateTime *time.Time `json:"departure_date_time,omitempty"`
	ArrivalDateTime   *time.Time `json:"arrival_date_time,omitempty"`
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
	Type       string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	Email      string `json:"email,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	EntityType string `json:"entity_type,omitempty"`

	Identification *IdentificationRequest `json:"identification,omitempty"`
	Address        *PayerAddressRequest   `json:"address,omitempty"`
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

// TransactionDetailsRequest represents transaction details request within Request.
type TransactionDetailsRequest struct {
	FinancialInstitution string `json:"financial_institution,omitempty"`
}

// PointOfInteractionRequest represents point of interaction request within Request.
type PointOfInteractionRequest struct {
	LinkedTo string `json:"linked_to,omitempty"`
	Type     string `json:"type,omitempty"`
}

// PaymentMethodRequest represents payment method request within Request.
type PaymentMethodRequest struct {
	Data *DataRequest `json:"data,omitempty"`
}

// DataRequest represents payment data request within PaymentMethodRequest.
type DataRequest struct {
	Rules *RulesRequest `json:"rules,omitempty"`
}

// RulesRequest represents payment rules request within DataRequest.
type RulesRequest struct {
	Fine      *FeeRequest       `json:"fine,omitempty"`
	Interest  *FeeRequest       `json:"interest,omitempty"`
	Discounts []DiscountRequest `json:"discounts,omitempty"`
}

// FeeRequest represents fee request within RulesRequest.
type FeeRequest struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// DiscountRequest represents discount request within RulesRequest.
type DiscountRequest struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`

	LimitDate *time.Time `json:"limit_date,omitempty"`
}

// TaxRequest represents tax request within Request.
type TaxRequest struct {
	Type       string  `json:"type,omitempty"`
	Value      float64 `json:"value,omitempty"`
	Percentage bool    `json:"percentage,omitempty"`
}
