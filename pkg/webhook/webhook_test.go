package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"testing"
	"time"
)

const (
	secret       = "your_secret_key_here"
	requestID    = "2066ca19-c6f1-498a-be75-1923005edd06"
	dataIDRaw    = "ORD01JQ4S4KY8HWQ6NA5PXB65B3D3"
	dataIDLower  = "ord01jq4s4ky8hwq6na5pxb65b3d3"
	ts           = "1742505638683"
	tsNum  int64 = 1742505638683
)

func computeHash(dataID, requestID, ts, secret string) string {
	var parts []string
	if dataID != "" {
		parts = append(parts, "id:"+dataID)
	}
	if requestID != "" {
		parts = append(parts, "request-id:"+requestID)
	}
	parts = append(parts, "ts:"+ts)
	manifest := strings.Join(parts, ";") + ";"
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(manifest))
	return hex.EncodeToString(mac.Sum(nil))
}

func buildHeader(hash string, optTs ...string) string {
	t := ts
	if len(optTs) > 0 {
		t = optTs[0]
	}
	return "ts=" + t + ",v1=" + hash
}

func expectReason(t *testing.T, err error, want Reason) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var sigErr *SignatureError
	if !errors.As(err, &sigErr) {
		t.Fatalf("expected *SignatureError, got %T", err)
	}
	if sigErr.Reason != want {
		t.Errorf("Reason: got %q, want %q", sigErr.Reason, want)
	}
}

// case 1
func TestValidateSignature_HappyPathLowercase(t *testing.T) {
	h := computeHash(dataIDLower, requestID, ts, secret)
	if err := ValidateSignature(buildHeader(h), requestID, dataIDLower, secret); err != nil {
		t.Fatal(err)
	}
}

// case 2
func TestValidateSignature_UppercaseDataIDIsPreserved(t *testing.T) {
	h := computeHash(dataIDRaw, requestID, ts, secret)
	if err := ValidateSignature(buildHeader(h), requestID, dataIDRaw, secret); err != nil {
		t.Fatal(err)
	}
}

// case 3
func TestValidateSignature_MalformedHeader(t *testing.T) {
	err := ValidateSignature("this-is-garbage", requestID, dataIDLower, secret)
	expectReason(t, err, ReasonMalformedSignatureHeader)
}

// case 4
func TestValidateSignature_MissingHeader(t *testing.T) {
	err := ValidateSignature("", requestID, dataIDLower, secret)
	expectReason(t, err, ReasonMissingSignatureHeader)
}

// case 5
func TestValidateSignature_MissingTimestamp(t *testing.T) {
	h := computeHash(dataIDLower, requestID, ts, secret)
	err := ValidateSignature("v1="+h, requestID, dataIDLower, secret)
	expectReason(t, err, ReasonMissingTimestamp)
}

// case 6
func TestValidateSignature_MissingV1(t *testing.T) {
	err := ValidateSignature("ts="+ts, requestID, dataIDLower, secret)
	expectReason(t, err, ReasonMissingHash)
}

// case 7
func TestValidateSignature_TamperedHash(t *testing.T) {
	h := computeHash(dataIDLower, requestID, ts, secret)
	tampered := h[:len(h)-2] + "00"
	if strings.HasSuffix(h, "00") {
		tampered = h[:len(h)-2] + "ff"
	}
	err := ValidateSignature(buildHeader(tampered), requestID, dataIDLower, secret)
	expectReason(t, err, ReasonSignatureMismatch)
}

// case 8
func TestValidateSignature_OutsideTolerance(t *testing.T) {
	h := computeHash(dataIDLower, requestID, ts, secret)
	far := func() time.Time { return time.UnixMilli(tsNum + 10*60*1000) }
	err := ValidateSignature(buildHeader(h), requestID, dataIDLower, secret,
		WithTolerance(60*time.Second), WithNow(far))
	expectReason(t, err, ReasonTimestampOutOfTolerance)
}

func TestValidateSignature_WithinTolerance(t *testing.T) {
	h := computeHash(dataIDLower, requestID, ts, secret)
	near := func() time.Time { return time.UnixMilli(tsNum + 30*1000) }
	if err := ValidateSignature(buildHeader(h), requestID, dataIDLower, secret,
		WithTolerance(60*time.Second), WithNow(near)); err != nil {
		t.Fatal(err)
	}
}

// case 9
func TestValidateSignature_DataIDAbsent(t *testing.T) {
	h := computeHash("", requestID, ts, secret)
	if err := ValidateSignature(buildHeader(h), requestID, "", secret); err != nil {
		t.Fatal(err)
	}
}

// case 10
func TestValidateSignature_RequestIDAbsent(t *testing.T) {
	h := computeHash(dataIDLower, "", ts, secret)
	if err := ValidateSignature(buildHeader(h), "", dataIDLower, secret); err != nil {
		t.Fatal(err)
	}
}

// case 11
func TestValidateSignature_BothAbsent(t *testing.T) {
	h := computeHash("", "", ts, secret)
	if err := ValidateSignature(buildHeader(h), "  ", "", secret); err != nil {
		t.Fatal(err)
	}
}

// case 12
func TestValidateSignature_NonPaymentTopic(t *testing.T) {
	orderID := "ord01abc123"
	h := computeHash(orderID, requestID, ts, secret)
	if err := ValidateSignature(buildHeader(h), requestID, orderID, secret); err != nil {
		t.Fatal(err)
	}
}

// supportedVersions
func TestValidateSignature_SupportsV1WhenBothPresent(t *testing.T) {
	h := computeHash(dataIDLower, requestID, ts, secret)
	header := "ts=" + ts + ",v1=" + h + ",v2=aaaa"
	if err := ValidateSignature(header, requestID, dataIDLower, secret); err != nil {
		t.Fatal(err)
	}
}

func TestValidateSignature_OnlyV2InHeaderOnlyV1Supported(t *testing.T) {
	header := "ts=" + ts + ",v2=somehash"
	err := ValidateSignature(header, requestID, dataIDLower, secret)
	expectReason(t, err, ReasonMissingHash)
}

func TestValidateSignature_ErrorsAsSignatureError(t *testing.T) {
	err := ValidateSignature("garbage", requestID, dataIDLower, secret)
	var sigErr *SignatureError
	if !errors.As(err, &sigErr) {
		t.Errorf("expected errors.As(err, &SignatureError) == true, got false")
	}
}
