package config

import (
	"github.com/mercadopago/sdk-go/pkg/internal/requester"
)

type Config struct {
	accessToken   string
	corporationID string
	integratorID  string
	platformID    string
	requester     requester.Requester
}

// New returns a new Config.
func New(accessToken string, opts ...Option) (*Config, error) {
	cfg := &Config{
		accessToken: accessToken,
		requester:   requester.Default(),
	}

	// Apply all the functional options to configure the client.
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg, nil
}

func (c *Config) GetAccessToken() string {
	return c.accessToken
}

func (c *Config) GetCorporationID() string {
	return c.corporationID
}

func (c *Config) GetIntegratorID() string {
	return c.integratorID
}

func (c *Config) GetPlatformID() string {
	return c.platformID
}

func (c *Config) GetHTTPClient() requester.Requester {
	return c.requester
}
