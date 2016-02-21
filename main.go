package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

var (
	SshKeyFlag       = cli.StringFlag{Name: "ssh-key, i", Usage: "your private ssh key", EnvVar: "YD_SSH_KEY"}
	SshHostFlag      = cli.StringFlag{Name: "ssh-host, H", Usage: "target server hostname or ip", EnvVar: "YD_SSH_HOST"}
	SshPortFlag      = cli.StringFlag{Name: "ssh-port, P", Value: "22", Usage: "target server port", EnvVar: "YD_SSH_PORT"}
	SshUserFlag      = cli.StringFlag{Name: "ssh-user, U", Value: "ubuntu", Usage: "ssh user name", EnvVar: "YD_SSH_USER"}
	TunnelLocalPort  = cli.StringFlag{Name: "local-port, l", Value: "8080", Usage: "tunnel local port"}
	TunnelRemotePort = cli.StringFlag{Name: "remote-port, r", Value: "8080", Usage: "tunnel remote port"}
)

func Ping(c *cli.Context) {
	data := ReadConnectionData(c)

	if err := DefaultSsh.Ping(data); err != nil {
		logger.Fatalf("error while connecting to %v'\n%v", c, err)
	} else {
		logger.Printf("ping successful.")
	}
}

func Connect(c *cli.Context) {
	data := ReadConnectionData(c)
	tunnel := ReadTunnelPorts(c)

	// establish the tunnel and block.
	DefaultSsh.Tunnel(data, tunnel)
}

func main() {
	app := cli.NewApp()
	app.Name = "yd"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:    "ping",
			Aliases: []string{"g"},
			Usage:   "checks that the remote connection is working properly",
			Flags:   []cli.Flag{SshKeyFlag, SshHostFlag, SshPortFlag, SshUserFlag},
			Action:  Ping,
		},
		{
			Name:   "connect",
			Usage:  "connects to the remote host ",
			Flags:  []cli.Flag{SshKeyFlag, SshHostFlag, SshPortFlag, SshUserFlag, TunnelLocalPort, TunnelRemotePort},
			Action: Connect,
		},
	}
	app.Run(os.Args)
}
