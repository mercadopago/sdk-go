package httpclient

import (
	"net/http"
	"testing"
)

func TestMakeParams(t *testing.T) {
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
			name: "should_replace_one_path_param_and_add_query_param",
			args: args{
				url: "http://localhost/payments/:payment_id/search",
				params: map[string]string{
					"payment_id":        "1234567890",
					"external_resource": "abcd12345",
				},
			},
			wantURL: "http://localhost/payments/1234567890/search?external_resource=abcd12345",
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

//
//func TestMakeParams(t *testing.T) {
//	type args struct {
//		url    string
//		params map[string]string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantURL string
//		wantErr error
//	}{
//		{
//			name: "test_replace_one_path_param",
//			args: args{
//				url: "http://localhost/payments/:payment_id",
//				params: map[string]string{
//					"payment_id": "1234567890",
//				},
//			},
//			wantURL: "http://localhost/payments/1234567890",
//		},
//		{
//			name: "should_return_path_param_not_informed",
//			args: args{
//				url: "http://localhost/customers/:customer_id",
//				params: map[string]string{
//					"payment_id": "1234567890",
//				},
//			},
//			wantErr: errors.New("path parameters not informed: customer_id"),
//		},
//		{
//			name: "should_return_two_path_params_not_informed",
//			args: args{
//				url: "http://localhost/tests/:test_id/units/:unit_id",
//				params: map[string]string{
//					"integrate_id": "1234567890",
//				},
//			},
//			wantErr: errors.New("path parameters not informed: test_id,unit_id"),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			req, _ := http.NewRequest(http.MethodPost, tt.args.url, nil)
//			gotErr := makeParams(req, tt.args.params)
//
//			if tt.wantErr != nil && gotErr != tt.wantErr.Error() {
//				t.Errorf("makeParams() wantErr = %v, gotErr = %v", tt.wantErr, gotErr)
//			}
//
//			if tt.wantErr == nil && tt.wantURL != req.URL.String() {
//				t.Errorf("makeParams() wantURL = %v, gotURL %v", tt.wantURL, req.URL.String())
//			}
//		})
//	}
//}
