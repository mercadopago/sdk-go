package refund

import "time"

// Response is the response from the Refund's API.
type Response struct {
	Source      SourceResponse `json:"source"`
	DateCreated time.Time      `json:"date_created"`

	Status               string  `json:"status"`
	RefundMode           string  `json:"refund_mode"`
	Reason               string  `json:"reason"`
	UniqueSequenceNumber string  `json:"unique_sequence_number"`
	Amount               float64 `json:"amount"`
	AdjustmentAmount     float64 `json:"adjustment_amount"`
	ID                   int     `json:"id"`
	PaymentID            int     `json:"payment_id"`
}

// SourceResponse represents the data to identify who originated the refund
type SourceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
