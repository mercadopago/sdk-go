package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/internal/requester"
)

func send[T any](requester requester.Requester, req *http.Request) (*T, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &ResponseError{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode > 399 {
		return nil, &ResponseError{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return makeResponse[T](response)
}

func makeResponse[T any](b []byte) (*T, error) {
	var response *T
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	return response, nil
}
