package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/oauth"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := oauth.NewClient(cfg)

	authorizationCode := "{{AUTHORIZATION_CODE}}"
	redirectURI := "{{REDIRECT_URI}}"

	resource, err := client.Create(context.Background(), authorizationCode, redirectURI)
	if err != nil {
		fmt.Println(err)
		return
	}

	refreshToken := resource.RefreshToken
	resource, err = client.Refresh(context.Background(), refreshToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
