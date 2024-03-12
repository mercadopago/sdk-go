package paymentmethod

type Response struct {
	Settings              []SettingsResponse             `json:"settings"`
	FinancialInstitutions []FinancialInstitutionResponse `json:"financial_institutions"`

	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	PaymentTypeID        string   `json:"payment_type_id"`
	Status               string   `json:"status"`
	SecureThumbnail      string   `json:"secure_thumbnail"`
	Thumbnail            string   `json:"thumbnail"`
	DeferredCapture      string   `json:"deferred_capture"`
	AdditionalInfoNeeded []string `json:"additional_info_needed"`
	ProcessingModes      []string `json:"processing_modes"`
	AccreditationTime    int64    `json:"accreditation_time"`
	MinAllowedAmount     float64  `json:"min_allowed_amount"`
	MaxAllowedAmount     float64  `json:"max_allowed_amount"`
}

// SettingsResponse represents payment method settings.
type SettingsResponse struct {
	Bin          SettingsBinResponse          `json:"bin"`
	CardNumber   SettingsCardNumberResponse   `json:"card_number"`
	SecurityCode SettingsSecurityCodeResponse `json:"security_code"`
}

// SettingsBinResponse represents BIN (Bank Identification Number) settings.
type SettingsBinResponse struct {
	Pattern             string `json:"pattern"`
	ExclusionPattern    string `json:"exclusion_pattern"`
	InstallmentsPattern string `json:"installments_pattern"`
}

// SettingsCardNumberResponse represents card number settings.
type SettingsCardNumberResponse struct {
	Length     int    `json:"length"`
	Validation string `json:"validation"`
}

// SettingsSecurityCodeResponse represents security code settings.
type SettingsSecurityCodeResponse struct {
	Mode         string `json:"mode"`
	Length       int    `json:"length"`
	CardLocation string `json:"card_location"`
}

// FinancialInstitutionResponse represents financial institution settings.
type FinancialInstitutionResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
