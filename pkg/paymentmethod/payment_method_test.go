package paymentmethod

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/option"
)

func TestClientList(t *testing.T) {
	type fields struct {
		config []option.HTTPOption
	}
	type args struct {
		ctx context.Context
		cdt credential.Credential
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Response
		wantErr error
	}{
		{
			name:   "should_return_error_creating_request",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				cdt: "TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800",
			},
			want:    nil,
			wantErr: fmt.Errorf("abc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.config...)
			got, err := c.List(tt.args.ctx, tt.args.cdt)
			if err != tt.wantErr {
				t.Errorf("client.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
