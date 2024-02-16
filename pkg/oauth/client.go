package oauth

import (
	"context"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	baseUrl  = "https://api.mercadopago.com/oauth/token"
	urlOAuth = "https://auth.mercadopago.com/authorization"
)

// Client contains the method to interact with the OAuth API.
type Client interface {

	// Create Oauth credentials to operate on behalf of a seller
	// It is a Post request to the endpoint: "https://api.mercadopago.com/oauth/token"
	// Reference: https://www.mercadopago.com.br/developers/en/reference/oauth/_oauth_token/post
	Create(ctx context.Context, authorizationCode, redirectURI string) (*Response, error)

	// Get URL for Oauth authorization.
	GetAuthorizationURL(ctx context.Context, clientID, redirectURI string) string

	// Refresh token received when you create credentials.
	// It is a Post request to the endpoint: "https://api.mercadopago.com/oauth/token"
	// Reference: https://www.mercadopago.com.br/developers/en/reference/oauth/_oauth_token/post
	Refresh(ctx context.Context, refreshToken string) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new User API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, authorizationCode, redirectURI string) (*Response, error) {
	request := &Request{
		ClientSecret: c.cfg.AccessToken,
		Code:         authorizationCode,
		RedirectURI:  redirectURI,
		GrantType:    "authorization_code",
	}

	res, err := httpclient.Post[Response](ctx, c.cfg, baseUrl, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetAuthorizationURL(ctx context.Context, clientID, redirectURI string) string {
	params := map[string]string{
		"client_id":     clientID,
		"response_type": "code",
		"platform_id":   "mp",
		"redirect_uri":  redirectURI,
	}

	host, err := url.Parse(urlOAuth)
	if err != nil {
		return ""
	}

	queryParams := url.Values{}

	for k, v := range params {
		queryParams.Add(k, v)
	}

	host.RawQuery = queryParams.Encode()

	return host.String()
}

func (c *client) Refresh(ctx context.Context, refreshToken string) (*Response, error) {
	request := &RefreshTokenRequest{
		ClientSecret: c.cfg.AccessToken,
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}

	res, err := httpclient.Post[Response](ctx, c.cfg, baseUrl, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
