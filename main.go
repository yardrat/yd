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

func Ping(c *cli.Context) {
	data := ReadConnectionData(c)

	if err := DefaultSsh.Ping(data); err != nil {
		logger.Fatalf("error while connecting to %v'\n%v", c, err)
	} else {
		logger.Printf("OK")
	}
}

func Connect(c *cli.Context) {
	data := ReadConnectionData(c)
	logger.Printf("connected to %s", data)
}

func main() {
	app := cli.NewApp()
	app.Name = "yd"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{SshKeyFlag, SshHostFlag, SshPortFlag, SshUserFlag}
	app.Commands = []cli.Command{
		{
			Name:    "ping",
			Aliases: []string{"g"},
			Usage:   "checks that the remote connection is working properly",
			Flags:   []cli.Flag{SshKeyFlag, SshHostFlag, SshPortFlag, SshUserFlag},
			Action:  Ping,
		},
		{
			Name:    "connect",
			Aliases: []string{"add"},
			Usage:   "connects to the remote host ",
			Flags:   []cli.Flag{SshKeyFlag, SshHostFlag, SshPortFlag, SshUserFlag},
			Action:  Connect,
		},
	}
	app.Run(os.Args)
}
