package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/user"
)

func TestUser(t *testing.T) {
	t.Run("should_get_user_information", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := user.NewClient(cfg)

		resource, err := client.Get(context.Background())

		if resource == nil {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
