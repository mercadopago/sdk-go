package baseclient

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/requester"
	"github.com/stretchr/testify/assert"
)

func TestMakePathParams(t *testing.T) {
	type args struct {
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantURL    string
		wantMsgErr string
	}{
		{
			name: "should_replace_one_path_param",
			args: args{
				url: "http://localhost/payments/:payment_id",
				params: map[string]string{
					"payment_id": "1234567890",
				},
			},
			wantURL: "http://localhost/payments/1234567890",
		},
		{
			name: "should_return_path_param_not_informed",
			args: args{
				url: "http://localhost/customers/:customer_id",
				params: map[string]string{
					"payment_id": "1234567890",
				},
			},
			wantMsgErr: "path parameters not informed: customer_id",
		},
		{
			name: "should_return_two_path_params_not_informed",
			args: args{
				url: "http://localhost/tests/:test_id/units/:unit_id",
				params: map[string]string{
					"integrate_id": "1234567890",
				},
			},
			wantMsgErr: "path parameters not informed: test_id,unit_id",
		},
		{
			name: "should_return_the_same_path_url",
			args: args{
				url: "http://localhost/tests/",
				params: map[string]string{
					"integrate_id": "1234567890",
				},
			},
			wantURL: "http://localhost/tests/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.args.url, nil)

			err := makePathParams(req, tt.args.params)
			if err != nil && err.Error() != tt.wantMsgErr {
				t.Errorf("makeParams() msgError = %v, wantMsgErr %v", err.Error(), tt.wantMsgErr)
				return
			}

			if err == nil && tt.wantURL != req.URL.String() {
				t.Errorf("makeParams() wantURL = %v, gotURL %v", tt.wantURL, req.URL.String())
			}
		})
	}
}

func TestMakeQueryParams(t *testing.T) {
	type args struct {
		url    string
		params map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantURL string
	}{
		{
			name: "should_add_one_query_param",
			args: args{
				url: "http://localhost/payments/1234567890/search",
				params: map[string]string{
					"external_resource": "as2f12345",
				},
			},
			wantURL: "http://localhost/payments/1234567890/search?external_resource=as2f12345",
		},
		{
			name: "should_add_two_query_params",
			args: args{
				url: "http://localhost/payments/1234567890/search",
				params: map[string]string{
					"external_resource": "as2f12345",
					"offset":            "2",
				},
			},
			wantURL: "http://localhost/payments/1234567890/search?external_resource=as2f12345&offset=2",
		},
		{
			name: "should_return_the_same_path_url",
			args: args{
				url:    "http://localhost/tests/",
				params: map[string]string{},
			},
			wantURL: "http://localhost/tests/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.args.url, nil)

			makeQueryParams(req, tt.args.params)
			if tt.wantURL != req.URL.String() {
				t.Errorf("makeQueryParams() wantURL = %v, gotURL %v", tt.wantURL, req.URL.String())
			}
		})
	}
}

func Test_makeRequest(t *testing.T) {
	type args struct {
		cfg    *config.Config
		method string
		url    string
		body   any
		opts   []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "should_create_http_request_success",
			args: args{
				cfg: &config.Config{
					AccessToken: "",
					Requester:   requester.Default(),
				},
				method: http.MethodGet,
				url:    "https://test.com/tests/:id",
				body:   nil,
				opts: []Option{
					WithPathParams(map[string]string{
						"id": "123",
					}),
				},
			},
			want: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Scheme: "https",
					Host:   "test.com",
					Path:   "/tests/123",
				},
				Host: "test.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := makeRequest(ctx, tt.args.cfg, tt.args.method, tt.args.url, tt.args.body, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.URL.String(), got.URL.String())
			assert.NotEmpty(t, got.Header.Get("X-Idempotency-Key"))
			assert.NotEmpty(t, got.Header.Get("X-Request-Id"))
		})
	}
}
