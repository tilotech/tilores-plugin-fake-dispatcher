package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/tilotech/tilores-plugin-api/dispatcher"
	"github.com/tilotech/tilores-plugin-fake-dispatcher/pkg"
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
