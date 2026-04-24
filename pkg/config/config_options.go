package config

import (
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/requester"
)

// Option is a functional option that configures a [Config] during construction
// via [New]. Each option returns an error if validation fails, allowing [New] to
// surface misconfiguration early.
type Option func(*Config) error

// WithHTTPClient returns an [Option] that replaces the default HTTP transport
// with the provided [requester.Requester]. This is useful for injecting custom
// timeouts, middleware, or a mock implementation during tests.
// It returns an error if r is nil.
func WithHTTPClient(r requester.Requester) Option {
	return func(c *Config) error {
		if r == nil {
			return fmt.Errorf("received http client can't be nil")
		}
		c.Requester = r
		return nil
	}
}

// WithCorporationID returns an [Option] that sets the corporation identifier
// sent in the X-Corporation-Id request header, used by MercadoPago to associate
// API calls with a specific corporation in marketplace integrations.
func WithCorporationID(value string) Option {
	return func(c *Config) error {
		c.CorporationID = value
		return nil
	}
}

// WithIntegratorID returns an [Option] that sets the integrator identifier sent
// in the X-Integrator-Id request header, used by MercadoPago to track which
// certified integrator originated the request.
func WithIntegratorID(i string) Option {
	return func(c *Config) error {
		c.IntegratorID = i
		return nil
	}
}

// WithPlatformID returns an [Option] that sets the platform identifier sent in
// the X-Platform-Id request header, used by MercadoPago to identify the
// e-commerce platform on behalf of which the request is made.
func WithPlatformID(p string) Option {
	return func(c *Config) error {
		c.PlatformID = p
		return nil
	}
}

// WithExpandNodes returns an [Option] that sets the comma-separated list of
// response nodes the API should expand inline. The value is sent in the
// X-Expand-Responde-Nodes request header.
func WithExpandNodes(nodes string) Option {
	return func(c *Config) error {
		c.ExpandNodes = nodes
		return nil
	}
}
