package invoice

import "time"

// Response is the response from the Invoice API.
type Response struct {
	DateCreated   *time.Time      `json:"date_created"`
	DebitDate     *time.Time      `json:"debit_date"`
	LastModified  *time.Time      `json:"last_modified"`
	NextRetryDate *time.Time      `json:"next_retry_date"`
	Payment       PaymentResponse `json:"payment"`

	CurrencyID        string  `json:"currency_id"`
	ExternalReference string  `json:"external_reference"`
	PaymentMethodID   string  `json:"payment_method_id"`
	PreapprovalID     string  `json:"preapproval_id"`
	Reason            string  `json:"reason"`
	Status            string  `json:"status"`
	Summarized        string  `json:"summarized"`
	Type              string  `json:"type"`
	ID                int     `json:"id"`
	RetryAttempt      int     `json:"retry_attempt"`
	TransactionAmount float64 `json:"transaction_amount"`
}

// PaymentResponse contains information about payment.
type PaymentResponse struct {
	Status       string `json:"status"`
	StatusDetail string `json:"status_detail"`
	ID           int    `json:"id"`
}
