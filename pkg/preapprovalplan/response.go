package preapprovalplan

import "time"

// Response represents the MercadoPago API response for a pre-approval plan (subscription
// template) resource. It is returned by [Client.Create], [Client.Get], and [Client.Update],
// and is also used as the element type inside [SearchResponse.Results].
type Response struct {
	// AutoRecurring contains the recurring billing configuration for this plan.
	AutoRecurring AutoRecurringResponse `json:"auto_recurring"`
	// PaymentMethodsAllowed lists the payment types and methods accepted by this plan.
	PaymentMethodsAllowed PaymentMethodsAllowedResponse `json:"payment_methods_allowed"`
	// DateCreated is the timestamp when the plan was created.
	DateCreated time.Time `json:"date_created"`
	// LastModified is the timestamp of the last modification to the plan.
	LastModified time.Time `json:"last_modified"`

	// ID is the unique identifier of the plan assigned by MercadoPago.
	ID string `json:"id"`
	// BackURL is the URL the payer is redirected to after completing the subscription flow.
	BackURL string `json:"back_url"`
	// AutoReturn indicates the auto-return behavior after checkout (e.g., "approved").
	AutoReturn string `json:"auto_return"`
	// Reason is the description of the plan shown to payers.
	Reason string `json:"reason"`
	// Status is the current state of the plan (e.g., "active", "inactive").
	Status string `json:"status"`
	// InitPoint is the URL to redirect payers to subscribe to this plan.
	InitPoint string `json:"init_point"`
	// CollectorID is the unique numeric identifier of the seller (collector) who owns the plan.
	CollectorID int64 `json:"collector_id"`
	// ApplicationID is the identifier of the MercadoPago application that created the plan.
	ApplicationID int `json:"application_id"`
}

// AutoRecurringResponse represents the recurring billing configuration returned by the
// MercadoPago API for a pre-approval plan. It mirrors the settings defined in
// [AutoRecurringRequest] but reflects the server-side state.
type AutoRecurringResponse struct {
	// FreeTrial contains the free trial configuration, if one was set.
	FreeTrial FreeTrialResponse `json:"free_trial"`

	// CurrencyID is the ISO 4217 currency code (e.g., "ARS", "BRL").
	CurrencyID string `json:"currency_id"`
	// FrequencyType is the time unit for the billing cycle (e.g., "months", "days").
	FrequencyType string `json:"frequency_type"`
	// TransactionAmount is the amount charged per billing cycle.
	TransactionAmount float64 `json:"transaction_amount"`
	// Frequency is the number of FrequencyType units between each charge.
	Frequency int `json:"frequency"`
	// Repetitions is the total number of billing cycles before the subscription ends.
	// A value of 0 means indefinite recurrence.
	Repetitions int `json:"repetitions"`
	// BillingDay is the day of the month on which the payer is charged (1-28).
	BillingDay int `json:"billing_day"`
	// BillingDayProportional indicates whether the first charge was prorated.
	BillingDayProportional bool `json:"billing_day_proportional"`
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

// PaymentMethodsAllowedResponse represents the payment types and methods enabled for
// payers subscribing to this plan. It reflects the configuration set in
// [PaymentMethodsAllowedRequest].
type PaymentMethodsAllowedResponse struct {
	// PaymentTypes is the list of accepted payment type categories.
	PaymentTypes []PaymentTypeResponse `json:"payment_types"`
	// PaymentMethods is the list of specific payment methods accepted.
	PaymentMethods []PaymentMethodResponse `json:"payment_methods"`
}

// PaymentTypeResponse identifies a payment type category allowed in the subscription plan.
type PaymentTypeResponse struct {
	// ID is the payment type identifier (e.g., "credit_card", "debit_card", "ticket").
	ID string `json:"id"`
}

// PaymentMethodResponse identifies a specific payment method allowed in the subscription plan.
type PaymentMethodResponse struct {
	// ID is the payment method identifier (e.g., "visa", "master", "amex").
	ID string `json:"id"`
}
