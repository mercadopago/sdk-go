package payment

import (
	"time"
)

// Response is the response from the Payments API.
type Response struct {
	Payer              PayerResponse              `json:"payer"`
	AdditionalInfo     AdditionalInfoResponse     `json:"additional_info"`
	Order              OrderResponse              `json:"order"`
	TransactionDetails TransactionDetailsResponse `json:"transaction_details"`
	Card               CardResponse               `json:"card"`
	PointOfInteraction PointOfInteractionResponse `json:"point_of_interaction"`
	PaymentMethod      PaymentMethodResponse      `json:"payment_method"`
	ThreeDSInfo        ThreeDSInfoResponse        `json:"three_ds_info"`
	DateCreated        time.Time                  `json:"date_created"`
	DateApproved       time.Time                  `json:"date_approved"`
	DateLastUpdated    time.Time                  `json:"date_last_updated"`
	DateOfExpiration   time.Time                  `json:"date_of_expiration"`
	MoneyReleaseDate   time.Time                  `json:"money_release_date"`
	FeeDetails         []FeeDetailResponse        `json:"fee_details"`
	Taxes              []TaxResponse              `json:"taxes"`
	Refunds            []RefundResponse           `json:"refunds"`

	DifferentialPricingID     string         `json:"differential_pricing_id"`
	MoneyReleaseSchema        string         `json:"money_release_schema"`
	OperationType             string         `json:"operation_type"`
	IssuerID                  string         `json:"issuer_id"`
	PaymentMethodID           string         `json:"payment_method_id"`
	PaymentTypeID             string         `json:"payment_type_id"`
	Status                    string         `json:"status"`
	StatusDetail              string         `json:"status_detail"`
	CurrencyID                string         `json:"currency_id"`
	Description               string         `json:"description"`
	AuthorizationCode         string         `json:"authorization_code"`
	IntegratorID              string         `json:"integrator_id"`
	PlatformID                string         `json:"platform_id"`
	CorporationID             string         `json:"corporation_id"`
	NotificationURL           string         `json:"notification_url"`
	CallbackURL               string         `json:"callback_url"`
	ProcessingMode            string         `json:"processing_mode"`
	MerchantAccountID         string         `json:"merchant_account_id"`
	MerchantNumber            string         `json:"merchant_number"`
	CouponCode                string         `json:"coupon_code"`
	ExternalReference         string         `json:"external_reference"`
	PaymentMethodOptionID     string         `json:"payment_method_option_id"`
	PosID                     string         `json:"pos_id"`
	StoreID                   string         `json:"store_id"`
	DeductionSchema           string         `json:"deduction_schema"`
	CounterCurrency           string         `json:"counter_currency"`
	CallForAuthorizeID        string         `json:"call_for_authorize_id"`
	StatementDescriptor       string         `json:"statement_descriptor"`
	MoneyReleaseStatus        string         `json:"money_release_status"`
	Installments              int            `json:"installments"`
	ID                        int64          `json:"id"`
	SponsorID                 int64          `json:"sponsor_id"`
	CollectorID               int64          `json:"collector_id"`
	TransactionAmount         float64        `json:"transaction_amount"`
	TransactionAmountRefunded float64        `json:"transaction_amount_refunded"`
	CouponAmount              float64        `json:"coupon_amount"`
	TaxesAmount               float64        `json:"taxes_amount"`
	ShippingAmount            float64        `json:"shipping_amount"`
	NetAmount                 float64        `json:"net_amount"`
	LiveMode                  bool           `json:"live_mode"`
	Captured                  bool           `json:"captured"`
	BinaryMode                bool           `json:"binary_mode"`
	Metadata                  map[string]any `json:"metadata"`
	InternalMetadata          map[string]any `json:"internal_metadata"`
}

// PayerResponse represents the payer of the payment.
type PayerResponse struct {
	Identification IdentificationResponse `json:"identification"`

	Type       string `json:"type"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	EntityType string `json:"entity_type"`
}

// IdentificationResponse represents payer's personal identification.
type IdentificationResponse struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// AdditionalInfoResponse represents additional information about a payment.
type AdditionalInfoResponse struct {
	Payer     AdditionalInfoPayerResponse `json:"payer"`
	Shipments ShipmentsResponse           `json:"shipments"`
	Items     []ItemResponse              `json:"items"`

	IPAddress string `json:"ip_address"`
}

// ItemResponse represents an item.
type ItemResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PictureURL  string  `json:"picture_url"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

// AdditionalInfoPayerResponse represents payer's additional information.
type AdditionalInfoPayerResponse struct {
	Phone            PhoneResponse   `json:"phone"`
	Address          AddressResponse `json:"address"`
	RegistrationDate time.Time       `json:"registration_date"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// PhoneResponse represents phone information.
type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

// AddressResponse represents address information.
type AddressResponse struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}

// ShipmentsResponse represents shipment information.
type ShipmentsResponse struct {
	ReceiverAddress ReceiverAddressResponse `json:"receiver_address"`
}

// ReceiverAddressResponse represents the receiver's address within ShipmentsResponse.
type ReceiverAddressResponse struct {
	Address AddressResponse `json:"address"`

	StateName string `json:"state_name"`
	CityName  string `json:"city_name"`
	Floor     string `json:"floor"`
	Apartment string `json:"apartment"`
}

// OrderResponse represents order information.
type OrderResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// TransactionDetailsResponse represents transaction details.
type TransactionDetailsResponse struct {
	Barcode BarcodeResponse `json:"barcode"`

	FinancialInstitution     string  `json:"financial_institution"`
	ExternalResourceURL      string  `json:"external_resource_url"`
	PaymentMethodReferenceID string  `json:"payment_method_reference_id"`
	AcquirerReference        string  `json:"acquirer_reference"`
	TransactionID            string  `json:"transaction_id"`
	DigitableLine            string  `json:"digitable_line"`
	VerificationCode         string  `json:"verification_code"`
	PayableDeferralPeriod    string  `json:"payable_deferral_period"`
	BankTransferID           string  `json:"bank_transfer_id"`
	NetReceivedAmount        float64 `json:"net_received_amount"`
	TotalPaidAmount          float64 `json:"total_paid_amount"`
	InstallmentAmount        float64 `json:"installment_amount"`
	OverpaidAmount           float64 `json:"overpaid_amount"`
}

// CardResponse represents card information.
type CardResponse struct {
	Cardholder      CardholderResponse `json:"cardholder"`
	DateCreated     time.Time          `json:"date_created"`
	DateLastUpdated time.Time          `json:"date_last_updated"`

	ID              string `json:"id"`
	LastFourDigits  string `json:"last_four_digits"`
	FirstSixDigits  string `json:"first_six_digits"`
	ExpirationYear  int    `json:"expiration_year"`
	ExpirationMonth int    `json:"expiration_month"`
}

// CardholderResponse represents cardholder information.
type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification"`

	Name string `json:"name"`
}

// PointOfInteractionResponse represents point of interaction information.
type PointOfInteractionResponse struct {
	ApplicationData ApplicationDataResponse `json:"application_data"`
	TransactionData TransactionDataResponse `json:"transaction_data"`

	Type     string `json:"type"`
	SubType  string `json:"sub_type"`
	LinkedTo string `json:"linked_to"`
}

// ApplicationDataResponse represents application data within PointOfInteractionResponse.
type ApplicationDataResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// TransactionDataResponse represents transaction data within PointOfInteractionResponse.
type TransactionDataResponse struct {
	BankInfo             BankInfoResponse             `json:"bank_info"`
	SubscriptionSequence SubscriptionSequenceResponse `json:"subscription_sequence"`
	InvoicePeriod        InvoicePeriodResponse        `json:"invoice_period"`
	PaymentReference     PaymentReferenceResponse     `json:"payment_reference"`

	QRCode               string `json:"qr_code"`
	QRCodeBase64         string `json:"qr_code_base64"`
	TransactionID        string `json:"transaction_id"`
	TicketURL            string `json:"ticket_url"`
	SubscriptionID       string `json:"subscription_id"`
	BillingDate          string `json:"billing_date"`
	BankTransferID       int64  `json:"bank_transfer_id"`
	FinancialInstitution int64  `json:"financial_institution"`
	FirstTimeUse         bool   `json:"first_time_use"`
}

// BankInfoResponse represents bank information.
type BankInfoResponse struct {
	Payer     BankInfoPayerResponse     `json:"payer"`
	Collector BankInfoCollectorResponse `json:"collector"`

	IsSameBankAccountOwner string `json:"is_same_bank_account_owner"`
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
	Data DataResponse `json:"data"`

	ID       string `json:"id"`
	Type     string `json:"type"`
	IssuerID string `json:"issuer_id"`
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
	LimitDate time.Time `json:"limit_date"`

	Type  string  `json:"type"`
	Value float64 `json:"value"`
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
	Source      SourceResponse `json:"source"`
	DateCreated time.Time      `json:"date_created"`

	Status               string  `json:"status"`
	RefundMode           string  `json:"refund_mode"`
	Reason               string  `json:"reason"`
	UniqueSequenceNumber string  `json:"unique_sequence_number"`
	ID                   int64   `json:"id"`
	PaymentID            int64   `json:"payment_id"`
	Amount               float64 `json:"amount"`
	AdjustmentAmount     float64 `json:"adjustment_amount"`
}

// SourceResponse represents source information.
type SourceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// BarcodeResponse represents barcode information.
type BarcodeResponse struct {
	Content string `json:"content"`
}
