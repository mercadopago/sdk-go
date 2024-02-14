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

	userClient := user.NewClient(cfg)
	user, err := userClient.Get(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}
