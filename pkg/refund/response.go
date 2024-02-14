package refund

import (
	"time"
)

// Response is the response from the Refund's API.
type Response struct {
	ID                   int64      `json:"id"`
	PaymentID            int64      `json:"payment_id"`
	Amount               float64    `json:"amount"`
	AdjustmentAmount     float64    `json:"adjustment_amount"`
	Status               string     `json:"status"`
	RefundMode           string     `json:"refund_mode"`
	DateCreated          *time.Time `json:"date_created"`
	Reason               string     `json:"reason"`
	UniqueSequenceNumber string     `json:"unique_sequence_number"`

	Source Source `json:"source"`
}

// Source represents the data to identify who originated the refund
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
