// Package user provides a client for retrieving the authenticated MercadoPago user's information.
//
// The User API returns account details such as name, email, country, and site for the
// user associated with the access token in the SDK configuration.
package user

import (
	"context"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/users/me"

// Client defines the interface for interacting with the MercadoPago User API.
type Client interface {
	// Get retrieves the account information of the authenticated MercadoPago user.
	// The returned [Response] contains the user's name, email, country, and site information.
	//
	// It performs a GET request to: https://api.mercadopago.com/users/me
	Get(ctx context.Context) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new User API [Client] configured with the provided [config.Config].
// The config must contain a valid access token for authenticating requests to the MercadoPago API.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context) (*Response, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    url,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
