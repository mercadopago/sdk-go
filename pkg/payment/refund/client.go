package refund

import (
	"context"
	"strconv"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseURL = "https://api.mercadopago.com/v1/payments/{id}/"
	getURL  = baseURL + "refunds/{refund_id}"
	listURL = baseURL + "refunds"
	postURL = baseURL + "refunds"
)

// Client contains the methods to interact with the Payment's refund API.
type Client interface {
	// Get gets a specific refund by payment id and refund id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds/{refund_id}
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/_payments_id_refunds_refund_id/get
	Get(ctx context.Context, paymentID, refundID int64) (*Response, error)

	// List gets a refund list by payment id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/_payments_id_refunds/get
	List(ctx context.Context, paymentID int64) ([]Response, error)

	// Create create a refund by payment id.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/_payments_id_refunds/post
	Create(ctx context.Context, request Request, paymentID int64) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Payment's refund API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, request Request, paymentID int64) (*Response, error) {
	conv := strconv.Itoa(int(paymentID))

	res, err := httpclient.Post[Response](ctx, c.cfg, strings.Replace(postURL, "{id}", conv, 1), request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Get(ctx context.Context, paymentID, refundID int64) (*Response, error) {
	conv := strconv.Itoa(int(paymentID))
	refundConv := strconv.Itoa(int(refundID))

	url := strings.NewReplacer("{id}", conv, "{refund_id}", refundConv).Replace(getURL)

	res, err := httpclient.Get[Response](ctx, c.cfg, url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) List(ctx context.Context, paymentID int64) ([]Response, error) {
	conv := strconv.Itoa(int(paymentID))

	res, err := httpclient.Get[[]Response](ctx, c.cfg, strings.Replace(listURL, "{id}", conv, 1))
	if err != nil {
		return nil, err
	}

	return *res, nil
}
