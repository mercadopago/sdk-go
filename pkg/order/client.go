// Package order provides a client for the MercadoPago Orders API (v1).
//
// The Orders API allows creating, processing, capturing, cancelling, and refunding
// payment orders. An order groups one or more payment transactions together with
// item details, payer information, and configuration options such as checkout URLs
// and installment settings.
//
// Use [NewClient] to create a client, then call its methods to interact with the API.
//
//	cfg, _ := config.New("ACCESS_TOKEN")
//	client := order.NewClient(cfg)
//	resp, err := client.Create(ctx, order.Request{...})
package order

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase              = "https://api.mercadopago.com/v1/orders"
	urlWithOrderID       = urlBase + "/" + "{orderID}"
	urlTransaction       = urlWithOrderID + "/transactions"
	urlProcess           = urlWithOrderID + "/process"
	urlUpdateTransaction = urlTransaction + "/{transactionID}"
	urlCapture           = urlWithOrderID + "/capture"
	urlCancel            = urlWithOrderID + "/cancel"
	urlRefund            = urlWithOrderID + "/refund"
	urlDeleteTransaction = urlTransaction + "/{transactionID}"
)

// Client contains the methods to interact with the MercadoPago Orders API.
type Client interface {
	// Create creates a new order.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a POST request to the endpoint: https://api.mercadopago.com/v1/orders
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api/create-order/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Get retrieves an existing order by its unique identifier.
	// It is a GET request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}
	Get(ctx context.Context, orderID string) (*Response, error)

	// Process triggers the processing of an order so that its payment transactions are executed.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a POST request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/process
	Process(ctx context.Context, orderID string) (*Response, error)

	// Cancel cancels an order that has not yet been fully processed.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a POST request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/cancel
	Cancel(ctx context.Context, orderID string) (*Response, error)

	// Capture captures a previously authorized order, settling its payment transactions.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a POST request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/capture
	Capture(ctx context.Context, orderID string) (*Response, error)

	// Refund initiates a full or partial refund for an order. Pass a nil [RefundRequest]
	// for a full refund, or specify individual transaction amounts for a partial refund.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a POST request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/refund
	Refund(ctx context.Context, orderID string, request *RefundRequest) (*Response, error)

	// CreateTransaction adds a new payment transaction to an existing order.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a POST request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/transactions
	CreateTransaction(ctx context.Context, orderID string, request TransactionRequest) (*TransactionResponse, error)

	// UpdateTransaction modifies an existing payment transaction within an order.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	// It is a PUT request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/transactions/{transactionID}
	UpdateTransaction(ctx context.Context, orderID string, transactionID string, request PaymentRequest) (*PaymentResponse, error)

	// DeleteTransaction removes a payment transaction from an order.
	// It is a DELETE request to the endpoint: https://api.mercadopago.com/v1/orders/{orderID}/transactions/{transactionID}
	DeleteTransaction(ctx context.Context, orderID string, transactionID string) error

	// Search finds orders that match the given filter criteria and pagination parameters.
	// It is a GET request to the endpoint: https://api.mercadopago.com/v1/orders
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient returns a new Orders API [Client] configured with the provided [config.Config].
// The config must contain a valid access token for authenticating requests.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlBase,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, orderID string) (*Response, error) {
	pathParams := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlWithOrderID,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Process(ctx context.Context, orderID string) (*Response, error) {
	pathParams := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlProcess,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CreateTransaction(ctx context.Context, orderID string, request TransactionRequest) (*TransactionResponse, error) {
	pathParams := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlTransaction,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*TransactionResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) UpdateTransaction(ctx context.Context, orderID string, transactionID string, request PaymentRequest) (*PaymentResponse, error) {
	pathParams := map[string]string{
		"orderID":       orderID,
		"transactionID": transactionID,
	}
	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlUpdateTransaction,
		PathParams: pathParams,
	}

	resource, err := httpclient.DoRequest[*PaymentResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Cancel(ctx context.Context, orderID string) (*Response, error) {
	pathParam := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlCancel,
		PathParams: pathParam,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Capture(ctx context.Context, orderID string) (*Response, error) {
	pathParam := map[string]string{
		"orderID": orderID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlCapture,
		PathParams: pathParam,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Refund(ctx context.Context, orderID string, request *RefundRequest) (*Response, error) {
	pathParam := map[string]string{
		"orderID": orderID,
	}

	var requestBody any = nil
	if request != nil {
		requestBody = request
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlRefund,
		PathParams: pathParam,
		Body:       requestBody,
	}

	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) DeleteTransaction(ctx context.Context, orderID string, transactionID string) error {
	pathParam := map[string]string{
		"orderID":       orderID,
		"transactionID": transactionID,
	}
	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlDeleteTransaction,
		PathParams: pathParam,
		Body:       nil,
	}

	_, err := httpclient.DoRequest[any](ctx, c.cfg, requestData)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		Method:      http.MethodGet,
		URL:         urlBase,
		QueryParams: queryParams,
	}

	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
