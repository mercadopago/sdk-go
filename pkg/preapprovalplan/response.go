package preapprovalplan

import "time"

// Response represents the response from the pre-approval plan endpoint.
type Response struct {
	AutoRecurring         AutoRecurringResponse         `json:"auto_recurring"`
	PaymentMethodsAllowed PaymentMethodsAllowedResponse `json:"payment_methods_allowed"`
	DateCreated           time.Time                     `json:"date_created"`
	LastModified          time.Time                     `json:"last_modified"`

	ID            string `json:"id"`
	BackURL       string `json:"back_url"`
	AutoReturn    string `json:"auto_return"`
	Reason        string `json:"reason"`
	Status        string `json:"status"`
	InitPoint     string `json:"init_point"`
	CollectorID   int    `json:"collector_id"`
	ApplicationID int    `json:"application_id"`
}

// AutoRecurringResponse represents the recurrence settings.
type AutoRecurringResponse struct {
	FreeTrial FreeTrialResponse `json:"free_trial"`

	CurrencyID             string  `json:"currency_id"`
	FrequencyType          string  `json:"frequency_type"`
	Frequency              int     `json:"frequency"`
	Repetitions            int     `json:"repetitions"`
	BillingDay             int     `json:"billing_day"`
	TransactionAmount      float64 `json:"transaction_amount"`
	BillingDayProportional bool    `json:"billing_day_proportional"`
}

// FreeTrialResponse represents the free trial settings.
type FreeTrialResponse struct {
	FrequencyType      string `json:"frequency_type"`
	Frequency          int    `json:"frequency"`
	FirstInvoiceOffset int    `json:"first_invoice_offset"`
}

// PaymentMethodsAllowedResponse represents the Payment Methods enabled at checkout.
type PaymentMethodsAllowedResponse struct {
	PaymentTypes   []PaymentTypeResponse   `json:"payment_types"`
	PaymentMethods []PaymentMethodResponse `json:"payment_methods"`
}

// PaymentTypeResponse represents the Payment Types allowed in the payment flow.
type PaymentTypeResponse struct {
	ID string `json:"id"`
}

// PaymentMethodResponse represents the Payment Methods allowed in the payment flow.
type PaymentMethodResponse struct {
	ID string `json:"id"`
}
