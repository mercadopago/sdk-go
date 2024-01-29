package config

import (
	"github.com/mercadopago/sdk-go/pkg/internal/requester"
)

type Option func(*Config)

// WithHTTTPClient allow to do requests using received http client.
func WithHTTTPClient(r requester.Requester) Option {
	return func(c *Config) {
		if r != nil {
			c.requester = r
		}
	}
}

// WithCorporationID send corporation id to api by headers.
func WithCorporationID(value string) Option {
	return func(c *Config) {
		c.corporationID = value
	}
}

// WithIntegratorID send integrator id to api by headers.
func WithIntegratorID(i string) Option {
	return func(c *Config) {
		c.integratorID = i
	}
}

// WithPlatformID send platform id to api by headers.
func WithPlatformID(p string) Option {
	return func(c *Config) {
		c.platformID = p
	}
}
