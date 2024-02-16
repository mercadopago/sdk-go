package payment

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/payments"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
)

// Client contains the methods to interact with the Payments API.
type Client interface {
	// Create creates a new payment.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments/post/
	Create(ctx context.Context, request Request) (*Response, error)

	// Search searches for payments.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/search
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_search/get/
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// Get gets a payment by its ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_id/get/
	Get(ctx context.Context, id int64) (*Response, error)

	// Cancel cancels a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	Cancel(ctx context.Context, id int64) (*Response, error)

	// Capture captures a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	Capture(ctx context.Context, id int64) (*Response, error)

	// CaptureAmount captures amount of a payment by its ID.
	// It is a put request to the endpoint: https://api.mercadopago.com/v1/payments/{id}
	CaptureAmount(ctx context.Context, id int64, amount float64) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Payments API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	res, err := baseclient.Post[*Response](ctx, c.cfg, urlBase, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Search(ctx context.Context, dto SearchRequest) (*SearchResponse, error) {
	params := dto.Parameters()

	url, err := url.Parse(urlSearch)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	url.RawQuery = params

	res, err := baseclient.Get[*SearchResponse](ctx, c.cfg, url.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Get(ctx context.Context, id int64) (*Response, error) {
	conv := strconv.Itoa(int(id))

	res, err := baseclient.Get[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", conv, 1))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Cancel(ctx context.Context, id int64) (*Response, error) {
	dto := &CancelRequest{Status: "cancelled"}
	conv := strconv.Itoa(int(id))

	res, err := baseclient.Put[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", conv, 1), dto)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Capture(ctx context.Context, id int64) (*Response, error) {
	dto := &CaptureRequest{Capture: true}
	conv := strconv.Itoa(int(id))

	res, err := baseclient.Put[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", conv, 1), dto)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) CaptureAmount(ctx context.Context, id int64, amount float64) (*Response, error) {
	dto := &CaptureRequest{TransactionAmount: amount, Capture: true}
	conv := strconv.Itoa(int(id))

	res, err := baseclient.Put[*Response](ctx, c.cfg, strings.Replace(urlWithID, "{id}", conv, 1), dto)
	if err != nil {
		return nil, err
	}

	return res, nil
}
