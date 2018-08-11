package main

import (
	"github.com/coreos/dex/server"
	"github.com/lstoll/dex-oidc-staticgroups/internal/oidcstatic"
)

func init() {
	server.ConnectorsConfig["oidc-staticgroups"] = func() server.ConnectorConfig { return new(oidcstatic.Config) }
}
