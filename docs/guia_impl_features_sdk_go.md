## Guia prático — Implementar novas features no SDK Go (ex.: Orders 3DS)

Este guia descreve, de forma objetiva, como identificar, desenhar e implementar novas features no SDK Go, usando 3DS (3D Secure) em Orders como caso de referência.

Referências cruzadas (alinhamento multiplataforma):
- Java — ["[Feature] Orders 3DS"](https://github.com/mercadopago/sdk-java/pull/323)
- Node.js — ["feat: adding 3ds feature"](https://github.com/mercadopago/sdk-nodejs/pull/402)
- .NET — [PR relacionado a Orders 3DS](https://github.com/mercadopago/sdk-dotnet/pull/222)

---

### 1) Monitoramento de Mudanças

- O que observar
  - Anúncios de produto e changelogs de APIs.
  - Itens de backlog (Jira/Boards), RFCs e comunicações do time de pagamentos/fraud.
  - PRs em outros SDKs (Java/Node/.NET) para manter paridade de recursos.

- Ações rápidas
  - Abrir issue interna descrevendo a feature, links de referência e decisões.
  - Criar branch temática: `feature/<produto>_<feature>` (ex.: `feature/orders_3ds`).
  - Definir critérios de aceite (request/response, exemplos, testes, docs, changelog).

---

### 2) Verificação no Fury Spec Hub

- O que confirmar na spec
  - Endpoints e métodos (ex.: `POST /v1/orders`).
  - Estrutura de Request e Response (campos, tipos, obrigatoriedade, valores válidos).
  - Exemplos de payloads (mínimos e completos).
  - Fluxos de estado relevantes (ex.: `processed`, `action_required` + `pending_challenge`).
  - Versionamento da API e estabilidade (ex.: Stable).

- Saídas esperadas
  - JSON de referência para Request e Response anotados.
  - Lista de campos novos e seu mapeamento para o SDK Go (nomes, tipos, opcionalidade).

---

### 3) Análise da Spec e Exemplos

- Inspiração em outros SDKs (garantir consistência):
  - Java: request `config.online.transaction_security` e response `payment_method.transaction_security` com `url` e metadados.
  - Node.js e .NET: validar naming, campos opcionais e exemplos de uso.

- Decisões para o SDK Go
  - Manter nomes em PascalCase nas structs e `json:"snake_case"` nas tags.
  - Campos novos como opcionais (ponteiros ou `omitempty`) para retrocompatibilidade.
  - Evitar enums rígidos; usar `string` e documentar valores aceitos.

---

### 4) Entendendo a Estrutura de Pastas (SDK Go)

- `pkg/<recurso>/`
  - `client.go`: cliente HTTP do recurso, orquestra chamadas.
  - `request.go`: modelos de entrada (Request) e subestruturas.
  - `response.go`: modelos de saída (Response) e subestruturas.

- `examples/apis/<recurso>/<operacao>/main.go`
  - Exemplos executáveis de uso real da API.

- `test/integration/<recurso>/`
  - Testes de integração ponta a ponta por operação.

- `resources/mocks/<recurso>/`
  - JSONs de exemplo para respostas conhecidas (quando aplicável).

---

### 5) Mapeamento JSON e exemplo (Orders 3DS)

- Request (ex.: `config.online.transaction_security`)
```go
type OnlineConfigRequest struct {
    CallbackURL         string                      `json:"callback_url,omitempty"`
    SuccessURL          string                      `json:"success_url,omitempty"`
    PendingURL          string                      `json:"pending_url,omitempty"`
    FailureURL          string                      `json:"failure_url,omitempty"`
    AutoReturnURL       string                      `json:"auto_return_url,omitempty"`
    DifferentialPricing *DifferentialPricingRequest `json:"differential_pricing,omitempty"`
    TransactionSecurity *TransactionSecurityRequest `json:"transaction_security,omitempty"`
}

type TransactionSecurityRequest struct {
    Validation     string `json:"validation,omitempty"`      // "always" | "on_fraud_risk" | "never"
    LiabilityShift string `json:"liability_shift,omitempty"` // "required" | "preferred"
}
```

- Response (ex.: `payment_method.transaction_security`)
```go
type PaymentMethodResponse struct {
    // ... campos existentes ...
    TransactionSecurity *TransactionSecurityResponse `json:"transaction_security,omitempty"`
}

type TransactionSecurityResponse struct {
    URL            string `json:"url,omitempty"`
    Validation     string `json:"validation,omitempty"`
    LiabilityShift string `json:"liability_shift,omitempty"`
}
```

- Exemplo de criação com 3DS
```go
req := order.Request{
    Type:        "online",
    TotalAmount: "100.00",
    Transactions: &order.TransactionRequest{Payments: []order.PaymentRequest{{
        Amount: "100.00",
        PaymentMethod: &order.PaymentMethodRequest{
            ID:           "master",
            Token:        "{{CARD_TOKEN}}",
            Type:         "credit_card",
            Installments: 1,
        },
    }}},
    Payer: &order.PayerRequest{Email: "buyer@example.com"},
    Config: &order.ConfigRequest{Online: &order.OnlineConfigRequest{
        TransactionSecurity: &order.TransactionSecurityRequest{
            Validation:     "on_fraud_risk",
            LiabilityShift: "required",
        },
    }},
}

res, err := client.Create(ctx, req)
if err != nil { /* tratar */ }

if res.Status == "action_required" && res.StatusDetail == "pending_challenge" {
    ts := res.Transactions.Payments[0].PaymentMethod.TransactionSecurity
    if ts != nil && ts.URL != "" {
        // direcionar o comprador para completar o challenge 3DS (iframe/redirect)
    }
}
```

---

### 6) Passo a passo de implementação

1. Criar branch: `feature/<recurso>_<feature>`
2. Models de Request: adicionar campos/estruturas novas em `pkg/<recurso>/request.go`
3. Models de Response: adicionar campos/estruturas novas em `pkg/<recurso>/response.go`
4. Cliente (`client.go`): avaliar se há novos endpoints/ações a suportar
5. Exemplo: criar `examples/apis/<recurso>/<operacao>` com `main.go`
6. Testes de integração: cobrir caso mínimo de sucesso e cenários-chave
7. Documentação: criar/atualizar `docs/*.md` com guia e exemplos
8. Changelog: descrever a feature (escopo, breaking changes = nenhum)
9. Rodar testes: `go test ./...` e validar exemplos com `go run`

---

### 7) Boas práticas

- Retrocompatibilidade: novos campos como opcionais (`omitempty`, ponteiros quando necessário).
- Tipagem: usar `string` para valores enumerados; documentar valores suportados.
- Erros: propagar erros do cliente e mensagens claras.
- Contexto/Timeout: usar `context.Context` e tempos configuráveis no http.Client, quando aplicável.
- Consistência: seguir padrões de nome do repo e de outros SDKs.
- PRs pequenos e focados; incluir exemplos e testes juntos.

---

### 8) Checklist de revisão

- Spec conferida (endpoints, modelos, estados, exemplos)
- Mapeamento Request/Response fiel e opcionalidade correta
- Cliente atualizado (se necessário)
- Exemplo funcional incluído
- Testes de integração executando e cobrindo o fluxo mínimo
- Documentação e changelog atualizados
- Paridade com outros SDKs validada (Java/Node/.NET)

---

### 9) Como validar

- Gerar token de cartão (ex.: exemplo em `examples/apis/cardtoken/create`).
- Executar exemplo da feature: `go run ./examples/apis/<recurso>/<operacao>`.
- Rodar testes: `go test ./... -run Test<Order|Feature>`.
- Verificar estados de Order e presença de `transaction_security.url` quando challenge for requerido.

---

### 10) Publicação

- Atualizar `CHANGELOG.md` com a nova versão e feature.
- Criar tag e release.
- Opcional: destacar no `README.md` e em exemplos.


