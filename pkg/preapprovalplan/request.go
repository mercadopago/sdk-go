package preapprovalplan

// Request represents the body sent to create or update a pre-approval plan (subscription
// template) in the MercadoPago Subscriptions API. Fields tagged with "omitempty" are optional
// and will be omitted from the JSON payload when their zero value is set.
//
// Use this struct with [Client.Create] and [Client.Update].
type Request struct {
	// AutoRecurring defines the recurring billing configuration for the plan, including
	// frequency, amount, currency, and optional free trial settings.
	AutoRecurring *AutoRecurringRequest `json:"auto_recurring,omitempty"`
	// PaymentMethodsAllowed restricts which payment types and methods are available to
	// payers subscribing to this plan.
	PaymentMethodsAllowed *PaymentMethodsAllowedRequest `json:"payment_methods_allowed,omitempty"`

	// BackURL is the URL the payer is redirected to after completing the subscription flow.
	BackURL string `json:"back_url,omitempty"`
	// Reason is a short description of the plan, displayed to payers during checkout.
	Reason string `json:"reason,omitempty"`
}

// AutoRecurringRequest represents the recurring billing configuration for a pre-approval
// plan. It defines how often and how much payers are charged when they subscribe.
type AutoRecurringRequest struct {
	// FreeTrial configures an optional free trial period before billing begins.
	FreeTrial *FreeTrialRequest `json:"free_trial,omitempty"`

	// CurrencyID is the ISO 4217 currency code for recurring charges (e.g., "ARS", "BRL").
	CurrencyID string `json:"currency_id,omitempty"`
	// FrequencyType is the time unit for the billing cycle (e.g., "months", "days").
	FrequencyType string `json:"frequency_type,omitempty"`
	// TransactionAmount is the amount charged in each billing cycle.
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	// Frequency is the number of FrequencyType units between each charge (e.g., 1 for monthly).
	Frequency int `json:"frequency,omitempty"`
	// Repetitions is the total number of billing cycles before the subscription ends.
	// A value of 0 means the subscription recurs indefinitely.
	Repetitions int `json:"repetitions,omitempty"`
	// BillingDay is the day of the month on which the payer is charged (1-28).
	BillingDay int `json:"billing_day,omitempty"`
	// BillingDayProportional indicates whether the first charge should be prorated based
	// on the remaining days until the next BillingDay.
	BillingDayProportional bool `json:"billing_day_proportional,omitempty"`
}

// FreeTrialRequest represents the free trial period configuration for a subscription plan.
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

// PaymentMethodsAllowedRequest specifies which payment types and methods are accepted for
// subscriptions created under this plan. This allows sellers to restrict the checkout
// experience to specific payment instruments.
type PaymentMethodsAllowedRequest struct {
	// PaymentTypes is the list of accepted payment type categories (e.g., credit_card, debit_card).
	PaymentTypes []PaymentTypeRequest `json:"payment_types,omitempty"`
	// PaymentMethods is the list of specific payment methods accepted (e.g., visa, master).
	PaymentMethods []PaymentMethodRequest `json:"payment_methods,omitempty"`
}

// PaymentTypeRequest identifies a payment type category allowed in the subscription plan.
type PaymentTypeRequest struct {
	// ID is the payment type identifier (e.g., "credit_card", "debit_card", "ticket").
	ID string `json:"id,omitempty"`
}

// PaymentMethodRequest identifies a specific payment method allowed in the subscription plan.
type PaymentMethodRequest struct {
	// ID is the payment method identifier (e.g., "visa", "master", "amex").
	ID string `json:"id,omitempty"`
}
