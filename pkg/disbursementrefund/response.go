package disbursementrefund

import "time"

// Response represents a disbursement refund returned by the MercadoPago API.
type Response struct {
	// DateCreated is the date when the refund was created.
	DateCreated time.Time `json:"date_created"`

	// Status is the current status of the refund.
	Status string `json:"status"`

	// Amount is the refunded amount.
	Amount float64 `json:"amount"`

	// AdvancedPaymentID is the identifier of the parent advanced payment.
	AdvancedPaymentID int `json:"advanced_payment_id"`

	// DisbursementID is the identifier of the disbursement that was refunded.
	DisbursementID int `json:"disbursement_id"`

	// ID is the unique refund identifier assigned by MercadoPago.
	ID int `json:"id"`
}
