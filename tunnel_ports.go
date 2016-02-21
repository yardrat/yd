package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strconv"
)

type TunnelPorts struct {
	local  string
	remote string
}

func ReadTunnelPorts(c *cli.Context) *TunnelPorts {
	local := c.String("local-port")
	_, err := strconv.Atoi(local)
	if err != nil {
		logger.Fatalf("%s is not a valid port number", local)
	}
	remote := c.String("remote-port")
	_, err = strconv.Atoi(remote)
	if err != nil {
		logger.Fatalf("%s is not a valid port number", remote)
	}
	return &TunnelPorts{local, remote}
}

func (t *TunnelPorts) RemoteConnectionString() string {
	return fmt.Sprintf("0.0.0.0:%s", t.remote)
}

func (t *TunnelPorts) LocalConnectionString() string {
	return fmt.Sprintf("localhost:%s", t.local)
}
