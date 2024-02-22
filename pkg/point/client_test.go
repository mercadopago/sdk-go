package point

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	createResponseJSON, _ = os.Open("../../resources/mocks/point/create_response.json")
	createResponse, _     = io.ReadAll(createResponseJSON)

	searchResponseJSON, _ = os.Open("../../resources/mocks/point/search_response.json")
	searchResponse, _     = io.ReadAll(searchResponseJSON)

	cancelResponseJSON, _ = os.Open("../../resources/mocks/point/cancel_response.json")
	cancelResponse, _     = io.ReadAll(cancelResponseJSON)

	getDevicesResponseJSON, _ = os.Open("../../resources/mocks/point/get_devices_response.json")
	getDevicesResponse, _     = io.ReadAll(getDevicesResponseJSON)

	updateDeviceOperationModeResponseJSON, _ = os.Open("../../resources/mocks/point/update_device_operation_mode_response.json")
	updateDeviceOperationModeResponse, _     = io.ReadAll(updateDeviceOperationModeResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx      context.Context
		deviceID string
		request  Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_create",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				deviceID: "any",
				request:  Request{},
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_create_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(createResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				deviceID: "any",
				request: Request{
					Amount:      1500,
					Description: "your payment intent description",
					AdditionalInfo: AdditionalInfo{
						PrintOnTerminal:   false,
						ExternalReference: "4561ads-das4das4-das4754-das456",
					},
					Payment: Payment{
						Installments:     1,
						Type:             "credit_card",
						InstallmentsCost: "seller",
					},
				},
			},
			want: &CreateResponse{
				ID:          "7f25f9aa-eea6-4f9c-bf16-a341f71ba2f1",
				DeviceID:    "PAX_A910__SMARTPOS1234345545",
				Amount:      1500,
				Description: "your payment intent description",
				Payment: CreatePaymentResponse{
					Installments:     1,
					Type:             "credit_card",
					InstallmentsCost: "seller",
				},
				AdditionalInfo: AdditionalInfo{
					ExternalReference: "someone-reference-from-your-application",
					PrintOnTerminal:   true,
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Create(tt.args.ctx, tt.args.deviceID, tt.args.request)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx             context.Context
		paymentIntentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SearchResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_search",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx:             context.Background(),
				paymentIntentID: "any",
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_search_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(searchResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:             context.Background(),
				paymentIntentID: "any",
			},
			want: &SearchResponse{
				ID:       "7f25f9aa-eea6-4f9c-bf16-a341f71ba2f1",
				State:    "FINISHED",
				Amount:   1500,
				DeviceID: "PAX_A910__SMARTPOS1234345545",
				Payment: PaymentResponse{
					ID: 16499678033,
				},
				AdditionalInfo: AdditionalInfo{
					ExternalReference: "some-reference-from-your-application",
					PrintOnTerminal:   true,
					TicketNumber:      "S0392JED",
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Search(tt.args.ctx, tt.args.paymentIntentID)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancel(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx             context.Context
		deviceID        string
		paymentIntentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CancelResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_cancel",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx:             context.Background(),
				deviceID:        "any",
				paymentIntentID: "any",
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_cancel_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(cancelResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:             context.Background(),
				deviceID:        "any",
				paymentIntentID: "any",
			},
			want: &CancelResponse{
				ID: "d71e88d6-6281-416b-b8ed-592c27352c99",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Cancel(tt.args.ctx, tt.args.deviceID, tt.args.paymentIntentID)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDevices(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *DevicesResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_get_devices",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_get_devices_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(getDevicesResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: &DevicesResponse{
				Devices: []Device{
					{
						ID:            "PAX_A910__SMARTPOS1234345545",
						PosID:         47792476,
						StoreID:       47792478,
						ExternalPosID: "SUC0101POS",
						OperatingMode: "PDV",
					},
				},
				Paging: Paging{
					Total:  1,
					Offset: 0,
					Limit:  50,
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.GetDevices(tt.args.ctx)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateDeviceOperationMode(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx      context.Context
		deviceID string
		request  UpdateDeviceOperatingModeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OperationModeResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_update_device_operation_mode",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							return nil, fmt.Errorf("some error")
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				deviceID: "any",
				request:  UpdateDeviceOperatingModeRequest{},
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_update_device_operation_mode_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(updateDeviceOperationModeResponse))
							stringReadCloser := io.NopCloser(stringReader)
							return &http.Response{
								Body: stringReadCloser,
							}, nil
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				deviceID: "any",
			},
			want: &OperationModeResponse{
				OperatingMode: "PDV",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.UpdateDeviceOperationMode(tt.args.ctx, tt.args.deviceID, tt.args.request)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
