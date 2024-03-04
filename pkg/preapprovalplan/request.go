package preapprovalplan

// Request represents a request for creating a pre approval plan.
type Request struct {
	BackURL               string                `json:"back_url,omitempty"`
	Reason                string                `json:"reason,omitempty"`
	AutoRecurring         AutoRecurring         `json:"auto_recurring,omitempty"`
	PaymentMethodsAllowed PaymentMethodsAllowed `json:"payment_methods_allowed,omitempty"`
}

type AutoRecurring struct {
	Frequency              int       `json:"frequency,omitempty"`
	FrequencyType          string    `json:"frequency_type,omitempty"`
	TransactionAmount      int       `json:"transaction_amount,omitempty"`
	CurrencyID             string    `json:"currency_id,omitempty"`
	Repetitions            int       `json:"repetitions,omitempty"`
	BillingDay             int       `json:"billing_day,omitempty"`
	BillingDayProportional bool      `json:"billing_day_proportional,omitempty"`
	FreeTrial              FreeTrial `json:"free_trial,omitempty"`
}

type FreeTrial struct {
	Frequency          int    `json:"frequency"`
	FrequencyType      string `json:"frequency_type"`
	FirstInvoiceOffset int    `json:"first_invoice_offset,omitempty"`
}

type PaymentMethodsAllowed struct {
	PaymentTypes   []PaymentType   `json:"payment_types,omitempty"`
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty"`
}

type PaymentType struct {
	ID string `json:"id,omitempty"`
}

type PaymentMethod struct {
	ID string `json:"id,omitempty"`
}
