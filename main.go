package main

import (
	"github.com/hashicorp/go-plugin"
	"gitlab.com/tilotech/tilores-plugin-api/dispatcher"
	"gitlab.com/tilotech/tilores-plugin-fake-dispatcher/pkg"
)

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"dispatcher": &dispatcher.Plugin{
			Impl: &pkg.FakeDispatcher{},
		},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: dispatcher.Handshake,
		Plugins:         pluginMap,
	})
}
