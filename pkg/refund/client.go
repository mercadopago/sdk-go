// Package refund provides a client for the MercadoPago Refunds API.
//
// It supports creating full and partial refunds, listing refunds for a payment,
// and retrieving individual refund details. Refunds are always associated with
// an existing payment identified by its numeric ID.
// Use [NewClient] to obtain a [Client] and interact with the API.
//
// For full API documentation see https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/create-refund/post
package refund

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/payments/{id}/refunds"
	urlWithID = urlBase + "/{refund_id}"
)

// Client provides methods to interact with the MercadoPago Refunds API.
// Refunds are scoped to a specific payment and can be full or partial.
// Create one via [NewClient].
type Client interface {
	// Get retrieves a specific refund by its payment ID and refund ID.
	//
	// GET https://api.mercadopago.com/v1/payments/{id}/refunds/{refund_id}
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/get-refund/get
	Get(ctx context.Context, paymentID, refundID int) (*Response, error)

	// List retrieves all refunds associated with the given payment ID.
	//
	// GET https://api.mercadopago.com/v1/payments/{id}/refunds
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/get-refunds/get
	List(ctx context.Context, paymentID int) ([]Response, error)

	// Create issues a full refund for the given payment ID, returning the full
	// transaction amount to the payer.
	//
	// POST https://api.mercadopago.com/v1/payments/{id}/refunds
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/create-refund/post
	Create(ctx context.Context, paymentID int) (*Response, error)

	// CreatePartialRefund issues a partial refund for the given payment ID,
	// returning the specified amount to the payer. Multiple partial refunds
	// can be created until the total transaction amount is fully refunded.
	//
	// POST https://api.mercadopago.com/v1/payments/{id}/refunds
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/create-refund/post
	CreatePartialRefund(ctx context.Context, paymentID int, amount float64) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates a new Refunds API client using the provided [config.Config].
// The configuration must include valid credentials for authenticating with the MercadoPago API.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context, paymentID, refundID int) (*Response, error) {
	pathParams := map[string]string{
		"id":        strconv.Itoa(paymentID),
		"refund_id": strconv.Itoa(refundID),
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

func (c *client) List(ctx context.Context, paymentID int) ([]Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(paymentID),
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlBase,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[[]Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Create(ctx context.Context, paymentID int) (*Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(paymentID),
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlBase,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CreatePartialRefund(ctx context.Context, paymentID int, amount float64) (*Response, error) {
	request := &Request{Amount: amount}
	pathParams := map[string]string{
		"id": strconv.Itoa(paymentID),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlBase,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
