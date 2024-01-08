package api

import (
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
	"github.com/mercadopago/sdk-go/pkg/internal"
)

func Send(requester httpclient.Requester, req *http.Request, opts ...RequestOption) ([]byte, error) {
	options := requestOptions{}
	for _, opt := range opts {
		opt.applyRequestOption(&options)
	}

	if options.Requester != nil {
		requester = options.Requester
	}
	if options.customHeaders != nil {
		for k, v := range options.customHeaders {
			canonicalKey := http.CanonicalHeaderKey(k)
			req.Header[canonicalKey] = v
		}
	}

	return internal.Send(requester, req)
}
