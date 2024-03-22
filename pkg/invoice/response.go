package invoice

import "time"

// Response is the response from the Invoice API.
type Response struct {
	Payment       PaymentResponse `json:"payment"`         // payment's data
	DateCreated   time.Time       `json:"date_created"`    // creation data
	DebitDate     time.Time       `json:"debit_date"`      // debit date
	LastModified  time.Time       `json:"last_modified"`   // last modified date
	NextRetryDate time.Time       `json:"next_retry_date"` // next retry date

	CurrencyID        string  `json:"currency_id"`        // invoice currency
	ExternalReference string  `json:"external_reference"` // external reference sent by integrator
	PaymentMethodID   string  `json:"payment_method_id"`  // used payment method
	PreapprovalID     string  `json:"preapproval_id"`     // invoice's preapproval id
	Reason            string  `json:"reason"`             // invoice reason
	Status            string  `json:"status"`             // invoice status
	Summarized        string  `json:"summarized"`
	Type              string  `json:"type"`               // invoice type
	TransactionAmount float64 `json:"transaction_amount"` // invoice amount
	ID                int     `json:"id"`                 // invoice id
	RetryAttempt      int     `json:"retry_attempt"`      // invoice's retry attempt
}

// PaymentResponse contains information about payment.
type PaymentResponse struct {
	Status       string `json:"status"`        // payment's status
	StatusDetail string `json:"status_detail"` // payment's status detail
	ID           int    `json:"id"`            // payment's id
}
