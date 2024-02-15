package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/identificationtype"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := identificationtype.NewClient(cfg)
	identificationTypes, err := client.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range identificationTypes {
		fmt.Println(v)
	}
}
