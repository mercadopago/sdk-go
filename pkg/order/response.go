package order

// API version: 5d077b6f-61b2-4b3a-8333-7a64ee547448
type Response struct {
	ID                  string              `json:"id"`
	Type                string              `json:"type"`
	ExternalReference   string              `json:"external_reference"`
	CountryCode         string              `json:"country_code"`
	Status              string              `json:"status"`
	StatusDetail        string              `json:"status_detail"`
	CaptureMode         string              `json:"capture_mode"`
	UserID              string              `json:"user_id,omitempty"`
	ClientToken         string              `json:"client_token,omitempty"`
	TotalAmount         string              `json:"total_amount"`
	ProcessingMode      string              `json:"processing_mode"`
	Description         string              `json:"description,omitempty"`
	Marketplace         string              `json:"marketplace,omitempty"`
	MarketplaceFee      string              `json:"marketplace_fee,omitempty"`
	CheckoutAvailableAt string              `json:"checkout_available_at,omitempty"`
	ExpirationTime      string              `json:"expiration_time,omitempty"`
	Transactions        TransactionResponse `json:"transactions"`
	Items               []ItemsResponse     `json:"items,omitempty"`
	IntegrationData     IntegrationData     `json:"integration_data,omitempty"`
	Config              ConfigResponse      `json:"config,omitempty"`
}

type TransactionResponse struct {
	Payments []PaymentResponse `json:"payments"`
	Refunds  []RefundResponse  `json:"refunds,omitempty"`
}

type PaymentResponse struct {
	ID                string                    `json:"id"`
	ReferenceID       string                    `json:"reference_id"`
	Status            string                    `json:"status"`
	StatusDetail      string                    `json:"status_detail,omitempty"`
	Amount            string                    `json:"amount"`
	PaymentMethod     *PaymentMethodResponse    `json:"payment_method,omitempty"`
	AutomaticPayments *AutomaticPaymentResponse `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredentialResponse `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionDataResponse `json:"subscription_data,omitempty"`
}

type PaymentMethodResponse struct {
	ID                   string `json:"id,omitempty"`
	CardID               string `json:"card_id,omitempty"`
	Type                 string `json:"type,omitempty"`
	Token                string `json:"token,omitempty"`
	StatementDescriptor  string `json:"statement_descriptor,omitempty"`
	Installments         int    `json:"installments,omitempty"`
	TicketUrl            string `json:"ticket_url,omitempty"`
	BarcodeContent       string `json:"barcode_content,omitempty"`
	Reference            string `json:"reference,omitempty"`
	VerificationCode     string `json:"verification_code,omitempty"`
	FinancialInstitution string `json:"financial_institution,omitempty"`
	QrCode               string `json:"qr_code,omitempty"`
	QrCodeBase64         string `json:"qr_code_base64,omitempty"`
	DigitableLine        string `json:"digitable_line,omitempty"`
	NotAllowedIds        string `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes      string `json:"not_allowed_types,omitempty"`
	DefaultId            string `json:"default_id,omitempty"`
	MaxInstallments      int    `json:"max_installments,omitempty"`
	DefaultInstallments  int    `json:"default_installments,omitempty"`
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
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	ReferenceId   string `json:"reference_id"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
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
	PictureUrl   string `json:"picture_url"`
	Quantity     int    `json:"quantity"`
}

type IntegrationData struct {
	CorporationId string  `json:"corporation_id"`
	ApplicationId string  `json:"application_id"`
	IntegratorId  string  `json:"integrator_id"`
	PlatformId    string  `json:"platform_id"`
	Sponsor       Sponsor `json:"sponsor"`
}

type Sponsor struct {
	ID string `json:"id"`
}
type ConfigResponse struct {
	PaymentMethodResponse PaymentMethodResponse `json:"payment_method"`
	Online                Online                `json:"online"`
}
