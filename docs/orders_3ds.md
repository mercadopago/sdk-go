### Implementação de Suporte a 3DS (3D Secure) em Orders — SDK Go

Este documento guia a implementação e o uso de 3DS (Transaction Security) no módulo de Orders do SDK Go, espelhando a estrutura e abrangência do guia feito para o SDK PHP [link](https://github.com/mercadopago/sdk-php/pull/574).

### 1. Processo de Descoberta e Implementação

- Monitorar anúncios de mudanças da API.
- Verificar a spec estável (swagger) no hub interno.
- Identificar endpoints e exemplos relevantes (ex.: "POST /v1/orders - Online with 3DS").
- Mapear campos novos no request/response e seu impacto no SDK.

### 2. Verificação na Spec (Request e Response)

- Request: novo nó `config.online.transaction_security` com campos:
  - `validation`: `always`, `on_fraud_risk`, `never`.
  - `liability_shift`: `required`, `preferred`.
- Response: novo nó `payment_method.transaction_security` contendo, entre outros:
  - `url`: URL do Challenge 3DS (quando challenge é necessário).
  - `validation`, `liability_shift`, `type`, `status`.

### 3. Impacto no SDK Go

- Request (já suportado):
  - `order.TransactionSecurityRequest` em `order.OnlineConfigRequest` via `Config.Online.TransactionSecurity`.
- Response (adicionado):
  - `order.PaymentMethodResponse.TransactionSecurity *TransactionSecurityResponse` retornado no pagamento.

### 4. Estrutura do SDK Go (onde fica cada parte)

- `pkg/order/request.go`: Tipos de request, incluindo `ConfigRequest`, `OnlineConfigRequest` e `TransactionSecurityRequest`.
- `pkg/order/response.go`: Tipos de response, incluindo `PaymentMethodResponse` e `TransactionSecurityResponse`.
- `pkg/order/client.go`: Métodos públicos (`Create`, `Get`, `Process`, `Capture`, `Refund`, etc.).
- `examples/apis/order/create3ds/`: Exemplo prático de criação de Order com 3DS.

### 5. Alterações Técnicas (Go)

- `pkg/order/response.go`: adição de `TransactionSecurityResponse` e campo `TransactionSecurity` em `PaymentMethodResponse`.
- `test/integration/order/order_test.go`: cenário validando criação com `Config.Online.TransactionSecurity`.
- `examples/apis/order/create`: exemplo atualizado mostrando uso de `Config.Online.TransactionSecurity`.

### 6. Fluxo de Integração (passo a passo)

1) Gerar Card Token do comprador (ex.: `examples/apis/cardtoken/create`).
2) Montar `order.Request` com:
   - `Type: "online"`, valores de `TotalAmount`, `ExternalReference` e `Payer.Email`.
   - `Transactions.Payments[].PaymentMethod` com `id`, `token`, `type`, `installments`.
   - `Config.Online.TransactionSecurity` com `validation` e `liability_shift` conforme estratégia.
3) Chamar `orderClient.Create(ctx, request)`.
4) Ler a resposta e verificar `payment_method.transaction_security.url` no primeiro pagamento; se presente, conduzir o challenge 3DS (redireciono/iframe/webview).
5) Após retorno (redirect/webhook), decidir processar/capturar conforme `processing_mode` e `capture_mode`.
6) Tratar falhas (exibir mensagem, permitir nova tentativa, cancelar) e usar `Refund/Cancel/DeleteTransaction` quando aplicável.

### 7. Exemplo (Go) — habilitar 3DS na criação

```go
request := order.Request{
    Type:              "online",
    TotalAmount:       "100.00",
    ExternalReference: "ext_ref_123",
    Payer: &order.PayerRequest{ Email: "buyer@example.com" },
    Transactions: &order.TransactionRequest{
        Payments: []order.PaymentRequest{
            {
                Amount: "100.00",
                PaymentMethod: &order.PaymentMethodRequest{
                    ID:           "master",
                    Token:        "{{CARD_TOKEN}}",
                    Type:         "credit_card",
                    Installments: 1,
                },
            },
        },
    },
    Config: &order.ConfigRequest{
        Online: &order.OnlineConfigRequest{
            TransactionSecurity: &order.TransactionSecurityRequest{
                Validation:     "on_fraud_risk", // ou "always" / "never"
                LiabilityShift: "required",      // ou "preferred"
            },
            // (Opcional) URLs de retorno e callback
            // SuccessURL:  "https://example.com/success",
            // FailureURL:  "https://example.com/failure",
            // PendingURL:  "https://example.com/pending",
            // CallbackURL: "https://example.com/webhook",
        },
    },
}
```

### 8. Exemplo (Go) — interpretar a resposta e challenge URL

```go
res, err := client.Create(ctx, request)
if err != nil {
    // tratar erro
}

if len(res.Transactions.Payments) > 0 {
    p := res.Transactions.Payments[0]
    ts := p.PaymentMethod.TransactionSecurity
    if ts != nil && ts.URL != "" {
        // Redirecionar comprador para ts.URL (navegador/iframe/webview)
        // Aguardar retorno (redirect/webhook) e seguir com Process/Capture conforme necessário
    }
}
```

### 9. Valores Disponíveis e Estratégia

- `validation`:
  - `on_fraud_risk`: executa 3DS conforme risco (recomendado)
  - `always`: sempre tenta 3DS
  - `never`: nunca executa 3DS
- `liability_shift`:
  - `required`: prioriza transferência de responsabilidade (pode reduzir aprovação)
  - `preferred`: prioriza aprovação (transferência quando possível)

### 10. Status possíveis (pagamento)

| status           | status_detail                 | descrição                                       |
| ---------------- | ------------------------------ | ----------------------------------------------- |
| processed        | accredited                     | transação aprovada sem autenticação             |
| action_required  | pending_challenge              | transação pendente de autenticação (3DS)        |
| failed           | cc_rejected_3ds_challenge*     | challenge falhou ou não concluído               |
| cancelled        | expired                        | challenge expirado                              |

Observação: os valores podem variar por emissor/meio; sempre valide contra a spec vigente.

### 11. Modos de processamento e captura

- `ProcessingMode`: `automatic` (aut./captura automatizadas) vs `manual` (você controla `Process`/`Capture`).
- `CaptureMode`: `manual` exige `Capture` explícito mesmo com processamento automático.

### 12. Boas práticas

- Use `on_fraud_risk` para equilibrar segurança e aprovação.
- Configure URLs de retorno e webhooks para estados robustos.
- Registre `type`, `status` e resultado do challenge para análise/suporte.
- Revalide estado com `Get`/`Process`/`Capture` após o retorno do challenge.

### 13. Troubleshooting

- Sem `url` no `transaction_security`: challenge não requerido.
- Card token inválido/expirado: gere e use imediatamente.
- Status divergente: garanta processamento de webhooks/redirect e revalide com `Get`.

### 14. Referências

- Guia SDK PHP (base conceitual): [added 3ds — PR #574](https://github.com/mercadopago/sdk-php/pull/574)
- Tipos de Request Go: `pkg/order/request.go`
- Tipos de Response Go: `pkg/order/response.go`
- Exemplo Go: `examples/apis/order/create3ds/main.go`



### 15. Checklist para Novas Features (SDK Go)

- Verificar anúncio da API e spec estável (swagger) para confirmar campos e exemplos
- Mapear/request: adicionar/confirmar structs em `pkg/order/request.go`
- Mapear/response: adicionar/confirmar structs em `pkg/order/response.go`
- Atualizar/criar exemplo em `examples/` (aqui: `examples/apis/order/create3ds/`)
- Cobrir cenário em teste de integração (aqui: `TestOrderTransactionSecurity` em `test/integration/order/order_test.go`)
- Atualizar documentação em `docs/` e linkar no `README.md`
- Atualizar `CHANGELOG.md` com a nova feature

Baseado no guia do SDK PHP [PR #574](https://github.com/mercadopago/sdk-php/pull/574).


### 16. Comportamento Padrão e Diferenças Request vs Response (Go)

- Padrão quando `transaction_security` não é enviado: assume validação "never" (não executa 3DS) — confirme na spec vigente.
- Go usa structs tipados tanto para request quanto para response. A serialização/deserialização é feita via tags `json`.
- Acesso típico no Go (primeiro pagamento):

```go
res, _ := client.Create(ctx, request)
if len(res.Transactions.Payments) > 0 {
    p := res.Transactions.Payments[0]
    ts := p.PaymentMethod.TransactionSecurity
    switch {
    case ts != nil && ts.URL != "":
        // action_required/pending_challenge — redirecionar ou exibir iframe/webview
    default:
        // sem challenge — seguir fluxo normal (processed/failed)
    }
}
```


### 17. Como rodar exemplos e testes

1) Exemplo 3DS (edição rápida dos placeholders):

- Abra `examples/apis/order/create3ds/main.go` e troque os placeholders: `{{ACCESS_TOKEN}}`, `{{TOTAL_AMOUNT}}`, `{{EXTERNAL_REFERENCE}}`, `{{EMAIL}}`, `{{AMOUNT}}`, `{{PAYMENT_METHOD_ID}}`, `{{CARD_TOKEN}}`, `{{TYPE}}`.
- Rode:

```sh
go run ./examples/apis/order/create3ds
```

2) Teste de integração focado em 3DS (requer credenciais válidas):

- Exporte o token de acesso (se aplicável ao seu ambiente de testes):

```sh
export ACCESS_TOKEN="<SEU_ACCESS_TOKEN>"
```

- Execute apenas o teste de 3DS:

```sh
go test -v ./test/integration/order -run TestOrderTransactionSecurity
```

3) Dicas:

- Se o teste não retornar `url` em `transaction_security`, o emissor/bandeira pode não exigir challenge para o cenário. Valide o status e `status_detail` retornados.
- Para o fluxo completo (redirect/webhook), valide `Get`/`Process`/`Capture` após o retorno, conforme seu `processing_mode`/`capture_mode`.
