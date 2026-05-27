// Package advancedpayment provides a client for the MercadoPago Advanced Payments API.
//
// Advanced payments enable marketplace integrations to collect a single payment and
// distribute funds among multiple sellers (disbursements). The package supports
// two-step flows (authorise → capture) and individual disbursement release-date control.
//
// For more information, see the MercadoPago Marketplace documentation:
// https://www.mercadopago.com/developers/en/reference
package advancedpayment

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase          = "https://api.mercadopago.com/v1/advanced_payments"
	urlSearch        = urlBase + "/search"
	urlWithID        = urlBase + "/{id}"
	urlDisburses     = urlBase + "/{id}/disburses"
)

// Client defines the interface for interacting with the MercadoPago Advanced Payments API.
type Client interface {
	// Create submits a new advanced (split) payment.
	//
	// POST https://api.mercadopago.com/v1/advanced_payments
	Create(ctx context.Context, request Request) (*Response, error)

	// Get retrieves an advanced payment by its unique ID.
	//
	// GET https://api.mercadopago.com/v1/advanced_payments/{id}
	Get(ctx context.Context, id int) (*Response, error)

	// Search retrieves a paginated list of advanced payments matching the given criteria.
	//
	// GET https://api.mercadopago.com/v1/advanced_payments/search
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Update updates an existing advanced payment with arbitrary fields.
	//
	// PUT https://api.mercadopago.com/v1/advanced_payments/{id}
	Update(ctx context.Context, id int, request UpdateRequest) (*Response, error)

	// Capture finalises a previously authorised advanced payment (two-step flow).
	//
	// PUT https://api.mercadopago.com/v1/advanced_payments/{id}
	Capture(ctx context.Context, id int) (*Response, error)

	// Cancel transitions a pending advanced payment to the "cancelled" status.
	//
	// PUT https://api.mercadopago.com/v1/advanced_payments/{id}
	Cancel(ctx context.Context, id int) (*Response, error)

	// UpdateReleaseDate changes the money release date for all disbursements.
	//
	// POST https://api.mercadopago.com/v1/advanced_payments/{id}/disburses
	UpdateReleaseDate(ctx context.Context, id int, releaseDate string) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new Advanced Payments API [Client] configured with
// the provided [config.Config].
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
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

func (c *client) Get(ctx context.Context, id int) (*Response, error) {
	pathParams := map[string]string{"id": strconv.Itoa(id)}
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

func (c *client) Update(ctx context.Context, id int, request UpdateRequest) (*Response, error) {
	pathParams := map[string]string{"id": strconv.Itoa(id)}
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
	pathParams := map[string]string{"id": strconv.Itoa(id)}
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

func (c *client) Cancel(ctx context.Context, id int) (*Response, error) {
	request := &CancelRequest{Status: "cancelled"}
	pathParams := map[string]string{"id": strconv.Itoa(id)}
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

func (c *client) UpdateReleaseDate(ctx context.Context, id int, releaseDate string) (*Response, error) {
	request := &UpdateReleaseDateRequest{MoneyReleaseDate: releaseDate}
	pathParams := map[string]string{"id": strconv.Itoa(id)}
	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlDisburses,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
