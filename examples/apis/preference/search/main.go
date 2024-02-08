package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	c, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}
	client := preference.NewClient(c)

	filter := make(map[string]string)
	filter["external_reference"] = "wee3rffee23"
	filters := preference.SearchRequest{
		Limit:   "10",
		Offset:  "10",
		Filters: filter,
	}

	res, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
