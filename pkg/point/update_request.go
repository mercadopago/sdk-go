package point

// OperatingModeRequest represents the operation mode to be changed of the device.
// The options are: PDV, which is when the device is used in integrated mode with our API, and
// STANDALONE, which is used when you want to process payments without our API.
type OperatingModeRequest struct {
	OperatingMode string `json:"operating_mode,omitempty"`
}
