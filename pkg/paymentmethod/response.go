package paymentmethod

// Response represents a payment method returned by the MercadoPago Payment Methods API.
// It includes configuration details such as allowed amount ranges, processing modes,
// accreditation times, and card validation settings.
type Response struct {
	Settings              []SettingResponse              `json:"settings"`
	FinancialInstitutions []FinancialInstitutionResponse `json:"financial_institutions"`

	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	PaymentTypeID        string   `json:"payment_type_id"`
	Status               string   `json:"status"`
	SecureThumbnail      string   `json:"secure_thumbnail"`
	Thumbnail            string   `json:"thumbnail"`
	DeferredCapture      string   `json:"deferred_capture"`
	MinAllowedAmount     float64  `json:"min_allowed_amount"`
	MaxAllowedAmount     float64  `json:"max_allowed_amount"`
	AccreditationTime    int      `json:"accreditation_time"`
	AdditionalInfoNeeded []string `json:"additional_info_needed"`
	ProcessingModes      []string `json:"processing_modes"`
}

// SettingResponse contains validation settings for a payment method, including BIN patterns,
// card number length, and security code requirements.
type SettingResponse struct {
	Bin          BinResponse          `json:"bin"`
	CardNumber   CardNumberResponse   `json:"card_number"`
	SecurityCode SecurityCodeResponse `json:"security_code"`
}

// BinResponse contains BIN (Bank Identification Number) pattern rules used to identify
// which card numbers belong to this payment method.
type BinResponse struct {
	Pattern             string `json:"pattern"`
	ExclusionPattern    string `json:"exclusion_pattern"`
	InstallmentsPattern string `json:"installments_pattern"`
}

// CardNumberResponse defines the expected card number validation algorithm and length
// for this payment method.
type CardNumberResponse struct {
	Validation string `json:"validation"`
	Length     int    `json:"length"`
}

// SecurityCodeResponse defines the security code (CVV/CVC) requirements for this
// payment method, including its display mode, physical location on the card, and length.
type SecurityCodeResponse struct {
	Mode         string `json:"mode"`
	CardLocation string `json:"card_location"`
	Length       int    `json:"length"`
}

// FinancialInstitutionResponse identifies a financial institution (e.g., a bank)
// associated with the payment method.
type FinancialInstitutionResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
