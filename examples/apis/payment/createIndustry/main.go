package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{Access_Token}}"
	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println("Config error:", err)
		return
	}

	client := payment.NewClient(cfg)

	request := payment.Request{
		TransactionAmount:   190.00,
		Description:         "Teste dados da industria",
		PaymentMethodID:     "master",
		Token:               "{{card_token}}",
		Installments:        1,
		BinaryMode:          false,
		Capture:             true,
		ExternalReference:   "Pedido01",
		NotificationURL:     "{{notification_url}}",
		StatementDescriptor: "LOJA 123",
		Payer: &payment.PayerRequest{
			FirstName: "Name",
			LastName:  "LasName",
			Email:     "teste_dados_industria@testuser.com",
			Identification: &payment.IdentificationRequest{
				Type:   "CPF",
				Number: "{{Number_CPF}}",
			},
			Phone: &payment.PhoneRequest{
				AreaCode: "{{DDD}}",
				Number:   "{{phone_number}}",
			},
			Address: &payment.AddressRequest{
				Neighborhood: "Teste Neighborhood",
				City:         "Teste City",
				FederalUnit:  "SP",
				ZipCode:      "{{CEP}}",
				StreetName:   "Teste StreetName",
				StreetNumber: "123",
			},
		},
		Metadata: map[string]any{
			"order_number": "order_01",
		},
		ForwardData: &payment.ForwardDataRequest{
			SubMerchant: &payment.SubMerchantRequest{
				SubMerchantID:     "1234183712",
				MCC:               "5462",
				Country:           "BRA",
				AddressDoorNumber: 123,
				ZIP:               "222222",
				DocumentNumber:    "222222222222222",
				City:              "SÃO PAULO",
				LegalName:         "LOJINHA DO ZÉ",
				AddressStreet:     "RUA A",
				RegionCodeISO:     "BR-MG",
				RegionCode:        "BR",
				DocumentType:      "CNPJ",
				Phone:             "123123123",
				URL:               "{{URL}}",
			},
		},
		AdditionalInfo: &payment.AdditionalInfoRequest{
			Items: []payment.ItemRequest{
				{
					ID:          "1941",
					Title:       "25/08/2022 | Pista Inteira5 lote - GREEN VALLEY GRAMADO 2022",
					Description: "25/08/2022 | Pista Inteira5 lote - GREEN VALLEY GRAMADO 2022",
					CategoryID:  "Tickets",
					Quantity:    1,
					UnitPrice:   1000.00,
					EventDate:   "2019-10-25T19:30:00.000-03:00",
					PictureURL:  "{{url_image}}",
					Type:        "my_items_type",
					Warranty:    true,
					CategoryDescriptor: payment.CategoryDescriptorRequest{
						Passenger: &payment.PassengerRequest{
							FirstName: "Guest Nome",
							LastName:  "Guest Sobrenome",
							Identification: &payment.IdentificationRequest{
								Type:   "DNI",
								Number: "012345678",
							},
						},
						Route: &payment.RouteRequest{
							ArrivalDateTime:   "2019-12-25T19:30:00.000-03:00",
							Company:           "Companhia",
							Departure:         "Osasco",
							DepartureDateTime: "2022-03-12T12:58:41.425-04:00",
							Destination:       "Sao Paulo",
						},
					},
				},
			},
			Payer: &payment.AdditionalInfoPayerRequest{
				FirstName:             "Nome",
				LastName:              "Sobrenome",
				IsPrimeUser:           true,
				IsFirstPurchaseOnline: true,
				LastPurchase:          "2019-10-25T19:30:00.000-03:00",
				Phone: &payment.AdditionalInfoPayerPhoneRequest{
					AreaCode: "{{DDD}}",
					Number:   "{{Phone_number}}",
				},
				Address: &payment.AdditionalInfoPayerAddressRequest{
					StreetName:   "Av. das Nações Unidas",
					StreetNumber: "3003",
					ZipCode:      "206233-2002",
				},
				RegistrationDate:   "2020-08-06T09:25:04.000-03:00",
				AuthenticationType: "Gmail",
			},
			Shipments: &payment.ShipmentsRequest{
				LocalPickup:     true,
				ExpressShipment: true,
				ReceiverAddress: &payment.ReceiverAddressRequest{
					ZipCode:      "306233-2003",
					StreetName:   "Av. das Nações Unidas",
					StreetNumber: "3003",
					Floor:        "5",
					Apartment:    "502",
					StateName:    "DF",
					CityName:     "Bogota",
				},
			},
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println("Payment error:", err)
		return
	}

	fmt.Println("Payment ID:", resource.ID)
}
