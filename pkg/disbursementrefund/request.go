package disbursementrefund

// Request is the body sent when creating a disbursement refund.
type Request struct {
	// Amount is the amount to refund. When zero, the full disbursement amount is refunded.
	Amount float64 `json:"amount,omitempty"`
}
