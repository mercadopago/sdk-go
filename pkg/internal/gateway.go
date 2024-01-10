package internal

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

func Send(requester httpclient.Requester, req *http.Request) ([]byte, error) {
	setDefaultHeaders(req)

	res, err := do(requester, req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return mountResponse(res)
}

func do(requester httpclient.Requester, req *http.Request) (*http.Response, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %s", err.Error())
	}

	if res == nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error getting response",
		}
	}

	return res, nil
}

func mountResponse(res *http.Response) ([]byte, error) {
	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &httpclient.ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    "error reading response body: " + err.Error(),
			Headers:    res.Header,
		}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, &httpclient.ErrorResponse{
			StatusCode: res.StatusCode,
			Message:    string(response),
			Headers:    res.Header,
		}
	}

	return response, nil
}
