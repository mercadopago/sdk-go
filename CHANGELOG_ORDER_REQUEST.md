# Changelog Order Request

## Análise da estrutura atual vs. JSON exemplo

### Estrutura Principal (`Request`)
- ✅ Todos os campos do JSON exemplo estão presentes
- ✅ Campos obsoletos estão marcados como deprecated:
  - `CheckoutAvailableAt`
  - `Config`
  - `AdditionalInfo`

### Sub-estruturas

#### PaymentMethodRequest
- ✅ Todos os campos do JSON exemplo estão presentes:
  - `id`
  - `type`
  - `token`
  - `installments`
  - `statement_descriptor`

#### PayerRequest
- ✅ Todos os campos do JSON exemplo estão presentes:
  - `entity_type`
  - `email`
  - `first_name`
  - `last_name`
  - `identification`
  - `phone`
  - `address`

#### PayerAddressRequest
- ✅ Todos os campos do JSON exemplo estão presentes
- 🔄 Alterações realizadas:
  - Renomeado campo `Zipcode` para `ZipCode`
  - Alterado tag json de `zipcode` para `zip_code` para manter consistência com a API

#### ItemsRequest
- ✅ Todos os campos do JSON exemplo estão presentes:
  - `title`
  - `unit_price`
  - `quantity`
  - `description`
  - `external_code`
  - `category_id`
  - `type`
  - `picture_url`
- ℹ️ Campos adicionais (opcionais):
  - `warranty`
  - `event_date`

## JSON Exemplo
```json
{
  "type": "online",
  "total_amount": "1000.00",
  "external_reference": "ext_ref_1234",
  "capture_mode": "automatic_async",
  "transactions": {
    "payments": [
      {
        "amount": "1000.00",
        "expiration_time": "P1D",
        "payment_method": {
          "id": "master",
          "type": "credit_card",
          "token": "677859ef5f18ea7e3a87c41d02c3fbe3",
          "installments": 1,
          "statement_descriptor": "LOJA X"
        }
      }
    ]
  },
  "processing_mode": "automatic",
  "description": "some description",
  "payer": {
    "entity_type": "individual",
    "email": "test_123@testuser.com",
    "first_name": "John",
    "last_name": "Doe",
    "identification": {
      "type": "CPF",
      "number": "15635614680"
    },
    "phone": {
      "area_code": "55",
      "number": "99999999999"
    },
    "address": {
      "street_name": "R. Ângelo Piva",
      "street_number": "144",
      "zip_code": "06210110",
      "neighborhood": "Presidente Altino",
      "city": "Osasco",
      "state": "SP",
      "complement": "303"
    }
  },
  "marketplace": "NONE",
  "marketplace_fee": "10.00",
  "items": [
    {
      "title": "Some item title",
      "unit_price": "1000.00",
      "quantity": 1,
      "description": "Some item description",
      "external_code": "item_external_code",
      "category_id": "category_id",
      "type": "item type",
      "picture_url": "https://mysite.com/img/item.jpg"
    }
  ],
  "expiration_time": "P3D"
}
``` 