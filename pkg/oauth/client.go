// Package oauth provides a client for interacting with the MercadoPago OAuth API.
//
// OAuth is used to obtain credentials that allow operating on behalf of a seller
// (marketplace or platform scenario). The flow involves redirecting the seller to
// the MercadoPago authorization URL, receiving an authorization code, and exchanging
// it for access and refresh tokens.
//
// For more information, see the MercadoPago OAuth API reference:
// https://www.mercadopago.com/developers/en/reference/authentication/oauth/_oauth_token/post
package oauth

import (
	"context"
	"net/http"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase = "https://api.mercadopago.com/oauth/token"
	urlAuth = "https://auth.mercadopago.com/authorization"
)

// Client defines the interface for interacting with the MercadoPago OAuth API.
// It provides methods to generate authorization URLs, create OAuth credentials,
// and refresh expired tokens.
type Client interface {
	// Create exchanges an authorization code for OAuth credentials (access token and refresh token)
	// that allow operating on behalf of a seller. The authorizationCode is obtained from the
	// OAuth authorization flow, and redirectURI must match the one used when generating the
	// authorization URL.
	//
	// It performs a POST request to: https://api.mercadopago.com/oauth/token
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/authentication/oauth/_oauth_token/post
	Create(ctx context.Context, authorizationCode, redirectURI string) (*Response, error)

	// GetAuthorizationURL builds and returns the MercadoPago authorization URL that the seller
	// must visit to grant permissions. The clientID identifies the application, redirectURI is
	// where the seller is sent after authorization, and state is a value preserved across the
	// redirect to prevent CSRF attacks.
	GetAuthorizationURL(clientID, redirectURI, state string) string

	// Refresh exchanges a refresh token for a new set of OAuth credentials, extending the
	// session without requiring the seller to re-authorize. The refreshToken is obtained
	// from the original [Client.Create] response.
	//
	// It performs a POST request to: https://api.mercadopago.com/oauth/token
	//
	// Reference: https://www.mercadopago.com/developers/en/reference/authentication/oauth/_oauth_token/post
	Refresh(ctx context.Context, refreshToken string) (*Response, error)
}

// client is the unexported implementation of [Client].
type client struct {
	cfg *config.Config
}

// NewClient creates and returns a new OAuth API [Client] configured with the provided [config.Config].
// The config must contain a valid access token for authenticating requests to the MercadoPago API.
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

	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlBase,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) GetAuthorizationURL(clientID, redirectURI, state string) string {
	params := map[string]string{
		"client_id":     clientID,
		"response_type": "code",
		"platform_id":   "mp",
		"state":         state,
		"redirect_uri":  redirectURI,
	}

	parsedURL, err := url.Parse(urlAuth)
	if err != nil {
		return ""
	}

	queryParams := url.Values{}

	for k, v := range params {
		queryParams.Add(k, v)
	}

	parsedURL.RawQuery = queryParams.Encode()

	return parsedURL.String()
}

func (c *client) Refresh(ctx context.Context, refreshToken string) (*Response, error) {
	request := &RefreshTokenRequest{
		ClientSecret: c.cfg.AccessToken,
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}

	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlBase,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
