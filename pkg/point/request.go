package point

// Request is used to create a payment intention, sending the details of a transaction.
type Request struct {
	AdditionalInfo *AdditionalInfoRequest `json:"additional_info"`
	Payment        *PaymentRequest        `json:"payment"`

	Description string `json:"description"`
	Amount      int    `json:"amount"`
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

// UpdateDeviceOperatingModeRequest represents the operation mode to be changed of the device.
// The options are: PDV, which is when the device is used in integrated mode with our API, and
// STANDALONE, which is used when you want to process payments without our API.
type UpdateDeviceOperatingModeRequest struct {
	OperatingMode string `json:"operating_mode"`
}
