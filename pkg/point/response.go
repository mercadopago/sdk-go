package point

// Response contains the response of a payment intention.
type Response struct {
	Payment        PaymentResponse        `json:"payment"`
	AdditionalInfo AdditionalInfoResponse `json:"additional_info"`

	ID          string `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
	DeviceID    string `json:"device_id"`
	Amount      int    `json:"amount"`
}

type AdditionalInfoResponse struct {
	ExternalReference string `json:"external_reference"`
	TicketNumber      string `json:"ticket_number"`
	PrintOnTerminal   bool   `json:"print_on_terminal"`
}

type PaymentResponse struct {
	Type             string `json:"type"`
	InstallmentsCost string `json:"installments_cost"`
	VoucherType      string `json:"voucher_type"`
	ID               int    `json:"id"`
	Installments     int    `json:"installments"`
}

type CancelResponse struct {
	ID string `json:"id"`
}

type DevicesResponse struct {
	Paging  PagingResponse   `json:"paging"`
	Devices []DeviceResponse `json:"devices"`
}

type DeviceResponse struct {
	ID            string `json:"id"`
	ExternalPosID string `json:"external_pos_id"`
	OperatingMode string `json:"operating_mode"`
	PosID         int    `json:"pos_id"`
	StoreID       string `json:"store_id"`
}

type PagingResponse struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OperatingModeResponse struct {
	OperatingMode string `json:"operating_mode"`
}
