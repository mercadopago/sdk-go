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
		c, err := config.New(os.Getenv("at"))
		if err != nil {
			t.Fatal(err)
		}

		pmc := user.NewClient(c)
		res, err := pmc.Get(context.Background())

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
