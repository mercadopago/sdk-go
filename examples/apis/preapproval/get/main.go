package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preapproval.NewClient(cfg)

	preApprovalID := "123"

	result, err := client.Get(context.Background(), preApprovalID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
