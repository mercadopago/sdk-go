// Package chargeback provides a client for the MercadoPago Chargebacks API.
//
// Chargebacks are dispute records initiated by cardholders through their issuing bank.
// This package provides read-only access to retrieve and search chargebacks.
//
// For more information, see:
// https://www.mercadopago.com.br/developers/en/reference/chargebacks/
package chargeback

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase   = "https://api.mercadopago.com/v1/chargebacks"
	urlSearch = urlBase + "/search"
	urlWithID = urlBase + "/{id}"
)

// Client defines the interface for interacting with the MercadoPago Chargebacks API.
// It provides read-only methods to retrieve and search chargeback disputes.
type Client interface {
	// Get retrieves a chargeback by its unique identifier.
	//
	// GET https://api.mercadopago.com/v1/chargebacks/{id}
	//
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/
	Get(ctx context.Context, id string) (*Response, error)

	// Search finds chargebacks matching the filters specified in [SearchRequest].
	// Results are paginated and returned as a [SearchResponse].
	//
	// GET https://api.mercadopago.com/v1/chargebacks/search
	//
	// Reference: https://www.mercadopago.com.br/developers/en/reference/chargebacks/
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new Chargebacks API [Client] configured with
// the provided [config.Config]. The config must contain a valid access token.
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	pathParams := map[string]string{"id": id}

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
