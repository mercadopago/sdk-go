package preapproval

import "time"

// Request represents a request for creating a pre approval.
type Request struct {
	AutoRecurring *AutoRecurringRequest `json:"auto_recurring,omitempty"`

	CardTokenID       string `json:"card_token_id,omitempty"`
	PreapprovalPlanID string `json:"preapproval_plan_id,omitempty"`
	PayerEmail        string `json:"payer_email,omitempty"`
	BackURL           string `json:"back_url,omitempty"`
	CollectorID       string `json:"collector_id,omitempty"`
	Reason            string `json:"reason,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	Status            string `json:"status,omitempty"`
}

// AutoRecurringRequest represents the recurrence settings.
type AutoRecurringRequest struct {
	StartDate *time.Time        `json:"start_date,omitempty"`
	EndDate   *time.Time        `json:"end_date,omitempty"`
	FreeTrial *FreeTrialRequest `json:"free_trial,omitempty"`

	CurrencyID        string  `json:"currency_id,omitempty"`
	FrequencyType     string  `json:"frequency_type,omitempty"`
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	Frequency         int     `json:"frequency,omitempty"`
}

// FreeTrialRequest represents the free trial settings.
type FreeTrialRequest struct {
	FrequencyType      string `json:"frequency_type,omitempty"`
	Frequency          int    `json:"frequency,omitempty"`
	FirstInvoiceOffset int    `json:"first_invoice_offset,omitempty"`
}
