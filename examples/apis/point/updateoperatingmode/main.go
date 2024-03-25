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

	operatingMode := "PDV" // PDV or STANDALONE

	resource, err := client.UpdateOperatingMode(context.Background(), "{{DEVICE_ID}}", operatingMode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
