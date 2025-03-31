package order

// API version: d0494f1c-8d81-4c76-ae1d-0c65bb8ef6de

type Response struct {
	ID                  string                  `json:"id"`
	Type                string                  `json:"type"`
	ExternalReference   string                  `json:"external_reference"`
	CountryCode         string                  `json:"country_code"`
	Status              string                  `json:"status"`
	StatusDetail        string                  `json:"status_detail"`
	CaptureMode         string                  `json:"capture_mode"`
	UserID              string                  `json:"user_id,omitempty"`
	ClientToken         string                  `json:"client_token,omitempty"`
	TotalAmount         string                  `json:"total_amount"`
	TotalPaidAmount     string                  `json:"total_paid_amount,omitempty"`
	ProcessingMode      string                  `json:"processing_mode"`
	Description         string                  `json:"description,omitempty"`
	Marketplace         string                  `json:"marketplace,omitempty"`
	MarketplaceFee      string                  `json:"marketplace_fee,omitempty"`
	CheckoutAvailableAt string                  `json:"checkout_available_at,omitempty"`
	ExpirationTime      string                  `json:"expiration_time,omitempty"`
	CreatedDate         string                  `json:"created_date,omitempty"`
	LastUpdatedDate     string                  `json:"last_updated_date,omitempty"`
	Transactions        TransactionResponse     `json:"transactions"`
	Items               []ItemsResponse         `json:"items,omitempty"`
	IntegrationData     IntegrationDataResponse `json:"integration_data,omitempty"`
	Config              ConfigResponse          `json:"config,omitempty"`
	Payer               PayerResponse           `json:"payer"`
}

type TransactionResponse struct {
	Payments []PaymentResponse `json:"payments"`
	Refunds  []RefundResponse  `json:"refunds,omitempty"`
}

type PaymentResponse struct {
	ID                string                   `json:"id"`
	ReferenceID       string                   `json:"reference_id"`
	Status            string                   `json:"status"`
	StatusDetail      string                   `json:"status_detail,omitempty"`
	Amount            string                   `json:"amount"`
	PaidAmount        string                   `json:"paid_amount,omitempty"`
	DateOfExpiration  string                   `json:"date_of_expiration,omitempty"`
	ExpirationTime    string                   `json:"expiration_time,omitempty"`
	AttemptNumber     int                      `json:"attempt_number,omitempty"`
	PaymentMethod     PaymentMethodResponse    `json:"payment_method"`
	AutomaticPayments AutomaticPaymentResponse `json:"automatic_payments,omitempty"`
	StoredCredential  StoredCredentialResponse `json:"stored_credential,omitempty"`
	SubscriptionData  SubscriptionDataResponse `json:"subscription_data,omitempty"`
	Attempts          []AttemptResponse        `json:"attempts,omitempty"`
}

type PaymentMethodResponse struct {
	ID                   string `json:"id,omitempty"`
	CardID               string `json:"card_id,omitempty"`
	Type                 string `json:"type,omitempty"`
	Token                string `json:"token,omitempty"`
	StatementDescriptor  string `json:"statement_descriptor,omitempty"`
	Installments         int    `json:"installments,omitempty"`
	TicketURL            string `json:"ticket_url,omitempty"`
	BarcodeContent       string `json:"barcode_content,omitempty"`
	Reference            string `json:"reference,omitempty"`
	ReferenceID          string `json:"reference_id,omitempty"`
	VerificationCode     string `json:"verification_code,omitempty"`
	FinancialInstitution string `json:"financial_institution,omitempty"`
	QrCode               string `json:"qr_code,omitempty"`
	QrCodeBase64         string `json:"qr_code_base64,omitempty"`
	DigitableLine        string `json:"digitable_line,omitempty"`
}

type AutomaticPaymentResponse struct {
	PaymentProfileID string `json:"payment_profile_id"`
	ScheduleDate     string `json:"schedule_date"`
	DueDate          string `json:"due_date"`
	Retries          int    `json:"retries"`
}

type StoredCredentialResponse struct {
	PaymentInitiator   string `json:"payment_initiator"`
	Reason             string `json:"reason"`
	StorePaymentMethod bool   `json:"store_payment_method"`
	FirstPayment       bool   `json:"first_payment"`
}

type SubscriptionDataResponse struct {
	InvoiceID            string                       `json:"invoice_id"`
	BillingDate          string                       `json:"billing_date"`
	SubscriptionSequence SubscriptionSequenceResponse `json:"subscription_sequence"`
	InvoicePeriod        InvoicePeriodResponse        `json:"invoice_period"`
}

type SubscriptionSequenceResponse struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

type InvoicePeriodResponse struct {
	Type   string `json:"type"`
	Period int    `json:"period"`
}

type RefundResponse struct {
	ID            string          `json:"id"`
	TransactionID string          `json:"transaction_id"`
	ReferenceID   string          `json:"reference_id"`
	Status        string          `json:"status"`
	Amount        string          `json:"amount"`
	Items         []ItemsResponse `json:"items,omitempty"`
}

type RefundReferenceResponse struct {
	ID       string `json:"id"`
	SourceID string `json:"source_id"`
}

type ItemsResponse struct {
	Title        string `json:"title"`
	UnitPrice    string `json:"unit_price"`
	ExternalCode string `json:"external_code"`
	Description  string `json:"description"`
	CategoryID   string `json:"category_id"`
	PictureURL   string `json:"picture_url"`
	Quantity     int    `json:"quantity"`
}

type IntegrationDataResponse struct {
	CorporationID string          `json:"corporation_id"`
	ApplicationID string          `json:"application_id"`
	IntegratorID  string          `json:"integrator_id"`
	PlatformID    string          `json:"platform_id"`
	Sponsor       SponsorResponse `json:"sponsor"`
}

type SponsorResponse struct {
	ID string `json:"id"`
}

type ConfigResponse struct {
	PaymentMethodResponse PaymentMethodConfigResponse `json:"payment_method,omitempty"`
	Online                OnlineConfigResponse        `json:"online,omitempty"`
}

type PaymentMethodConfigResponse struct {
	NotAllowedIDs       []string `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes     []string `json:"not_allowed_types,omitempty"`
	DefaultID           string   `json:"default_id,omitempty"`
	MaxInstallments     int      `json:"max_installments,omitempty"`
	DefaultInstallments int      `json:"default_installments,omitempty"`
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
	CustomerID string `json:"customer_id"`
}

type AttemptResponse struct {
	ID            string                `json:"id,omitempty"`
	Status        string                `json:"status,omitempty"`
	StatusDetail  string                `json:"status_detail,omitempty"`
	PaymentMethod PaymentMethodResponse `json:"payment_method,omitempty"`
}
