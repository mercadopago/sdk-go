package preference

import (
	"time"
)

type Response struct {
	ID                 string                    `json:"id"`
	Items              []PreferenceItem          `json:"items"`
	Payer              PreferencePayer           `json:"payer"`
	ClientID           string                    `json:"client_id"`
	PaymentMethods     PreferencePaymentMethods   `json:"payment_methods"`
	BackUrls           PreferenceBackUrls         `json:"back_urls"`
	Shipments          PreferenceShipments        `json:"shipments"`
	NotificationURL    string                    `json:"notification_url"`
	StatementDescriptor string                    `json:"statement_descriptor"`
	ExternalReference  string                    `json:"external_reference"`
	Expires            bool                      `json:"expires"`
	DateOfExpiration   time.Time                 `json:"date_of_expiration"`
	ExpirationDateFrom time.Time                 `json:"expiration_date_from"`
	ExpirationDateTo   time.Time                 `json:"expiration_date_to"`
	CollectorID        int64                     `json:"collector_id"`
	Marketplace        string                    `json:"marketplace"`
	MarketplaceFee     float64                   `json:"marketplace_fee"`
	AdditionalInfo     string                    `json:"additional_info"`
	AutoReturn         string                    `json:"auto_return"`
	OperationType      string                    `json:"operation_type"`
	DifferentialPricing PreferenceDifferentialPricing `json:"differential_pricing"`
	ProcessingModes    []string                  `json:"processing_modes"`
	BinaryMode         bool                      `json:"binary_mode"`
	Taxes              []PreferenceTax           `json:"taxes"`
	Tracks             []PreferenceTrack         `json:"tracks"`
	Metadata           map[string]interface{}    `json:"metadata"`
	InitPoint          string                    `json:"init_point"`
	SandboxInitPoint   string                    `json:"sandbox_init_point"`
	DateCreated        time.Time                 `json:"date_created"`
}

// PreferenceItem representa um item comprado.
type PreferenceItem struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	PictureURL  string      `json:"picture_url"`
	CategoryID  string      `json:"category_id"`
	Quantity    int         `json:"quantity"`
	UnitPrice   float64  `json:"unit_price"`
	CurrencyID  string      `json:"currency_id"`
}

// PreferencePayer contém informações do pagador na preferência.
type PreferencePayer struct {
	Name          string          `json:"name"`
	Surname       string          `json:"surname"`
	Email         string          `json:"email"`
	Phone         Phone           `json:"phone"`
	Identification Identification `json:"identification"`
	Address       Address         `json:"address"`
	DateCreated   time.Time       `json:"date_created"`
	LastPurchase  time.Time       `json:"last_purchase"`
}

// PreferencePaymentMethods contém informações sobre métodos de pagamento na preferência.
type PreferencePaymentMethods struct {
	ExcludedPaymentMethods   []PreferencePaymentMethod `json:"excluded_payment_methods"`
	ExcludedPaymentTypes     []PreferencePaymentType   `json:"excluded_payment_types"`
	DefaultPaymentMethodID   string                    `json:"default_payment_method_id"`
	Installments             int                       `json:"installments"`
	DefaultInstallments      int                       `json:"default_installments"`
}

// PreferencePaymentMethod contém informações sobre o método de pagamento na preferência.
type PreferencePaymentMethod struct {
	ID string `json:"id"`
}

// PreferencePaymentType contém informações sobre o tipo de pagamento na preferência.
type PreferencePaymentType struct {
	ID string `json:"id"`
}

// PreferenceBackUrls contém URLs de retorno da preferência.
type PreferenceBackUrls struct {
	Success string `json:"success"`
	Pending string `json:"pending"`
	Failure string `json:"failure"`
}

// PreferenceShipments contém informações de envio da preferência.
type PreferenceShipments struct {
	Mode                  string                     `json:"mode"`
	LocalPickup           bool                       `json:"local_pickup"`
	Dimensions            string                     `json:"dimensions"`
	DefaultShippingMethod string                     `json:"default_shipping_method"`
	FreeMethods           []PreferenceFreeMethod      `json:"free_methods"`
	Cost                  float64                  `json:"cost"`
	FreeShipping          bool                       `json:"free_shipping"`
	ReceiverAddress       PreferenceReceiverAddress   `json:"receiver_address"`
	ExpressShipment       bool                       `json:"express_shipment"`
}

// PreferenceFreeMethod contém informações sobre métodos de envio gratuitos.
type PreferenceFreeMethod struct {
	ID int64 `json:"id"`
}

// PreferenceReceiverAddress representa um endereço de envio.
type PreferenceReceiverAddress struct {
	Address
	CountryName string `json:"country_name"`
	StateName   string `json:"state_name"`
	Floor       string `json:"floor"`
	Apartment   string `json:"apartment"`
	CityName    string `json:"city_name"`
}

// PreferenceDifferentialPricing contém informações sobre a configuração de precificação diferencial na preferência.
type PreferenceDifferentialPricing struct {
	ID int64 `json:"id"`
}

// PreferenceTax contém informações sobre impostos na preferência.
type PreferenceTax struct {
	Type  string     `json:"type"`
	Value float64 `json:"value"`
}

// PreferenceTrack representa um rastreamento a ser executado durante a interação do usuário no fluxo de Checkout.
type PreferenceTrack struct {
	Type   string                         `json:"type"`
	Values PreferenceTrackValuesResponse   `json:"values"`
}

// PreferenceTrackValuesRequest contém os valores dos rastreamentos a serem executados durante a interação do usuário no fluxo de Checkout.
type PreferenceTrackValuesResponse struct {
	ConversionID    string `json:"conversion_id"`
	ConversionLabel string `json:"conversion_label"`
	PixelID         string `json:"pixel_id"`
}

// Phone representa um número de telefone.
type Phone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

// Identification é um tipo base que representa identificações, como identificação do cliente.
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// Address representa um endereço.
type Address struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
}

// PreferenceSearch contém informações sobre uma preferência para pesquisa.
type SearchResponse struct {
	ID                  string        `json:"id,omitempty"`
	ClientID            string        `json:"client_id,omitempty"`
	CollectorID         int64         `json:"collector_id,omitempty"`
	DateCreated         time.Time     `json:"date_created,omitempty"`
	ExpirationDateFrom  time.Time     `json:"expiration_date_from,omitempty"`
	ExpirationDateTo    time.Time     `json:"expiration_date_to,omitempty"`
	Expires             bool          `json:"expires,omitempty"`
	ExternalReference   string        `json:"external_reference,omitempty"`
	Items               []string      `json:"items,omitempty"`
	LastUpdated         time.Time     `json:"last_updated,omitempty"`
	LiveMode            bool          `json:"live_mode,omitempty"`
	Marketplace         string        `json:"marketplace,omitempty"`
	OperationType       string        `json:"operation_type,omitempty"`
	PayerEmail          string        `json:"payer_email,omitempty"`
	PayerID             string        `json:"payer_id,omitempty"`
	PlatformID          string        `json:"platform_id,omitempty"`
	ProcessingModes     []string      `json:"processing_modes,omitempty"`
	ProductID           string        `json:"product_id,omitempty"`
	Purpose             string        `json:"purpose,omitempty"`
	SiteID              string        `json:"site_id,omitempty"`
	SponsorID           int64         `json:"sponsor_id,omitempty"`
	ShippingMode        string        `json:"shipping_mode,omitempty"`
}

// MPElementsResourcesPage é uma página de pesquisa que contém elementos.
type SearchResponsePage struct {
	Total      int         `json:"total,omitempty"`
	NextOffset int         `json:"next_offset,omitempty"`
	Elements   []SearchResponse `json:"elements,omitempty"`
}