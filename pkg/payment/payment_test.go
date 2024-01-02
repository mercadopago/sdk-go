package payment

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/mp/rest"
)

func TestClientCreate(t *testing.T) {
	type fields struct {
		rc rest.Client
	}
	type args struct {
		dto  Request
		opts []rest.Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr string
	}{
		{
			name:   "should_return_marshal_error",
			fields: fields{},
			args: args{
				dto: Request{
					TransactionAmount: math.Inf(1),
				},
			},
			want:    nil,
			wantErr: "error marshaling request body: json: unsupported value: +Inf",
		},
		{
			name: "should_return_send_error",
			fields: fields{
				rc: &rest.Mock{
					SendMock: func(req *http.Request, opts ...rest.Option) ([]byte, error) {
						return nil, fmt.Errorf("some error")
					},
				},
			},
			args:    args{},
			want:    nil,
			wantErr: "some error",
		},
		{
			name: "should_return_unmarshal_error",
			fields: fields{
				rc: &rest.Mock{
					SendMock: func(req *http.Request, opts ...rest.Option) ([]byte, error) {
						return []byte("malformed json"), nil
					},
				},
			},
			args:    args{},
			want:    nil,
			wantErr: "invalid character 'm' looking for beginning of value",
		},
		{
			name: "should_return_success",
			fields: fields{
				rc: &rest.Mock{
					SendMock: func(req *http.Request, opts ...rest.Option) ([]byte, error) {
						return []byte(`{"transaction_amount": 123.5}`), nil
					},
				},
			},
			args:    args{},
			want:    &Response{TransactionAmount: 123.5},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				rc: tt.fields.rc,
			}
			got, err := c.Create(tt.args.dto, tt.args.opts...)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("client.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
