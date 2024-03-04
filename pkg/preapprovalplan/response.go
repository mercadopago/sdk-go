package preapprovalplan

import "time"

type Response struct {
	ID                    string                        `json:"id"`
	BackURL               string                        `json:"back_url"`
	AutoReturn            string                        `json:"auto_return"`
	CollectorID           int                           `json:"collector_id"`
	ApplicationID         int                           `json:"application_id"`
	Reason                string                        `json:"reason"`
	Status                string                        `json:"status"`
	DateCreated           *time.Time                    `json:"date_created"`
	LastModified          *time.Time                    `json:"last_modified"`
	InitPoint             string                        `json:"init_point"`
	AutoRecurring         AutoRecurringResponse         `json:"auto_recurring"`
	PaymentMethodsAllowed PaymentMethodsAllowedResponse `json:"payment_methods_allowed"`
}

type AutoRecurringResponse struct {
	Frequency              int               `json:"frequency"`
	FrequencyType          string            `json:"frequency_type"`
	TransactionAmount      float64           `json:"transaction_amount"`
	CurrencyID             string            `json:"currency_id"`
	Repetitions            int               `json:"repetitions"`
	BillingDay             int               `json:"billing_day"`
	BillingDayProportional bool              `json:"billing_day_proportional"`
	FreeTrial              FreeTrialResponse `json:"free_trial"`
}

type FreeTrialResponse struct {
	Frequency          int    `json:"frequency"`
	FrequencyType      string `json:"frequency_type"`
	FirstInvoiceOffset int    `json:"first_invoice_offset"`
}

type PaymentMethodsAllowedResponse struct {
	PaymentTypes   []PaymentTypeResponse   `json:"payment_types"`
	PaymentMethods []PaymentMethodResponse `json:"payment_methods"`
}

type PaymentTypeResponse struct {
	ID string `json:"id"`
}

type PaymentMethodResponse struct {
	ID string `json:"id"`
}
