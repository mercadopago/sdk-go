package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/oauth"
)

func TestOauth(t *testing.T) {
	t.Run("should_create_credentials", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := oauth.NewClient(cfg)

		authorizationCode := "authorization_code"
		redirectURI := "redirect_uri"

		cred, err := client.Create(context.Background(), authorizationCode, redirectURI)

		if cred == nil {
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

		authorizationCode := "authorization_code"
		redirectURI := "redirect_uri"

		cred, err := client.Create(context.Background(), authorizationCode, redirectURI)

		if cred == nil {
			t.Error("credential can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		refreshToken := cred.RefreshToken
		cred, err = client.Refresh(context.Background(), refreshToken)

		if cred == nil {
			t.Error("credential can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
