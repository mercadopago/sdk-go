package customercard

type Request struct {
	Issuer *IssuerRequest `json:"issuer,omitempty"`

	Token           string `json:"token,omitempty"`
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}

type IssuerRequest struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
