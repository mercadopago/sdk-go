package integration

import (
	"context"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/invoice"
	"os"
	"testing"
)

func TestInvoice(t *testing.T) {
	t.Run("should_get_invoice", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := invoice.NewClient(cfg)

		inv, err := client.Get(context.Background(), "id")
		if inv == nil {
			t.Error("invoice can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_invoice", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		filters := invoice.SearchRequest{
			Limit:  "10",
			Offset: "10",
			Filters: map[string]string{
				"preapproval_id": "id",
			},
		}

		client := invoice.NewClient(cfg)
		inv, err := client.Search(context.Background(), filters)

		if inv == nil {
			t.Error("invoice can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
