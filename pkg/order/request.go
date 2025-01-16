package order

type Request struct {
	Type              string             `json:"type"`
	TotalAmount       string             `json:"total_amount"`
	ExternalReference string             `json:"external_reference"`
	CaptureMode       string             `json:"capture_mode,omitempty"`
	ProcessingMode    string             `json:"processing_mode,omitempty"`
	Description       string             `json:"description,omitempty"`
	Marketplace       string             `json:"marketplace,omitempty"`
	MarketPlaceFee    string             `json:"marketplace_fee,omitempty"`
	ExpirationTime    string             `json:"expiration_time,omitempty"`
	Transactions      TransactionRequest `json:"transactions"`
	Payer             PayerRequest       `json:"payer"`
	Items             []ItemsRequest     `json:"items,omitempty"`
}

type TransactionRequest struct {
	Payments []PaymentRequest `json:"payments"`
}

type PaymentRequest struct {
	Amount            string                   `json:"amount"`
	PaymentMethod     PaymentMethodRequest     `json:"payment_method"`
	AutomaticPayments *AutomaticPaymentRequest `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredentialRequest `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionDataRequest `json:"subscription_data,omitempty"`
}

type PaymentMethodRequest struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Token               string `json:"token"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments"`
}

type AutomaticPaymentRequest struct {
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
	InvoiceID            string                      `json:"invoice_id"`
	BillingDate          string                      `json:"billing_date"`
	SubscriptionSequence SubscriptionSequenceRequest `json:"subscription_sequence"`
	InvoicePeriod        InvoicePeriodRequest        `json:"invoice_period"`
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
	CustomerID     *string                `json:"customer_id,omitempty"`
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
}

type ItemsRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UnitPrice   string `json:"unit_price"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
	Type        string `json:"type"`
	PictureUrl  string `json:"picture_url"`
	Quantity    int    `json:"quantity"`
}
