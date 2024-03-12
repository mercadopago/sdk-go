package cardtoken

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

var (
	cardTokenResponseJSON, _ = os.Open("../../resources/mocks/cardtoken/response.json")
	cardTokenResponse, _     = io.ReadAll(cardTokenResponseJSON)
)

func TestCreate(t *testing.T) {
	type fields struct {
		cfg *config.Config
	}
	type args struct {
		ctx     context.Context
		request Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr string
	}{
		{
			name: "should_fail_create_card_token",
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
			name: "should_create_card_token",
			fields: fields{
				cfg: &config.Config{
					Requester: &httpclient.Mock{
						DoMock: func(req *http.Request) (*http.Response, error) {
							stringReader := strings.NewReader(string(cardTokenResponse))
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
			want:    mockCardToken(),
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				cfg: tt.fields.cfg,
			}
			got, err := c.Create(tt.args.ctx, tt.args.request)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("card token client.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("card token client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockCardToken() *Response {
	return &Response{
		ID:              "3d40b34eb41a6d0923e5bc545927c2e9",
		FirstSixDigits:  "503143",
		ExpirationMonth: 11,
		ExpirationYear:  2025,
		LastFourDigits:  "6351",
		Cardholder: CardholderResponse{
			Identification: IdentificationResponse{
				Number: "70383868084",
				Type:   "CPF",
			},
			Name: "MASTER TEST",
		},
		Status:             "active",
		DateCreated:        parseDate("2024-02-08T09:05:42.725-04:00"),
		DateLastUpdated:    parseDate("2024-02-08T09:05:42.725-04:00"),
		DateDue:            parseDate("2024-02-16T09:05:42.725-04:00"),
		LuhnValidation:     true,
		LiveMode:           false,
		CardNumberLength:   16,
		SecurityCodeLength: 3,
	}
}

func parseDate(s string) time.Time {
	d, _ := time.Parse(time.RFC3339, s)
	return d
}
