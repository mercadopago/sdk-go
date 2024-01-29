package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/option"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(
		accessToken,
		option.WithCorporationID("1yuy811998tt11199"),
		option.WithIntegratorID("6888999999000001"),
		option.WithPlatformID("prd_02647ea11edb6888682a831752a"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := paymentmethod.NewClient(cfg)
	paymentMethods, err := client.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range paymentMethods {
		fmt.Println(v)
	}
}
