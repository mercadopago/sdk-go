package httpclient

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/internal/requester"
)

func Send(requester requester.Requester, req *http.Request) ([]byte, error) {
	result, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	defer func() { _ = result.Body.Close() }()

	response, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, &ResponseError{
			StatusCode: result.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    result.Header,
		}
	}

	if result.StatusCode > 399 {
		return nil, &ResponseError{
			StatusCode: result.StatusCode,
			Message:    string(response),
			Headers:    result.Header,
		}
	}

	return response, nil
}
