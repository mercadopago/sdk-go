// Package payment provides a client for the MercadoPago Payments API.
//
// It supports creating, searching, retrieving, cancelling, and capturing payments.
// Use [NewClient] to obtain a [Client] and interact with the API.
//
// For full API documentation see https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/create-payment/post.
package payment

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/payments"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
)

// Client provides methods to interact with the MercadoPago Payments API.
// Create one via [NewClient].
type Client interface {
	// Create submits a new payment to the MercadoPago Payments API.
	// The returned [Response] contains the full payment resource including its assigned ID and status.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	//
	// POST https://api.mercadopago.com/v1/payments
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/create-payment/post
	Create(ctx context.Context, request Request) (*Response, error)

	// Search retrieves a paginated list of payments that match the criteria specified in [SearchRequest].
	//
	// GET https://api.mercadopago.com/v1/payments/search
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/search-payments/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Get retrieves a single payment by its unique numeric ID.
	//
	// GET https://api.mercadopago.com/v1/payments/{id}
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/get-payment/get
	Get(ctx context.Context, id int) (*Response, error)

	// Cancel transitions a payment to the "cancelled" status.
	// Only payments that have not yet been approved can be cancelled.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	//
	// PUT https://api.mercadopago.com/v1/payments/{id}
	Cancel(ctx context.Context, id int) (*Response, error)

	// Capture confirms a previously authorized payment, settling the full transaction amount.
	// This is used when the payment was created with capture=false (two-step flow).
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	//
	// PUT https://api.mercadopago.com/v1/payments/{id}
	Capture(ctx context.Context, id int) (*Response, error)

	// CaptureAmount confirms a previously authorized payment for a specific amount,
	// which may differ from the originally authorized transaction amount.
	// This is used in the two-step capture flow when a partial capture is needed.
	// To set a custom X-Idempotency-Key, attach one to ctx with [requestoptions.WithIdempotencyKey].
	//
	// PUT https://api.mercadopago.com/v1/payments/{id}
	CaptureAmount(ctx context.Context, id int, amount float64) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates a new Payments API client using the provided [config.Config].
// The configuration must include valid credentials for authenticating with the MercadoPago API.
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

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		Method:      http.MethodGet,
		URL:         urlSearch,
		QueryParams: queryParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, id int) (*Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(id),
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Cancel(ctx context.Context, id int) (*Response, error) {
	request := &CancelRequest{Status: "cancelled"}

	pathParams := map[string]string{
		"id": strconv.Itoa(id),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Capture(ctx context.Context, id int) (*Response, error) {
	request := &CaptureRequest{Capture: true}

	pathParams := map[string]string{
		"id": strconv.Itoa(id),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CaptureAmount(ctx context.Context, id int, amount float64) (*Response, error) {
	request := &CaptureRequest{TransactionAmount: amount, Capture: true}

	pathParams := map[string]string{
		"id": strconv.Itoa(id),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPut,
		URL:        urlWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
