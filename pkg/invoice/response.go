package invoice

import "time"

type Response struct {
	CurrencyID        string          `json:"currency_id"`
	DateCreated       *time.Time      `json:"date_created"`
	DebitDate         *time.Time      `json:"debit_date"`
	ExternalReference string          `json:"external_reference"`
	ID                int             `json:"id"`
	LastModified      *time.Time      `json:"last_modified"`
	NextRetryDate     *time.Time      `json:"next_retry_date"`
	Payment           PaymentResponse `json:"payment"`
	PaymentMethodID   string          `json:"payment_method_id"`
	PreApprovalID     string          `json:"preapproval_id"`
	Reason            string          `json:"reason"`
	RetryAttempt      int             `json:"retry_attempt"`
	Status            string          `json:"status"`
	Summarized        string          `json:"summarized"`
	TransactionAmount float64         `json:"transaction_amount"`
	Type              string          `json:"type"`
}

// PaymentResponse contains information about payment.
type PaymentResponse struct {
	ID           int    `json:"id"`
	Status       string `json:"status"`
	StatusDetail string `json:"status_detail"`
}
