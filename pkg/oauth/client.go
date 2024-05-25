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

// Client contains the method to interact with the Oauth API.
type Client interface {

	// Create oauth credentials to operate on behalf of a seller
	// It is a post request to the endpoint: "https://api.mercadopago.com/oauth/token"
	// Reference: https://www.mercadopago.com/developers/en/reference/oauth/_oauth_token/post
	Create(ctx context.Context, clientID, authorizationCode, redirectURI string) (*Response, error)

	// GetAuthorizationURL gets url for oauth authorization.
	GetAuthorizationURL(clientID, redirectURI, state string) string

	// Refresh token received when you create credentials.
	// It is a post request to the endpoint: "https://api.mercadopago.com/oauth/token"
	// Reference: https://www.mercadopago.com/developers/en/reference/oauth/_oauth_token/post
	Refresh(ctx context.Context, refreshToken string) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Oauth API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Create(ctx context.Context, clientID, authorizationCode, redirectURI string) (*Response, error) {
	request := &Request{
		ClientID:     clientID,
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
