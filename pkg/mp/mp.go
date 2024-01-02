package mp

import (
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/mp/rest"
)

// NewRestClient returns a new rest client.
func NewRestClient(accessToken string) rest.Client {
	return rest.NewClient(accessToken)
}

// SetAccessToken sets the access token.
func SetAccessToken(at string) {
	rest.SetAT(at)
}

// SetCustomHTTPClient sets a custom http client.
func SetCustomHTTPClient(hc *http.Client) {
	rest.SetHC(hc)
}

// SetCustomRetryClient sets a custom retry client.
func SetCustomRetryClient(rc rest.RetryClient) {
	rest.SetRC(rc)
}
