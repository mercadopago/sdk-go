package invoice

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/authorized_payments"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
)

// Client contains the methods to interact with the Invoice API.
type Client interface {
	// Get finds an invoice by ID.
	// It is a get request to the endpoint: https://api.mercadopago.com/authorized_payments/{id}
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_authorized_payments_id/get
	Get(ctx context.Context, id string) (*Response, error)

	// Search the invoices for a subscriptions by different parameters.
	// It is a get request to the endpoint: https://api.mercadopago.com/authorized_payments/search
	// Reference: https://www.mercadopago.com/developers/en/reference/subscriptions/_authorized_payments_search/get
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Invoice API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		PathParams: pathParams,
		Method:     http.MethodGet,
		URL:        urlWithID,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	queryParameters := request.GetParams()

	requestData := httpclient.RequestData{
		QueryParams: queryParameters,
		Method:      http.MethodGet,
		URL:         urlSearch,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
