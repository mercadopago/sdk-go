package disbursementrefund

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
	getResponseJSON, _ = os.Open("../../resources/mocks/disbursementrefund/get_response.json")
	getResponse, _     = io.ReadAll(getResponseJSON)
)

func TestCreate(t *testing.T) {
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
							StatusCode: http.StatusCreated,
							Body:       io.NopCloser(bytes.NewReader(getResponse)),
						}, nil
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.cfg)
			result, err := c.Create(context.Background(), 20458724, 123456, Request{Amount: 50.0})
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

func TestCreateAll(t *testing.T) {
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
							StatusCode: http.StatusCreated,
							Body:       io.NopCloser(bytes.NewReader(getResponse)),
						}, nil
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.cfg)
			result, err := c.CreateAll(context.Background(), 20458724, Request{Amount: 50.0})
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

func TestListAll(t *testing.T) {
	listResponse := []byte(`[{"id":78901234,"advanced_payment_id":20458724,"disbursement_id":123456,"amount":50.0,"status":"approved"}]`)
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
							Body:       io.NopCloser(bytes.NewReader(listResponse)),
						}, nil
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.cfg)
			result, err := c.ListAll(context.Background(), 20458724)
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
