package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
)

func main() {
	accessToken := "APP_USR-4849723703374061-053108-80867bffcb4a85cda0cd797f6c40cf28-1340175910"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := merchantorder.NewClient(cfg)

	var merchantOrderID int64 = 8416510703

	merchant, err := client.Get(context.Background(), merchantOrderID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(merchant)
}
