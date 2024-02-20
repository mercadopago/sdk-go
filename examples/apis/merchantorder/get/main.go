package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := merchantorder.NewClient(cfg)

	var merchantOrderID int64 = 1234566788

	order, err := client.Get(context.Background(), merchantOrderID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(order)
}
