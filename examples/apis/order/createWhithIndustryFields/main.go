package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "APP_USR-874202490252970-100714-e890db6519b0dceb4ef24ef41ed816e4-2021490138"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	request := order.Request{
		Type:              "online",
		TotalAmount:       "1000.00",
		ProcessingMode:    "automatic",
		Marketplace:       "NONE",
		ExternalReference: "ext_ref_1234",
		CaptureMode:       "automatic_async",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1000.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "master",
						Token:        "6176fb9fe529922e788c3db01d9a8738",
						Type:         "credit_card",
						Installments: 1,
					},
				},
			},
		},

		Payer: &order.PayerRequest{
			Email:     "anderson@testuser.com",
			FirstName: "Anderson",
			LastName:  "TestUser",
			Identification: &order.IdentificationRequest{
				Type:   "CPF",
				Number: "000000000",
			},
			Phone: &order.PhoneRequest{
				AreaCode: "55",
				Number:   "9881107889",
			},
			Address: &order.PayerAddress{
				StreetName:   "TestUser",
				StreetNumber: "125",
				ZipCode:      "123",
				Neighborhood: "TestUser",
				City:         "TestUser",
				State:        "TestUser",
				Complement:   "TestUser",
			},
		},

		Items: []order.ItemsRequest{
			{
				ExternalCode: "ext_ref_1234",
				Title:        "Test",
				Description:  "Test",
				PictureURL:   "http://picture.testuser.com",
				CategoryID:   "test",
				Quantity:     1,
				Type:         "travel",
				UnitPrice:    "1000.00",
				Warranty:     "true",
				Event_date:   "2020-01-01",
			},
		},
	}
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resource)
}
