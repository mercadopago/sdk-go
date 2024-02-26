package payment

import (
	"time"
)

// Response is the response from the Payments API.
type Response struct {
	DifferentialPricingID     string         `json:"differential_pricing_id"`
	MoneyReleaseSchema        string         `json:"money_release_schema"`
	OperationType             string         `json:"operation_type"`    // operation type (regular payment, subscription payment, etc)
	IssuerID                  string         `json:"issuer_id"`         // card brand identification
	PaymentMethodID           string         `json:"payment_method_id"` // payment method that was used by the payer
	PaymentTypeID             string         `json:"payment_type_id"`   // payment type (credit card, debit card, money, etc)
	Status                    string         `json:"status"`            // status (approved, pending, rejected, etc)
	StatusDetail              string         `json:"status_detail"`     // status detail complements status (can be for instance: status approved and status detail partially_refunded)
	CurrencyID                string         `json:"currency_id"`
	Description               string         `json:"description"` // payment description, can be useful for during/after payment experience (for payers and sellers)
	AuthorizationCode         string         `json:"authorization_code"`
	IntegratorID              string         `json:"integrator_id"`
	PlatformID                string         `json:"platform_id"`
	CorporationID             string         `json:"corporation_id"`
	NotificationURL           string         `json:"notification_url"` // every time that you create or update a payment, a request will be sent to this url, for more details see your app notification section
	CallbackURL               string         `json:"callback_url"`     // some payment methods have redirects after payment, here you can set to where payer will be redirected
	ProcessingMode            string         `json:"processing_mode"`  // for payments using cards this field says the processing mode
	MerchantAccountID         string         `json:"merchant_account_id"`
	MerchantNumber            string         `json:"merchant_number"`
	CouponCode                string         `json:"coupon_code"`
	ExternalReference         string         `json:"external_reference"`       // a payment identification used by the integrator, will have the value sent on payment creation
	PaymentMethodOptionID     string         `json:"payment_method_option_id"` // useful for not instantaneous payments, change a option to where payer should realize the payment. Example: the payment_method_id is X and should be paid on Y (can be a virtual or in-person place)
	PosID                     string         `json:"pos_id"`
	StoreID                   string         `json:"store_id"`
	DeductionSchema           string         `json:"deduction_schema"`
	CounterCurrency           string         `json:"counter_currency"`
	CallForAuthorizeID        string         `json:"call_for_authorize_id"`
	StatementDescriptor       string         `json:"statement_descriptor"`
	MoneyReleaseStatus        string         `json:"money_release_status"`
	Installments              int            `json:"installments"` // number of installments
	ID                        int64          `json:"id"`           // created payment identification
	SponsorID                 int64          `json:"sponsor_id"`
	CollectorID               int64          `json:"collector_id"`                // receiver identification
	TransactionAmount         float64        `json:"transaction_amount"`          // payment amount
	TransactionAmountRefunded float64        `json:"transaction_amount_refunded"` // payment refunded amount (will be > 0 if a refund occurs)
	CouponAmount              float64        `json:"coupon_amount"`
	TaxesAmount               float64        `json:"taxes_amount"`
	ShippingAmount            float64        `json:"shipping_amount"`
	NetAmount                 float64        `json:"net_amount"`
	LiveMode                  bool           `json:"live_mode"`
	Captured                  bool           `json:"captured"`
	BinaryMode                bool           `json:"binary_mode"`
	Metadata                  map[string]any `json:"metadata"`          // occasional data sent to the payment
	InternalMetadata          map[string]any `json:"internal_metadata"` // occasional internal data sent to the payment

	DateCreated        *time.Time                 `json:"date_created"`       // creation date
	DateApproved       *time.Time                 `json:"date_approved"`      // approved date (filled when the payment is set to approved)
	DateLastUpdated    *time.Time                 `json:"date_last_updated"`  // last updated date
	DateOfExpiration   *time.Time                 `json:"date_of_expiration"` // expiration date
	MoneyReleaseDate   *time.Time                 `json:"money_release_date"` // money release date
	Payer              PayerResponse              `json:"payer"`              // payer's data
	AdditionalInfo     AdditionalInfoResponse     `json:"additional_info"`    // additional info data
	Order              OrderResponse              `json:"order"`
	TransactionDetails TransactionDetailsResponse `json:"transaction_details"`
	Card               CardResponse               `json:"card"` // card data
	PointOfInteraction PointOfInteractionResponse `json:"point_of_interaction"`
	PaymentMethod      PaymentMethodResponse      `json:"payment_method"`
	ThreeDSInfo        ThreeDSInfoResponse        `json:"three_ds_info"` // useful for payments using 3DS authentication, see: https://www.mercadopago.com/developers/en/docs/checkout-api/how-tos/integrate-3ds
	FeeDetails         []FeeDetailResponse        `json:"fee_details"`
	Taxes              []TaxResponse              `json:"taxes"`
	Refunds            []RefundResponse           `json:"refunds"`
}

// PayerResponse represents the payer of the payment.
type PayerResponse struct {
	Type       string `json:"type"`       // good for differentiating customer & cards payments: https://www.mercadopago.com/developers/en/docs/checkout-api/customer-management
	ID         string `json:"id"`         // can be useful when the payments was created by customer & cards feature: https://www.mercadopago.com/developers/en/docs/checkout-api/customer-management
	Email      string `json:"email"`      // email can be useful to notify payer when a payment change occurs
	FirstName  string `json:"first_name"` // first name
	LastName   string `json:"last_name"`  // last name
	EntityType string `json:"entity_type"`

	Identification IdentificationResponse `json:"identification"` // identification
}

// IdentificationResponse represents payer's personal identification.
type IdentificationResponse struct {
	Type   string `json:"type"`   // type (can change depending on the country)
	Number string `json:"number"` // number (its format can change depending on the country)
}

// AdditionalInfoResponse receives non required data on payment operations.
type AdditionalInfoResponse struct {
	IPAddress string `json:"ip_address"`

	Payer     AdditionalInfoPayerResponse `json:"payer"` // payer's payment additional data
	Shipments ShipmentsResponse           `json:"shipments"`
	Items     []ItemResponse              `json:"items"`
}

// ItemResponse represents an item.
type ItemResponse struct {
	ID          string  `json:"id"`          // identification
	Title       string  `json:"title"`       // title
	Description string  `json:"description"` // more detailed text about the item
	PictureURL  string  `json:"picture_url"` // item picture, this picture will be used on during/after payment
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`   // quantity
	UnitPrice   float64 `json:"unit_price"` // it will not be used for calculate the final price, it's only for reference
}

// AdditionalInfoPayerResponse represents payer's additional information.
type AdditionalInfoPayerResponse struct {
	FirstName string `json:"first_name"` // first name
	LastName  string `json:"last_name"`  // last name

	RegistrationDate *time.Time      `json:"registration_date"` // registration date
	Phone            PhoneResponse   `json:"phone"`             // phone information
	Address          AddressResponse `json:"address"`           // address information
}

// PhoneResponse represents phone information.
type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"` // phone number
}

// AddressResponse represents address information.
type AddressResponse struct {
	ZipCode      string `json:"zip_code"`      // address zip code
	StreetName   string `json:"street_name"`   // street name
	StreetNumber string `json:"street_number"` // place's number
}

// ShipmentsResponse represents shipment information.
type ShipmentsResponse struct {
	ReceiverAddress ReceiverAddressResponse `json:"receiver_address"` // receiver address
}

// ReceiverAddressResponse represents the receiver's address within ShipmentsResponse.
type ReceiverAddressResponse struct {
	StateName string `json:"state_name"` // street name
	CityName  string `json:"city_name"`  // city name
	Floor     string `json:"floor"`      // floor (when it is relevant)
	Apartment string `json:"apartment"`  // apartment (when it is relevant)

	Address AddressResponse `json:"address"` // address
}

// OrderResponse represents order information.
type OrderResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// TransactionDetailsResponse represents transaction details.
type TransactionDetailsResponse struct {
	FinancialInstitution     string  `json:"financial_institution"`
	ExternalResourceURL      string  `json:"external_resource_url"` // can be useful for no card payments, depending on the payment method saves an useful payment experience
	PaymentMethodReferenceID string  `json:"payment_method_reference_id"`
	AcquirerReference        string  `json:"acquirer_reference"`
	TransactionID            string  `json:"transaction_id"` // BACEN identification for Pix payments (Brazil)
	NetReceivedAmount        float64 `json:"net_received_amount"`
	TotalPaidAmount          float64 `json:"total_paid_amount"`
	InstallmentAmount        float64 `json:"installment_amount"` // installment amount
	OverpaidAmount           float64 `json:"overpaid_amount"`    // overpaid amount is > 0 when payer paids more than transaction_amount (it is possible on some payment methods)
}

// CardResponse represents card information.
type CardResponse struct {
	ID              string `json:"id"`               // card identification
	LastFourDigits  string `json:"last_four_digits"` // last four digits
	FirstSixDigits  string `json:"first_six_digits"` // first six digits
	ExpirationYear  int    `json:"expiration_year"`  // expiration year
	ExpirationMonth int    `json:"expiration_month"` // expiration month

	DateCreated     *time.Time         `json:"date_created"`      // creation date
	DateLastUpdated *time.Time         `json:"date_last_updated"` // last update date
	Cardholder      CardholderResponse `json:"cardholder"`        // cardholder data
}

// CardholderResponse represents cardholder information.
type CardholderResponse struct {
	Name string `json:"name"` // name

	Identification IdentificationResponse `json:"identification"` // identification
}

// PointOfInteractionResponse represents point of interaction information.
type PointOfInteractionResponse struct {
	Type     string `json:"type"`
	SubType  string `json:"sub_type"`
	LinkedTo string `json:"linked_to"`

	ApplicationData ApplicationDataResponse `json:"application_data"`
	TransactionData TransactionDataResponse `json:"transaction_data"`
}

// ApplicationDataResponse represents application data within PointOfInteractionResponse.
type ApplicationDataResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// TransactionDataResponse represents transaction data within PointOfInteractionResponse.
type TransactionDataResponse struct {
	QRCode               string `json:"qr_code"`
	QRCodeBase64         string `json:"qr_code_base64"`
	TransactionID        string `json:"transaction_id"`
	TicketURL            string `json:"ticket_url"`
	SubscriptionID       string `json:"subscription_id"`
	BillingDate          string `json:"billing_date"`
	BankTransferID       int64  `json:"bank_transfer_id"`
	FinancialInstitution int64  `json:"financial_institution"`
	FirstTimeUse         bool   `json:"first_time_use"`

	BankInfo             BankInfoResponse             `json:"bank_info"`
	SubscriptionSequence SubscriptionSequenceResponse `json:"subscription_sequence"`
	InvoicePeriod        InvoicePeriodResponse        `json:"invoice_period"`
	PaymentReference     PaymentReferenceResponse     `json:"payment_reference"`
}

// BankInfoResponse represents bank information.
type BankInfoResponse struct {
	IsSameBankAccountOwner string `json:"is_same_bank_account_owner"`

	Payer     BankInfoPayerResponse     `json:"payer"`
	Collector BankInfoCollectorResponse `json:"collector"`
}

// SubscriptionSequenceResponse represents subscription sequence.
type SubscriptionSequenceResponse struct {
	Number int64 `json:"number"`
	Total  int64 `json:"total"`
}

// InvoicePeriodResponse represents invoice period.
type InvoicePeriodResponse struct {
	Type   string `json:"type"`
	Period int64  `json:"period"`
}

// PaymentReferenceResponse represents payment reference.
type PaymentReferenceResponse struct {
	ID string `json:"id"`
}

// BankInfoPayerResponse represents payer information within BankInfoResponse.
type BankInfoPayerResponse struct {
	Email     string `json:"email"`
	LongName  string `json:"long_name"`
	AccountID int64  `json:"account_id"`
}

// BankInfoCollectorResponse represents collector information within BankInfoResponse.
type BankInfoCollectorResponse struct {
	LongName  string `json:"long_name"`
	AccountID int64  `json:"account_id"`
}

// PaymentMethodResponse represents payment method information.
type PaymentMethodResponse struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	IssuerID string `json:"issuer_id"`

	Data DataResponse `json:"data"`
}

// DataResponse represents data within PaymentMethodResponse.
type DataResponse struct {
	Rules RulesResponse `json:"rules"`
}

// RulesResponse represents payment rules.
type RulesResponse struct {
	Fine      FeeResponse        `json:"fine"`
	Interest  FeeResponse        `json:"interest"`
	Discounts []DiscountResponse `json:"discounts"`
}

// DiscountResponse represents payment discount information.
type DiscountResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`

	LimitDate *time.Time `json:"limit_date"`
}

// FeeResponse represents payment fee information.
type FeeResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// ThreeDSInfoResponse represents 3DS (Three-Domain Secure) information.
type ThreeDSInfoResponse struct {
	ExternalResourceURL string `json:"external_resource_url"`
	Creq                string `json:"creq"`
}

// FeeDetailResponse represents payment fee detail information.
type FeeDetailResponse struct {
	Type     string  `json:"type"`
	FeePayer string  `json:"fee_payer"`
	Amount   float64 `json:"amount"`
}

// TaxResponse represents tax information.
type TaxResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// RefundResponse represents refund information.
type RefundResponse struct {
	Status               string  `json:"status"`
	RefundMode           string  `json:"refund_mode"`
	Reason               string  `json:"reason"`
	UniqueSequenceNumber string  `json:"unique_sequence_number"`
	ID                   int64   `json:"id"`
	PaymentID            int64   `json:"payment_id"`
	Amount               float64 `json:"amount"`
	AdjustmentAmount     float64 `json:"adjustment_amount"`

	DateCreated *time.Time     `json:"date_created"`
	Source      SourceResponse `json:"source"`
}

// SourceResponse represents source information.
type SourceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
