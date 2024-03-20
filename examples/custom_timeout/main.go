package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := paymentmethod.NewClient(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resources, err := client.List(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resources {
		fmt.Println(v)
	}
}
