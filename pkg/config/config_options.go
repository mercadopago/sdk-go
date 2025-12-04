package config

import (
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/requester"
)

type Option func(*Config) error

// WithHTTPClient allow to do requests using received http client.
func WithHTTPClient(r requester.Requester) Option {
	return func(c *Config) error {
		if r == nil {
			return fmt.Errorf("received http client can't be nil")
		}
		c.Requester = r
		return nil
	}
}

// WithCorporationID send corporation id to api by headers.
func WithCorporationID(value string) Option {
	return func(c *Config) error {
		c.CorporationID = value
		return nil
	}
}

// WithIntegratorID send integrator id to api by headers.
func WithIntegratorID(i string) Option {
	return func(c *Config) error {
		c.IntegratorID = i
		return nil
	}
}

// WithPlatformID send platform id to api by headers.
func WithPlatformID(p string) Option {
	return func(c *Config) error {
		c.PlatformID = p
		return nil
	}
}

// WithExpandNodes send nodes to be expanded in the response.
func WithExpandNodes(nodes string) Option {
	return func(c *Config) error {
		c.ExpandNodes = nodes
		return nil
	}
}
