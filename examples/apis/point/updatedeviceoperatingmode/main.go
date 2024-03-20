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

	req := point.OperatingModeRequest{
		OperatingMode: "PDV", // PDV or STANDALONE
	}

	client := point.NewClient(cfg)
	opMode, err := client.UpdateOperatingMode(context.Background(), "{{DEVICE_ID}}", req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(opMode)
}
