// This package is useful for manage payments.
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

// Client contains the methods to interact with the Payments API.
type Client interface {
	// Create a new payment.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments.
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments/post/.
	Create(ctx context.Context, request Request) (*Response, error)

	// Search for payments.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/search.
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_search/get/.
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Get a payment by id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}.
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_id/get/.
	Get(ctx context.Context, id int) (*Response, error)

	// Cancel a payment by id.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}.
	Cancel(ctx context.Context, id int) (*Response, error)

	// Capture a payment by id.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}.
	Capture(ctx context.Context, id int) (*Response, error)

	// CaptureAmount captures amount of a payment by id.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}.
	CaptureAmount(ctx context.Context, id int, amount float64) (*Response, error)
}

type client struct {
	cfg *config.Config
}

// NewClient returns an implementation of Client.
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
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParams := request.GetParams()

	requestData := httpclient.RequestData{
		QueryParams: queryParams,
		Method:      http.MethodGet,
		URL:         urlSearch,
	}
	result, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Get(ctx context.Context, id int) (*Response, error) {
	pathParams := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	requestData := httpclient.RequestData{
		PathParams: pathParams,
		Method:     http.MethodGet,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Cancel(ctx context.Context, id int) (*Response, error) {
	request := &CancelRequest{Status: "cancelled"}

	pathParams := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		PathParams: pathParams,
		Method:     http.MethodPut,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) Capture(ctx context.Context, id int) (*Response, error) {
	request := &CaptureRequest{Capture: true}

	pathParams := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		PathParams: pathParams,
		Method:     http.MethodPut,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) CaptureAmount(ctx context.Context, id int, amount float64) (*Response, error) {
	request := &CaptureRequest{TransactionAmount: amount, Capture: true}

	pathParams := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	requestData := httpclient.RequestData{
		Body:       request,
		PathParams: pathParams,
		Method:     http.MethodPut,
		URL:        urlWithID,
	}
	result, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
