package internal

import (
	"errors"
	"io"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

func Send(requester httpclient.Requester, req *http.Request) ([]byte, error) {
	setDefaultHeaders(req)

	res, err := requester.Do(req)
	if err != nil {
		if res == nil {
			return nil, err
		}

		return nil, &httpclient.ErrorResponse{
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
