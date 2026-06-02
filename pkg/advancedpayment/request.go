package advancedpayment

// Request represents the body sent to create an advanced payment.
type Request struct {
	Disbursements []DisbursementRequest `json:"disbursements,omitempty"`
	Payments      []PaymentRequest      `json:"payments,omitempty"`
	Payer         *PayerRequest         `json:"payer,omitempty"`

	ApplicationID     string `json:"application_id,omitempty"`
	Description       string `json:"description,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`

	BinaryMode bool `json:"binary_mode,omitempty"`
	Capture    bool `json:"capture,omitempty"`
}

// DisbursementRequest defines a single receiver (collector) and amount in a create request.
type DisbursementRequest struct {
	ExternalReference string  `json:"external_reference,omitempty"`
	MoneyReleaseDate  string  `json:"money_release_date,omitempty"`
	ApplicationFee    float64 `json:"application_fee,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	CollectorID       int     `json:"collector_id,omitempty"`
}

// PaymentRequest defines a single payment source within an advanced payment.
type PaymentRequest struct {
	DateOfExpiration    string  `json:"date_of_expiration,omitempty"`
	Description         string  `json:"description,omitempty"`
	ExternalReference   string  `json:"external_reference,omitempty"`
	PaymentMethodID     string  `json:"payment_method_id,omitempty"`
	PaymentTypeID       string  `json:"payment_type_id,omitempty"`
	ProcessingMode      string  `json:"processing_mode,omitempty"`
	StatementDescriptor string  `json:"statement_descriptor,omitempty"`
	Token               string  `json:"token,omitempty"`
	TransactionAmount   float64 `json:"transaction_amount,omitempty"`
	Installments        int     `json:"installments,omitempty"`
}

// PayerRequest contains buyer information for an advanced payment.
type PayerRequest struct {
	Email string `json:"email,omitempty"`
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
}

// UpdateRequest is the body sent when updating an advanced payment.
type UpdateRequest struct {
	Status  string `json:"status,omitempty"`
	Capture *bool  `json:"capture,omitempty"`
}

// CancelRequest sets the status to "cancelled".
type CancelRequest struct {
	Status string `json:"status"`
}

// CaptureRequest sets capture to true to finalise a two-step payment.
type CaptureRequest struct {
	Capture bool `json:"capture"`
}

// UpdateReleaseDateRequest sets a new money release date for all disbursements.
type UpdateReleaseDateRequest struct {
	MoneyReleaseDate string `json:"money_release_date"`
}
