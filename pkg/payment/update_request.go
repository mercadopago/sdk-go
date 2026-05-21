package payment

// CancelRequest is the internal request body sent by [Client.Cancel] to transition
// a payment to the "cancelled" status. It is not intended for direct use by callers.
type CancelRequest struct {
	Status string `json:"status,omitempty"`
}

// CaptureRequest is the internal request body sent by [Client.Capture] and
// [Client.CaptureAmount] to confirm an authorized payment. When TransactionAmount
// is set, it performs a partial capture for that specific amount.
type CaptureRequest struct {
	TransactionAmount float64 `json:"transaction_amount,omitempty"`
	Capture           bool    `json:"capture,omitempty"`
}
