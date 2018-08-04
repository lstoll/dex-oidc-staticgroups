package ksconnector

import (
	"github.com/coreos/dex/connector"
	"github.com/sirupsen/logrus"
)

// Config holds configuration options for kubernetes secret oriented logins.
type Config struct {
}

// Open returns a strategy for logging in through GitHub.
func (c *Config) Open(id string, logger logrus.FieldLogger) (connector.Connector, error) {
	h := &Connector{
		Config: c,
	}

	return &h, nil
}
