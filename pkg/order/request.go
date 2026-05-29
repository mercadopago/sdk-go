package order

import (
	"strconv"
	"strings"
)

// Request represents the body sent to the MercadoPago Orders API when creating a new order.
// It contains all the information needed to describe the order, including total amount,
// payment transactions, payer details, line items, and checkout configuration.
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
	Currency            string                 `json:"currency,omitempty"`
	Transactions        *TransactionRequest    `json:"transactions,omitempty"`
	Payer               *PayerRequest          `json:"payer,omitempty"`
	Items               []ItemsRequest         `json:"items,omitempty"`
	Config              *ConfigRequest         `json:"config,omitempty"`
	Shipment            *ShipmentRequest       `json:"shipment,omitempty"`
	AdditionalInfo      *AdditionalInfoRequest `json:"additional_info,omitempty"`
	// IntegrationData contains integration metadata identifying the integrator, platform,
	// corporation, and sponsor.
	IntegrationData *IntegrationDataRequest `json:"integration_data,omitempty"`
}

// TravelPassengerRequest represents a passenger in a travel-related order.
// It is used within [AdditionalInfoRequest] to provide fraud-prevention data
// for airline or bus ticket purchases.
type TravelPassengerRequest struct {
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
}

// TravelRouteRequest represents a travel route segment for airline or bus ticket orders.
// It is used within [AdditionalInfoRequest] to supply itinerary details for risk analysis.
type TravelRouteRequest struct {
	Departure         string `json:"departure,omitempty"`
	Destination       string `json:"destination,omitempty"`
	DepartureDateTime string `json:"departure_date_time,omitempty"`
	ArrivalDateTime   string `json:"arrival_date_time,omitempty"`
	Company           string `json:"company,omitempty"`
}

// AdditionalInfoRequest represents supplementary data sent with an order to improve
// fraud analysis and risk scoring. It includes payer behavior details, shipment metadata,
// platform seller information, and travel itinerary data. The flat JSON key structure
// (e.g., "payer.authentication_type") matches the MercadoPago API's dot-notation format.
type AdditionalInfoRequest struct {
	PayerAuthenticationType            string                    `json:"payer.authentication_type,omitempty"`
	PayerRegistrationDate              string                    `json:"payer.registration_date,omitempty"`
	PayerIsPrimeUser                   bool                      `json:"payer.is_prime_user,omitempty"`
	PayerIsFirstPurchaseOnLine         bool                      `json:"payer.is_first_purchase_online,omitempty"`
	PayerLastPurchase                  string                    `json:"payer.last_purchase,omitempty"`
	PayerIPAddress                     string                    `json:"payer.ip_address,omitempty"`
	ShipmentExpress                    bool                      `json:"shipment.express,omitempty"`
	ShipmentLocalPickup                bool                      `json:"shipment.local_pickup,omitempty"`
	PlatFormShipmentDeliveryPromise    string                    `json:"platform.shipment.delivery_promise,omitempty"`
	PlatFormShipmentDropShipping       string                    `json:"platform.shipment.drop_shipping,omitempty"`
	PlatformShipmentSafety             string                    `json:"platform.shipment.safety,omitempty"`
	PlatformShipmentTrackingCode       string                    `json:"platform.shipment.tracking.code,omitempty"`
	PlatformShipmentTrackingStatus     string                    `json:"platform.shipment.tracking.status,omitempty"`
	PlatformShipmentWithdrawn          bool                      `json:"platform.shipment.withdrawn,omitempty"`
	PlatformSellerID                   string                    `json:"platform.seller.id,omitempty"`
	PlatformSellerName                 string                    `json:"platform.seller.name,omitempty"`
	PlatformSellerEmail                string                    `json:"platform.seller.email,omitempty"`
	PlatformSellerStatus               string                    `json:"platform.seller.status,omitempty"`
	PlatformSellerReferralURL          string                    `json:"platform.seller.referral_url,omitempty"`
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

// TravelRouterRequest represents a travel route segment used within [AdditionalInfoRequest].
// It captures departure and destination details for travel-related orders.
type TravelRouterRequest struct {
	Departure         string `json:"departure,omitempty"`
	Destination       string `json:"destination,omitempty"`
	DepartureDateTime string `json:"departure_date_time,omitempty"`
	ArrivalDateTime   string `json:"arrival_date_time,omitempty"`
	Company           string `json:"company,omitempty"`
}

// PayerAddressRequest represents the payer's address when creating or updating an order.
// It is used within [PayerRequest] to provide the payer's billing or residential address.
type PayerAddressRequest struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Country      string `json:"country,omitempty"`
}

// ShipmentRequest represents shipping information for an order, including the delivery address.
type ShipmentRequest struct {
	Address *AddressRequest `json:"address,omitempty"`
}

// TransactionRequest represents the collection of payment transactions to associate with an order.
// Each order may contain one or more payments, enabling split-payment scenarios.
type TransactionRequest struct {
	Payments []PaymentRequest `json:"payments,omitempty"`
}

// PaymentRequest represents an individual payment transaction within an order.
// It specifies the amount, payment method, and optional configuration for automatic
// payments, stored credentials, and subscription billing.
type PaymentRequest struct {
	Amount            string                    `json:"amount,omitempty"`
	ExpirationTime    string                    `json:"expiration_time,omitempty"`
	DateOfExpiration  string                    `json:"date_of_expiration,omitempty"`
	PaymentMethod     *PaymentMethodRequest     `json:"payment_method,omitempty"`
	AutomaticPayments *AutomaticPaymentsRequest `json:"automatic_payments,omitempty"`
	StoredCredential  *StoredCredentialRequest  `json:"stored_credential,omitempty"`
	SubscriptionData  *SubscriptionDataRequest  `json:"subscription_data,omitempty"`
}

// PaymentMethodRequest represents the payment method details for a transaction.
// It identifies how the payer intends to pay (e.g., credit card, debit card, boleto)
// and includes tokenized card data when applicable.
type PaymentMethodRequest struct {
	ID                   string `json:"id,omitempty"`
	Type                 string `json:"type,omitempty"`
	Token                string `json:"token,omitempty"`
	StatementDescriptor  string `json:"statement_descriptor,omitempty"`
	Installments         int    `json:"installments,omitempty"`
	FinancialInstitution string `json:"financial_institution,omitempty"`
}

// AutomaticPaymentsRequest represents configuration for recurring automatic payment scheduling.
// It is used within [PaymentRequest] to define when and how automatic charges should occur.
type AutomaticPaymentsRequest struct {
	PaymentProfileID string `json:"payment_profile_id,omitempty"`
	ScheduleDate     string `json:"schedule_date,omitempty"`
	DueDate          string `json:"due_date,omitempty"`
	Retries          int    `json:"retries,omitempty"`
}

// StoredCredentialRequest represents stored credential information for a payment, used
// to indicate whether the payment method should be saved for future use and whether this
// is the initial or subsequent payment in a recurring series.
type StoredCredentialRequest struct {
	PaymentInitiator   string `json:"payment_initiator,omitempty"`
	Reason             string `json:"reason,omitempty"`
	StorePaymentMethod bool   `json:"store_payment_method,omitempty"`
	FirstPayment       bool   `json:"first_payment,omitempty"`
	PrevTransactionRef string `json:"prev_transaction_ref,omitempty"`
}

// SubscriptionDataRequest represents subscription billing details for a payment transaction.
// It links the payment to a specific invoice and billing cycle within a subscription.
type SubscriptionDataRequest struct {
	InvoiceID            string                       `json:"invoice_id,omitempty"`
	BillingDate          string                       `json:"billing_date,omitempty"`
	SubscriptionSequence *SubscriptionSequenceRequest `json:"subscription_sequence,omitempty"`
	InvoicePeriod        *InvoicePeriodRequest        `json:"invoice_period,omitempty"`
}

// SubscriptionSequenceRequest represents the position of a payment within a subscription series.
// Number indicates the current installment, and Total indicates the total number of installments.
type SubscriptionSequenceRequest struct {
	Number int `json:"number,omitempty"`
	Total  int `json:"total,omitempty"`
}

// InvoicePeriodRequest represents the billing period for a subscription invoice.
// Type indicates the period unit (e.g., "monthly", "yearly") and Period indicates
// the number of such units.
type InvoicePeriodRequest struct {
	Type   string `json:"type,omitempty"`
	Period int    `json:"period,omitempty"`
}

// PayerRequest represents the payer associated with an order. It includes personal
// identification, contact information, and address details used for payment processing
// and fraud prevention.
type PayerRequest struct {
	Email          string                 `json:"email,omitempty"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	CustomerID     string                 `json:"customer_id,omitempty"`
	EntityType     string                 `json:"entity_type,omitempty"`
	Identification *IdentificationRequest `json:"identification,omitempty"`
	Phone          *PhoneRequest          `json:"phone,omitempty"`
	Address        *PayerAddressRequest   `json:"address,omitempty"`
}

// PayerAddress represents a payer's address. It is similar to [PayerAddressRequest]
// but may be used in different contexts within the order lifecycle.
type PayerAddress struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
}

// IdentificationRequest represents a payer's identification document (e.g., CPF, DNI, CURP).
// Type specifies the document type and Number contains the document value.
type IdentificationRequest struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

// PhoneRequest represents a phone number, split into area code and number components.
type PhoneRequest struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

// AddressRequest represents a generic address used for shipment destinations within an order.
type AddressRequest struct {
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Complement   string `json:"complement,omitempty"`
}

// ItemsRequest represents a line item in an order. Each item describes a product or service
// being purchased, including its price, quantity, and descriptive metadata.
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

// RefundRequest represents the body sent to the Orders API to initiate a partial refund.
// It specifies which transactions to refund and the corresponding amounts.
// For a full refund, pass nil instead of a [RefundRequest] to [Client.Refund].
type RefundRequest struct {
	Transactions []RefundTransaction `json:"transactions,omitempty"`
}

// RefundTransaction identifies a single transaction to refund and the amount to return.
// ID is the transaction identifier and Amount is the value to refund as a decimal string.
type RefundTransaction struct {
	ID     string `json:"id,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// ConfigRequest represents order-level configuration options, including payment method
// restrictions and online checkout redirect URLs.
type ConfigRequest struct {
	PaymentMethod *PaymentMethodConfigRequest `json:"payment_method,omitempty"`
	Online        *OnlineConfigRequest        `json:"online,omitempty"`
}

// PaymentMethodConfigRequest represents constraints and defaults applied to the payment
// methods available for an order. It allows blocking specific methods, setting a default,
// and controlling installment limits.
type PaymentMethodConfigRequest struct {
	NotAllowedIDs       []string `json:"not_allowed_ids,omitempty"`
	NotAllowedTypes     []string `json:"not_allowed_types,omitempty"`
	DefaultID           string   `json:"default_id,omitempty"`
	MaxInstallments     int      `json:"max_installments,omitempty"`
	DefaultInstallments int      `json:"default_installments,omitempty"`
}

// OnlineConfigRequest represents online checkout configuration for an order, including
// redirect URLs for different payment outcomes and optional differential pricing or
// 3D Secure transaction security settings.
type OnlineConfigRequest struct {
	CallbackURL         string                      `json:"callback_url,omitempty"`
	SuccessURL          string                      `json:"success_url,omitempty"`
	PendingURL          string                      `json:"pending_url,omitempty"`
	FailureURL          string                      `json:"failure_url,omitempty"`
	AutoReturnURL       string                      `json:"auto_return_url,omitempty"`
	DifferentialPricing *DifferentialPricingRequest `json:"differential_pricing,omitempty"`
	TransactionSecurity *TransactionSecurityRequest `json:"transaction_security,omitempty"`
}

// DifferentialPricingRequest represents a differential pricing configuration identified by its ID.
// Differential pricing allows offering different prices based on the payment method chosen.
type DifferentialPricingRequest struct {
	ID int `json:"id,omitempty"`
}

// TransactionSecurityRequest represents 3D Secure (3DS) configuration for a payment transaction.
// Validation specifies the security validation mode, and LiabilityShift indicates who bears
// liability for chargebacks.
type TransactionSecurityRequest struct {
	Validation     string `json:"validation,omitempty"`
	LiabilityShift string `json:"liability_shift,omitempty"`
}

// SearchRequest contains the parameters for searching orders via the [Client.Search] method.
// Filters is a key-value map of field names to expected values. Limit defaults to 30 if zero.
type SearchRequest struct {
	Limit   int
	Offset  int
	Filters map[string]string
}

// GetParams converts the [SearchRequest] into a flat map of query parameters suitable
// for URL encoding. Filter keys are lowercased automatically. If Limit is zero it
// defaults to 30.
func (sr *SearchRequest) GetParams() map[string]string {
	params := map[string]string{}
	for k, v := range sr.Filters {
		key := strings.ToLower(k)
		params[key] = v
	}

	if sr.Limit == 0 {
		sr.Limit = 30
	}
	params["limit"] = strconv.Itoa(sr.Limit)
	params["offset"] = strconv.Itoa(sr.Offset)

	return params
}

// IntegrationDataRequest contains integration metadata for an order. Identifies the integrator,
// platform, and corporation associated with the integration, as well as any sponsoring
// marketplace owner.
type IntegrationDataRequest struct {
	// IntegratorID is the identifier of the certified integrator. Type: string.
	IntegratorID string `json:"integrator_id,omitempty"`
	// PlatformID is the platform identifier assigned by MercadoPago. Type: string.
	PlatformID string `json:"platform_id,omitempty"`
	// CorporationID is the corporation identifier for multi-account setups. Type: string.
	CorporationID string `json:"corporation_id,omitempty"`
	// Sponsor contains sponsoring marketplace owner information.
	Sponsor *SponsorRequest `json:"sponsor,omitempty"`
}

// SponsorRequest represents the sponsoring marketplace owner associated with an order's
// integration metadata.
type SponsorRequest struct {
	// ID is the MercadoPago user ID of the sponsoring marketplace owner. Type: string.
	ID string `json:"id,omitempty"`
}
