package main

import (
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

	clientID := "{{CLIENT_ID}}"
	redirectURI := "{{REDIRECT_URI}}"
	state := "state"

	url := client.GetAuthorizationURL(clientID, redirectURI, state)

	fmt.Println(url)
}
