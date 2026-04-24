package preapproval

import "time"

// Request represents the body sent to create a new pre-approval (subscription) in the
// MercadoPago Subscriptions API. Fields tagged with "omitempty" are optional and will be
// omitted from the JSON payload when their zero value is set.
//
// Use this struct with [Client.Create].
type Request struct {
	// AutoRecurring defines the recurrence configuration for the subscription, including
	// billing frequency, currency, amount, and optional free trial settings.
	AutoRecurring *AutoRecurringRequest `json:"auto_recurring,omitempty"`

	// CardTokenID is the token that identifies the payer's card, obtained through the
	// MercadoPago card token API.
	CardTokenID string `json:"card_token_id,omitempty"`
	// PreapprovalPlanID links this subscription to an existing pre-approval plan template.
	// When set, plan-level settings are inherited.
	PreapprovalPlanID string `json:"preapproval_plan_id,omitempty"`
	// PayerEmail is the email address of the subscriber.
	PayerEmail string `json:"payer_email,omitempty"`
	// BackURL is the URL the payer is redirected to after completing the subscription flow.
	BackURL string `json:"back_url,omitempty"`
	// CollectorID is the identifier of the seller (collector) who owns this subscription.
	CollectorID string `json:"collector_id,omitempty"`
	// Reason is a short description of the subscription, displayed to the payer.
	Reason string `json:"reason,omitempty"`
	// ExternalReference is a reference value that can be used to synchronize the subscription
	// with an entity in an external system (e.g., an order or membership ID).
	ExternalReference string `json:"external_reference,omitempty"`
	// Status is the desired initial status of the subscription (e.g., "pending", "authorized").
	Status string `json:"status,omitempty"`
}

// AutoRecurringRequest represents the recurring billing configuration sent when creating
// a pre-approval. It defines how often and how much the payer is charged.
type AutoRecurringRequest struct {
	// FreeTrial configures an optional free trial period before billing begins.
	FreeTrial *FreeTrialRequest `json:"free_trial,omitempty"`
	// StartDate is the date when the subscription's recurring charges begin.
	StartDate *time.Time `json:"start_date,omitempty"`
	// EndDate is the date when the subscription's recurring charges end.
	EndDate *time.Time `json:"end_date,omitempty"`

	// CurrencyID is the ISO 4217 currency code for the recurring charges (e.g., "ARS", "BRL").
	CurrencyID string `json:"currency_id,omitempty"`
	// FrequencyType is the time unit for the billing cycle (e.g., "months", "days").
	FrequencyType string `json:"frequency_type,omitempty"`
	// TransactionAmount is the amount charged in each billing cycle.
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	// Frequency is the number of FrequencyType units between each charge (e.g., 1 for monthly).
	Frequency int `json:"frequency,omitempty"`
}

// FreeTrialRequest represents the free trial period configuration for a subscription.
// During the free trial, the payer is not charged.
type FreeTrialRequest struct {
	// FrequencyType is the time unit for the free trial duration (e.g., "months", "days").
	FrequencyType string `json:"frequency_type,omitempty"`
	// Frequency is the number of FrequencyType units the free trial lasts.
	Frequency int `json:"frequency,omitempty"`
	// FirstInvoiceOffset is the number of billing cycles to skip before the first charge
	// after the free trial ends.
	FirstInvoiceOffset int `json:"first_invoice_offset,omitempty"`
}
