package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "<ACCESS_TOKEN>"
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
						Token:        "<CARD_TOKEN>",
						Type:         "credit_card",
						Installments: 1,
					},
				},
			},
		},

		Payer: &order.PayerRequest{
			Email:     "<EMAIL_ADDRESS>",
			FirstName: "<FIRST_NAME>",
			LastName:  "<LAST_NAME>",
			Identification: &order.IdentificationRequest{
				Type:   "<CPF>",
				Number: "<NUMBER>",
			},
			Phone: &order.PhoneRequest{
				AreaCode: "55",
				Number:   "9881107889",
			},
			Address: &order.PayerAddressRequest{
				StreetName:   "<STREET_NAME>",
				StreetNumber: "<STREET_NUMBER>",
				Neighborhood: "	BAIRRO",
				City:         "CIDADE",
				State:        "SP",
				Complement:   "303",
			},
		},
		Items: []order.ItemsRequest{
			{
				ExternalCode: "ext_ref_1234",
				Title:        "Passagem Para SP",
				Description:  "Test",
				PictureURL:   "http://picture.testuser.com",
				CategoryID:   "test",
				Quantity:     1,
				Type:         "travel",
				UnitPrice:    "1000.00",
				Warranty:     true,
				Event_date:   "2023-10-10T00:00:00Z",
			},
		},
		Shipment: &order.ShipmentRequest{
			Address: &order.AddressRequest{
				StreetName:   "<STREET_NAME>",
				StreetNumber: "129",
				ZipCode:      "06233903",
				Neighborhood: "BAIRRO",
				City:         "Osasco",
				State:        "SP",
				Complement:   "2",
			},
		},

		AdditionalInfo: &order.AdditionalInfoRequest{
			PayerAuthenticationType:            "MOBILE",
			PayerRegistrationDate:              "2020-08-06T09:25:04.000-03:00",
			PayerIsPrimeUser:                   true,
			PayerIsFirstPurchaseOnLine:         false,
			PayerLastPurchase:                  "2020-08-06T09:25:04.000-03:00",
			ShipmentExpress:                    true,
			ShipmentLocalPickup:                true,
			PlatFormShipmentDeliveryPromise:    "2024-12-31T23:59:59Z",
			PlatFormShipmentDropShipping:       "string",
			PlatformShipmentSafety:             "string",
			PlatformShipmentTrackingCode:       "1234",
			PlatformShipmentTrackingStatus:     "Em rota",
			PlatformShipmentWithdrawn:          true,
			PlatformSellerId:                   "123456",
			PlatformSellerName:                 "NAME",
			PlatformSellerEmail:                "<EMAIL>",
			PlatformSellerStatus:               "Active",
			PlatformSellerReferralUrl:          "https://www.testuser.com/seller/123456",
			PlatformSellerRegistrationDate:     "2020-01-01T00:00:00.000-03:00",
			PlatformSellerHiredPlan:            "Premium",
			PlatformSellerBusinessType:         "E-commerce",
			PlatformSellerAddressZipCode:       "123456",
			PlatformSellerAddressStreetName:    "NAME",
			PlatformSellerAddressStreetNumber:  "125",
			PlatformSellerAddressCity:          "São Paulo",
			PlatformSellerAddressState:         "SP",
			PlatformSellerAddressComplement:    "101",
			PlatformSellerAddressCountry:       "Brasil",
			PlatformSellerIdentificationType:   "CNPJ",
			PlatformSellerIdentificationNumber: "<NUMBER>",
			PlatformSellerPhoneNumber:          "NUMBER",
			PlatformSellerPhoneAreaCode:        "11",
			PlatformAuthentication:             "string",
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resource)
}
