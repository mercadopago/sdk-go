package payment

import (
	"time"
)

// Request contains every field accepted by Payments API.
type Request struct {
	AdditionalInfo     *AdditionalInfoRequest     `json:"additional_info,omitempty"` // additional data of the payment, complete this can help to improve payment approval rate
	MerchantServices   *MerchantServicesRequest   `json:"merchant_services,omitempty"`
	Order              *OrderRequest              `json:"order,omitempty"`
	Payer              *PayerRequest              `json:"payer,omitempty"` // payer's payment data
	TransactionDetails *TransactionDetailsRequest `json:"transaction_details,omitempty"`
	PointOfInteraction *PointOfInteractionRequest `json:"point_of_interaction,omitempty"`
	PaymentMethod      *PaymentMethodRequest      `json:"payment_method,omitempty"`
	DateOfExpiration   *time.Time                 `json:"date_of_expiration,omitempty"` // expiration date of the payment
	Taxes              []TaxRequest               `json:"taxes,omitempty"`

	CallbackURL           string         `json:"callback_url,omitempty"` // some payment methods have redirects after payment, here you can set to where payer will be redirected
	CouponCode            string         `json:"coupon_code,omitempty"`
	Description           string         `json:"description,omitempty"`        // payment description, can be useful for during/after payment experience (for payers and sellers)
	ExternalReference     string         `json:"external_reference,omitempty"` // a payment identification sent by the integrator, can be anything that you use as an identifier
	IssuerID              string         `json:"issuer_id,omitempty"`          // card brand identification
	MerchantAccountID     string         `json:"merchant_account_id,omitempty"`
	NotificationURL       string         `json:"notification_url,omitempty"`         // every time that you create or update a payment, a request will be sent to this url, for more details see your app notification section
	PaymentMethodID       string         `json:"payment_method_id,omitempty"`        // payment method that would be used by the payer (can change depending on the country)
	ProcessingMode        string         `json:"processing_mode,omitempty"`          // for payments using cards this field says the processing mode, if you don't know the different modes, just don't send it
	Token                 string         `json:"token,omitempty"`                    // for payments using cards this field receives the generated card token
	PaymentMethodOptionID string         `json:"payment_method_option_id,omitempty"` // useful for not instantaneous payments, change a option to where payer should realize the payment. Example: the payment_method_id is X and should be paid on Y (can be a virtual or in-person place)
	StatementDescriptor   string         `json:"statement_descriptor,omitempty"`
	ThreeDSecureMode      string         `json:"three_d_secure_mode,omitempty"` // useful for payments using 3DS authentication, see: https://www.mercadopago.com/developers/en/docs/checkout-api/how-tos/integrate-3ds
	ApplicationFee        float64        `json:"application_fee,omitempty"`
	CouponAmount          float64        `json:"coupon_amount,omitempty"`
	NetAmount             float64        `json:"net_amount,omitempty"`
	TransactionAmount     float64        `json:"transaction_amount,omitempty"` // amount to be paid
	Installments          int            `json:"installments,omitempty"`       // number of installments
	CampaignID            int            `json:"campaign_id,omitempty"`
	DifferentialPricingID int            `json:"differential_pricing_id,omitempty"`
	SponsorID             int            `json:"sponsor_id,omitempty"`
	BinaryMode            bool           `json:"binary_mode,omitempty"`
	Capture               bool           `json:"capture,omitempty"`  // useful for reserve values feature: https://www.mercadopago.com/developers/en/docs/checkout-api/payment-management/make-value-reserve
	Metadata              map[string]any `json:"metadata,omitempty"` // occasional data sent to the payment
}

// AdditionalInfoRequest allows sent non required data on payment operations.
// Complete this can help to improve payment approval rate.
type AdditionalInfoRequest struct {
	Payer     *AdditionalInfoPayerRequest   `json:"payer,omitempty"` // payer's payment additional data
	Shipments *ShipmentsRequest             `json:"shipments,omitempty"`
	Barcode   *AdditionalInfoBarcodeRequest `json:"barcode,omitempty"`
	Items     []ItemRequest                 `json:"items,omitempty"`

	IPAddress string `json:"ip_address,omitempty"`
}

// AdditionalInfoPayerRequest is the payer's payment additional data.
type AdditionalInfoPayerRequest struct {
	Phone            *AdditionalInfoPayerPhoneRequest   `json:"phone,omitempty"`             // phone information
	Address          *AdditionalInfoPayerAddressRequest `json:"address,omitempty"`           // address information
	RegistrationDate *time.Time                         `json:"registration_date,omitempty"` // registration date
	LastPurchase     *time.Time                         `json:"last_purchase,omitempty"`

	FirstName             string `json:"first_name,omitempty"` // first name
	LastName              string `json:"last_name,omitempty"`  // last name
	AuthenticationType    string `json:"authentication_type,omitempty"`
	IsPrimeUser           bool   `json:"is_prime_user,omitempty"`
	IsFirstPurchaseOnline bool   `json:"is_first_purchase_online,omitempty"`
}

// AdditionalInfoPayerPhoneRequest is the payer's phone on payment additional data.
type AdditionalInfoPayerPhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"` // phone number
}

// AdditionalInfoPayerAddressRequest is the payer's address on payment additional data.
type AdditionalInfoPayerAddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`      // address zip code
	StreetName   string `json:"street_name,omitempty"`   // street name
	StreetNumber string `json:"street_number,omitempty"` // place's number
}

// ShipmentsRequest represents shipments request within AdditionalInfoRequest.
type ShipmentsRequest struct {
	ReceiverAddress *ReceiverAddressRequest `json:"receiver_address,omitempty"`

	LocalPickup     bool `json:"local_pickup,omitempty"`
	ExpressShipment bool `json:"express_shipment,omitempty"`
}

// ReceiverAddressRequest represents receiver address request within ShipmentsRequest.
type ReceiverAddressRequest struct {
	StateName    string `json:"state_name,omitempty"`    // state name
	CityName     string `json:"city_name,omitempty"`     // city name
	Floor        string `json:"floor,omitempty"`         // floor (when it is relevant)
	Apartment    string `json:"apartment,omitempty"`     // apartment (when it is relevant)
	ZipCode      string `json:"zip_code,omitempty"`      // zip code
	StreetName   string `json:"street_name,omitempty"`   // street name
	StreetNumber string `json:"street_number,omitempty"` // place's number
}

// AdditionalInfoBarcodeRequest represents barcode request within AdditionalInfoRequest.
type AdditionalInfoBarcodeRequest struct {
	Type    string  `json:"type,omitempty"`    // type
	Content string  `json:"content,omitempty"` // value
	Width   float64 `json:"width,omitempty"`   // width
	Height  float64 `json:"height,omitempty"`  // height
}

// ItemRequest represents an item request within AdditionalInfoRequest.
type ItemRequest struct {
	CategoryDescriptor *CategoryDescriptorRequest `json:"category_descriptor,omitempty"`
	EventDate          *time.Time                 `json:"event_date,omitempty"`

	ID          string  `json:"id,omitempty"`          // identification
	Title       string  `json:"title,omitempty"`       // title
	Description string  `json:"description,omitempty"` // more detailed text about the item
	PictureURL  string  `json:"picture_url,omitempty"` // the url sent here should has a saved picture, this picture will be used on during/after payment
	CategoryID  string  `json:"category_id,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"` // it will not be used for calculate the final price, it's only for reference
	Quantity    int     `json:"quantity,omitempty"`   // quantity
	Warranty    bool    `json:"warranty,omitempty"`
}

// CategoryDescriptorRequest represents category descriptor request within ItemRequest.
type CategoryDescriptorRequest struct {
	Passenger *PassengerRequest `json:"passenger,omitempty"`
	Route     *RouteRequest     `json:"route,omitempty"`
}

// PassengerRequest represents passenger request within CategoryDescriptorRequest.
type PassengerRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"` // identification

	FirstName string `json:"first_name,omitempty"` // first name
	LastName  string `json:"last_name,omitempty"`  // last name
}

// IdentificationRequest represents identification request within PaymentPassengerRequest.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`   // type (can change depending on the country)
	Number string `json:"number,omitempty"` // number (its format can change depending on the country)
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
	ID   int    `json:"id,omitempty"`
}

// PayerRequest represents payer request within Request.
type PayerRequest struct {
	Identification *IdentificationRequest `json:"identification,omitempty"` // identification
	Address        *PayerAddressRequest   `json:"address,omitempty"`        // address
	Phone          *PayerPhoneRequest     `json:"phone,omitempty"`          // phone

	Type       string `json:"type,omitempty"`       // it is useful for customer & cards feature: https://www.mercadopago.com/developers/en/docs/checkout-api/customer-management
	ID         string `json:"id,omitempty"`         // it is useful for customer & cards feature (receives customer id): https://www.mercadopago.com/developers/en/docs/checkout-api/customer-management
	Email      string `json:"email,omitempty"`      // it is required for payments that don't have an assigned customer
	FirstName  string `json:"first_name,omitempty"` // first name
	LastName   string `json:"last_name,omitempty"`  // last name
	EntityType string `json:"entity_type,omitempty"`
}

// PayerAddressRequest represents payer address request within PayerRequest.
type PayerAddressRequest struct {
	Neighborhood string `json:"neighborhood,omitempty"`  // neighborhood
	City         string `json:"city,omitempty"`          // city
	FederalUnit  string `json:"federal_unit,omitempty"`  // federal unit (normally it is an acronym)
	ZipCode      string `json:"zip_code,omitempty"`      // zip code (each country has an own system)
	StreetName   string `json:"street_name,omitempty"`   // street name
	StreetNumber string `json:"street_number,omitempty"` // place's number
}

// PayerPhoneRequest represents payer phone request within PayerRequest.
type PayerPhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"` // number
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
	SubscriptionSequence *SubscriptionSequenceRequest `json:"subscription_sequence,omitempty"` // subscription sequence is useful for subscriptions feature: https://www.mercadopago.com/developers/en/docs/subscriptions/landing
	InvoicePeriod        *InvoicePeriodRequest        `json:"invoice_period,omitempty"`
	PaymentReference     *PaymentReferenceRequest     `json:"payment_reference,omitempty"`

	SubscriptionID string `json:"subscription_id,omitempty"` // subscription id is useful for subscriptions feature: https://www.mercadopago.com/developers/en/docs/subscriptions/landing
	BillingDate    string `json:"billing_date,omitempty"`    // billing date is useful for subscriptions feature: https://www.mercadopago.com/developers/en/docs/subscriptions/landing
	FirstTimeUse   bool   `json:"first_time_use,omitempty"`
}

type SubscriptionSequenceRequest struct {
	Number int `json:"number,omitempty"`
	Total  int `json:"total,omitempty"`
}

type InvoicePeriodRequest struct {
	Type   string `json:"type,omitempty"`
	Period int    `json:"period,omitempty"`
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
