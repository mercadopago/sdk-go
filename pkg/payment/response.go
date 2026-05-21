package payment

import (
	"time"
)

// Response represents the full payment resource returned by the MercadoPago Payments API.
// It is returned by [Client.Create], [Client.Get], [Client.Cancel], [Client.Capture],
// and [Client.CaptureAmount].
type Response struct {
	Payer              PayerResponse              `json:"payer"`
	ForwardData        ForwardDataResponse        `json:"forward_data,omitempty"`
	AdditionalInfo     AdditionalInfoResponse     `json:"additional_info"`
	Order              OrderResponse              `json:"order"`
	TransactionDetails TransactionDetailsResponse `json:"transaction_details"`
	Card               CardResponse               `json:"card"`
	PointOfInteraction PointOfInteractionResponse `json:"point_of_interaction"`
	PaymentMethod      PaymentMethodResponse      `json:"payment_method"`
	ThreeDSInfo        ThreeDSInfoResponse        `json:"three_ds_info"`
	BackURLs           BackURLsResponse           `json:"back_urls,omitempty"`
	Expanded           *ExpandedResponse          `json:"expanded,omitempty"`
	DateCreated        time.Time                  `json:"date_created"`
	DateApproved       time.Time                  `json:"date_approved"`
	DateLastUpdated    time.Time                  `json:"date_last_updated"`
	DateOfExpiration   time.Time                  `json:"date_of_expiration"`
	MoneyReleaseDate   time.Time                  `json:"money_release_date"`
	FeeDetails         []FeeDetailResponse        `json:"fee_details"`
	Taxes              []TaxResponse              `json:"taxes"`
	Refunds            []RefundResponse           `json:"refunds"`
	Amounts            *AmountsResponse           `json:"amounts,omitempty"`
	CounterCurrency    *CounterCurrencyResponse   `json:"counter_currency,omitempty"`

	DifferentialPricingID     int            `json:"differential_pricing_id"`
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
	CallForAuthorizeID        string         `json:"call_for_authorize_id"`
	StatementDescriptor       string         `json:"statement_descriptor"`
	MoneyReleaseStatus        string         `json:"money_release_status"`
	TransactionAmount         float64        `json:"transaction_amount"`
	TransactionAmountRefunded float64        `json:"transaction_amount_refunded"`
	CouponAmount              float64        `json:"coupon_amount"`
	TaxesAmount               float64        `json:"taxes_amount"`
	ShippingAmount            float64        `json:"shipping_amount"`
	NetAmount                 float64        `json:"net_amount"`
	Installments              int            `json:"installments"`
	ID                        int            `json:"id"`
	SponsorID                 int            `json:"sponsor_id"`
	CollectorID               int64          `json:"collector_id"`
	LiveMode                  bool           `json:"live_mode"`
	Captured                  bool           `json:"captured"`
	BinaryMode                bool           `json:"binary_mode"`
	Metadata                  map[string]any `json:"metadata"`
	InternalMetadata          map[string]any `json:"internal_metadata"`
	DeviceIdentifier          string         `json:"device_identifier"`
	DeviceID                  string         `json:"device_id,omitempty"`
}

// PayerResponse contains the identification, contact, and authentication details
// of the person or entity who made the payment.
type PayerResponse struct {
	Type                  string `json:"type"`
	ID                    string `json:"id"`
	Email                 string `json:"email"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	EntityType            string `json:"entity_type"`
	AuthenticationType    string `json:"authentication_type"`
	IsPrimeUser           bool   `json:"is_prime_user"`
	IsFirstPurchaseOnline bool   `json:"is_first_purchase_online"`
	RegistrationDate      string `json:"registration_date"`
	LastPurchaseDate      string `json:"last_purchase_date"`
	DateCreated           string `json:"date_created"`

	Identification IdentificationResponse `json:"identification"`
	Phone          PhoneResponse          `json:"phone"`
	Address        AddressResponse        `json:"address"`
}

// ForwardDataResponse contains data forwarded to acquirers or processors, such as
// sub-merchant information for payment facilitator flows.
type ForwardDataResponse struct {
	SubMerchant SubMerchantResponse `json:"sub_merchant,omitempty"`
}

// SubMerchantResponse contains the sub-merchant details returned by the API when
// operating in a payment facilitator or marketplace model.
type SubMerchantResponse struct {
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

// IdentificationResponse contains the payer's personal identification document type
// (e.g., CPF, DNI) and number.
type IdentificationResponse struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// AdditionalInfoResponse contains supplementary payment data including payer details,
// shipment information, purchased items, and the originating IP address.
type AdditionalInfoResponse struct {
	Payer     AdditionalInfoPayerResponse `json:"payer"`
	Shipments ShipmentsResponse           `json:"shipments"`
	Items     []ItemResponse              `json:"items"`

	IPAddress string `json:"ip_address"`
}

// ItemResponse represents a purchased item or service associated with the payment.
type ItemResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`
	CategoryID  string `json:"category_id"`
	CurrencyID  string `json:"currency_id"`
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unit_price"`
	Warranty    string `json:"warranty"`

	CategoryDescriptor CategoryDescriptorResponse `json:"category_descriptor"`
}

// CategoryDescriptorResponse provides category-specific metadata for an item, such as
// passenger and route details for travel-related payments.
type CategoryDescriptorResponse struct {
	Passenger PassengerResponse `json:"passenger"`
	Route     RouteResponse     `json:"route"`
	EventDate string            `json:"event_date"`
	Type      string            `json:"type"`
}

// PassengerResponse contains passenger identification data returned for travel-related payments.
type PassengerResponse struct {
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	IdentificationType   string `json:"identification_type"`
	IdentificationNumber string `json:"identification_number"`
}

// RouteResponse describes a travel route with departure, destination, and carrier details.
type RouteResponse struct {
	Departure         string    `json:"departure"`
	Destination       string    `json:"destination"`
	DepartureDataTime time.Time `json:"departure_data_time"`
	ArrivalDateTime   time.Time `json:"arrival_date_time"`
	Company           string    `json:"company"`
}

// AdditionalInfoPayerResponse contains the payer's supplementary details returned
// in the additional_info section, including contact info and registration date.
type AdditionalInfoPayerResponse struct {
	Phone            PhoneResponse   `json:"phone"`
	Address          AddressResponse `json:"address"`
	RegistrationDate time.Time       `json:"registration_date"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
}

// PhoneResponse contains a phone number split into area code and number.
type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

// AddressResponse contains a street address with zip code, street name, and street number.
type AddressResponse struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}

// ShipmentsResponse contains the shipment details associated with the payment,
// including the receiver's address.
type ShipmentsResponse struct {
	ReceiverAddress ReceiverAddressResponse `json:"receiver_address"`
}

// ReceiverAddressResponse contains the full destination address for the shipment,
// including geographic details and express shipment information.
type ReceiverAddressResponse struct {
	Address AddressResponse `json:"address"`

	ZipCode           string `json:"zip_code"`
	StateName         string `json:"state_name"`
	CityName          string `json:"city_name"`
	StreetName        string `json:"street_name"`
	StreetNumber      string `json:"street_number"`
	CountryName       string `json:"country_name"`
	Floor             string `json:"floor"`
	Apartment         string `json:"apartment"`
	LocalPicURL       string `json:"local_pic_url"`
	ExpressShipmentID string `json:"express_shipment_id"`
}

// OrderResponse contains the order identifier and type associated with the payment.
type OrderResponse struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// TransactionDetailsResponse contains financial and processing details of the transaction,
// including amounts, external URLs, barcodes, and acquirer references.
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
	NetReceivedAmount        float64 `json:"net_received_amount"`
	TotalPaidAmount          float64 `json:"total_paid_amount"`
	InstallmentAmount        float64 `json:"installment_amount"`
	OverpaidAmount           float64 `json:"overpaid_amount"`
	BankTransferID           int     `json:"bank_transfer_id"`
}

// CardResponse contains the payment card details including masked card numbers,
// expiration dates, and cardholder information.
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

// CardholderResponse contains the cardholder's name and identification document.
type CardholderResponse struct {
	Identification IdentificationResponse `json:"identification"`

	Name string `json:"name"`
}

// PointOfInteractionResponse describes the context in which the payment was initiated,
// including application data, transaction details, and the interaction type (e.g., QR, POS).
type PointOfInteractionResponse struct {
	ApplicationData ApplicationDataResponse `json:"application_data"`
	TransactionData TransactionDataResponse `json:"transaction_data"`

	Type     string `json:"type"`
	SubType  string `json:"sub_type"`
	LinkedTo string `json:"linked_to"`
}

// ApplicationDataResponse identifies the application that originated the payment interaction.
type ApplicationDataResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// TransactionDataResponse contains transaction-specific data from the point of interaction,
// including QR codes, ticket URLs, bank information, and subscription details.
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
	BankTransferID       int    `json:"bank_transfer_id"`
	FinancialInstitution int    `json:"financial_institution"`
	FirstTimeUse         bool   `json:"first_time_use"`
}

// BankInfoResponse contains banking details for both the payer and the collector,
// including whether they share the same bank account ownership.
type BankInfoResponse struct {
	Payer     BankInfoPayerResponse     `json:"payer"`
	Collector BankInfoCollectorResponse `json:"collector"`

	IsSameBankAccountOwner bool `json:"is_same_bank_account_owner"`
}

// SubscriptionSequenceResponse tracks the current position and total count of payments
// within a recurring subscription.
type SubscriptionSequenceResponse struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

// InvoicePeriodResponse describes the billing cycle type and duration for subscription payments.
type InvoicePeriodResponse struct {
	Type   string `json:"type"`
	Period int    `json:"period"`
}

// PaymentReferenceResponse contains a reference identifier linking to a related payment.
type PaymentReferenceResponse struct {
	ID string `json:"id"`
}

// BankInfoPayerResponse contains the payer's bank account details within [BankInfoResponse].
type BankInfoPayerResponse struct {
	Email     string `json:"email"`
	LongName  string `json:"long_name"`
	AccountID int    `json:"account_id"`
}

// BankInfoCollectorResponse contains the collector's bank account details within [BankInfoResponse].
type BankInfoCollectorResponse struct {
	LongName  string `json:"long_name"`
	AccountID int    `json:"account_id"`
}

// PaymentMethodResponse contains the payment method used for the transaction,
// including type, issuer, and associated data.
type PaymentMethodResponse struct {
	Data DataResponse `json:"data"`

	ID       string `json:"id"`
	Type     string `json:"type"`
	IssuerID string `json:"issuer_id"`
}

// DataResponse contains extended payment method data including rules, reference IDs,
// and external resource URLs.
type DataResponse struct {
	Rules               RulesResponse `json:"rules"`
	ReferenceID         string        `json:"reference_id"`
	ExternalReferenceID string        `json:"external_reference_id"`
	ExternalResourceURL string        `json:"external_resource_url"`
}

// RulesResponse contains the fine, interest, and discount rules applied to the payment method.
type RulesResponse struct {
	Fine      FeeResponse        `json:"fine"`
	Interest  FeeResponse        `json:"interest"`
	Discounts []DiscountResponse `json:"discounts"`
}

// DiscountResponse contains details of an early-payment discount including its deadline.
type DiscountResponse struct {
	LimitDate time.Time `json:"limit_date"`

	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// FeeResponse contains a fine or interest fee applied to the payment.
type FeeResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// ThreeDSInfoResponse contains 3DS (Three-Domain Secure) challenge information
// returned when additional cardholder authentication is required.
type ThreeDSInfoResponse struct {
	ExternalResourceURL string `json:"external_resource_url"`
	Creq                string `json:"creq"`
}

// FeeDetailResponse describes a fee charged on the payment, including who pays it
// (payer or collector) and the amount.
type FeeDetailResponse struct {
	Type     string  `json:"type"`
	FeePayer string  `json:"fee_payer"`
	Amount   float64 `json:"amount"`
}

// TaxResponse contains tax type and value applied to the payment.
type TaxResponse struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// RefundResponse contains details of a refund applied to the payment, including
// the refunded amount, status, and originating source.
type RefundResponse struct {
	Source      SourceResponse `json:"source"`
	DateCreated time.Time      `json:"date_created"`

	Status               string  `json:"status"`
	RefundMode           string  `json:"refund_mode"`
	Reason               string  `json:"reason"`
	UniqueSequenceNumber string  `json:"unique_sequence_number"`
	Amount               float64 `json:"amount"`
	AdjustmentAmount     float64 `json:"adjustment_amount"`
	ID                   int     `json:"id"`
	PaymentID            int     `json:"payment_id"`
}

// SourceResponse identifies the origin of an action (e.g., a refund), including
// the source ID, name, and type.
type SourceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// BarcodeResponse contains the barcode content for ticket-based or offline payment methods.
type BarcodeResponse struct {
	Content string `json:"content"`
}

// BackURLsResponse contains the redirect URLs for the buyer after the payment flow completes,
// with separate URLs for success, pending, and failure outcomes.
type BackURLsResponse struct {
	Success string `json:"success,omitempty"`
	Pending string `json:"pending,omitempty"`
	Failure string `json:"failure,omitempty"`
}

// ExpandedResponse contains optional expanded data returned when specific query parameters
// are requested, such as gateway-level details.
type ExpandedResponse struct {
	Gateway *GatewayResponse `json:"gateway,omitempty"`
}

// GatewayResponse contains gateway-level processing details within [ExpandedResponse].
type GatewayResponse struct {
	Reference *ReferenceResponse `json:"reference,omitempty"`
}

// ReferenceResponse contains network-level transaction references from the payment gateway.
type ReferenceResponse struct {
	NetworkTransactionID string `json:"network_transaction_id,omitempty"`
}

// AmountsResponse contains currency-specific transaction amounts for both the collector
// and the payer, used in cross-border or cross-currency payment scenarios.
type AmountsResponse struct {
	Collector CollectorAmountResponse `json:"collector,omitempty"`
	Payer     PayerAmountResponse     `json:"payer,omitempty"`
}

// PayerAmountResponse contains the payer-side transaction and total paid amounts in their currency.
type PayerAmountResponse struct {
	CurrencyID  string  `json:"currency_id,omitempty"`
	Transaction float64 `json:"transaction,omitempty"`
	TotalPaid   float64 `json:"total_paid,omitempty"`
}

// CollectorAmountResponse contains the collector-side transaction and net received amounts in their currency.
type CollectorAmountResponse struct {
	CurrencyID  string  `json:"currency_id,omitempty"`
	Transaction float64 `json:"transaction,omitempty"`
	NetReceived float64 `json:"net_received,omitempty"`
}

// CounterCurrencyResponse contains the counter currency conversion details including
// the exchange rate and converted amounts.
type CounterCurrencyResponse struct {
	CurrencyID     string  `json:"currency_id,omitempty"`
	Rate           float64 `json:"rate,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	AmountRefunded float64 `json:"amount_refunded,omitempty"`
}
