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

	getResponseJSON, _ = os.Open("../../resources/mocks/point/get_response.json")
	getResponse, _     = io.ReadAll(getResponseJSON)

	cancelResponseJSON, _ = os.Open("../../resources/mocks/point/cancel_response.json")
	cancelResponse, _     = io.ReadAll(cancelResponseJSON)

	listDevicesResponseJSON, _ = os.Open("../../resources/mocks/point/list_devices_response.json")
	listDevicesResponse, _     = io.ReadAll(listDevicesResponseJSON)

	UpdateDeviceOperatingModeResponseJSON, _ = os.Open("../../resources/mocks/point/update_device_operating_mode_response.json")
	UpdateDeviceOperatingModeResponse, _     = io.ReadAll(UpdateDeviceOperatingModeResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx      context.Context
		deviceID string
		request  CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_create_point_transaction_intent",
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
				request:  CreateRequest{},
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_create_point_transaction_intent_success",
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
				request: CreateRequest{
					Amount:      1500,
					Description: "your payment intent description",
					AdditionalInfo: &AdditionalInfo{
						PrintOnTerminal:   false,
						ExternalReference: "4561ads-das4das4-das4754-das456",
					},
					Payment: &PaymentRequest{
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

func TestGet(t *testing.T) {
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
		want    *GetResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_get_point_transaction_intent",
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
			name: "should_get_point_transaction_intent_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(getResponse))
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
			want: &GetResponse{
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
			got, err := c.Get(tt.args.ctx, tt.args.paymentIntentID)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Get() = %v, want %v", got, tt.want)
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
			name: "should_return_error_when_cancel_point_transaction_intent",
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
			name: "should_cancel_point_transaction_intent_success",
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
				t.Errorf("client.Cancel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Cancel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListDevices(t *testing.T) {
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
							stringReader := strings.NewReader(string(listDevicesResponse))
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
			got, err := c.ListDevices(tt.args.ctx)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.ListDevices() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.ListDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateDeviceOperatingMode(t *testing.T) {
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
							stringReader := strings.NewReader(string(UpdateDeviceOperatingModeResponse))
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
			got, err := c.UpdateDeviceOperatingMode(tt.args.ctx, tt.args.deviceID, tt.args.request)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.UpdateDeviceOperatingMode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.UpdateDeviceOperatingMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
