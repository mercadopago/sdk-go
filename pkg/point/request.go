package point

// Request is used to create a payment intention, sending the details of a transaction.
type Request struct {
	AdditionalInfo *AdditionalInfoRequest `json:"additional_info,omitempty"`
	Payment        *PaymentRequest        `json:"payment,omitempty"`

	Description string `json:"description,omitempty"`
	Amount      int    `json:"amount,omitempty"`
}

// AdditionalInfoRequest contains the additional payment intent information.
type AdditionalInfoRequest struct {
	ExternalReference string `json:"external_reference,omitempty"`
	TicketNumber      string `json:"ticket_number,omitempty"`
	PrintOnTerminal   bool   `json:"print_on_terminal,omitempty"`
}

// PaymentRequest contains properties of Payment Intent.
type PaymentRequest struct {
	Type             string `json:"type,omitempty"`
	InstallmentsCost string `json:"installments_cost,omitempty"`
	ID               int    `json:"id,omitempty"`
	Installments     int    `json:"installments,omitempty"`
}
