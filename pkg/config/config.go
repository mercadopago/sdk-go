package config

import (
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/internal/defaultrequester"
	"github.com/mercadopago/sdk-go/pkg/requester"
)

// Config allows you to send custom settings and API authentication
type Config struct {
	Requester requester.Requester

	AccessToken   string
	CorporationID string
	IntegratorID  string
	PlatformID    string
	ExpandNodes   string
}

// New returns a new Config.
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
