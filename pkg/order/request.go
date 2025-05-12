package order

// API version: d0494f1c-8d81-4c76-ae1d-0c65bb8ef6de

type Request struct {
	Type                string                `json:"type,omitempty"`
	TotalAmount         string                `json:"total_amount,omitempty"`
	ExternalReference   string                `json:"external_reference,omitempty"`
	CaptureMode         string                `json:"capture_mode,omitempty"`
	ProcessingMode      string                `json:"processing_mode,omitempty"`
	Description         string                `json:"description,omitempty"`
	Marketplace         string                `json:"marketplace,omitempty"`
	MarketPlaceFee      string                `json:"marketplace_fee,omitempty"`
	ExpirationTime      string                `json:"expiration_time,omitempty"`
	CheckoutAvailableAt string                `json:"checkout_available_at,omitempty"`
	Transactions        *TransactionRequest   `json:"transactions,omitempty"`
	Payer               *PayerRequest         `json:"payer,omitempty"`
	Items               []ItemsRequest        `json:"items,omitempty"`
	Config              *ConfigRequest        `json:"config,omitempty"`
	Address             []AddressRequest      `json:"adresses,omitempty"`
	AdditionalInfo      AdditionalInfoRequest `json:"additional_info,omitempty"`
}

type AdditionalInfoRequest struct {
	Payer    *AdditionalInfoPayer    `json:"payer,omitempty"`
	Shipment *AdditionalInfoShipment `json:"shipment,omitempty"`
	Platform *AdditionalInfoPlatform `json:"platform,omitempty"`
	Travel   *AdditionalInfoTravel   `json:"travel,omitempty"`
}

type AdditionalInfoPayer struct {
	AuthenticationType    string        `json:"authentication_type,omitempty"`
	RegistrationDate      string        `json:"registration_date,omitempty"`
	IsPrimeUser           bool          `json:"is_prime_user,omitempty"`
	IsFirstPurchaseOnline bool          `json:"is_first_purchase_online,omitempty"`
	LastPurchase          string        `json:"last_purchase,omitempty"`
	Address               *PayerAddress `json:"address,omitempty"`
}

type AdditionalInfoShipment struct {
	Express     bool `json:"express,omitempty"`
	LocalPickup bool `json:"local_pickup,omitempty"`
}

type AdditionalInfoPlatform struct {
	Shipment       *AdditionalInfoPlatformShipment `json:"shipment,omitempty"`
	Seller         *AdditionalInfoPlatformSeller   `json:"seller,omitempty"`
	Authentication string                          `json:"authentication,omitempty"`
}

type AdditionalInfoPlatformShipment struct {
	DeliveryPromise string                          `json:"delivery_promise,omitempty"`
	DropShipping    bool                            `json:"drop_shipping,omitempty"`
	Safety          string                          `json:"safety,omitempty"`
	Tracking        *AdditionalInfoPlatformTracking `json:"tracking,omitempty"`
	Withdrawn       bool                            `json:"withdrawn,omitempty"`
}

type AdditionalInfoPlatformTracking struct {
	Code   string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
}

type AdditionalInfoPlatformSeller struct {
	ID                 string                 `json:"id,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Email              string                 `json:"email,omitempty"`
	Status             string                 `json:"status,omitempty"`
	ReferralURL        string                 `json:"referral_url,omitempty"`
	RegistrationDate   string                 `json:"registration_date,omitempty"`
	HiredPlan          string                 `json:"hired_plan,omitempty"`
	BusinessType       string                 `json:"business_type,omitempty"`
	Address            *AddressRequest        `json:"address,omitempty"`
	IdentificationType *IdentificationRequest `json:"identification,omitempty"`
	Phone              *PhoneRequest          `json:"phone,omitempty"`
}

type AdditionalInfoTravel struct {
	Passengers []string `json:"passengers,omitempty"`
	Routes     []string `json:"routes,omitempty"`
}
type TransactionRequest struct {
	Payments []PaymentRequest `json:"payments,omitempty"`
}

type PaymentRequest struct {
	Amount            string                    `json:"amount,omitempty"`
	ExpirationTime    string                    `json:"expiration_time,omitempty"`
	PaymentMethod     *PaymentMethodRequest     `json:"payment_method,omitempty"`
	AutomaticPayments *AutomaticPaymentsRequest `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredentialRequest  `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionDataRequest  `json:"subscription_data,omitempty"`
}

type PaymentMethodRequest struct {
	ID                  string `json:"id,omitempty"`
	Type                string `json:"type,omitempty"`
	Token               string `json:"token,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	Installments        int    `json:"installments,omitempty"`
}

type AutomaticPaymentsRequest struct {
	PaymentProfileID string `json:"payment_profile_id,omitempty"`
	ScheduleDate     string `json:"schedule_date,omitempty"`
	DueDate          string `json:"due_date,omitempty"`
	Retries          int    `json:"retries,omitempty"`
}

type StoredCredentialRequest struct {
	PaymentInitiator   string `json:"payment_initiator,omitempty"`
	Reason             string `json:"reason,omitempty"`
	StorePaymentMethod bool   `json:"store_payment_method,omitempty"`
	FirstPayment       bool   `json:"first_payment,omitempty"`
}

type SubscriptionDataRequest struct {
	InvoiceID            string                       `json:"invoice_id,omitempty"`
	BillingDate          string                       `json:"billing_date,omitempty"`
	SubscriptionSequence *SubscriptionSequenceRequest `json:"subscription_sequence,omitempty"`
	InvoicePeriod        *InvoicePeriodRequest        `json:"invoice_period,omitempty"`
}

type SubscriptionSequenceRequest struct {
	Number int `json:"number,omitempty"`
	Total  int `json:"total,omitempty"`
}

type InvoicePeriodRequest struct {
	Type   string `json:"type,omitempty"`
	Period int    `json:"period,omitempty"`
}

type PayerRequest struct {
	Email          string                 `json:"email,omitempty"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	CustomerID     string                 `json:"customer_id,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`
	Address        *PayerAddress          `json:"address,omitempty"`
}
type PayerAddress struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
}

type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

type AddressRequest struct {
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Complement   string `json:"complement,omitempty"`
}

type ItemsRequest struct {
	Title        string `json:"title,omitempty"`
	Type         string `json:"type,omitempty"`
	Warranty     string `json:"warranty,omitempty"`
	Event_date   string `json:"event_date,omitempty"`
	UnitPrice    string `json:"unit_price,omitempty"`
	ExternalCode string `json:"external_code,omitempty"`
	CategoryID   string `json:"category_id,omitempty"`
	Description  string `json:"description,omitempty"`
	PictureURL   string `json:"picture_url,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
}

type RefundRequest struct {
	Transactions []RefundTransaction `json:"transactions,omitempty"`
}

type RefundTransaction struct {
	ID     string `json:"id,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type ConfigRequest struct {
	PaymentMethod *PaymentMethodConfigRequest `json:"payment_method,omitempty"`
	Online        *OnlineConfigRequest        `json:"online,omitempty"`
}

type PaymentMethodConfigRequest struct {
	NotAllowedIDs       []string `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes     []string `json:"not_allowed_types,omitempty"`
	DefaultID           string   `json:"default_id,omitempty"`
	MaxInstallments     int      `json:"max_installments,omitempty"`
	DefaultInstallments int      `json:"default_installments,omitempty"`
}

type OnlineConfigRequest struct {
	CallbackURL         string                      `json:"callback_url,omitempty"`
	SuccessURL          string                      `json:"success_url,omitempty"`
	PendingURL          string                      `json:"pending_url,omitempty"`
	FailureURL          string                      `json:"failure_url,omitempty"`
	AutoReturnURL       string                      `json:"auto_return_url,omitempty"`
	DifferentialPricing *DifferentialPricingRequest `json:"differential_pricing,omitempty"`
}

type DifferentialPricingRequest struct {
	ID int `json:"id,omitempty"`
}
