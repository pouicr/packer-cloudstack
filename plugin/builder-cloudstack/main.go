package main

import (
	"github.com/schubergphilis/packer-cloudstack"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(cloudstack.Builder))
	server.Serve()
}
