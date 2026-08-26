package payment

import (
	"encoding/json"
	"testing"
)

func TestNetworkDataIsNestedInTransactionData(t *testing.T) {
	request := Request{PointOfInteraction: &PointOfInteractionRequest{TransactionData: &TransactionDataRequest{
		NetworkTransactionID: "network-transaction-id",
		NetworkData: &NetworkDataRequest{
			TransactionID:     "VISA-TID-ABC123",
			TransactionLinkID: "550e8400-e29b-41d4-a716-446655440000",
		},
	}}}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal payload: %v", err)
	}
	poi := payload["point_of_interaction"].(map[string]any)
	transactionData := poi["transaction_data"].(map[string]any)
	networkData := transactionData["network_data"].(map[string]any)
	if networkData["transaction_id"] != "VISA-TID-ABC123" || networkData["transaction_link_id"] != "550e8400-e29b-41d4-a716-446655440000" {
		t.Fatalf("unexpected network data: %#v", networkData)
	}
	if _, found := poi["network_data"]; found {
		t.Fatal("network_data must not be serialized at point_of_interaction level")
	}
}

func TestNetworkDataIsDeserializedFromTransactionData(t *testing.T) {
	var response Response
	if err := json.Unmarshal([]byte(`{"point_of_interaction":{"transaction_data":{"network_data":{"transaction_id":"VISA-TID-ABC123","transaction_link_id":"550e8400-e29b-41d4-a716-446655440000"}}}}`), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if response.PointOfInteraction.TransactionData.NetworkData.TransactionID != "VISA-TID-ABC123" {
		t.Fatal("transaction_id was not deserialized")
	}
}

func TestExpandedGatewayReferenceNetworkDataDeserializes(t *testing.T) {
	var response Response
	if err := json.Unmarshal([]byte(`{"expanded":{"gateway":{"reference":{"network_data":{"transaction_id":"ABC123","transaction_link_id":"550e8400"}}}}}`), &response); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if response.Expanded.Gateway.Reference.NetworkData.TransactionID != "ABC123" {
		t.Fatal("expanded gateway network_data was not deserialized")
	}
}
