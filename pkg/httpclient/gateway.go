package httpclient

import (
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/internal"
)

func Send(requester Requester, req *http.Request, opts ...RequestOption) ([]byte, error) {
	options := requestOptions{}
	for _, opt := range opts {
		opt.applyRequest(&options)
	}

	if options.RequestRequester != nil {
		requester = options.RequestRequester
	}
	if options.CustomHeaders != nil {
		for k, v := range options.CustomHeaders {
			canonicalKey := http.CanonicalHeaderKey(k)
			req.Header[canonicalKey] = v
		}
	}
	internal.SetDefaultHeaders(req)

	res, err := requester.Do(req)
	if err != nil {
		status := 0
		if res != nil {
			status = res.StatusCode
		}

		return nil, &ErrorResponse{
			StatusCode: status,
			Message:    "error sending request: " + err.Error(),
		}
	}

	defer res.Body.Close()

	return mountResponse(res)
}

func mountResponse(res *http.Response) ([]byte, error) {
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}
