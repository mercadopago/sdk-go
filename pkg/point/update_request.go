package point

// OperatingModeRequest represents the body payload sent to the MercadoPago Point Integration
// API when changing a device's operating mode via [Client.UpdateOperatingMode].
//
// The supported operating modes are:
//   - "PDV": Integrated mode where the device is controlled by the API for payment processing.
//   - "STANDALONE": Standalone mode where payments are processed directly on the device without API integration.
type OperatingModeRequest struct {
	// OperatingMode is the target operating mode for the device: "PDV" or "STANDALONE".
	OperatingMode string `json:"operating_mode,omitempty"`
}
