package preapprovalplan

// Request represents a request for creating a pre approval plan.
type Request struct {
	AutoRecurring         *AutoRecurringRequest         `json:"auto_recurring,omitempty"`
	PaymentMethodsAllowed *PaymentMethodsAllowedRequest `json:"payment_methods_allowed,omitempty"`

	BackURL string `json:"back_url,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// AutoRecurringRequest represents the recurrence settings.
type AutoRecurringRequest struct {
	FreeTrial *FreeTrialRequest `json:"free_trial,omitempty"`

	CurrencyID             string  `json:"currency_id,omitempty"`
	FrequencyType          string  `json:"frequency_type,omitempty"`
	TransactionAmount      float64 `json:"transaction_amount,omitempty"`
	Frequency              int     `json:"frequency,omitempty"`
	Repetitions            int     `json:"repetitions,omitempty"`
	BillingDay             int     `json:"billing_day,omitempty"`
	BillingDayProportional bool    `json:"billing_day_proportional,omitempty"`
}

// FreeTrialRequest represents the free trial settings.
type FreeTrialRequest struct {
	FrequencyType      string `json:"frequency_type,omitempty"`
	Frequency          int    `json:"frequency,omitempty"`
	FirstInvoiceOffset int    `json:"first_invoice_offset,omitempty"`
}

// PaymentMethodsAllowedRequest represents the Payment Methods enabled at checkout.
type PaymentMethodsAllowedRequest struct {
	PaymentTypes   []PaymentTypeRequest   `json:"payment_types,omitempty"`
	PaymentMethods []PaymentMethodRequest `json:"payment_methods,omitempty"`
}

// PaymentTypeRequest represents the Payment Types allowed in the payment flow.
type PaymentTypeRequest struct {
	ID string `json:"id,omitempty"`
}

// PaymentMethodRequest represents the Payment Methods allowed in the payment flow.
type PaymentMethodRequest struct {
	ID string `json:"id,omitempty"`
}
