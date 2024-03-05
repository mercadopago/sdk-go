package preapprovalplan

// Request represents a request for creating a pre approval plan.
type Request struct {
	AutoRecurring         AutoRecurring         `json:"auto_recurring,omitempty"`
	PaymentMethodsAllowed PaymentMethodsAllowed `json:"payment_methods_allowed,omitempty"`

	BackURL string `json:"back_url,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// AutoRecurring represents the recurrence settings.
type AutoRecurring struct {
	FreeTrial FreeTrial `json:"free_trial,omitempty"`

	CurrencyID             string `json:"currency_id,omitempty"`
	FrequencyType          string `json:"frequency_type,omitempty"`
	TransactionAmount      int    `json:"transaction_amount,omitempty"`
	Frequency              int    `json:"frequency,omitempty"`
	Repetitions            int    `json:"repetitions,omitempty"`
	BillingDay             int    `json:"billing_day,omitempty"`
	BillingDayProportional bool   `json:"billing_day_proportional,omitempty"`
}

// FreeTrial represents the free trial settings.
type FreeTrial struct {
	FrequencyType      string `json:"frequency_type,omitempty"`
	Frequency          int    `json:"frequency,omitempty"`
	FirstInvoiceOffset int    `json:"first_invoice_offset,omitempty"`
}

// PaymentMethodsAllowed represents the Payment Methods enabled at checkout.
type PaymentMethodsAllowed struct {
	PaymentTypes   []PaymentType   `json:"payment_types,omitempty"`
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty"`
}

// PaymentType represents the Payment Types allowed in the payment flow.
type PaymentType struct {
	ID string `json:"id,omitempty"`
}

// PaymentMethod represents the Payment Methods allowed in the payment flow.
type PaymentMethod struct {
	ID string `json:"id,omitempty"`
}
