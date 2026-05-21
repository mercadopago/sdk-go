package point

// Response represents the payment intent resource returned by the MercadoPago Point
// Integration API. It is returned by [Client.Create] and [Client.Get].
type Response struct {
	// Payment contains the payment type and installment details.
	Payment PaymentResponse `json:"payment"`

	// AdditionalInfo contains optional metadata about the payment intent.
	AdditionalInfo AdditionalInfoResponse `json:"additional_info"`

	// ID is the unique payment intent identifier assigned by MercadoPago.
	ID string `json:"id"`

	// Description is the payment description shown on the Point device.
	Description string `json:"description"`

	// State is the current state of the payment intent (e.g., "OPEN", "ON_TERMINAL", "PROCESSED", "CANCELED", "ERROR", "EXPIRED").
	State string `json:"state"`

	// DeviceID is the identifier of the Point device assigned to this payment intent.
	DeviceID string `json:"device_id"`

	// Amount is the total transaction amount in the currency's smallest unit.
	Amount int `json:"amount"`
}

// AdditionalInfoResponse contains optional metadata returned within a payment intent [Response].
type AdditionalInfoResponse struct {
	// ExternalReference is the integrator-provided external identifier.
	ExternalReference string `json:"external_reference"`

	// TicketNumber is the integrator-assigned ticket or receipt number.
	TicketNumber string `json:"ticket_number"`

	// PrintOnTerminal indicates whether a receipt was configured to print on the device.
	PrintOnTerminal bool `json:"print_on_terminal"`
}

// PaymentResponse contains payment type and installment details returned within
// a payment intent [Response].
type PaymentResponse struct {
	// Type is the payment type (e.g., "credit_card", "debit_card").
	Type string `json:"type"`

	// InstallmentsCost defines who absorbs the installment cost: "buyer" or "seller".
	InstallmentsCost string `json:"installments_cost"`

	// VoucherType is the voucher type used, if applicable (e.g., for meal or food vouchers).
	VoucherType string `json:"voucher_type"`

	// ID is the payment identifier, populated after the payment is processed.
	ID int `json:"id"`

	// Installments is the number of installments for the payment.
	Installments int `json:"installments"`
}

// CancelResponse represents the response returned by [Client.Cancel] after
// successfully cancelling a payment intent on a Point device.
type CancelResponse struct {
	// ID is the identifier of the cancelled payment intent.
	ID string `json:"id"`
}

// DevicesResponse represents the paginated list of Point devices returned by [Client.ListDevices].
type DevicesResponse struct {
	// Paging contains pagination metadata for the device list.
	Paging PagingResponse `json:"paging"`

	// Devices is the list of Point devices associated with the authenticated account.
	Devices []DeviceResponse `json:"devices"`
}

// DeviceResponse represents a single MercadoPago Point device returned within a [DevicesResponse].
type DeviceResponse struct {
	// ID is the unique device identifier assigned by MercadoPago.
	ID string `json:"id"`

	// ExternalPosID is the integrator-assigned point-of-sale identifier linked to this device.
	ExternalPosID string `json:"external_pos_id"`

	// OperatingMode is the current operating mode of the device: "PDV" (integrated) or "STANDALONE".
	OperatingMode string `json:"operating_mode"`

	// PosID is the MercadoPago point-of-sale identifier linked to this device.
	PosID int `json:"pos_id"`

	// StoreID is the MercadoPago store identifier where this device is located.
	StoreID string `json:"store_id"`
}

// PagingResponse contains pagination metadata for a list of Point devices.
type PagingResponse struct {
	// Total is the total number of devices matching the query.
	Total int `json:"total"`

	// Offset is the number of devices skipped in the current page.
	Offset int `json:"offset"`

	// Limit is the maximum number of devices returned per page.
	Limit int `json:"limit"`
}

// OperatingModeResponse represents the response returned by [Client.UpdateOperatingMode]
// after successfully changing a Point device's operating mode.
type OperatingModeResponse struct {
	// OperatingMode is the new operating mode of the device: "PDV" (integrated) or "STANDALONE".
	OperatingMode string `json:"operating_mode"`
}
