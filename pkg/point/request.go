package point

type CreateRequest struct {
	AdditionalInfo *AdditionalInfoRequest `json:"additional_info"`
	Payment        *PaymentRequest        `json:"payment"`

	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type AdditionalInfoRequest struct {
	ExternalReference string `json:"external_reference,omitempty"`
	TicketNumber      string `json:"ticket_number,omitempty"`
	PrintOnTerminal   bool   `json:"print_on_terminal,omitempty"`
}

type PaymentRequest struct {
	Type             string `json:"type,omitempty"`
	InstallmentsCost string `json:"installments_cost,omitempty"`
	ID               int    `json:"id,omitempty"`
	Installments     int    `json:"installments,omitempty"`
}

type UpdateDeviceOperatingModeRequest struct {
	OperatingMode string `json:"operating_mode"`
}
