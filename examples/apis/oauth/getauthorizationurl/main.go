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

	clientID := "175307012991241"
	redirectURI := "https://httpdump.app/inspect/e340c21c-3afa-45a2-a6c7-a859cafcb7d2"

	url := client.GetAuthorizationURL(context.Background(), clientID, redirectURI)

	fmt.Println(url)
}
