package invoice

import "time"

// Response represents an invoice (authorized payment) resource returned by the MercadoPago
// Authorized Payments API. It is returned by [Client.Get] and as elements within [SearchResponse].
//
// Each invoice corresponds to a billing cycle of a subscription and includes the payment
// amount, status, retry information, and a reference to the associated subscription (preapproval).
type Response struct {
	// Payment contains the status and identifier of the payment associated with this invoice.
	Payment PaymentResponse `json:"payment"`

	// DateCreated is the date when the invoice was created.
	DateCreated time.Time `json:"date_created"`

	// DebitDate is the scheduled date for the payment debit.
	DebitDate time.Time `json:"debit_date"`

	// LastModified is the date when the invoice was last modified.
	LastModified time.Time `json:"last_modified"`

	// NextRetryDate is the date of the next automatic retry if the payment failed.
	NextRetryDate time.Time `json:"next_retry_date"`

	// CurrencyID is the ISO 4217 currency code for the invoice amount (e.g., "BRL", "ARS").
	CurrencyID string `json:"currency_id"`

	// ExternalReference is the integrator-provided external identifier for reconciliation.
	ExternalReference string `json:"external_reference"`

	// PaymentMethodID is the identifier of the payment method used for this invoice.
	PaymentMethodID string `json:"payment_method_id"`

	// PreapprovalID is the subscription (preapproval) identifier that generated this invoice.
	PreapprovalID string `json:"preapproval_id"`

	// Reason is the description or reason for the invoice charge.
	Reason string `json:"reason"`

	// Status is the invoice status (e.g., "scheduled", "processed", "recycling", "cancelled").
	Status string `json:"status"`

	// Summarized provides a summary description of the invoice status.
	Summarized string `json:"summarized"`

	// Type is the invoice type.
	Type string `json:"type"`

	// TransactionAmount is the amount to be charged for this invoice.
	TransactionAmount float64 `json:"transaction_amount"`

	// ID is the unique invoice identifier assigned by MercadoPago.
	ID int `json:"id"`

	// RetryAttempt is the current retry attempt number if the payment failed and is being retried.
	RetryAttempt int `json:"retry_attempt"`
}

// PaymentResponse contains the payment status and identifier associated with an invoice,
// as returned within a [Response].
type PaymentResponse struct {
	// Status is the payment status (e.g., "approved", "rejected", "pending").
	Status string `json:"status"`

	// StatusDetail provides additional detail about the payment status (e.g., "accredited", "cc_rejected_other_reason").
	StatusDetail string `json:"status_detail"`

	// ID is the unique payment identifier assigned by MercadoPago.
	ID int `json:"id"`
}
