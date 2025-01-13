## VERSION 1.0.8

Fix `LastChargedAmount` to float64 in `Response`

## VERSION 1.0.7

Fix `CardID` to string in `Response` from the `pre-approval flow` 

Fix `UserID` to int64

## VERSION 1.0.6

Include `reference_id` in `DataResponse` in `PaymentMethodResponse` in `Payment`.

Include `external_reference_id` in `DataResponse` in `PaymentMethodResponse` in `Payment`.

Include `external_resource_url` in `DataResponse` in `PaymentMethodResponse` in `Payment`.

## VERSION 1.0.5

Fix `is_same_bank_account_owner` type. Now it is `bool` instead of `string`.

## VERSION 1.0.4

Include `sub_type` in `PointOfInteractionRequest` in `Payment`.

## VERSION 1.0.3

Include `sub_merchant` in `ForwardData` in `Payment`.

## VERSION 1.0.2

Fix `differential_pricing_id` type. Now it is `int` instead of `string`.

## VERSION 1.0.1

Solution for issue [56](https://github.com/mercadopago/sdk-go/issues/56).

## VERSION 1.0.0

Initial release
