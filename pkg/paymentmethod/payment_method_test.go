package paymentmethod

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/option"
)

var (
	cdt, _ = credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")

	listResponseJSON, _ = os.Open("../../resources/mocks/payment_method_list.json")
	listResponse, _     = io.ReadAll(listResponseJSON)
)

func TestList(t *testing.T) {
	type fields struct {
		credential *credential.Credential
		config     *option.ClientOptions
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Response
		wantErr string
	}{
		{
			name: "should_return_error_when_creating_request",
			fields: fields{
				credential: cdt,
				config:     option.ApplyClientOptions(),
			},
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: "error creating request: net/http: nil Context",
		},
		{
			name: "should_return_error_when_send_request",
			fields: fields{
				credential: cdt,
				config: option.ApplyClientOptions(
					option.WithCustomClient(
						&httpclient.Mock{
							DoMock: func(req *http.Request) (*http.Response, error) {
								return nil, fmt.Errorf("some error")
							},
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "transport level error: some error",
		},
		{
			name: "should_return_error_unmarshal_response",
			fields: fields{
				credential: cdt,
				config: option.ApplyClientOptions(
					option.WithCustomClient(
						&httpclient.Mock{
							DoMock: func(req *http.Request) (*http.Response, error) {
								stringReader := strings.NewReader("invalid json")
								stringReadCloser := io.NopCloser(stringReader)
								return &http.Response{
									Body: stringReadCloser,
								}, nil
							},
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: "invalid character 'i' looking for beginning of value",
		},
		{
			name: "should_return_formatted_response",
			fields: fields{
				credential: cdt,
				config: option.ApplyClientOptions(
					option.WithCustomClient(
						&httpclient.Mock{
							DoMock: func(req *http.Request) (*http.Response, error) {
								stringReader := strings.NewReader(string(listResponse))
								stringReadCloser := io.NopCloser(stringReader)
								return &http.Response{
									Body: stringReadCloser,
								}, nil
							},
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want: []Response{
				{
					ID:              "debmaster",
					Name:            "Mastercard Débito",
					PaymentTypeID:   "debit_card",
					Status:          "testing",
					SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif",
					Thumbnail:       "https://www.mercadopago.com/org-img/MP3/API/logos/debmaster.gif",
					DeferredCapture: "unsupported",
					Settings: []SettingsResponse{
						{
							CardNumber: &SettingsCardNumberResponse{
								Validation: "standard",
								Length:     16,
							},
							Bin: &SettingsBinResponse{
								Pattern:             "^(502121|536106)",
								InstallmentsPattern: "",
								ExclusionPattern:    "",
							},
							SecurityCode: &SettingsSecurityCodeResponse{
								Length:       3,
								CardLocation: "back",
								Mode:         "mandatory",
							},
						},
					},
					AdditionalInfoNeeded: []string{
						"cardholder_name",
						"cardholder_identification_type",
						"cardholder_identification_number",
					},
					MinAllowedAmount:      0.5,
					MaxAllowedAmount:      60000,
					AccreditationTime:     1440,
					FinancialInstitutions: []FinancialInstitutionResponse{},
					ProcessingModes: []string{
						"aggregator",
					},
				},
				{
					ID:              "cabal",
					Name:            "Cabal",
					PaymentTypeID:   "credit_card",
					Status:          "testing",
					SecureThumbnail: "https://www.mercadopago.com/org-img/MP3/API/logos/cabal.gif",
					Thumbnail:       "https://www.mercadopago.com/org-img/MP3/API/logos/cabal.gif",
					DeferredCapture: "supported",
					Settings:        []SettingsResponse{},
					AdditionalInfoNeeded: []string{
						"cardholder_name",
						"cardholder_identification_type",
						"cardholder_identification_number",
					},
					MinAllowedAmount:      0.5,
					MaxAllowedAmount:      60000,
					AccreditationTime:     2880,
					FinancialInstitutions: []FinancialInstitutionResponse{},
					ProcessingModes: []string{
						"aggregator",
					},
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				credential: tt.fields.credential,
				config:     tt.fields.config,
			}
			got, err := c.List(tt.args.ctx)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.List() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
