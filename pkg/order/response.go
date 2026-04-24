package order

// API version: d0494f1c-8d81-4c76-ae1d-0c65bb8ef6de

// Response represents the full order object returned by the MercadoPago Orders API.
// It is the primary return type for most [Client] methods and contains the order's
// current state, including status, amounts, transactions, items, and configuration.
type Response struct {
	ID                  string                  `json:"id,omitempty"`
	Type                string                  `json:"type,omitempty"`
	ExternalReference   string                  `json:"external_reference,omitempty"`
	CountryCode         string                  `json:"country_code,omitempty"`
	Status              string                  `json:"status,omitempty"`
	StatusDetail        string                  `json:"status_detail,omitempty"`
	CaptureMode         string                  `json:"capture_mode,omitempty"`
	UserID              string                  `json:"user_id,omitempty"`
	ClientToken         string                  `json:"client_token,omitempty"`
	TotalAmount         string                  `json:"total_amount,omitempty"`
	TotalPaidAmount     string                  `json:"total_paid_amount,omitempty"`
	ProcessingMode      string                  `json:"processing_mode,omitempty"`
	Description         string                  `json:"description,omitempty"`
	Marketplace         string                  `json:"marketplace,omitempty"`
	MarketplaceFee      string                  `json:"marketplace_fee,omitempty"`
	CheckoutAvailableAt string                  `json:"checkout_available_at,omitempty"`
	ExpirationTime      string                  `json:"expiration_time,omitempty"`
	CreatedDate         string                  `json:"created_date,omitempty"`
	LastUpdatedDate     string                  `json:"last_updated_date,omitempty"`
	Currency            string                  `json:"currency,omitempty"`
	Transactions        TransactionResponse     `json:"transactions,omitempty"`
	Items               []ItemsResponse         `json:"items,omitempty"`
	IntegrationData     IntegrationDataResponse `json:"integration_data,omitempty"`
	Config              ConfigResponse          `json:"config,omitempty"`
	Payer               PayerResponse           `json:"payer,omitempty"`
	Taxes               []TaxResponse           `json:"taxes,omitempty"`
	Discounts           *DiscountsResponse      `json:"discounts,omitempty"`
	TypeResponse        *TypeResponse           `json:"type_response,omitempty"`
}

// TransactionResponse contains the financial transactions associated with an order,
// including payments, refunds, and chargebacks.
type TransactionResponse struct {
	Payments    []PaymentResponse    `json:"payments,omitempty"`
	Refunds     []RefundResponse     `json:"refunds,omitempty"`
	Chargebacks []ChargebackResponse `json:"chargebacks,omitempty"`
}

// PaymentResponse represents an individual payment transaction returned by the API.
// It includes the payment's current status, amounts, method details, and any
// previous processing attempts.
type PaymentResponse struct {
	ID                string                   `json:"id,omitempty"`
	ReferenceID       string                   `json:"reference_id,omitempty"`
	Status            string                   `json:"status,omitempty"`
	StatusDetail      string                   `json:"status_detail,omitempty"`
	Amount            string                   `json:"amount,omitempty"`
	PaidAmount        string                   `json:"paid_amount,omitempty"`
	DateOfExpiration  string                   `json:"date_of_expiration,omitempty"`
	ExpirationTime    string                   `json:"expiration_time,omitempty"`
	AttemptNumber     int                      `json:"attempt_number,omitempty"`
	PaymentMethod     PaymentMethodResponse    `json:"payment_method,omitempty"`
	AutomaticPayments AutomaticPaymentResponse `json:"automatic_payments,omitempty"`
	StoredCredential  StoredCredentialResponse `json:"stored_credential,omitempty"`
	SubscriptionData  SubscriptionDataResponse `json:"subscription_data,omitempty"`
	Attempts          []AttemptResponse        `json:"attempts,omitempty"`
	RefundedAmount    string                   `json:"refunded_amount,omitempty"`
	Provider          string                   `json:"provider,omitempty"`
	Discounts         []DiscountResponse       `json:"discounts,omitempty"`
}

// PaymentMethodResponse represents the resolved payment method details returned by the API.
// It includes method-specific fields such as ticket URLs for offline payments, QR codes
// for Pix, barcode content for boleto, and 3D Secure challenge data.
type PaymentMethodResponse struct {
	ID                   string                       `json:"id,omitempty"`
	CardID               string                       `json:"card_id,omitempty"`
	Type                 string                       `json:"type,omitempty"`
	Token                string                       `json:"token,omitempty"`
	StatementDescriptor  string                       `json:"statement_descriptor,omitempty"`
	Installments         int                          `json:"installments,omitempty"`
	TicketURL            string                       `json:"ticket_url,omitempty"`
	BarcodeContent       string                       `json:"barcode_content,omitempty"`
	Reference            string                       `json:"reference,omitempty"`
	ReferenceID          string                       `json:"reference_id,omitempty"`
	VerificationCode     string                       `json:"verification_code,omitempty"`
	FinancialInstitution string                       `json:"financial_institution,omitempty"`
	QrCode               string                       `json:"qr_code,omitempty"`
	QrCodeBase64         string                       `json:"qr_code_base64,omitempty"`
	DigitableLine        string                       `json:"digitable_line,omitempty"`
	TransactionSecurity  *TransactionSecurityResponse `json:"transaction_security,omitempty"`
	E2eID                string                       `json:"e2e_id,omitempty"`
	RedirectURL          string                       `json:"redirect_url,omitempty"`
}

// TransactionSecurityResponse represents 3D Secure (3DS) authentication data returned
// by the API for a payment method. When a challenge is required, URL contains the
// address where the payer must complete authentication.
type TransactionSecurityResponse struct {
	ID             string `json:"id,omitempty"`
	URL            string `json:"url,omitempty"`
	Validation     string `json:"validation,omitempty"`
	LiabilityShift string `json:"liability_shift,omitempty"`
	Type           string `json:"type,omitempty"`
	Status         string `json:"status,omitempty"`
}

// AutomaticPaymentResponse represents the automatic payment scheduling information
// returned by the API for a recurring payment transaction.
type AutomaticPaymentResponse struct {
	PaymentProfileID string `json:"payment_profile_id,omitempty"`
	ScheduleDate     string `json:"schedule_date,omitempty"`
	DueDate          string `json:"due_date,omitempty"`
	Retries          int    `json:"retries,omitempty"`
}

// StoredCredentialResponse represents the stored credential metadata returned by the API,
// indicating whether the payment method was stored and the initiator of the transaction.
type StoredCredentialResponse struct {
	PaymentInitiator   string `json:"payment_initiator,omitempty"`
	Reason             string `json:"reason,omitempty"`
	StorePaymentMethod bool   `json:"store_payment_method,omitempty"`
	FirstPayment       bool   `json:"first_payment,omitempty"`
}

// SubscriptionDataResponse represents the subscription billing details returned by the API
// for a payment that belongs to a recurring subscription.
type SubscriptionDataResponse struct {
	InvoiceID            string                       `json:"invoice_id,omitempty"`
	BillingDate          string                       `json:"billing_date,omitempty"`
	SubscriptionSequence SubscriptionSequenceResponse `json:"subscription_sequence,omitempty"`
	InvoicePeriod        InvoicePeriodResponse        `json:"invoice_period,omitempty"`
}

// SubscriptionSequenceResponse represents the position of a payment within its subscription
// series as returned by the API. Number is the current installment and Total is the
// total number of planned installments.
type SubscriptionSequenceResponse struct {
	Number int `json:"number,omitempty"`
	Total  int `json:"total,omitempty"`
}

// InvoicePeriodResponse represents the billing period for a subscription invoice
// as returned by the API. Type describes the period unit and Period is the count.
type InvoicePeriodResponse struct {
	Type   string `json:"type,omitempty"`
	Period int    `json:"period,omitempty"`
}

// RefundResponse represents a refund operation associated with an order transaction,
// including its status and the individual refund items.
type RefundResponse struct {
	ID            string               `json:"id,omitempty"`
	TransactionID string               `json:"transaction_id,omitempty"`
	ReferenceID   string               `json:"reference_id,omitempty"`
	Status        string               `json:"status,omitempty"`
	Amount        string               `json:"amount,omitempty"`
	Items         []RefundItemResponse `json:"items,omitempty"`
}

// RefundItemResponse represents a single item within a refund, identified by its ID
// and the refunded amount. E2eID is the end-to-end identifier assigned by the payment network.
type RefundItemResponse struct {
	ID     string `json:"id,omitempty"`
	E2eID  string `json:"e2e_id,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// RefundReferenceResponse represents a cross-reference between a refund and its source
// transaction, enabling traceability of refund operations.
type RefundReferenceResponse struct {
	ID       string `json:"id,omitempty"`
	SourceID string `json:"source_id,omitempty"`
}

// ItemsResponse represents a line item within an order as returned by the API.
// It includes product details, pricing, and optional external category classifications.
type ItemsResponse struct {
	Title              string                     `json:"title,omitempty"`
	UnitPrice          string                     `json:"unit_price,omitempty"`
	ExternalCode       string                     `json:"external_code,omitempty"`
	Description        string                     `json:"description,omitempty"`
	CategoryID         string                     `json:"category_id,omitempty"`
	PictureURL         string                     `json:"picture_url,omitempty"`
	Quantity           int                        `json:"quantity,omitempty"`
	Type               string                     `json:"type,omitempty"`
	Warranty           string                     `json:"warranty,omitempty"`
	EventDate          string                     `json:"event_date,omitempty"`
	UnitMeasure        string                     `json:"unit_measure,omitempty"`
	ExternalCategories []ExternalCategoryResponse `json:"external_categories,omitempty"`
}

// IntegrationDataResponse represents integration metadata returned by the API, identifying
// the corporation, application, integrator, and platform associated with the order.
type IntegrationDataResponse struct {
	CorporationID string          `json:"corporation_id,omitempty"`
	ApplicationID string          `json:"application_id,omitempty"`
	IntegratorID  string          `json:"integrator_id,omitempty"`
	PlatformID    string          `json:"platform_id,omitempty"`
	Sponsor       SponsorResponse `json:"sponsor,omitempty"`
}

// SponsorResponse represents the sponsor associated with an integration, identified by
// the sponsor's MercadoPago account ID.
type SponsorResponse struct {
	ID string `json:"id,omitempty"`
}

// ConfigResponse represents the order-level configuration returned by the API, including
// payment method restrictions and online checkout redirect settings.
type ConfigResponse struct {
	PaymentMethodResponse PaymentMethodConfigResponse `json:"payment_method,omitempty"`
	Online                OnlineConfigResponse        `json:"online,omitempty"`
}

// PaymentMethodConfigResponse represents the payment method constraints and installment
// settings returned by the API for an order.
type PaymentMethodConfigResponse struct {
	NotAllowedIDs       []string              `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes     []string              `json:"not_allowed_types,omitempty"`
	DefaultID           string                `json:"default_id,omitempty"`
	MaxInstallments     int                   `json:"max_installments,omitempty"`
	DefaultInstallments int                   `json:"default_installments,omitempty"`
	DefaultType         string                `json:"default_type,omitempty"`
	InstallmentsCost    string                `json:"installments_cost,omitempty"`
	Installments        *InstallmentsResponse `json:"installments,omitempty"`
	MinInstallments     int                   `json:"min_installments,omitempty"`
}

// OnlineConfigResponse represents the online checkout configuration returned by the API,
// including redirect URLs for payment outcomes and differential pricing settings.
type OnlineConfigResponse struct {
	CallbackURL         string                      `json:"callback_url,omitempty"`
	SuccessURL          string                      `json:"success_url,omitempty"`
	PendingURL          string                      `json:"pending_url,omitempty"`
	FailureURL          string                      `json:"failure_url,omitempty"`
	AutoReturnURL       string                      `json:"auto_return_url,omitempty"`
	DifferentialPricing DifferentialPricingResponse `json:"differential_pricing,omitempty"`
}

// DifferentialPricingResponse represents a differential pricing configuration returned
// by the API, identified by its unique ID.
type DifferentialPricingResponse struct {
	ID string `json:"id,omitempty"`
}

// PayerResponse represents the payer information returned by the API for an order.
// It includes the customer identifier and entity type (e.g., individual or association).
type PayerResponse struct {
	CustomerID string `json:"customer_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}

// AttemptResponse represents a single payment processing attempt within a transaction.
// Multiple attempts may exist when a payment is retried with different methods or
// after a failure.
type AttemptResponse struct {
	ID            string                `json:"id,omitempty"`
	Status        string                `json:"status,omitempty"`
	StatusDetail  string                `json:"status_detail,omitempty"`
	PaymentMethod PaymentMethodResponse `json:"payment_method,omitempty"`
}

// TypeResponse represents type-specific data returned by the API. For QR-based orders,
// QrData contains the QR code payload.
type TypeResponse struct {
	QrData string `json:"qr_data,omitempty"`
}

// TaxResponse represents a tax applied to an order as returned by the API.
// It describes the tax type, the payer's fiscal condition, and the tax value.
type TaxResponse struct {
	PayerCondition string `json:"payer_condition,omitempty"`
	Type           string `json:"type,omitempty"`
	Value          string `json:"value,omitempty"`
}

// DiscountsResponse represents the collection of discounts applied to an order,
// grouped by payment method.
type DiscountsResponse struct {
	PaymentMethods []DiscountPaymentMethodResponse `json:"payment_methods,omitempty"`
}

// DiscountPaymentMethodResponse represents a discount available for a specific payment
// method type, including the resulting total amount after the discount is applied.
type DiscountPaymentMethodResponse struct {
	Type           string `json:"type,omitempty"`
	NewTotalAmount string `json:"new_total_amount,omitempty"`
}

// DiscountResponse represents a discount applied to an individual payment transaction.
type DiscountResponse struct {
	Type string `json:"type,omitempty"`
}

// InstallmentsResponse represents the installment options returned by the API for an order's
// payment method configuration, including interest-free and generally available plans.
type InstallmentsResponse struct {
	InterestFree *InstallmentsInterestFreeResponse `json:"interest_free,omitempty"`
	Available    *InstallmentsAvailableResponse    `json:"available,omitempty"`
}

// InstallmentsInterestFreeResponse represents interest-free installment options.
// Values contains the specific installment counts (e.g., 3, 6, 12) that qualify
// for interest-free financing.
type InstallmentsInterestFreeResponse struct {
	Type   string `json:"type,omitempty"`
	Values []int  `json:"values,omitempty"`
}

// InstallmentsAvailableResponse represents the available installment plan type
// returned by the API for a payment method configuration.
type InstallmentsAvailableResponse struct {
	Type string `json:"type,omitempty"`
}

// ChargebackResponse represents a chargeback dispute associated with a payment transaction.
// It includes the dispute case identifier and its current resolution status.
type ChargebackResponse struct {
	ID            string   `json:"id,omitempty"`
	TransactionID string   `json:"transaction_id,omitempty"`
	CaseID        string   `json:"case_id,omitempty"`
	Status        string   `json:"status,omitempty"`
	References    []string `json:"references,omitempty"`
}

// ExternalCategoryResponse represents an external product category classification
// associated with an order item.
type ExternalCategoryResponse struct {
	ID string `json:"id,omitempty"`
}

// SearchResponse represents the paginated result returned by the [Client.Search] method.
// Data contains the matching orders and Paging provides pagination metadata.
type SearchResponse struct {
	Data   []Response      `json:"data,omitempty"`
	Paging *PagingResponse `json:"paging,omitempty"`
}

// PagingResponse represents pagination metadata for a search result, indicating the
// total number of matching orders, the total number of pages, and the current offset and limit.
type PagingResponse struct {
	Total      string `json:"total,omitempty"`
	TotalPages string `json:"total_pages,omitempty"`
	Offset     string `json:"offset,omitempty"`
	Limit      string `json:"limit,omitempty"`
}
