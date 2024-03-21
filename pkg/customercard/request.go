package customercard

type Request struct {
	Issuer *IssuerRequest `json:"issuer,omitempty"` // issuer data

	Token           string `json:"token,omitempty"`             // card token
	PaymentMethodID string `json:"payment_method_id,omitempty"` // card's payment method
}

type IssuerRequest struct {
	ID   string `json:"id,omitempty"`   // issuer identification
	Name string `json:"name,omitempty"` // issuer name
}
