// Package point provides a client for interacting with the MercadoPago Point Integration API.
//
// The Point API enables in-person payment processing through MercadoPago Point devices
// (card readers). It allows creating payment intents that are sent to a physical device
// for card-present transactions, managing device configurations, and tracking payment
// intent status.
//
// For more information, see the MercadoPago Point Integration API reference:
// https://www.mercadopago.com.ar/developers/en/reference/in-person-payments/point/orders/create-order/post
package point

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase                = "https://api.mercadopago.com/point"
	urlDevices             = urlBase + "/integration-api/devices"
	urlPaymentIntent       = urlDevices + "/{device_id}/payment-intents"
	urlPaymentIntentGet    = urlBase + "/integration-api/payment-intents/{payment_intent_id}"
	urlPaymentIntentCancel = urlDevices + "/{device_id}/payment-intents/{payment_intent_id}"
	urlDevicesWithID       = urlDevices + "/{device_id}"
)

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// Client defines the interface for interacting with the MercadoPago Point Integration API.
// It provides methods to create, retrieve, and cancel payment intents on Point devices,
// as well as manage device listings and operating modes.
type Client interface {
	// Create creates a payment intent on the specified Point device. The payment intent
	// contains the transaction details (amount, description, payment method) and is sent
	// to the physical device for the buyer to complete the payment.
	//
	// It performs a POST request to: https://api.mercadopago.com/point/integration-api/devices/{device_id}/payment-intents
	//
	// Reference: https://www.mercadopago.com.ar/developers/en/reference/in-person-payments/point/orders/create-order/post
	Create(ctx context.Context, deviceID string, request Request) (*Response, error)

	// Get retrieves the current state of a payment intent by its unique identifier.
	// This is used to check the payment status after a payment intent has been created.
	//
	// It performs a GET request to: https://api.mercadopago.com/point/integration-api/payment-intents/{payment_intent_id}
	//
	// Reference: https://www.mercadopago.com.ar/developers/en/reference/in-person-payments/point/orders/get-order/get
	Get(ctx context.Context, paymentIntentID string) (*Response, error)

	// Cancel cancels a pending payment intent on a specific device. This is used when
	// a transaction needs to be aborted before the buyer completes the payment on the device.
	//
	// It performs a DELETE request to: https://api.mercadopago.com/point/integration-api/devices/{device_id}/payment-intents/{payment_intent_id}
	//
	// Reference: https://www.mercadopago.com.ar/developers/en/reference/in-person-payments/point/orders/cancel-order/post
	Cancel(ctx context.Context, deviceID, paymentIntentID string) (*CancelResponse, error)

	// ListDevices retrieves all Point devices associated with the authenticated account.
	// The response includes device identifiers, operating modes, and store associations.
	//
	// It performs a GET request to: https://api.mercadopago.com/point/integration-api/devices
	//
	// Reference: https://www.mercadopago.com.ar/developers/en/reference/in-person-payments/point/terminals/get-terminals/get
	ListDevices(ctx context.Context) (*DevicesResponse, error)

	// UpdateOperatingMode changes the operating mode of a specific Point device.
	// Use "PDV" for integrated mode (API-driven payments) or "STANDALONE" for
	// standalone mode (device-driven payments without API integration).
	//
	// It performs a PATCH request to: https://api.mercadopago.com/point/integration-api/devices/{device-id}
	//
	// Reference: https://www.mercadopago.com.ar/developers/en/reference/in-person-payments/point/terminals/update-operation-mode/patch
	UpdateOperatingMode(ctx context.Context, deviceID, operatingMode string) (*OperatingModeResponse, error)
}

// NewClient creates and returns a new Point Integration API [Client] configured with
// the provided [config.Config]. The config must contain a valid access token for
// authenticating requests to the MercadoPago API.
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, deviceID string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"device_id": deviceID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlPaymentIntent,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, paymentIntentID string) (*Response, error) {
	pathParams := map[string]string{
		"payment_intent_id": paymentIntentID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlPaymentIntentGet,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Cancel(ctx context.Context, deviceID string, paymentIntentID string) (*CancelResponse, error) {
	pathParams := map[string]string{
		"device_id":         deviceID,
		"payment_intent_id": paymentIntentID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlPaymentIntentCancel,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*CancelResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ListDevices(ctx context.Context) (*DevicesResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    urlDevices,
	}
	resource, err := httpclient.DoRequest[*DevicesResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) UpdateOperatingMode(ctx context.Context, deviceID, operatingMode string) (*OperatingModeResponse, error) {
	request := &OperatingModeRequest{OperatingMode: operatingMode}

	pathParams := map[string]string{
		"device_id": deviceID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPatch,
		URL:        urlDevicesWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*OperatingModeResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
