package point

type Response struct {
	Payment        PaymentResponse        `json:"payment"`
	AdditionalInfo AdditionalInfoResponse `json:"additional_info"`

	Amount      int    `json:"amount"`
	ID          string `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
	DeviceID    string `json:"device_id"`
}

type AdditionalInfoResponse struct {
	PrintOnTerminal   bool   `json:"print_on_terminal"`
	ExternalReference string `json:"external_reference"`
	TicketNumber      string `json:"ticket_number"`
}

type PaymentResponse struct {
	ID               int64  `json:"id"`
	Installments     int    `json:"installments"`
	Type             string `json:"type"`
	InstallmentsCost string `json:"installments_cost"`
	VoucherType      string `json:"voucher_type"`
}

type CancelResponse struct {
	ID string `json:"id"`
}

type DevicesResponse struct {
	Devices []DeviceResponse `json:"devices"`
	Paging  PagingResponse   `json:"paging"`
}

type DeviceResponse struct {
	PosID         int    `json:"pos_id"`
	StoreID       int    `json:"store_id"`
	ID            string `json:"id"`
	ExternalPosID string `json:"external_pos_id"`
	OperatingMode string `json:"operating_mode"`
}

type PagingResponse struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OperatingModeResponse struct {
	OperatingMode string `json:"operating_mode"`
}
