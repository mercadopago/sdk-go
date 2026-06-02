package chargeback

import "time"

// Response represents a chargeback dispute record returned by the MercadoPago API.
// A chargeback is initiated by a cardholder through their issuing bank when they
// dispute a payment.
type Response struct {
	// DateCreated is the date when the chargeback was created.
	DateCreated time.Time `json:"date_created"`

	// LastModified is the date of the last modification.
	LastModified time.Time `json:"last_modified"`

	// ID is the unique chargeback identifier assigned by MercadoPago.
	ID string `json:"id"`

	// CurrencyID is the ISO 4217 currency code of the disputed amount.
	CurrencyID string `json:"currency_id"`

	// ReasonID is the card-network reason code for the dispute.
	ReasonID string `json:"reason_id"`

	// Reason is the human-readable description of the dispute reason.
	Reason string `json:"reason"`

	// Status is the current status (e.g., "new", "in_review", "won", "lost").
	Status string `json:"status"`

	// Amount is the disputed amount.
	Amount float64 `json:"amount"`

	// PaymentID is the identifier of the payment that originated the dispute.
	PaymentID int `json:"payment_id"`
}
