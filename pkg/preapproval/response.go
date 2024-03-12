package preapproval

import "time"

// Response represents the response from the pre-approval endpoint.
type Response struct {
	AutoRecurring   AutoRecurringResponse `json:"auto_recurring"`
	Summarized      SummarizedResponse    `json:"summarized"`
	DateCreated     time.Time             `json:"date_created"`
	LastModified    time.Time             `json:"last_modified"`
	NextPaymentDate time.Time             `json:"next_payment_date"`

	ID                 string `json:"id"`
	PayerEmail         string `json:"payer_email"`
	Status             string `json:"status"`
	Reason             string `json:"reason"`
	ExternalReference  string `json:"external_reference"`
	InitPoint          string `json:"init_point"`
	SandboxInitPoint   string `json:"sandbox_init_point"`
	PaymentMethodID    string `json:"payment_method_id"`
	FirstInvoiceOffset string `json:"first_invoice_offset"`
	BackURL            string `json:"back_url"`
	PreapprovalPlanID  string `json:"preapproval_plan_id"`
	CardID             int    `json:"card_id"`
	Version            int    `json:"version"`
	PayerID            int    `json:"payer_id"`
	CollectorID        int    `json:"collector_id"`
	ApplicationID      int    `json:"application_id"`
	PayerFirstName     string `json:"payer_first_name"`
	PayerLastName      string `json:"payer_last_name"`
}

// AutoRecurringResponse represents the recurrence settings.
type AutoRecurringResponse struct {
	FreeTrial FreeTrialResponse `json:"free_trial"`
	StartDate time.Time         `json:"start_date"`
	EndDate   time.Time         `json:"end_date"`

	CurrencyID        string  `json:"currency_id"`
	FrequencyType     string  `json:"frequency_type"`
	Frequency         int     `json:"frequency"`
	TransactionAmount float64 `json:"transaction_amount"`
}

// FreeTrialResponse represents the free trial settings.
type FreeTrialResponse struct {
	FrequencyType      string `json:"frequency_type"`
	Frequency          int    `json:"frequency"`
	FirstInvoiceOffset int    `json:"first_invoice_offset"`
}

// SummarizedResponse contains summary information about invoices and subscription charges.
type SummarizedResponse struct {
	LastChargedDate   time.Time `json:"last_charged_date"`
	LastChargedAmount time.Time `json:"last_charged_amount"`

	Quotas                int     `json:"quotas"`
	PendingChargeQuantity int     `json:"pending_charge_quantity"`
	ChargedQuantity       int     `json:"charged_quantity"`
	PendingChargeAmount   float64 `json:"pending_charge_amount"`
	ChargedAmount         float64 `json:"charged_amount"`
	Semaphore             string  `json:"semaphore"`
}
