package advancedpayment

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	getResponseJSON, _    = os.Open("../../resources/mocks/advancedpayment/get_response.json")
	getResponse, _        = io.ReadAll(getResponseJSON)
	searchResponseJSON, _ = os.Open("../../resources/mocks/advancedpayment/search_response.json")
	searchResponse, _     = io.ReadAll(searchResponseJSON)
)

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr string
	}{
		{
			name: "should_return_error_when_send_request",
			cfg: &config.Config{
				Requester: &httpclient.Mock{
					DoMock: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("some error")
					},
				},
			},
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_response",
			cfg: &config.Config{
				Requester: &httpclient.Mock{
					DoMock: func(req *http.Request) (*http.Response, error) {
						resp := &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(getResponse)),
						}
						return resp, nil
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.cfg)
			result, err := c.Get(context.Background(), 20458724)
			if tt.wantErr != "" {
				if err == nil || err.Error() != tt.wantErr {
					t.Errorf("expected error %q, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result == nil {
				t.Fatal("expected result, got nil")
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr string
	}{
		{
			name: "should_return_error_when_send_request",
			cfg: &config.Config{
				Requester: &httpclient.Mock{
					DoMock: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("some error")
					},
				},
			},
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_response",
			cfg: &config.Config{
				Requester: &httpclient.Mock{
					DoMock: func(req *http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(searchResponse)),
						}, nil
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.cfg)
			result, err := c.Search(context.Background(), SearchRequest{})
			if tt.wantErr != "" {
				if err == nil || err.Error() != tt.wantErr {
					t.Errorf("expected error %q, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result == nil {
				t.Fatal("expected result, got nil")
			}
		})
	}
}
