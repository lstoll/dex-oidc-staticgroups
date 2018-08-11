package oidcstatic

import (
	"context"
	"net/http"

	"github.com/coreos/dex/connector"
)

var (
	_ connector.CallbackConnector = (*Connector)(nil)
	_ connector.RefreshConnector  = (*Connector)(nil)
)

type Connector struct {
	wrapped connector.Connector
}

func (c *Connector) LoginURL(s connector.Scopes, callbackURL, state string) (string, error) {
	conn, ok := c.wrapped.(connector.CallbackConnector)
	if !ok {
		panic("internal error - wrapped is not a callback connector")
	}
	return conn.LoginURL(s, callbackURL, state)
}

func (c *Connector) HandleCallback(s connector.Scopes, r *http.Request) (identity connector.Identity, err error) {
	conn, ok := c.wrapped.(connector.CallbackConnector)
	if !ok {
		panic("internal error - wrapped is not a callback connector")
	}

	// TODO - group check goes here

	return conn.HandleCallback(s, r)
}

// Refresh is implemented for backwards compatibility, even though it's a no-op.
func (c *Connector) Refresh(ctx context.Context, s connector.Scopes, identity connector.Identity) (connector.Identity, error) {
	// TODO - do we want a refresh mechanism?
	conn, ok := c.wrapped.(connector.RefreshConnector)
	if !ok {
		panic("internal error - wrapped is not a refresh connector")
	}
	return conn.Refresh(ctx, s, identity)
}
