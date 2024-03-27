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

// Client contains the methods to interact with the Payment's refund API.
type Client interface {
	// Get gets a specific refund by payment id and refund id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds/{refund_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds_refund_id/get
	Get(ctx context.Context, paymentID, refundID int) (*Response, error)

	// List gets a refund list by payment id.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds/get
	List(ctx context.Context, paymentID int) ([]Response, error)

	// Create a refund by payment id.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds/post
	Create(ctx context.Context, paymentID int) (*Response, error)

	// CreatePartialRefund create a partial refund by payment id.
	// It is a post request to the endpoint: https://api.mercadopago.com/v1/payments/{id}/refunds
	// Reference: https://www.mercadopago.com/developers/en/reference/chargebacks/_payments_id_refunds/post
	CreatePartialRefund(ctx context.Context, paymentID int, amount float64) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient contains the methods to interact with the Refund's API.
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
