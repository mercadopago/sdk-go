package customercard

// Request represents the body payload for creating or updating a saved customer card
// via the MercadoPago Customer Cards API. The Token field should contain a card token
// previously generated through the Card Token API.
type Request struct {
	// Issuer contains optional card-issuer details to associate with the card.
	Issuer *IssuerRequest `json:"issuer,omitempty"`

	// Token is the card token obtained from the MercadoPago Card Token API.
	// It is required when creating a new card.
	Token string `json:"token,omitempty"`
	// PaymentMethodID is the identifier of the payment method (e.g., "visa", "master").
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}

// IssuerRequest represents card-issuer information within a customer card [Request].
// It identifies the financial institution that issued the card.
type IssuerRequest struct {
	// ID is the issuer identifier in MercadoPago.
	ID string `json:"id,omitempty"`
	// Name is the issuer's human-readable name.
	Name string `json:"name,omitempty"`
}
