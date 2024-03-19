package customercard

// Request is the helper structure to build request.
type Request struct {
	Issuer *IssuerRequest `json:"issuer,omitempty"`

	Token           string `json:"token,omitempty"`
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}

// IssuerRequest is the helper structure to build issuer request.
type IssuerRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
