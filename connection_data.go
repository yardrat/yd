package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strconv"
)

type ConnectionData struct {
	user string
	key  string
	host string
	port string
}

func require(flagName string, c *cli.Context) {
	value := c.String(flagName)
	if len(value) <= 0 {
		logger.Fatalf("flag '%s' is missing", flagName)
	}
}

func ReadConnectionData(c *cli.Context) *ConnectionData {
	require("ssh-key", c)
	require("ssh-host", c)

	key := c.String("ssh-key")
	user := c.String("ssh-user")
	host := c.String("ssh-host")
	port := c.String("ssh-port")
	_, err := strconv.Atoi(port)

	if err != nil {
		logger.Fatalf("%s is not a valid port number", port)
	}

	return &ConnectionData{user, key, host, port}
}

func (data *ConnectionData) String() string {
	return fmt.Sprintf("%s:%s", data.host, data.port)
}
