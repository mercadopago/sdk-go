package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/user"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := user.NewClient(cfg)

	resource, err := client.Get(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
