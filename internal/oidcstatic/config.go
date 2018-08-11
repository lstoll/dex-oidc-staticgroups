package oidcstatic

import (
	"github.com/coreos/dex/connector"
	"github.com/coreos/dex/connector/oidc"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config holds configuration options for OpenID Connect logins.
type Config struct {
	// OIDCConfig maps the standard OIDC connector config in, to configure what we wrap.
	OIDCConfig oidc.Config `json:"oidc"`

	// GroupFile is the path to the file containing user/group mappings
	GroupFile string `json:"group_file"`
}

// Open returns a connector which can be used to login users through an upstream
// OpenID Connect provider.
func (c *Config) Open(id string, logger logrus.FieldLogger) (conn connector.Connector, err error) {
	wr, err := c.OIDCConfig.Open(id, logger)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening wrapped oidc connector")
	}

	m, err := LoadMappings(c.GroupFile)
	if err != nil {
		return nil, errors.Wrap(err, "Error loading mappings")
	}

	return &Connector{
		wrapped:  wr,
		mappings: m,
		logger:   logger,
	}, nil
}
