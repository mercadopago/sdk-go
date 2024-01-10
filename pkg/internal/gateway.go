package internal

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

func Send(requester httpclient.Requester, req *http.Request) ([]byte, error) {
	setDefaultHeaders(req)
	return send(requester, req)
}

func send(requester httpclient.Requester, req *http.Request) ([]byte, error) {
	res, err := requester.Do(req)
	if err != nil {
		return nil, fmt.Errorf("transport level error: %s", err.Error())
	}

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
