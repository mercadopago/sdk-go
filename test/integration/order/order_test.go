package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/order"
	"github.com/mercadopago/sdk-go/test"
)

var (
	cfg             = test.Config()
	orderClient     = order.NewClient(cfg)
	cardTokenClient = cardtoken.NewClient(cfg)
)

func TestOrder(t *testing.T) {
	t.Run("should_create_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_ref_1234",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func TestGetOrder(t *testing.T) {
	t.Run("should_create_and_get_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_ref_12345",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

		resource, err = orderClient.Get(ctx, resource.ID)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

	})
}

func TestProcessOrder(t *testing.T) {
	t.Run("should_create_and_get_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			ProcessingMode:    "manual",
			TotalAmount:       "100.00",
			ExternalReference: "ext_ref_12345",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "100.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

		resource, err = orderClient.Process(ctx, resource.ID)
		if resource == nil || resource.ID == "" {
			t.Errorf("error processing order: %v", err)
		}

		if resource.Status != "processed" {
			t.Errorf("expected order status to be 'processed', got %v", resource.Status)
		}

	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("should_create_transaction", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}
		request := order.Request{
			Type:              "online",
			ProcessingMode:    "manual",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_1234",
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

		requestTransaction := order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1000.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "master",
						Token:        token,
						Type:         "credit_card",
						Installments: 1,
					},
				},
			},
		}

		resp, err := orderClient.CreateTransaction(ctx, resource.ID, requestTransaction)
		if resp == nil || resp.Payments[0].ID == "" {
			t.Error("transaction can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("should_update_installments", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Fatalf("fail to generate card token: %v", err)
		}

		orderRequest := order.Request{
			Type:              "online",
			ProcessingMode:    "manual",
			TotalAmount:       "1000.00",
			ExternalReference: "ext_1234",
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 3,
						},
					},
				},
			},
		}

		resource, err := orderClient.Create(ctx, orderRequest)
		if err != nil {
			t.Fatalf("failed to create order: %v", err)
		}
		if resource == nil || resource.ID == "" {
			t.Fatalf("order can't be nil")
		}

		orderID := resource.ID
		fmt.Println("Order ID: ", orderID)
		transactionID := resource.Transactions.Payments[0].ID
		fmt.Println("transaction: ", transactionID)

		updateRequest := order.PaymentRequest{
			PaymentMethod: &order.PaymentMethodRequest{Installments: 12},
		}
		fmt.Printf("updateRequest: %+v\n", updateRequest)

		updateResp, err := orderClient.UpdateTransaction(ctx, orderID, transactionID, updateRequest)
		if err != nil {
			t.Fatalf("failed to update transaction: %v", err)
		}
		if updateResp == nil || updateResp.PaymentMethod.Installments != 12 {
			t.Fatalf("expected installments to be updated to 12, got %v", updateResp.PaymentMethod.Installments)
		}
	})
}

func TestCaptureOrder(t *testing.T) {
	t.Run("should_create_and_get_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		// Create order with capture_mode = manual
		request := order.Request{
			Type:              "online",
			ProcessingMode:    "automatic",
			CaptureMode:       "manual",
			TotalAmount:       "200.00",
			ExternalReference: "ext_ref_12345",
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "200.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
		}

		// Create an order
		resource, err := orderClient.Create(ctx, request)
		if err != nil {
			t.Fatal("failed to create order", err)
		}
		if resource == nil || resource.ID == "" {
			t.Fatal("order can't be nil")
		}

		time.Sleep(2 * time.Second)

		// Capture an order
		captureResp, err := orderClient.Capture(ctx, resource.ID)
		if err != nil {
			t.Fatalf("failed to capture order: %v", err)
		}

		if captureResp == nil || captureResp.Status != "processed" {
			t.Fatalf("expected order status to be 'processed', got %v", captureResp.Status)
		}
	})
}

func TestCancelOrder(t *testing.T) {
	t.Run("should_create_and_get_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		// Create order
		request := order.Request{
			Type:              "online",
			ProcessingMode:    "automatic",
			CaptureMode:       "manual",
			TotalAmount:       "200.00",
			ExternalReference: "ext_ref_12345",
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "200.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
		}

		// Create an order
		resource, err := orderClient.Create(ctx, request)
		if err != nil {
			t.Fatal("failed to create order", err)
		}
		if resource == nil || resource.ID == "" {
			t.Fatal("order can't be nil")
		}

		time.Sleep(2 * time.Second)

		// Capture an order
		captureResp, err := orderClient.Cancel(ctx, resource.ID)
		if err != nil {
			t.Fatalf("failed to capture order: %v", err)
		}

		if captureResp == nil || captureResp.Status != "canceled" {
			t.Fatalf("expected order status to be 'canceled', got %v", captureResp.Status)
		}
	})
}

/*func TestDeleteOrder(t *testing.T) {
	t.Run("should_create_and_get_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			ProcessingMode:    "manual",
			TotalAmount:       "100.00",
			ExternalReference: "ext_ref_12345",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "100.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if resource == nil || resource.ID == "" {
			t.Error("order can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}

		orderID := resource.ID
		transactionID := resource.Transactions.Payments[0].ID

		time.Sleep(3 * time.Second)

		err = orderClient.DeleteTransaction(ctx, orderID, transactionID)
		if err != nil {
			if err != nil {
				t.Fatalf("failed to delete transaction: %v", err)
			}
		}
	})
}*/

/*func TestRefundOrder(t *testing.T) {
	t.Run("should_create_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			TotalAmount:       "1000.00",
			ProcessingMode:    "automatic",
			ExternalReference: "ext_ref_1234",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if err != nil {
			t.Fatalf("failed to create order: %v", err)
		}
		if resource == nil || resource.ID == "" {
			t.Fatalf("order can't be nil")
		}

		time.Sleep(3 * time.Second)

		// Refund an order
		refundRequest := map[string]interface{}{
			"transactions": []map[string]interface{}{}, // gambiarra e ainda da errado
		}
		refundResp, err := orderClient.Refund(ctx, resource.ID, refundRequest)
		if err != nil {
			t.Fatalf("failed to refund order: %v", err)
		}

		if refundResp == nil || refundResp.Status != "refunded" {
			t.Fatalf("expected order status to be 'refunded', got %v", refundResp.Status)
		}
	})
}*/

/*func TestRefundPartialOrder(t *testing.T) {
	t.Run("should_create_order", func(t *testing.T) {
		ctx := context.Background()
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		request := order.Request{
			Type:              "online",
			TotalAmount:       "1000.00",
			ProcessingMode:    "automatic",
			ExternalReference: "ext_ref_1234",
			Transactions: &order.TransactionRequest{
				Payments: []order.PaymentRequest{
					{
						Amount: "1000.00",
						PaymentMethod: &order.PaymentMethodRequest{
							ID:           "master",
							Token:        token,
							Type:         "credit_card",
							Installments: 1,
						},
					},
				},
			},
			Payer: order.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New().String()[:7]),
			},
		}

		resource, err := orderClient.Create(ctx, request)
		if err != nil {
			t.Fatalf("failed to create order: %v", err)
		}
		if resource == nil || resource.ID == "" {
			t.Fatalf("order can't be nil")
		}

		time.Sleep(2 * time.Second)

		// Refund an order
		transactionIDs := []map[string]interface{}{
			{
				"id":     resource.Transactions.Payments[0].ID,
				"amount": "25.00",
			},
		}

		refundRequest := map[string]interface{}{
			"transactions": transactionIDs,
		}

		refundResp, err := orderClient.Refund(ctx, resource.ID, refundRequest)
		if err != nil {
			t.Fatalf("failed to refund order: %v", err)
		}

		if refundResp == nil || refundResp.Status != "partially_refunded" {
			t.Fatalf("expected order status to be `partially_refunded`, got %v", refundResp.StatusDetail)
		}
	})
}*/
