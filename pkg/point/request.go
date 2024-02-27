package point

type CreateRequest struct {
	AdditionalInfo *AdditionalInfo `json:"additional_info"`
	Payment        *PaymentRequest `json:"payment"`

	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type AdditionalInfo struct {
	PrintOnTerminal   bool   `json:"print_on_terminal,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
	TicketNumber      string `json:"ticket_number,omitempty"`
}

type PaymentRequest struct {
	ID               int64  `json:"id,omitempty"`
	Installments     int    `json:"installments,omitempty"`
	Type             string `json:"type,omitempty"`
	InstallmentsCost string `json:"installments_cost,omitempty"`
}

type UpdateDeviceOperatingModeRequest struct {
	OperatingMode string `json:"operating_mode"`
}
