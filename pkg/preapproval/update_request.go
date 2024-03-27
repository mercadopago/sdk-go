package preapproval

import "time"

// UpdateRequest represents a request for updating a pre approval.
type UpdateRequest struct {
	AutoRecurring *AutoRecurringUpdateRequest `json:"auto_recurring,omitempty"`

	CardTokenID       string `json:"card_token_id,omitempty"`
	PayerEmail        string `json:"payer_email,omitempty"`
	BackURL           string `json:"back_url,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Status            string `json:"status,omitempty"`
}

// AutoRecurringUpdateRequest represents the recurrence settings.
type AutoRecurringUpdateRequest struct {
	StartDate         *time.Time `json:"start_date,omitempty"`
	EndDate           *time.Time `json:"end_date,omitempty"`
	TransactionAmount float64    `json:"transaction_amount,omitempty"`
}
