package config

import (
	"github.com/mercadopago/sdk-go/pkg/internal/requester"
)

// Config allows you to send custom settings and API authentication
type Config struct {
	AccessToken   string
	CorporationID string
	IntegratorID  string
	PlatformID    string
	Requester     requester.Requester
}

// New returns a new Config.
func New(accessToken string, opts ...Option) (*Config, error) {
	cfg := &Config{
		AccessToken: accessToken,
		Requester:   requester.Default(),
	}

	// Apply all the functional options to configure the client.
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg, nil
}
