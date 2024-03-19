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

	request := point.UpdateDeviceOperatingModeRequest{
		OperatingMode: "PDV", // PDV or STANDALONE
	}

	client := point.NewClient(cfg)
	opMode, err := client.UpdateDeviceOperatingMode(context.Background(), "{{DEVICE_ID}}", request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(opMode)
}
