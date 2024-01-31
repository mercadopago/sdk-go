package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/users/me"

// Client contains the method to interact with the User API.
type Client interface {
	// Get get user information.
	// It is a get request to the endpoint: https://api.mercadopago.com/users/me
	Get(ctx context.Context) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	config *config.Config
}

// NewClient returns a new User API Client.
func NewClient(c *config.Config) Client {
	return &client{
		config: c,
	}
}

func (c *client) Get(ctx context.Context) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := httpclient.Send(ctx, c.config, req)
	if err != nil {
		return nil, err
	}

	var formatted *Response
	if err := json.Unmarshal(res, &formatted); err != nil {
		return nil, err
	}

	return formatted, nil
}
