package preapproval

import "time"

// Response represents the MercadoPago API response for a pre-approval (subscription) resource.
// It is returned by [Client.Create], [Client.Get], and [Client.Update], and is also used as
// the element type inside [SearchResponse.Results].
type Response struct {
	// AutoRecurring contains the recurring billing configuration for this subscription.
	AutoRecurring AutoRecurringResponse `json:"auto_recurring"`
	// Summarized contains aggregated billing and charge information for this subscription.
	Summarized SummarizedResponse `json:"summarized"`
	// DateCreated is the timestamp when the subscription was created.
	DateCreated time.Time `json:"date_created"`
	// LastModified is the timestamp of the last modification to the subscription.
	LastModified time.Time `json:"last_modified"`
	// NextPaymentDate is the scheduled date for the next recurring charge.
	NextPaymentDate time.Time `json:"next_payment_date"`

	// ID is the unique identifier of the subscription assigned by MercadoPago.
	ID string `json:"id"`
	// PayerEmail is the email address of the subscriber.
	PayerEmail string `json:"payer_email"`
	// Status is the current state of the subscription (e.g., "authorized", "paused", "cancelled").
	Status string `json:"status"`
	// Reason is the description of the subscription shown to the payer.
	Reason string `json:"reason"`
	// ExternalReference is the external reference value used to correlate this subscription
	// with an entity in the integrator's system.
	ExternalReference string `json:"external_reference"`
	// InitPoint is the URL to redirect the payer to complete the subscription in production.
	InitPoint string `json:"init_point"`
	// SandboxInitPoint is the URL to redirect the payer in the sandbox (testing) environment.
	SandboxInitPoint string `json:"sandbox_init_point"`
	// PaymentMethodID identifies the payment method selected by the payer.
	PaymentMethodID string `json:"payment_method_id"`
	// FirstInvoiceOffset is the number of billing cycles to skip before the first charge.
	FirstInvoiceOffset int `json:"first_invoice_offset"`
	// BackURL is the URL the payer is redirected to after completing the subscription flow.
	BackURL string `json:"back_url"`
	// PreapprovalPlanID is the identifier of the associated pre-approval plan, if any.
	PreapprovalPlanID string `json:"preapproval_plan_id"`
	// PayerFirstName is the payer's first name.
	PayerFirstName string `json:"payer_first_name"`
	// PayerLastName is the payer's last name.
	PayerLastName string `json:"payer_last_name"`
	// CardID is the identifier of the card used for recurring payments.
	CardID string `json:"card_id"`
	// Version is the resource version number, incremented on each update.
	Version int `json:"version"`
	// PayerID is the unique numeric identifier of the payer in MercadoPago.
	PayerID int64 `json:"payer_id"`
	// CollectorID is the unique numeric identifier of the seller (collector) in MercadoPago.
	CollectorID int64 `json:"collector_id"`
	// ApplicationID is the identifier of the MercadoPago application that created the subscription.
	ApplicationID int `json:"application_id"`
}

// AutoRecurringResponse represents the recurring billing configuration returned by the
// MercadoPago API for a pre-approval. It mirrors the settings defined in [AutoRecurringRequest]
// but reflects the server-side state.
type AutoRecurringResponse struct {
	// FreeTrial contains the free trial configuration, if one was set.
	FreeTrial FreeTrialResponse `json:"free_trial"`
	// StartDate is the date when recurring charges began or will begin.
	StartDate time.Time `json:"start_date"`
	// EndDate is the date when recurring charges end or will end.
	EndDate time.Time `json:"end_date"`

	// CurrencyID is the ISO 4217 currency code (e.g., "ARS", "BRL").
	CurrencyID string `json:"currency_id"`
	// FrequencyType is the time unit for the billing cycle (e.g., "months", "days").
	FrequencyType string `json:"frequency_type"`
	// TransactionAmount is the amount charged per billing cycle.
	TransactionAmount float64 `json:"transaction_amount"`
	// Frequency is the number of FrequencyType units between each charge.
	Frequency int `json:"frequency"`
}

// FreeTrialResponse represents the free trial period configuration returned by the
// MercadoPago API. During the free trial, the payer is not charged.
type FreeTrialResponse struct {
	// FrequencyType is the time unit for the free trial duration (e.g., "months", "days").
	FrequencyType string `json:"frequency_type"`
	// Frequency is the number of FrequencyType units the free trial lasts.
	Frequency int `json:"frequency"`
	// FirstInvoiceOffset is the number of billing cycles skipped after the trial ends
	// before the first charge.
	FirstInvoiceOffset int `json:"first_invoice_offset"`
}

// SummarizedResponse contains aggregated billing metrics for a subscription, including
// totals for charged and pending amounts. This information is useful for displaying
// subscription health and payment history to integrators.
type SummarizedResponse struct {
	// LastChargedDate is the date of the most recent successful charge.
	LastChargedDate time.Time `json:"last_charged_date"`
	// LastChargedAmount is the amount of the most recent successful charge.
	LastChargedAmount float64 `json:"last_charged_amount"`

	// Semaphore indicates the health status of the subscription billing
	// (e.g., "green", "yellow", "red").
	Semaphore string `json:"semaphore"`
	// PendingChargeAmount is the total monetary amount of charges that are still pending.
	PendingChargeAmount float64 `json:"pending_charge_amount"`
	// ChargedAmount is the total monetary amount successfully charged so far.
	ChargedAmount float64 `json:"charged_amount"`
	// Quotas is the total number of billing cycles (installments) for this subscription.
	Quotas int `json:"quotas"`
	// PendingChargeQuantity is the number of billing cycles that are still pending collection.
	PendingChargeQuantity int `json:"pending_charge_quantity"`
	// ChargedQuantity is the number of billing cycles that have been successfully charged.
	ChargedQuantity int `json:"charged_quantity"`
}
