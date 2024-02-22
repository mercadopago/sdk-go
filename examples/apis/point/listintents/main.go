package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/point"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := point.NewClient(cfg)
	intents, err := client.ListIntents(context.Background(), "2024-01-01", "2024-01-02")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range intents.Intent {
		fmt.Println(i)
	}
}
