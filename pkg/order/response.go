package order

type Response struct {
	ID                string              `json:"id"`
	Type              string              `json:"type"`
	TotalAmount       string              `json:"total_amount"`
	ExternalReference string              `json:"external_reference"`
	CountryCode       string              `json:"country_code"`
	Status            string              `json:"status"`
	StatusDetail      string              `json:"status_detail"`
	CaptureMode       string              `json:"capture_mode"`
	ProcessingMode    string              `json:"processing_mode"`
	Description       string              `json:"description,omitempty"`
	Marketplace       string              `json:"marketplace,omitempty"`
	MarketplaceFee    string              `json:"marketplace_fee,omitempty"`
	ExpirationTime    string              `json:"expiration_time,omitempty"`
	CreatedDate       string              `json:"created_date"`
	LastUpdatedDate   string              `json:"last_updated_date"`
	ClientID          string              `json:"client_id,omitempty"`
	CollectorID       string              `json:"collector_id,omitempty"`
	Transactions      TransactionResponse `json:"transactions"`
	Payer             PayerResponse       `json:"payer"`
	Items             []ItemsResponse     `json:"items,omitempty"`
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
	PaymentMethod     PaymentMethodResponse     `json:"payment_method"`
	AutomaticPayments *AutomaticPaymentResponse `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredentialResponse `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionDataResponse `json:"subscription_data,omitempty"`
}

type PaymentMethodResponse struct {
	ID                  string `json:"id"`
	CardID              string `json:"card_id,omitempty"`
	Type                string `json:"type"`
	Token               string `json:"token"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments"`
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
	Status        string `json:"status"`
	Amount        string `json:"amount"`
}

type RefundReferenceResponse struct {
	ID       string `json:"id"`
	SourceID string `json:"source_id"`
}

type PayerResponse struct {
	Email          string                  `json:"email"`
	FirstName      string                  `json:"first_name,omitempty"`
	LastName       string                  `json:"last_name,omitempty"`
	CustomerID     *string                 `json:"customer_id,omitempty"`
	Identification *IdentificationResponse `json:"identification,omitempty"`
	Phone          *PhoneResponse          `json:"phone,omitempty"`
	Address        *AddressResponse        `json:"address,omitempty"`
}

type IdentificationResponse struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type PhoneResponse struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type AddressResponse struct {
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	ZipCode      string `json:"zip_code"`
}

type ItemsResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UnitPrice   string `json:"unit_price"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
	Type        string `json:"type"`
	PictureUrl  string `json:"picture_url"`
	Quantity    int    `json:"quantity"`
}
