package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	at := "{{ACCESS_TOKEN}}"
	c, err := config.New(at)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := preference.NewClient(c)

	filters := preference.SearchRequest{
		Limit: 10,
		Offset: 10,
	}

	res, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}
	resJSON, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resJSON))
}
