package cardtoken

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseURL = "https://api.mercadopago.com/v1/"
	urlGet  = baseURL + "/v1/card_tokens/{id}"
	urlPost = baseURL + "/v1/card_tokens"
)

type Client interface {
	Get(ctx context.Context, id string) (*Response, error)
	Create(ctx context.Context, request Request) (*Response, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Get(ctx context.Context, id string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlGet, nil)
	if err != nil {
		return nil, err
	}

	r, err := httpclient.Send(ctx, c.cfg, req)
	if err != nil {
		return nil, err
	}

	var res *Response
	if err := json.Unmarshal(r, res); err != nil {
		return nil, fmt.Errorf("error unmarshaling card token response: %w", err)
	}

	return res, nil
}

func (c *client) Create(ctx context.Context, request Request) (*Response, error) {
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling card token request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlPost, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("error creating card token request: %w", err)
	}

	r, err := httpclient.Send(ctx, c.cfg, req)
	if err != nil {
		return nil, err
	}

	var res *Response
	if err := json.Unmarshal(r, res); err != nil {
		return nil, fmt.Errorf("error unmarshaling card token response: %w", err)
	}

	return res, nil
}
