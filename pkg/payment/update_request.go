package payment

// CancelRequest represents a payment cancellation request.
type CancelRequest struct {
	Status string `json:"status,omitempty"`
}

// CaptureRequest represents a payment capture request.
type CaptureRequest struct {
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	Capture           bool    `json:"capture"`
}
