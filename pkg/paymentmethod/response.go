package paymentmethod

// Response represents a detailed payment method.
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

// SettingResponse represents payment method settings.
type SettingResponse struct {
	Bin          BinResponse          `json:"bin"`
	CardNumber   CardNumberResponse   `json:"card_number"`
	SecurityCode SecurityCodeResponse `json:"security_code"`
}

// BinResponse represents BIN (Bank Identification Number) settings.
type BinResponse struct {
	Pattern             string `json:"pattern"`
	ExclusionPattern    string `json:"exclusion_pattern"`
	InstallmentsPattern string `json:"installments_pattern"`
}

// CardNumberResponse represents card number settings.
type CardNumberResponse struct {
	Validation string `json:"validation"`
	Length     int    `json:"length"`
}

// SecurityCodeResponse represents security code settings.
type SecurityCodeResponse struct {
	Mode         string `json:"mode"`
	CardLocation string `json:"card_location"`
	Length       int    `json:"length"`
}

// FinancialInstitutionResponse represents financial institution settings.
type FinancialInstitutionResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
