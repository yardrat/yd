package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

var (
	SshKeyFlag  = cli.StringFlag{Name: "ssh-key, i", Usage: "your private ssh key", EnvVar: "YD_SSH_KEY"}
	SshHostFlag = cli.StringFlag{Name: "ssh-host, H", Usage: "target server hostname or ip", EnvVar: "YD_SSH_HOST"}
	SshPortFlag = cli.StringFlag{Name: "ssh-port, P", Value: "22", Usage: "target server port", EnvVar: "YD_SSH_PORT"}
	SshUserFlag = cli.StringFlag{Name: "ssh-user, U", Value: "ubuntu", Usage: "ssh user name", EnvVar: "YD_SSH_USER"}
)

func require(flagName string, c *cli.Context) {
	value := c.String(flagName)
	if len(value) <= 0 {
		logger.Fatalf("flag '%s' is missing", flagName)
	}
}

func Ping(c *cli.Context) {
	require("ssh-key", c)
	require("ssh-host", c)

	key := c.String("ssh-key")
	user := c.String("ssh-user")
	host := c.String("ssh-host")
	port := c.String("ssh-port")

	if _, err := DefaultSsh.Connect(key, user, host, port); err != nil {
		logger.Fatalf("error while connecting to %s@'%s:%s'\n%v", user, host, port, err)
	} else {
		logger.Printf("OK")
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "yd"
	app.Commands = []cli.Command{
		{
			Name:    "ping",
			Aliases: []string{"g"},
			Usage:   "checks that the remote connection is working properly",
			Flags:   []cli.Flag{SshKeyFlag, SshHostFlag, SshPortFlag, SshUserFlag},
			Action:  Ping,
		},
	}
	app.Run(os.Args)
}
