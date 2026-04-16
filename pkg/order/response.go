package order

// API version: d0494f1c-8d81-4c76-ae1d-0c65bb8ef6de

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

type TransactionResponse struct {
	Payments   []PaymentResponse   `json:"payments,omitempty"`
	Refunds    []RefundResponse    `json:"refunds,omitempty"`
	Chargebacks []ChargebackResponse `json:"chargebacks,omitempty"`
}

type PaymentResponse struct {
	ID               string                   `json:"id,omitempty"`
	ReferenceID      string                   `json:"reference_id,omitempty"`
	Status           string                   `json:"status,omitempty"`
	StatusDetail     string                   `json:"status_detail,omitempty"`
	Amount           string                   `json:"amount,omitempty"`
	PaidAmount       string                   `json:"paid_amount,omitempty"`
	DateOfExpiration string                   `json:"date_of_expiration,omitempty"`
	ExpirationTime   string                   `json:"expiration_time,omitempty"`
	AttemptNumber    int                      `json:"attempt_number,omitempty"`
	PaymentMethod    PaymentMethodResponse    `json:"payment_method,omitempty"`
	AutomaticPayments AutomaticPaymentResponse `json:"automatic_payments,omitempty"`
	StoredCredential StoredCredentialResponse `json:"stored_credential,omitempty"`
	SubscriptionData SubscriptionDataResponse `json:"subscription_data,omitempty"`
	Attempts         []AttemptResponse        `json:"attempts,omitempty"`
	RefundedAmount   string                   `json:"refunded_amount,omitempty"`
	Provider         string                   `json:"provider,omitempty"`
	Discounts        []DiscountResponse       `json:"discounts,omitempty"`
}

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

// TransactionSecurityResponse represents 3DS-related information returned by the API
// for a payment method when a challenge may be required.
type TransactionSecurityResponse struct {
	ID             string `json:"id,omitempty"`
	URL            string `json:"url,omitempty"`
	Validation     string `json:"validation,omitempty"`
	LiabilityShift string `json:"liability_shift,omitempty"`
	Type           string `json:"type,omitempty"`
	Status         string `json:"status,omitempty"`
}

type AutomaticPaymentResponse struct {
	PaymentProfileID string `json:"payment_profile_id,omitempty"`
	ScheduleDate     string `json:"schedule_date,omitempty"`
	DueDate          string `json:"due_date,omitempty"`
	Retries          int    `json:"retries,omitempty"`
}

type StoredCredentialResponse struct {
	PaymentInitiator   string `json:"payment_initiator,omitempty"`
	Reason             string `json:"reason,omitempty"`
	StorePaymentMethod bool   `json:"store_payment_method,omitempty"`
	FirstPayment       bool   `json:"first_payment,omitempty"`
}

type SubscriptionDataResponse struct {
	InvoiceID            string                       `json:"invoice_id,omitempty"`
	BillingDate          string                       `json:"billing_date,omitempty"`
	SubscriptionSequence SubscriptionSequenceResponse `json:"subscription_sequence,omitempty"`
	InvoicePeriod        InvoicePeriodResponse        `json:"invoice_period,omitempty"`
}

type SubscriptionSequenceResponse struct {
	Number int `json:"number,omitempty"`
	Total  int `json:"total,omitempty"`
}

type InvoicePeriodResponse struct {
	Type   string `json:"type,omitempty"`
	Period int    `json:"period,omitempty"`
}

type RefundResponse struct {
	ID            string          `json:"id,omitempty"`
	TransactionID string          `json:"transaction_id,omitempty"`
	ReferenceID   string          `json:"reference_id,omitempty"`
	Status        string          `json:"status,omitempty"`
	Amount        string          `json:"amount,omitempty"`
	Items         []RefundItemResponse `json:"items,omitempty"`
}

type RefundItemResponse struct {
	ID     string `json:"id,omitempty"`
	E2eID  string `json:"e2e_id,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type RefundReferenceResponse struct {
	ID       string `json:"id,omitempty"`
	SourceID string `json:"source_id,omitempty"`
}

type ItemsResponse struct {
	Title              string                   `json:"title,omitempty"`
	UnitPrice          string                   `json:"unit_price,omitempty"`
	ExternalCode       string                   `json:"external_code,omitempty"`
	Description        string                   `json:"description,omitempty"`
	CategoryID         string                   `json:"category_id,omitempty"`
	PictureURL         string                   `json:"picture_url,omitempty"`
	Quantity           int                      `json:"quantity,omitempty"`
	Type               string                   `json:"type,omitempty"`
	Warranty           string                   `json:"warranty,omitempty"`
	EventDate          string                   `json:"event_date,omitempty"`
	UnitMeasure        string                   `json:"unit_measure,omitempty"`
	ExternalCategories []ExternalCategoryResponse `json:"external_categories,omitempty"`
}

type IntegrationDataResponse struct {
	CorporationID string          `json:"corporation_id,omitempty"`
	ApplicationID string          `json:"application_id,omitempty"`
	IntegratorID  string          `json:"integrator_id,omitempty"`
	PlatformID    string          `json:"platform_id,omitempty"`
	Sponsor       SponsorResponse `json:"sponsor,omitempty"`
}

type SponsorResponse struct {
	ID string `json:"id,omitempty"`
}

type ConfigResponse struct {
	PaymentMethodResponse PaymentMethodConfigResponse `json:"payment_method,omitempty"`
	Online                OnlineConfigResponse        `json:"online,omitempty"`
}

type PaymentMethodConfigResponse struct {
	NotAllowedIDs       []string                   `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes     []string                   `json:"not_allowed_types,omitempty"`
	DefaultID           string                     `json:"default_id,omitempty"`
	MaxInstallments     int                        `json:"max_installments,omitempty"`
	DefaultInstallments int                        `json:"default_installments,omitempty"`
	DefaultType         string                     `json:"default_type,omitempty"`
	InstallmentsCost    string                     `json:"installments_cost,omitempty"`
	Installments        *InstallmentsResponse      `json:"installments,omitempty"`
	MinInstallments     int                        `json:"min_installments,omitempty"`
}

type OnlineConfigResponse struct {
	CallbackURL         string                      `json:"callback_url,omitempty"`
	SuccessURL          string                      `json:"success_url,omitempty"`
	PendingURL          string                      `json:"pending_url,omitempty"`
	FailureURL          string                      `json:"failure_url,omitempty"`
	AutoReturnURL       string                      `json:"auto_return_url,omitempty"`
	DifferentialPricing DifferentialPricingResponse `json:"differential_pricing,omitempty"`
}

type DifferentialPricingResponse struct {
	ID string `json:"id,omitempty"`
}

type PayerResponse struct {
	CustomerID string `json:"customer_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}

type AttemptResponse struct {
	ID            string                `json:"id,omitempty"`
	Status        string                `json:"status,omitempty"`
	StatusDetail  string                `json:"status_detail,omitempty"`
	PaymentMethod PaymentMethodResponse `json:"payment_method,omitempty"`
}

type TypeResponse struct {
	QrData string `json:"qr_data,omitempty"`
}

type TaxResponse struct {
	PayerCondition string `json:"payer_condition,omitempty"`
	Type           string `json:"type,omitempty"`
	Value          string `json:"value,omitempty"`
}

type DiscountsResponse struct {
	PaymentMethods []DiscountPaymentMethodResponse `json:"payment_methods,omitempty"`
}

type DiscountPaymentMethodResponse struct {
	Type           string `json:"type,omitempty"`
	NewTotalAmount string `json:"new_total_amount,omitempty"`
}

type DiscountResponse struct {
	Type string `json:"type,omitempty"`
}

type InstallmentsResponse struct {
	InterestFree *InstallmentsInterestFreeResponse `json:"interest_free,omitempty"`
	Available    *InstallmentsAvailableResponse    `json:"available,omitempty"`
}

type InstallmentsInterestFreeResponse struct {
	Type   string `json:"type,omitempty"`
	Values []int  `json:"values,omitempty"`
}

type InstallmentsAvailableResponse struct {
	Type string `json:"type,omitempty"`
}

type ChargebackResponse struct {
	ID            string   `json:"id,omitempty"`
	TransactionID string   `json:"transaction_id,omitempty"`
	CaseID        string   `json:"case_id,omitempty"`
	Status        string   `json:"status,omitempty"`
	References    []string `json:"references,omitempty"`
}

type ExternalCategoryResponse struct {
	ID string `json:"id,omitempty"`
}

type SearchResponse struct {
	Data   []Response      `json:"data,omitempty"`
	Paging *PagingResponse `json:"paging,omitempty"`
}

type PagingResponse struct {
	Total      string `json:"total,omitempty"`
	TotalPages string `json:"total_pages,omitempty"`
	Offset     string `json:"offset,omitempty"`
	Limit      string `json:"limit,omitempty"`
}
