package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/mperror"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := payment.NewClient(cfg)

	invalidRequest := payment.Request{}

	resource, err := client.Create(context.Background(), invalidRequest)
	if err != nil {
		var mpErr *mperror.ResponseError
		if errors.As(err, &mpErr) {
			fmt.Printf("\nheaders: %s\nmessage: %s\nstatus code: %d", mpErr.Headers, mpErr.Message, mpErr.StatusCode)
			return
		}
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
