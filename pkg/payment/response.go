package payment

import (
	"time"
)

// Response is the response from the Payments API.
type Response struct {
	DifferentialPricingID     string         `json:"differential_pricing_id,omitempty"`
	MoneyReleaseSchema        string         `json:"money_release_schema,omitempty"`
	OperationType             string         `json:"operation_type,omitempty"`
	IssuerID                  string         `json:"issuer_id,omitempty"`
	PaymentMethodID           string         `json:"payment_method_id,omitempty"`
	PaymentTypeID             string         `json:"payment_type_id,omitempty"`
	Status                    string         `json:"status,omitempty"`
	StatusDetail              string         `json:"status_detail,omitempty"`
	CurrencyID                string         `json:"currency_id,omitempty"`
	Description               string         `json:"description,omitempty"`
	AuthorizationCode         string         `json:"authorization_code,omitempty"`
	IntegratorID              string         `json:"integrator_id,omitempty"`
	PlatformID                string         `json:"platform_id,omitempty"`
	CorporationID             string         `json:"corporation_id,omitempty"`
	NotificationURL           string         `json:"notification_url,omitempty"`
	CallbackURL               string         `json:"callback_url,omitempty"`
	ProcessingMode            string         `json:"processing_mode,omitempty"`
	MerchantAccountID         string         `json:"merchant_account_id,omitempty"`
	MerchantNumber            string         `json:"merchant_number,omitempty"`
	CouponCode                string         `json:"coupon_code,omitempty"`
	ExternalReference         string         `json:"external_reference,omitempty"`
	PaymentMethodOptionID     string         `json:"payment_method_option_id,omitempty"`
	PosID                     string         `json:"pos_id,omitempty"`
	StoreID                   string         `json:"store_id,omitempty"`
	DeductionSchema           string         `json:"deduction_schema,omitempty"`
	CounterCurrency           string         `json:"counter_currency,omitempty"`
	CallForAuthorizeID        string         `json:"call_for_authorize_id,omitempty"`
	StatementDescriptor       string         `json:"statement_descriptor,omitempty"`
	Installments              int            `json:"installments,omitempty"`
	ID                        int64          `json:"id,omitempty"`
	SponsorID                 int64          `json:"sponsor_id,omitempty"`
	CollectorID               int64          `json:"collector_id,omitempty"`
	TransactionAmount         float64        `json:"transaction_amount,omitempty"`
	TransactionAmountRefunded float64        `json:"transaction_amount_refunded,omitempty"`
	CouponAmount              float64        `json:"coupon_amount,omitempty"`
	TaxesAmount               float64        `json:"taxes_amount,omitempty"`
	ShippingAmount            float64        `json:"shipping_amount,omitempty"`
	NetAmount                 float64        `json:"net_amount,omitempty"`
	LiveMode                  bool           `json:"live_mode,omitempty"`
	Captured                  bool           `json:"captured,omitempty"`
	BinaryMode                bool           `json:"binary_mode,omitempty"`
	Metadata                  map[string]any `json:"metadata,omitempty"`
	InternalMetadata          map[string]any `json:"internal_metadata,omitempty"`

	DateCreated        *time.Time                  `json:"date_created,omitempty"`
	DateApproved       *time.Time                  `json:"date_approved,omitempty"`
	DateLastUpdated    *time.Time                  `json:"date_last_updated,omitempty"`
	DateOfExpiration   *time.Time                  `json:"date_of_expiration,omitempty"`
	MoneyReleaseDate   *time.Time                  `json:"money_release_date,omitempty"`
	Payer              *PayerResponse              `json:"payer,omitempty"`
	AdditionalInfo     *AdditionalInfoResponse     `json:"additional_info,omitempty"`
	Order              *OrderResponse              `json:"order,omitempty"`
	TransactionDetails *TransactionDetailsResponse `json:"transaction_details,omitempty"`
	Card               *CardResponse               `json:"card,omitempty"`
	PointOfInteraction *PointOfInteractionResponse `json:"point_of_interaction,omitempty"`
	PaymentMethod      *PaymentMethodResponse      `json:"payment_method,omitempty"`
	ThreeDSInfo        *ThreeDSInfoResponse        `json:"three_ds_info,omitempty"`
	FeeDetails         []FeeDetailResponse         `json:"fee_details,omitempty"`
	Taxes              []TaxResponse               `json:"taxes,omitempty"`
	Refunds            []RefundResponse            `json:"refunds,omitempty"`
}

// PayerResponse represents the payer of the payment.
type PayerResponse struct {
	Type       string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	Email      string `json:"email,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	EntityType string `json:"entity_type,omitempty"`

	Identification *IdentificationResponse `json:"identification,omitempty"`
}

// IdentificationResponse represents payer's personal identification.
type IdentificationResponse struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// AdditionalInfoResponse represents additional information about a payment.
type AdditionalInfoResponse struct {
	IPAddress string `json:"ip_address,omitempty"`

	Payer     *AdditionalInfoPayerResponse `json:"payer,omitempty"`
	Shipments *ShipmentsResponse           `json:"shipments,omitempty"`
	Items     []ItemResponse               `json:"items,omitempty"`
}

// ItemResponse represents an item.
type ItemResponse struct {
	ID          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	PictureURL  string  `json:"picture_url,omitempty"`
	CategoryID  string  `json:"category_id,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
}

// AdditionalInfoPayerResponse represents payer's additional information.
type AdditionalInfoPayerResponse struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`

	RegistrationDate *time.Time       `json:"registration_date,omitempty"`
	Phone            *PhoneResponse   `json:"phone,omitempty"`
	Address          *AddressResponse `json:"address,omitempty"`
}

// PhoneResponse represents phone information.
type PhoneResponse struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// AddressResponse represents address information.
type AddressResponse struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

// ShipmentsResponse represents shipment information.
type ShipmentsResponse struct {
	ReceiverAddress *ReceiverAddressResponse `json:"receiver_address,omitempty"`
}

// ReceiverAddressResponse represents the receiver's address within ShipmentsResponse.
type ReceiverAddressResponse struct {
	StateName string `json:"state_name,omitempty"`
	CityName  string `json:"city_name,omitempty"`
	Floor     string `json:"floor,omitempty"`
	Apartment string `json:"apartment,omitempty"`

	Address *AddressResponse `json:"address,omitempty"`
}

// OrderResponse represents order information.
type OrderResponse struct {
	ID   int    `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// TransactionDetailsResponse represents transaction details.
type TransactionDetailsResponse struct {
	FinancialInstitution     string  `json:"financial_institution,omitempty"`
	ExternalResourceURL      string  `json:"external_resource_url,omitempty"`
	PaymentMethodReferenceID string  `json:"payment_method_reference_id,omitempty"`
	AcquirerReference        string  `json:"acquirer_reference,omitempty"`
	NetReceivedAmount        float64 `json:"net_received_amount,omitempty"`
	TotalPaidAmount          float64 `json:"total_paid_amount,omitempty"`
	InstallmentAmount        float64 `json:"installment_amount,omitempty"`
	OverpaidAmount           float64 `json:"overpaid_amount,omitempty"`
}

// CardResponse represents card information.
type CardResponse struct {
	ID              string `json:"id,omitempty"`
	LastFourDigits  string `json:"last_four_digits,omitempty"`
	FirstSixDigits  string `json:"first_six_digits,omitempty"`
	ExpirationYear  int    `json:"expiration_year,omitempty"`
	ExpirationMonth int    `json:"expiration_month,omitempty"`

	DateCreated     *time.Time          `json:"date_created,omitempty"`
	DateLastUpdated *time.Time          `json:"date_last_updated,omitempty"`
	Cardholder      *CardholderResponse `json:"cardholder,omitempty"`
}

// CardholderResponse represents cardholder information.
type CardholderResponse struct {
	Name string `json:"name,omitempty"`

	Identification *IdentificationResponse `json:"identification,omitempty"`
}

// PointOfInteractionResponse represents point of interaction information.
type PointOfInteractionResponse struct {
	Type     string `json:"type,omitempty"`
	SubType  string `json:"sub_type,omitempty"`
	LinkedTo string `json:"linked_to,omitempty"`

	ApplicationData *ApplicationDataResponse `json:"application_data,omitempty"`
	TransactionData *TransactionDataResponse `json:"transaction_data,omitempty"`
}

// ApplicationDataResponse represents application data within PointOfInteractionResponse.
type ApplicationDataResponse struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// TransactionDataResponse represents transaction data within PointOfInteractionResponse.
type TransactionDataResponse struct {
	QRCode               string `json:"qr_code,omitempty"`
	QRCodeBase64         string `json:"qr_code_base64,omitempty"`
	TransactionID        string `json:"transaction_id,omitempty"`
	TicketURL            string `json:"ticket_url,omitempty"`
	BankTransferID       int64  `json:"bank_transfer_id,omitempty"`
	FinancialInstitution int64  `json:"financial_institution,omitempty"`

	BankInfo *BankInfoResponse `json:"bank_info,omitempty"`
}

// BankInfoResponse represents bank information.
type BankInfoResponse struct {
	IsSameBankAccountOwner string `json:"is_same_bank_account_owner,omitempty"`

	Payer     *BankInfoPayerResponse     `json:"payer,omitempty"`
	Collector *BankInfoCollectorResponse `json:"collector,omitempty"`
}

// BankInfoPayerResponse represents payer information within BankInfoResponse.
type BankInfoPayerResponse struct {
	Email     string `json:"email,omitempty"`
	LongName  string `json:"long_name,omitempty"`
	AccountID int64  `json:"account_id,omitempty"`
}

// BankInfoCollectorResponse represents collector information within BankInfoResponse.
type BankInfoCollectorResponse struct {
	LongName  string `json:"long_name,omitempty"`
	AccountID int64  `json:"account_id,omitempty"`
}

// PaymentMethodResponse represents payment method information.
type PaymentMethodResponse struct {
	Data *DataResponse `json:"data,omitempty"`
}

// DataResponse represents data within PaymentMethodResponse.
type DataResponse struct {
	Rules *RulesResponse `json:"rules,omitempty"`
}

// RulesResponse represents payment rules.
type RulesResponse struct {
	Fine      *FeeResponse       `json:"fine,omitempty"`
	Interest  *FeeResponse       `json:"interest,omitempty"`
	Discounts []DiscountResponse `json:"discounts,omitempty"`
}

// DiscountResponse represents payment discount information.
type DiscountResponse struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`

	LimitDate *time.Time `json:"limit_date,omitempty"`
}

// FeeResponse represents payment fee information.
type FeeResponse struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// ThreeDSInfoResponse represents 3DS (Three-Domain Secure) information.
type ThreeDSInfoResponse struct {
	ExternalResourceURL string `json:"external_resource_url,omitempty"`
	Creq                string `json:"creq,omitempty"`
}

// FeeDetailResponse represents payment fee detail information.
type FeeDetailResponse struct {
	Type     string  `json:"type,omitempty"`
	FeePayer string  `json:"fee_payer,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
}

// TaxResponse represents tax information.
type TaxResponse struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// RefundResponse represents refund information.
type RefundResponse struct {
	Status               string  `json:"status,omitempty"`
	RefundMode           string  `json:"refund_mode,omitempty"`
	Reason               string  `json:"reason,omitempty"`
	UniqueSequenceNumber string  `json:"unique_sequence_number,omitempty"`
	ID                   int64   `json:"id,omitempty"`
	PaymentID            int64   `json:"payment_id,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
	AdjustmentAmount     float64 `json:"adjustment_amount,omitempty"`

	DateCreated *time.Time      `json:"date_created,omitempty"`
	Source      *SourceResponse `json:"source,omitempty"`
}

// SourceResponse represents source information.
type SourceResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
