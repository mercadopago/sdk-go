package httpclient

import (
	"errors"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/internal"
)

func Send(requester Requester, req *http.Request, opts ...RequestOption) ([]byte, error) {
	options := requestOptions{}
	for _, opt := range opts {
		opt.applyRequest(&options)
	}

	if options.callRequester != nil {
		requester = options.callRequester
	}
	if options.customHeaders != nil {
		for k, v := range options.customHeaders {
			canonicalKey := http.CanonicalHeaderKey(k)
			req.Header[canonicalKey] = v
		}
	}
	internal.SetDefaultHeaders(req)

	res, err := requester.Do(req)
	if err != nil {
		if res == nil {
			return nil, err
		}

		return nil, &ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error sending request: " + err.Error(),
		}
	}
	if res == nil {
		return nil, errors.New("error getting response")
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
