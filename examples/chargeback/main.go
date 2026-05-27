package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mercadopago/sdk-go/pkg/chargeback"
	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	cfg, err := config.New("<YOUR_ACCESS_TOKEN>")
	if err != nil {
		log.Fatal(err)
	}

	client := chargeback.NewClient(cfg)

	// Get a chargeback by ID
	cb, err := client.Get(context.Background(), "CB-001-2022")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Chargeback ID: %s, Status: %s, Amount: %.2f\n", cb.ID, cb.Status, cb.Amount)

	// Search chargebacks by payment ID
	results, err := client.Search(context.Background(), chargeback.SearchRequest{
		Filters: map[string]string{"payment_id": "19951521071"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total chargebacks: %d\n", results.Paging.Total)
}
