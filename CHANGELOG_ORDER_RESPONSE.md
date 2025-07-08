# Changelog Order Response

## Comparação da estrutura atual vs. JSON exemplo

### Estrutura Atual (`pkg/order/response.go`)
```go
type Response struct {
    ID                  string                  `json:"id,omitempty"`
    Type                string                  `json:"type,omitempty"`
    ExternalReference   string                  `json:"external_reference,omitempty"`
    CountryCode         string                  `json:"country_code,omitempty"`
    Status              string                  `json:"status,omitempty"`
    StatusDetail        string                  `json:"status_detail,omitempty"`
    CaptureMode         string                  `json:"capture_mode,omitempty"`
    UserID              string                  `json:"user_id,omitempty"`
    ClientToken         string                  `json:"client_token,omitempty"`
    TotalAmount         string                  `json:"total_amount,omitempty"`
    TotalPaidAmount     string                  `json:"total_paid_amount,omitempty"`
    ProcessingMode      string                  `json:"processing_mode,omitempty"`
    Description         string                  `json:"description,omitempty"`
    Marketplace         string                  `json:"marketplace,omitempty"`
    MarketplaceFee      string                  `json:"marketplace_fee,omitempty"`
    CheckoutAvailableAt string                  `json:"checkout_available_at,omitempty"`
    ExpirationTime      string                  `json:"expiration_time,omitempty"`
    CreatedDate         string                  `json:"created_date,omitempty"`
    LastUpdatedDate     string                  `json:"last_updated_date,omitempty"`
    Transactions        TransactionResponse     `json:"transactions,omitempty"`
    Items               []ItemsResponse         `json:"items,omitempty"`
    IntegrationData     IntegrationDataResponse `json:"integration_data,omitempty"`
    Config              ConfigResponse          `json:"config,omitempty"`
    Payer               PayerResponse           `json:"payer,omitempty"`
}
```

### JSON Exemplo
```json
{
  "id": "ORD01HRYFWNYRE1MR1E60MW3X0T2P",
  "type": "online",
  "processing_mode": "automatic",
  "external_reference": "ext_ref_1234",
  "description": "some description",
  "marketplace": "NONE",
  "marketplace_fee": "10.00",
  "total_amount": "1000.00",
  "total_paid_amount": "1000.00",
  "country_code": "BRA",
  "user_id": "1245621468",
  "status": "processed",
  "status_detail": "accredited",
  "capture_mode": "automatic_async",
  "created_date": "2024-11-21T14:19:14.727Z",
  "last_updated_date": "2024-11-21T14:19:18.489Z",
  "integration_data": {
    "application_id": "130106526144588"
  },
  "transactions": {
    "payments": [
      {
        "id": "PAY01JD7HETD7WX4W31VA60R1KC8E",
        "amount": "1000.00",
        "paid_amount": "1000.00",
        "expiration_time": "P1D",
        "date_of_expiration": "2024-01-01T00:00:00.000-03:00",
        "reference_id": "22dvqmsf4yc",
        "status": "processed",
        "status_detail": "accredited",
        "payment_method": {
          "id": "master",
          "type": "credit_card",
          "token": "677859ef5f18ea7e3a87c41d02c3fbe3",
          "statement_descriptor": "LOJA X",
          "installments": 1
        }
      }
    ]
  },
  "items":[
    {
      "external_code": "item_external_code",
      "category_id": "category_id",
      "title": "Some item title",
      "description": "Some item description",
      "unit_price": "1000.00",
      "type": "item type",
      "picture_url": "https://mysite.com/img/item.jpg",
      "quantity": 1
    }
  ]
}
```

## Análise das Diferenças

### Campos Ausentes na Estrutura Atual
1. Nenhum campo está ausente na estrutura principal `Response`

### Campos Ausentes nas Sub-estruturas
1. Em `PaymentResponse`:
   - `date_of_expiration` (string)
   - `paid_amount` (string)

### Campos Obsoletos (presentes na estrutura mas não no JSON)
1. Em `Response`:
   - `client_token`
   - `checkout_available_at`
   - `expiration_time`

### Campos com Tipos Diferentes
Nenhum campo apresenta tipo diferente entre a estrutura atual e o JSON exemplo.

## Recomendações de Atualização

1. Adicionar os campos ausentes na estrutura `PaymentResponse`:
```go
type PaymentResponse struct {
    // ... campos existentes ...
    DateOfExpiration string `json:"date_of_expiration,omitempty"`
    PaidAmount       string `json:"paid_amount,omitempty"`
}
```

2. Manter os campos obsoletos por compatibilidade, mas documentá-los como deprecated:
```go
type Response struct {
    // ... campos existentes ...
    
    // Deprecated: campo não mais utilizado pela API
    ClientToken         string                  `json:"client_token,omitempty"`
    // Deprecated: campo não mais utilizado pela API
    CheckoutAvailableAt string                  `json:"checkout_available_at,omitempty"`
    // Deprecated: campo não mais utilizado pela API
    ExpirationTime      string                  `json:"expiration_time,omitempty"`
}
```

## Impacto das Mudanças
- As alterações são retrocompatíveis pois apenas adicionam novos campos e marcam campos existentes como deprecated
- Não há necessidade de alteração em código cliente existente
- Novos campos permitem melhor integração com a API atual 