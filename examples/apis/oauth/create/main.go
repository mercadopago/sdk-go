package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/oauth"
)

func main() {
	accessToken := "***REMOVED***"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := oauth.NewClient(cfg)
	authorizationCode := "TG-65cf441cb3cf3e000144dfbb-213761027"
	redirectURI := "https://httpdump.app/inspect/e340c21c-3afa-45a2-a6c7-a859cafcb7d2"

	credential, err := client.Create(context.Background(), authorizationCode, redirectURI)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(credential)
}
