package order

type Request struct {
	Type                string                 `json:"type,omitempty"`
	TotalAmount         string                 `json:"total_amount,omitempty"`
	ExternalReference   string                 `json:"external_reference,omitempty"`
	CaptureMode         string                 `json:"capture_mode,omitempty"`
	ProcessingMode      string                 `json:"processing_mode,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Marketplace         string                 `json:"marketplace,omitempty"`
	MarketPlaceFee      string                 `json:"marketplace_fee,omitempty"`
	ExpirationTime      string                 `json:"expiration_time,omitempty"`
	CheckoutAvailableAt string                 `json:"checkout_available_at,omitempty"`
	Transactions        *TransactionRequest    `json:"transactions,omitempty"`
	Payer               *PayerRequest          `json:"payer,omitempty"`
	Items               []ItemsRequest         `json:"items,omitempty"`
	Config              *ConfigRequest         `json:"config,omitempty"`
	Address             []AddressRequest       `json:"adresses,omitempty"`
	AdditionalInfo      *AdditionalInfoRequest `json:"additional_info,omitempty"`
	Shipment            *ShipmentRequest       `json:"shipment,omitempty"`
}

type TravelPassengerRequest struct {
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
}

type TravelRouteRequest struct {
	Departure         string `json:"departure,omitempty"`
	Destination       string `json:"destination,omitempty"`
	DepartureDateTime string `json:"departure_date_time,omitempty"`
	ArrivalDateTime   string `json:"arrival_date_time,omitempty"`
	Company           string `json:"company,omitempty"`
}

type AdditionalInfoRequest struct {
	PayerAuthenticationType            string                    `json:"payer.authentication_type,omitempty"`
	PayerRegistrationDate              string                    `json:"payer.registration_date,omitempty"`
	PayerIsPrimeUser                   bool                      `json:"payer.is_prime_user,omitempty"`
	PayerIsFirstPurchaseOnLine         bool                      `json:"payer.is_first_purchase_onLine,omitempty"`
	PayerLastPurchase                  string                    `json:"payer.last_purchase,omitempty"`
	ShipmentExpress                    bool                      `json:"shipment.express,omitempty"`
	ShipmentLocalPickup                bool                      `json:"shipment.local_pickup,omitempty"`
	PlatFormShipmentDeliveryPromise    string                    `json:"platform.shipment.delivery_promise,omitempty"`
	PlatFormShipmentDropShipping       string                    `json:"platform.shipment.drop_shipping,omitempty"`
	PlatformShipmentSafety             string                    `json:"platform.shipment.safety,omitempty"`
	PlatformShipmentTrackingCode       string                    `json:"platform.shipment.tracking.code,omitempty"`
	PlatformShipmentTrackingStatus     string                    `json:"platform.shipment.tracking.status,omitempty"`
	PlatformShipmentWithdrawn          bool                      `json:"platform.shipment.withdrawn,omitempty"`
	PlatformSellerId                   string                    `json:"platform.seller.id,omitempty"`
	PlatformSellerName                 string                    `json:"platform.seller.name,omitempty"`
	PlatformSellerEmail                string                    `json:"platform.seller.email,omitempty"`
	PlatformSellerStatus               string                    `json:"platform.seller.status,omitempty"`
	PlatformSellerReferralUrl          string                    `json:"platform.seller.referral_url,omitempty"`
	PlatformSellerRegistrationDate     string                    `json:"platform.seller.registration_date,omitempty"`
	PlatformSellerHiredPlan            string                    `json:"platform.seller.hired_plan,omitempty"`
	PlatformSellerBusinessType         string                    `json:"platform.seller.business_type,omitempty"`
	PlatformSellerAddressZipCode       string                    `json:"platform.seller.address.zip_code,omitempty"`
	PlatformSellerAddressStreetName    string                    `json:"platform.seller.address.street_name,omitempty"`
	PlatformSellerAddressStreetNumber  string                    `json:"platform.seller.address.street_number,omitempty"`
	PlatformSellerAddressCity          string                    `json:"platform.seller.address.city,omitempty"`
	PlatformSellerAddressState         string                    `json:"platform.seller.address.state,omitempty"`
	PlatformSellerAddressComplement    string                    `json:"platform.seller.address.complement,omitempty"`
	PlatformSellerAddressCountry       string                    `json:"platform.seller.address.country,omitempty"`
	PlatformSellerIdentificationType   string                    `json:"platform.seller.identification.type,omitempty"`
	PlatformSellerIdentificationNumber string                    `json:"platform.seller.identification.number,omitempty"`
	PlatformSellerPhoneAreaCode        string                    `json:"platform.seller.phone.area_code,omitempty"`
	PlatformSellerPhoneNumber          string                    `json:"platform.seller.phone.number,omitempty"`
	PlatformAuthentication             string                    `json:"platform.authentication,omitempty"`
	TravelPassengers                   *[]TravelPassengerRequest `json:"travel.passengers,omitempty"`
	TravelRoutes                       *[]TravelRouterRequest    `json:"travel.routes,omitempty"`
}

type TravelRouterRequest struct {
	Departure         string `json:"departure,omitempty"`
	Destination       string `json:"destination,omitempty"`
	DepartureDateTime string `json:"departure_date_time,omitempty"`
	ArrivalDateTime   string `json:"arrival_date_time,omitempty"`
	Company           string `json:"company,omitempty"`
}
type PayerAddressRequest struct {
	Zipcode      string `json:"zipcode,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Country      string `json:"country,omitempty"`
}

type ShipmentRequest struct {
	Address *AddressRequest `json:"address,omitempty"`
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
	Address        *PayerAddressRequest   `json:"address,omitempty"`
	EntityType     string                 `json:"entity_type,omitempty"`
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
	Warranty     bool   `json:"warranty,omitempty"`
	EventDate    string `json:"event_date,omitempty"`
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
