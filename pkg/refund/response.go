package refund

import "time"

// Response represents a refund resource returned by the MercadoPago Refunds API.
// It is returned by [Client.Create], [Client.CreatePartialRefund], [Client.Get],
// and as elements of the slice returned by [Client.List].
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

// SourceResponse identifies the origin of the refund, indicating who or what
// initiated it (e.g., the collector or an automated process).
type SourceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
