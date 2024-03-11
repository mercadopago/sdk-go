package preapproval

import "time"

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Results []SearchResults `json:"results"`
	Paging  PagingResponse  `json:"paging"`
}

// PagingResponse represents the paging information within SearchResponse.
type PagingResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// SearchResults contains the information of a preapproval.
type SearchResults struct {
	AutoRecurring   AutoRecurringResponse `json:"auto_recurring"`
	Summarized      SummarizedResponse    `json:"summarized"`
	DateCreated     *time.Time            `json:"date_created"`
	LastModified    *time.Time            `json:"last_modified"`
	NextPaymentDate *time.Time            `json:"next_payment_date"`

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
	PayerID            int    `json:"payer_id"`
	CollectorID        int    `json:"collector_id"`
	ApplicationID      int    `json:"application_id"`
	Version            int    `json:"version"`
	PayerFirstName     string `json:"payer_first_name"`
	PayerLastName      string `json:"payer_last_name"`
}
