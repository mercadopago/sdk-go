package paymentmethod

type Response struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	PaymentTypeID        string   `json:"payment_type_id,omitempty"`
	Status               string   `json:"status,omitempty"`
	SecureThumbnail      string   `json:"secure_thumbnail,omitempty"`
	Thumbnail            string   `json:"thumbnail,omitempty"`
	DeferredCapture      string   `json:"deferred_capture,omitempty"`
	AdditionalInfoNeeded []string `json:"additional_info_needed,omitempty"`
	ProcessingModes      []string `json:"processing_modes,omitempty"`
	AccreditationTime    int64    `json:"accreditation_time,omitempty"`
	MinAllowedAmount     float64  `json:"min_allowed_amount,omitempty"`
	MaxAllowedAmount     float64  `json:"max_allowed_amount,omitempty"`

	Settings              []SettingsResponse             `json:"settings,omitempty"`
	FinancialInstitutions []FinancialInstitutionResponse `json:"financial_institutions,omitempty"`
}

// SettingsResponse represents payment method settings.
type SettingsResponse struct {
	Bin          *SettingsBinResponse          `json:"bin,omitempty"`
	CardNumber   *SettingsCardNumberResponse   `json:"card_number,omitempty"`
	SecurityCode *SettingsSecurityCodeResponse `json:"security_code,omitempty"`
}

// SettingsBinResponse represents BIN (Bank Identification Number) settings.
type SettingsBinResponse struct {
	Pattern             string `json:"pattern,omitempty"`
	ExclusionPattern    string `json:"exclusion_pattern,omitempty"`
	InstallmentsPattern string `json:"installments_pattern,omitempty"`
}

// SettingsCardNumberResponse represents card number settings.
type SettingsCardNumberResponse struct {
	Length     int    `json:"length,omitempty"`
	Validation string `json:"validation,omitempty"`
}

// SettingsSecurityCodeResponse represents security code settings.
type SettingsSecurityCodeResponse struct {
	Mode         string `json:"mode,omitempty"`
	Length       int    `json:"length,omitempty"`
	CardLocation string `json:"card_location,omitempty"`
}

// FinancialInstitutionResponse represents financial institution settings.
type FinancialInstitutionResponse struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
