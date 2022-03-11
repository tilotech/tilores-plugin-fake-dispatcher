package main

import (
	"fmt"

	"github.com/tilotech/go-plugin"
	"github.com/tilotech/tilores-plugin-api/dispatcher"
	"github.com/tilotech/tilores-plugin-fake-dispatcher/pkg"
)

func main() {
	err := plugin.ListenAndServe(dispatcher.Provide(&pkg.FakeDispatcher{}))
	if err != nil {
		fmt.Println(err)
	}
}
