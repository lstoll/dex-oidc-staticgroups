package oidcstatic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/dex/connector"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	_ connector.CallbackConnector = (*Connector)(nil)
	_ connector.RefreshConnector  = (*Connector)(nil)
)

type Connector struct {
	wrapped connector.Connector

	mappings MappingFile

	logger logrus.FieldLogger
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

	identity, err = conn.HandleCallback(s, r)
	if err != nil {
		return identity, errors.Wrap(err, "Error handling callback in wrapped")
	}

	// Because we rely on the email to ensure the user maps, make sure it's actually verified
	if !identity.EmailVerified {
		c.logger.WithField("email", identity.Email).Warn("Email not verified")
		return identity, fmt.Errorf("Email %s not verified", identity.Email)
	}

	// check if a) the email exists and b) if we should attach the groups
	em, ok := c.mappings.Email[identity.Email]
	if !ok {
		c.logger.WithField("email", identity.Email).Warn("Email not in mapping")
		return identity, fmt.Errorf("Email %s not permitted access", identity.Email)
	}

	if s.Groups {
		identity.Groups = em.Groups
	}

	return identity, nil
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
