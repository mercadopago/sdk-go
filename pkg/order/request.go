package order

// API version: d0494f1c-8d81-4c76-ae1d-0c65bb8ef6de

type Request struct {
	Type                string              `json:"type"`
	TotalAmount         string              `json:"total_amount"`
	ExternalReference   string              `json:"external_reference"`
	CaptureMode         string              `json:"capture_mode,omitempty"`
	ProcessingMode      string              `json:"processing_mode,omitempty"`
	Description         string              `json:"description,omitempty"`
	Marketplace         string              `json:"marketplace,omitempty"`
	MarketPlaceFee      string              `json:"marketplace_fee,omitempty"`
	ExpirationTime      string              `json:"expiration_time,omitempty"`
	CheckoutAvailableAt string              `json:"checkout_available_at,omitempty"`
	Transactions        *TransactionRequest `json:"transactions,omitempty"`
	Payer               *PayerRequest       `json:"payer"`
	Items               []ItemsRequest      `json:"items,omitempty"`
	Config              *ConfigRequest      `json:"config,omitempty"`
}

type TransactionRequest struct {
	Payments []PaymentRequest `json:"payments"`
}

type PaymentRequest struct {
	Amount            string                    `json:"amount,omitempty"`
	ExpirationTime    string                    `json:"expiration_time,omitempty"`
	PaymentMethod     *PaymentMethodRequest     `json:"payment_method,omitempty"`
	AutomaticPayments *AutomaticPaymentsRequest `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredentialRequest  `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionDataRequest  `json:"subscription_data,omitempty"`
}

type PaymentMethodRequest struct {
	ID                  string `json:"id,omitempty"`
	Type                string `json:"type,omitempty"`
	Token               string `json:"token,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments,omitempty"`
}

type AutomaticPaymentsRequest struct {
	PaymentProfileID string `json:"payment_profile_id"`
	ScheduleDate     string `json:"schedule_date"`
	DueDate          string `json:"due_date"`
	Retries          int    `json:"retries"`
}

type StoredCredentialRequest struct {
	PaymentInitiator   string `json:"payment_initiator"`
	Reason             string `json:"reason"`
	StorePaymentMethod bool   `json:"store_payment_method"`
	FirstPayment       bool   `json:"first_payment"`
}

type SubscriptionDataRequest struct {
	InvoiceID            string                       `json:"invoice_id"`
	BillingDate          string                       `json:"billing_date"`
	SubscriptionSequence *SubscriptionSequenceRequest `json:"subscription_sequence,omitempty"`
	InvoicePeriod        *InvoicePeriodRequest        `json:"invoice_period,omitempty"`
}

type SubscriptionSequenceRequest struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

type InvoicePeriodRequest struct {
	Type   string `json:"type"`
	Period int    `json:"period"`
}

type PayerRequest struct {
	Email          string                 `json:"email"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	CustomerID     string                 `json:"customer_id,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`
	Address        *AddressRequest        `json:"address,omitempty"`
}

type IdentificationRequest struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type PhoneRequest struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type AddressRequest struct {
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	ZipCode      string `json:"zip_code"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state"`
	City         string `json:"city"`
	Complement   string `json:"complement"`
	Floor        string `json:"floor"`
}

type ItemsRequest struct {
	Title        string `json:"title"`
	UnitPrice    string `json:"unit_price"`
	ExternalCode string `json:"external_code"`
	CategoryID   string `json:"category_id"`
	Description  string `json:"description"`
	PictureURL   string `json:"picture_url"`
	Quantity     int    `json:"quantity"`
}

type RefundRequest struct {
	Transactions []RefundTransaction `json:"transactions"`
}

type RefundTransaction struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

type ConfigRequest struct {
	PaymentMethod *PaymentMethodConfigRequest `json:"payment_method,omitempty"`
	Online        *OnlineConfigRequest        `json:"online,omitempty"`
}

type PaymentMethodConfigRequest struct {
	NotAllowedIDs       []string `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes     []string `json:"not_allowed_types,omitempty"`
	DefaultID           string   `json:"default_id,omitempty"`
	MaxInstallments     int      `json:"max_installments,omitempty"`
	DefaultInstallments int      `json:"default_installments,omitempty"`
}

type OnlineConfigRequest struct {
	CallbackURL         string                      `json:"callback_url,omitempty"`
	SuccessURL          string                      `json:"success_url,omitempty"`
	PendingURL          string                      `json:"pending_url,omitempty"`
	FailureURL          string                      `json:"failure_url,omitempty"`
	AutoReturnURL       string                      `json:"auto_return_url,omitempty"`
	DifferentialPricing *DifferentialPricingRequest `json:"differential_pricing,omitempty"`
}

type DifferentialPricingRequest struct {
	ID int `json:"id,omitempty"`
}
