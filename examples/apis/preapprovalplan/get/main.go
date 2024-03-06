package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"

	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preapprovalplan.NewClient(cfg)

	preApprovalPlanID := "123"

	result, err := client.Get(context.Background(), preApprovalPlanID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
