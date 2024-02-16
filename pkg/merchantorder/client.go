package merchantorder

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
	urlBase   = "https://api.mercadopago.com/merchant_orders"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/:id"
)

// Client contains the methods to interact with the Payment's refund API.
type Client interface {
	// Get gets a specific refund by payment id and refund id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds/{refund_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds_refund_id/get
	Get(ctx context.Context, id int64) (*Response, error)

	// Search searches for payments.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/search
	// Reference: https://www.mercadopago.com/developers/en/reference/payments/_payments_search/get/
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)

	// List gets a refund list by payment id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds/get
	Update(ctx context.Context, paymentID int64) ([]Response, error)

	// Create create a refund by payment id.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds/post
	Create(ctx context.Context, request Request) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// Client contains the methods to interact with the Refund's API.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context, id int64) (*Response, error) {
	param := map[string]string{
		"id": strconv.Itoa(int(id)),
	}

	res, err := baseclient.Get[*Response](ctx, c.cfg, urlWithID, baseclient.WithPathParams(param))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	params := request.Parameters()

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

func (c *client) Update(ctx context.Context, paymentID int64) ([]Response, error) {
	convertedRefundID := strconv.Itoa(int(paymentID))

	res, err := baseclient.Get[[]Response](ctx, c.cfg, strings.Replace(urlBase, "{id}", convertedRefundID, 1))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	res, err := baseclient.Post[*Response](ctx, c.cfg, urlBase, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
