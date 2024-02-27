package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/identificationtype"
)

func TestIdentificationTypes(t *testing.T) {
	t.Run("should_list_identification_types", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := identificationtype.NewClient(cfg)
		result, err := client.List(context.Background())

		if result == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
