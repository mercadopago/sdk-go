package point

// Request represents the body payload sent to the MercadoPago Point Integration API when
// creating a new payment intent via [Client.Create]. It contains the transaction details
// that will be displayed on the Point device for the buyer to complete the payment.
type Request struct {
	// AdditionalInfo contains optional metadata about the payment intent, such as
	// external references and terminal printing preferences.
	AdditionalInfo *AdditionalInfoRequest `json:"additional_info,omitempty"`

	// Payment configures the payment type and installment details for the transaction.
	Payment *PaymentRequest `json:"payment,omitempty"`

	// Description is a text description of the payment shown to the buyer on the Point device.
	Description string `json:"description,omitempty"`

	// Amount is the total transaction amount in the currency's smallest unit (e.g., cents).
	Amount int `json:"amount,omitempty"`
}

// AdditionalInfoRequest contains optional metadata for a payment intent, allowing
// integrators to attach external references and configure terminal behavior.
type AdditionalInfoRequest struct {
	// ExternalReference is an external identifier for reconciliation with the integrator's system.
	ExternalReference string `json:"external_reference,omitempty"`

	// TicketNumber is the integrator-assigned ticket or receipt number for the transaction.
	TicketNumber string `json:"ticket_number,omitempty"`

	// PrintOnTerminal indicates whether a receipt should be printed on the Point device after payment.
	PrintOnTerminal bool `json:"print_on_terminal,omitempty"`
}

// PaymentRequest configures the payment type and installment details for a payment intent.
type PaymentRequest struct {
	// Type is the payment type (e.g., "credit_card", "debit_card").
	Type string `json:"type,omitempty"`

	// InstallmentsCost defines who absorbs the installment cost: "buyer" or "seller".
	InstallmentsCost string `json:"installments_cost,omitempty"`

	// ID is an optional payment identifier.
	ID int `json:"id,omitempty"`

	// Installments is the number of installments for the payment.
	Installments int `json:"installments,omitempty"`
}
