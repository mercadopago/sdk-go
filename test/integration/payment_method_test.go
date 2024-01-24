package integration

import (
	"context"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/credential"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func TestList(t *testing.T) {
	t.Run("should_list_brazil_payment_methods", func(t *testing.T) {
		cdt, err := credential.New("TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800")
		if err != nil {
			t.Fatal(err)
		}

		pmc := paymentmethod.NewClient(cdt)
		res, err := pmc.List(context.Background())

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
