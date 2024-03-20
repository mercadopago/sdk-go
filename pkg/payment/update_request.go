package payment

// CancelRequest represents a payment cancellation request.
type CancelRequest struct {
	Status string `json:"status,omitempty"` // status sent to update, in our cancel method we always sent as "cancelled"
}

// CaptureRequest represents a payment capture request.
type CaptureRequest struct {
	TransactionAmount float64 `json:"transaction_amount,omitempty"` // transaction amount to be capture, if no value is sent, the total amount will be capture
	Capture           bool    `json:"capture,omitempty"`            // need to be true to capture an authorized payment
}
