// Package config provides SDK-level configuration for the MercadoPago Go SDK.
//
// A [Config] holds the authentication credentials, optional partner identifiers,
// and the HTTP [requester.Requester] used by every service client in the SDK.
// Create one with [New] and pass functional [Option] values to customise behaviour:
//
//	cfg, err := config.New("MY_ACCESS_TOKEN",
//	    config.WithCorporationID("corp-123"),
//	    config.WithHTTPClient(myCustomRequester),
//	)
package config

import (
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/internal/defaultrequester"
	"github.com/mercadopago/sdk-go/pkg/requester"
)

// Config holds the authentication credentials and transport settings needed to
// interact with the MercadoPago API. It is the central configuration object
// shared by all resource clients in the SDK.
type Config struct {
	// Requester is the HTTP transport used to execute API calls.
	// When not overridden via [WithHTTPClient], a default implementation with
	// automatic retries and constant back-off is used.
	Requester requester.Requester

	// AccessToken is the OAuth Bearer token used to authenticate every request
	// against the MercadoPago API.
	AccessToken string

	// CorporationID is an optional identifier sent in the X-Corporation-Id header
	// to associate requests with a specific corporation in MercadoPago's
	// partner/marketplace model.
	CorporationID string

	// IntegratorID is an optional identifier sent in the X-Integrator-Id header
	// to track which certified integrator originated the request.
	IntegratorID string

	// PlatformID is an optional identifier sent in the X-Platform-Id header
	// to identify the e-commerce platform on behalf of which the request is made.
	PlatformID string

	// ExpandNodes is a comma-separated list of response nodes that the API should
	// expand inline, sent via the X-Expand-Responde-Nodes header.
	ExpandNodes string
}

// New creates a new [Config] initialised with the given accessToken and any
// supplied functional [Option] values. It returns an error if any option fails
// to apply (for example, if a nil requester is provided via [WithHTTPClient]).
func New(accessToken string, opts ...Option) (*Config, error) {
	cfg := &Config{
		AccessToken: accessToken,
		Requester:   defaultrequester.New(),
	}

	// Apply all the functional options to configure the client.
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, fmt.Errorf("fail to build config: %w", err)
		}
	}

	return cfg, nil
}
