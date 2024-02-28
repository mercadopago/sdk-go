package search

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/invoice"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := invoice.NewClient(cfg)

	filter := make(map[string]string)
	filter["preapproval_id"] = "preapproval_id"
	filters := invoice.SearchRequest{
		Limit:   "10",
		Offset:  "10",
		Filters: filter,
	}

	search, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, inv := range search.Results {
		fmt.Println(inv)
	}
}
