package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/oauth"
)

func TestUser(t *testing.T) {

	t.Run("should_create_credentails", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := oauth.NewClient(cfg)
		authorizationCode := "TG-65cf441cb3cf3e000144dfbb-213761027"
		redirectURI := "https://httpdump.app/inspect/e340c21c-3afa-45a2-a6c7-a859cafcb7d2"

		credential, err := client.Create(context.Background(), authorizationCode, redirectURI)

		if credential == nil {
			t.Error("credential can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_refresh_token", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := oauth.NewClient(cfg)
		authorizationCode := "TG-65cf441cb3cf3e000144dfbb-213761027"
		redirectURI := "https://httpdump.app/inspect/e340c21c-3afa-45a2-a6c7-a859cafcb7d2"

		credential, err := client.Create(context.Background(), authorizationCode, redirectURI)

		if credential == nil {
			t.Error("res can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		refreshToken := credential.RefreshToken
		credential, err = client.Refresh(context.Background(), refreshToken)

		if credential == nil {
			t.Error("credential can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
