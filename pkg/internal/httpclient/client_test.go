package httpclient

import (
	"net/http"
	"testing"
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
