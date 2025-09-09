package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	request := order.Request{
		Type:              "online",
		ProcessingMode:    "automatic",
		TotalAmount:       "1000.00",
		ExternalReference: "ext_ref_1234",
		Currency:          "BRL",
		Payer: &order.PayerRequest{
			Email:     "test_123@testuser.com",
			FirstName: "John",
			LastName:  "Doe",
			Identification: &order.IdentificationRequest{
				Type:   "CPF",
				Number: "15635614680",
			},
			Phone: &order.PhoneRequest{
				AreaCode: "55",
				Number:   "987654321",
			},
			Address: &order.PayerAddressRequest{
				StreetName:   "R. Ângelo Piva",
				StreetNumber: "144",
				ZipCode:      "06210110",
				Neighborhood: "Presidente Altino",
				City:         "Osasco",
				State:        "SP",
				Complement:   "303",
			},
		},
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1000.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "elo",
						Type:         "credit_card",
						Token:        "{{CARD_TOKEN}}",
						Installments: 1,
					},
				},
			},
		},
		Items: []order.ItemsRequest{
			{
				ExternalCode: "1",
				Title:        "Passagem Para SP",
				Description:  "Passagem Para SP",
				PictureURL:   "https://example_url.com/",
				CategoryID:   "travel",
				Quantity:     1,
				Type:         "travel",
				UnitPrice:    "10.00",
				Warranty:     true,
				EventDate:    "2014-06-28T16:53:03.176-04:00",
			},
		},
		AdditionalInfo: &order.AdditionalInfoRequest{
			PayerAuthenticationType:         "MOBILE",
			PayerRegistrationDate:           "2024-01-01T00:00:00",
			PayerIsPrimeUser:                true,
			PayerIsFirstPurchaseOnLine:      true,
			PayerLastPurchase:               "2024-01-01T00:00:00",
			ShipmentExpress:                 true,
			ShipmentLocalPickup:             true,
			PlatFormShipmentDeliveryPromise: "2024-12-31T23:59:59Z",
			PlatformSellerID:                "123456",
			PlatformSellerName:              "Example Seller",
			PlatformSellerEmail:             "seller@example.com",
			PlatformAuthentication:          "string",
			TravelPassengers: &[]order.TravelPassengerRequest{
				{
					FirstName: "jose da silva",
					LastName:  "ferreira",
					Identification: &order.IdentificationRequest{
						Type:   "CPF",
						Number: "11111111111",
					},
				},
			},
			TravelRoutes: &[]order.TravelRouterRequest{
				{
					Departure:         "GRU",
					Destination:       "CWB",
					DepartureDateTime: "2020-01-01T00:00:00.000-03:00",
					ArrivalDateTime:   "2020-01-01T00:00:00.000-03:00",
					Company:           "gol",
				},
				{
					Departure:         "GRU",
					Destination:       "CWB",
					DepartureDateTime: "2020-01-01T00:00:00.000-03:00",
					ArrivalDateTime:   "2020-01-01T00:00:00.000-03:00",
					Company:           "azul",
				},
			},
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order created successfully:")
	fmt.Printf("ID: %s\n", resource.ID)
	fmt.Printf("Status: %s\n", resource.Status)
	fmt.Printf("Total Amount: %s %s\n", resource.TotalAmount, request.Currency)
}
