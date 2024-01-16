package api

import (
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/internal"
)

// Send applies request options before call api.
func Send(cdt credential.Credential, config httpclient.Options, req *http.Request, opts ...RequestOption) ([]byte, error) {
	options := requestOptions{}
	for _, opt := range opts {
		opt.applyRequestOption(&options)
	}

	if options.httpClient != nil {
		config.HTTPClient = options.httpClient
	}
	if options.customHeaders != nil {
		for k, v := range options.customHeaders {
			canonicalKey := http.CanonicalHeaderKey(k)
			req.Header[canonicalKey] = v
		}
	}

	return internal.Send(cdt, req, config)
}
