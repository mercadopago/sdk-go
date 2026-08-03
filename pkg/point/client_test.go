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

	getPaymentIntentListResponseJSON, _ = os.Open("../../resources/mocks/point/get_payment_intent_list_response.json")
	getPaymentIntentListResponse, _     = io.ReadAll(getPaymentIntentListResponseJSON)

	getPaymentIntentStatusResponseJSON, _ = os.Open("../../resources/mocks/point/get_payment_intent_status_response.json")
	getPaymentIntentStatusResponse, _     = io.ReadAll(getPaymentIntentStatusResponseJSON)
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
		want    *Response
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
				request:  Request{},
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
				request: Request{
					Amount:      1500,
					Description: "your payment intent description",
					AdditionalInfo: &AdditionalInfoRequest{
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
			want: &Response{
				ID:          "7f25f9aa-eea6-4f9c-bf16-a341f71ba2f1",
				DeviceID:    "PAX_A910__SMARTPOS1234345545",
				Amount:      1500,
				Description: "your payment intent description",
				Payment: PaymentResponse{
					Installments:     1,
					Type:             "credit_card",
					InstallmentsCost: "seller",
				},
				AdditionalInfo: AdditionalInfoResponse{
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
		want    *Response
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
			want: &Response{
				ID:       "7f25f9aa-eea6-4f9c-bf16-a341f71ba2f1",
				State:    "FINISHED",
				Amount:   1500,
				DeviceID: "PAX_A910__SMARTPOS1234345545",
				Payment: PaymentResponse{
					ID: 16499678033,
				},
				AdditionalInfo: AdditionalInfoResponse{
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
				Devices: []DeviceResponse{
					{
						ID:            "PAX_A910__SMARTPOS1234345545",
						PosID:         47792476,
						StoreID:       "47792478",
						ExternalPosID: "SUC0101POS",
						OperatingMode: "PDV",
					},
				},
				Paging: PagingResponse{
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
		ctx           context.Context
		deviceID      string
		operatingMode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OperatingModeResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_update_device_operating_mode",
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
				ctx:           context.Background(),
				deviceID:      "any",
				operatingMode: "PDV",
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_update_device_operating_mode_success",
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
			want: &OperatingModeResponse{
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
			got, err := c.UpdateOperatingMode(tt.args.ctx, tt.args.deviceID, tt.args.operatingMode)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.UpdateOperatingMode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.UpdateOperatingMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPaymentIntentList(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		request PaymentIntentListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PaymentIntentListResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_get_payment_intent_list",
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
				ctx:     context.Background(),
				request: PaymentIntentListRequest{},
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_get_payment_intent_list_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(getPaymentIntentListResponse))
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
				request: PaymentIntentListRequest{
					StartDate: "2024-02-08T00:00:00.000-04:00",
					EndDate:   "2024-02-08T23:59:59.999-04:00",
				},
			},
			want: &PaymentIntentListResponse{
				Events: []PaymentIntentEventResponse{
					{
						PaymentIntentID: "7f25f9aa-eea6-4f9c-bf16-a341f71ba2f1",
						Status:          "FINISHED",
						CreatedOn:       "2024-02-08T09:05:42.725-04:00",
					},
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
			got, err := c.GetPaymentIntentList(tt.args.ctx, tt.args.request)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.GetPaymentIntentList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetPaymentIntentList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPaymentIntentStatus(t *testing.T) {
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
		want    *PaymentIntentStatusResponse
		wantErr string
	}{
		{
			name: "should_return_error_when_get_payment_intent_status",
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
			name: "should_get_payment_intent_status_success",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(getPaymentIntentStatusResponse))
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
				paymentIntentID: "7f25f9aa-eea6-4f9c-bf16-a341f71ba2f1",
			},
			want: &PaymentIntentStatusResponse{
				Status:    "FINISHED",
				CreatedOn: "2024-02-08T09:05:42.725-04:00",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.GetPaymentIntentStatus(tt.args.ctx, tt.args.paymentIntentID)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.GetPaymentIntentStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetPaymentIntentStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
