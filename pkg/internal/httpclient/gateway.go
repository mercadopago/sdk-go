package httpclient

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/mperror"
	"github.com/mercadopago/sdk-go/pkg/requester"
)

// Send executes the provided HTTP request using the given [requester.Requester]
// and returns the raw response body bytes. If the response status code is 400
// or above, it returns a [*mperror.ResponseError] containing the status code,
// headers, and body so the caller can inspect the API failure in a structured way.
func Send(requester requester.Requester, req *http.Request) ([]byte, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &mperror.ResponseError{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode > 399 {
		return nil, &mperror.ResponseError{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}
