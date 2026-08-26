package order

import (
	"encoding/json"
	"testing"
)

func TestAutomaticPaymentsSubscriptionSerializes(t *testing.T) {
	request := AutomaticPaymentsRequest{Subscription: &AutomaticPaymentsSubscriptionRequest{
		ID:       "subscription-1",
		Sequence: &SubscriptionSequenceRequest{Number: 1, Total: 12},
		Invoice:  &AutomaticPaymentsInvoiceRequest{ID: "invoice-1", BillingDate: "2026-08-26", Period: &AutomaticPaymentsPeriodRequest{Interval: 1, Type: "month"}},
	}}
	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}
	if string(body) != `{"subscription":{"id":"subscription-1","sequence":{"number":1,"total":12},"invoice":{"id":"invoice-1","billing_date":"2026-08-26","period":{"interval":1,"type":"month"}}}}` {
		t.Fatalf("unexpected payload: %s", body)
	}
}

func TestAutomaticPaymentsSubscriptionDeserializes(t *testing.T) {
	var response AutomaticPaymentResponse
	if err := json.Unmarshal([]byte(`{"subscription":{"id":"subscription-1","sequence":{"number":1,"total":12},"invoice":{"id":"invoice-1","billing_date":"2026-08-26","period":{"interval":1,"type":"month"}}}}`), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if response.Subscription.Invoice.Period.Type != "month" || response.Subscription.Sequence.Total != 12 {
		t.Fatalf("unexpected subscription: %#v", response.Subscription)
	}
}
