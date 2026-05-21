// Package webhook provides utilities to verify the authenticity of incoming
// MercadoPago webhook notifications.
//
// The package exposes a single entry point, [ValidateSignature], that
// recomputes the HMAC-SHA256 signature locally and compares it in constant
// time against the value carried in the x-signature header. The package
// is stateless, performs no outbound HTTP calls, and does not depend on
// any other SDK configuration; the integrator passes the secret signature
// explicitly on every call.
//
// QR Code notifications are not signed by MercadoPago — do not call this
// validator for those events; they will always fail signature verification.
package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Reason enumerates why [ValidateSignature] may reject a notification.
// Integrators are encouraged to log this value alongside the x-request-id
// for correlation against the MercadoPago notifications dashboard.
type Reason string

const (
	// ReasonMissingSignatureHeader: the x-signature header was missing, empty, or whitespace.
	ReasonMissingSignatureHeader Reason = "MissingSignatureHeader"

	// ReasonMalformedSignatureHeader: the header did not match the expected
	// ts=...,vN=... format and could not be parsed.
	ReasonMalformedSignatureHeader Reason = "MalformedSignatureHeader"

	// ReasonMissingTimestamp: the header parsed correctly but no ts= component was present.
	ReasonMissingTimestamp Reason = "MissingTimestamp"

	// ReasonMissingHash: no hash was found in the header for any of the supported versions.
	// Typically indicates MercadoPago has migrated to a new signature version and the
	// SDK needs to be upgraded.
	ReasonMissingHash Reason = "MissingHash"

	// ReasonSignatureMismatch: the computed HMAC did not match the value in the header.
	// Most often caused by an incorrect secret signature or a forged request.
	ReasonSignatureMismatch Reason = "SignatureMismatch"

	// ReasonTimestampOutOfTolerance: the header timestamp fell outside the configured
	// tolerance window against the current clock.
	ReasonTimestampOutOfTolerance Reason = "TimestampOutOfTolerance"
)

// SignatureError represents a webhook signature verification failure.
// Detect it with [errors.As]:
//
//	var sigErr *webhook.SignatureError
//	if errors.As(err, &sigErr) {
//	    log.Warn().Str("reason", string(sigErr.Reason)).Msg("rejected")
//	    w.WriteHeader(http.StatusUnauthorized)
//	}
type SignatureError struct {
	// Reason describes the specific failure mode.
	Reason Reason

	// RequestID echoes the x-request-id header value, when available.
	RequestID string

	// Timestamp echoes the ts value extracted from the x-signature header,
	// when parsing reached that point.
	Timestamp string
}

// Error implements the error interface.
func (e *SignatureError) Error() string {
	return fmt.Sprintf("webhook: invalid signature: %s", e.Reason)
}

// Option configures the behaviour of [ValidateSignature]. Following the
// functional-options pattern used elsewhere in this SDK (see pkg/config),
// each Option returns an error to allow validation of the supplied value.
type Option func(*options) error

type options struct {
	tolerance         time.Duration
	toleranceSet      bool
	supportedVersions []string
	now               func() time.Time
}

// WithTolerance enables the timestamp drift check with the given window.
// When omitted, no timestamp check is performed.
func WithTolerance(d time.Duration) Option {
	return func(o *options) error {
		if d < 0 {
			return fmt.Errorf("tolerance must be non-negative, got %s", d)
		}
		o.tolerance = d
		o.toleranceSet = true
		return nil
	}
}

// WithSupportedVersions overrides the default list of signature versions
// the validator will accept (default: []string{"v1"}). The validator
// iterates in order and uses the first version found in the header.
func WithSupportedVersions(versions ...string) Option {
	return func(o *options) error {
		if len(versions) == 0 {
			return fmt.Errorf("supportedVersions must not be empty")
		}
		o.supportedVersions = versions
		return nil
	}
}

// WithNow overrides the clock used for the tolerance check. Intended for tests.
func WithNow(now func() time.Time) Option {
	return func(o *options) error {
		if now == nil {
			return fmt.Errorf("now must not be nil")
		}
		o.now = now
		return nil
	}
}

var (
	defaultSupportedVersions = []string{"v1"}
	versionKeyRegex          = regexp.MustCompile(`^v\d+$`)
)

// ValidateSignature verifies a MercadoPago webhook notification.
//
// Inputs:
//
//   - xSignature: raw value of the x-signature header (e.g. "ts=...,v1=...").
//   - xRequestID: raw value of the x-request-id header. May be empty; in that
//     case the request-id: pair is omitted from the manifest.
//   - dataID: value of the data.id query parameter. May be empty; in that
//     case the id: pair is omitted. When present, it is lowercased before
//     being included in the manifest.
//   - secret: the secret signature configured for the application in
//     Tus Integraciones. Used as the HMAC key.
//
// On failure it returns a non-nil error that wraps [ErrInvalidSignature];
// use [errors.As] with [*SignatureError] to inspect the specific [Reason],
// [SignatureError.RequestID], and [SignatureError.Timestamp].
//
// The comparison is performed in constant time via [hmac.Equal] to mitigate
// timing attacks.
func ValidateSignature(xSignature, xRequestID, dataID, secret string, opts ...Option) error {
	o := &options{}
	for _, opt := range opts {
		if err := opt(o); err != nil {
			return fmt.Errorf("webhook: invalid option: %w", err)
		}
	}
	if len(o.supportedVersions) == 0 {
		o.supportedVersions = defaultSupportedVersions
	}
	if o.now == nil {
		o.now = time.Now
	}

	xSignature = normalize(xSignature)
	xRequestID = normalize(xRequestID)
	dataID = normalize(dataID)

	if xSignature == "" {
		return &SignatureError{Reason: ReasonMissingSignatureHeader, RequestID: xRequestID}
	}

	ts, hashes := parseSignatureHeader(xSignature)

	if ts == "" && len(hashes) == 0 {
		return &SignatureError{Reason: ReasonMalformedSignatureHeader, RequestID: xRequestID}
	}

	if ts == "" {
		return &SignatureError{Reason: ReasonMissingTimestamp, RequestID: xRequestID}
	}

	tsMs, parseErr := strconv.ParseInt(ts, 10, 64)
	if parseErr != nil {
		return &SignatureError{Reason: ReasonMalformedSignatureHeader, RequestID: xRequestID, Timestamp: ts}
	}

	var receivedHash string
	for _, v := range o.supportedVersions {
		if h, ok := hashes[v]; ok {
			receivedHash = h
			break
		}
	}

	if receivedHash == "" {
		return &SignatureError{Reason: ReasonMissingHash, RequestID: xRequestID, Timestamp: ts}
	}

	manifest := buildManifest(dataID, xRequestID, ts)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(manifest))
	computed := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(computed), []byte(receivedHash)) {
		return &SignatureError{Reason: ReasonSignatureMismatch, RequestID: xRequestID, Timestamp: ts}
	}

	if o.toleranceSet {
		nowMs := o.now().UnixMilli()
		drift := nowMs - tsMs
		if drift < 0 {
			drift = -drift
		}
		if drift > o.tolerance.Milliseconds() {
			return &SignatureError{Reason: ReasonTimestampOutOfTolerance, RequestID: xRequestID, Timestamp: ts}
		}
	}

	return nil
}

// normalize trims whitespace and treats empty values as missing.
func normalize(s string) string {
	return strings.TrimSpace(s)
}

// parseSignatureHeader extracts the ts and vN components from the x-signature header.
// Unknown keys are silently ignored.
func parseSignatureHeader(header string) (ts string, hashes map[string]string) {
	hashes = map[string]string{}
	for _, part := range strings.Split(header, ",") {
		rawKey, rawValue, ok := strings.Cut(part, "=")
		if !ok {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(rawKey))
		value := strings.TrimSpace(rawValue)
		if key == "" || value == "" {
			continue
		}
		if key == "ts" {
			ts = value
		} else if versionKeyRegex.MatchString(key) {
			hashes[key] = value
		}
	}
	return ts, hashes
}

// buildManifest assembles the HMAC manifest, omitting pairs whose value is empty.
func buildManifest(dataID, requestID, ts string) string {
	parts := make([]string, 0, 3)
	if dataID != "" {
		parts = append(parts, "id:"+strings.ToLower(dataID))
	}
	if requestID != "" {
		parts = append(parts, "request-id:"+requestID)
	}
	parts = append(parts, "ts:"+ts)
	return strings.Join(parts, ";") + ";"
}
