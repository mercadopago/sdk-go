// This package is useful to return identification types.
package identificationtype

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/v1/identification_types"

// Client contains the methods to interact with the Identification Types API.
type Client interface {
	// List all identification types available by country containing details of each one.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/identification_types.
	// Reference: https://www.mercadopago.com/developers/en/reference/identification_types/_identification_types/get.
	List(ctx context.Context) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Identification Types API Client.
func NewClient(cfg *config.Config) Client {
	return &client{cfg}
}

func (c *client) List(ctx context.Context) ([]Response, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    url,
	}
	result, err := httpclient.DoRequest[[]Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
