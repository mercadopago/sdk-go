package config

import (
	"github.com/mercadopago/sdk-go/pkg/internal/requester"
	"github.com/mercadopago/sdk-go/pkg/option"
)

type Config struct {
	AccessToken   string
	CorporationID string
	IntegratorID  string
	PlatformID    string

	HTTPClient option.Requester
}

// New returns a new Config.
func New(accessToken string, opts ...option.ConfigOption) (*Config, error) {
	options := option.ConfigOptions{
		HTTPClient: requester.Default(),
	}
	for _, opt := range opts {
		opt.Apply(&options)
	}

	return &Config{
		AccessToken:   accessToken,
		CorporationID: options.CorporationID,
		IntegratorID:  options.IntegratorID,
		PlatformID:    options.PlatformID,
		HTTPClient:    options.HTTPClient,
	}, nil
}
