package refund

// Request represents the body sent to the MercadoPago Refunds API when creating a partial refund.
// For full refunds the body is empty and this struct is not used. The Amount field specifies the
// portion of the original transaction amount to be refunded.
type Request struct {
	Amount float64 `json:"amount,omitempty"`
}
