package ksconnector

import (
	"context"

	"github.com/coreos/dex/connector"
)

var (
	// we are a connector
	_ connector.Connector = (*Connector)(nil)
	// of type PasswordConnector
	_ connector.PasswordConnector = (*Connector)(nil)
	// that supports refresh
	_ connector.RefreshConnector = (*Connector)(nil)
)

// connectorData is what we persist per connection
type connectorData struct{}

type Connector struct {
	Config *Config
}

func (c *Connector) Prompt() string {
	return "username"
}

func (c *Connector) Login(ctx context.Context, s connector.Scopes, username, password string) (identity connector.Identity, validPassword bool, err error) {
	return connector.Identity{}, false, nil
}

func (c *Connector) Refresh(ctx context.Context, s connector.Scopes, identity connector.Identity) (connector.Identity, error) {
	return connector.Identity{}, nil
}
