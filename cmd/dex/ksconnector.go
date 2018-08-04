package main

import (
	"github.com/coreos/dex/server"
)

func init() {
	server.ConnectorsConfig["kubesecrets"] = func() server.ConnectorConfig { return new(ksconnector.Config) }
}
