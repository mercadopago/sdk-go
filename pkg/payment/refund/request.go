package refund

// Request represents a request for creating a refund.
type Request struct {
	Amount float64 `json:"amount,omitempty"`
}
