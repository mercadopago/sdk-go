package point

type CreateResponse struct {
	Amount      int    `json:"amount"`
	ID          string `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
	DeviceID    string `json:"device_id"`

	Payment        CreatePaymentResponse `json:"payment"`
	AdditionalInfo AdditionalInfo        `json:"additional_info"`
}

type GetResponse struct {
	Amount   int    `json:"amount"`
	ID       string `json:"id"`
	State    string `json:"state"`
	DeviceID string `json:"device_id"`

	Payment        PaymentResponse `json:"payment"`
	AdditionalInfo AdditionalInfo  `json:"additional_info"`
}

type CreatePaymentResponse struct {
	Installments     int    `json:"installments"`
	Type             string `json:"type"`
	InstallmentsCost string `json:"installments_cost"`
}

type PaymentResponse struct {
	ID int64 `json:"id,omitempty"`
}

type CancelResponse struct {
	ID string `json:"id"`
}

type DevicesResponse struct {
	Devices []Device `json:"devices"`
	Paging  Paging   `json:"paging"`
}

type Device struct {
	PosID         int    `json:"pos_id"`
	StoreID       int    `json:"store_id"`
	ID            string `json:"id"`
	ExternalPosID string `json:"external_pos_id"`
	OperatingMode string `json:"operating_mode"`
}

type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OperationModeResponse struct {
	OperatingMode string `json:"operating_mode"`
}
