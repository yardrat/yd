package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var logger = log.New(os.Stdout, "", 0)

var (
	SshKeyFlag = cli.StringFlag{
		Name:   "ssh-key, i",
		Usage:  "your private ssh key",
		EnvVar: "YD_SSH_KEY",
	}
	SshHostFlag = cli.StringFlag{
		Name:   "ssh-host, H",
		Usage:  "target server hostname or ip",
		EnvVar: "YD_SSH_HOST",
	}
	SshPortFlag = cli.StringFlag{
		Name:   "ssh-port, P",
		Value:  "22",
		Usage:  "target server port",
		EnvVar: "YD_SSH_PORT",
	}
	SshUserFlag = cli.StringFlag{
		Name:   "ssh-user, U",
		Value:  "ubuntu",
		Usage:  "ssh user name",
		EnvVar: "YD_SSH_USER",
	}
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

	keyPath := c.String("ssh-key")

	buffer, err := ioutil.ReadFile(keyPath)
	if err != nil {
		logger.Fatalf("error while opening file %v", err)
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		logger.Panicf("error while parsing key %v", err)
	}
	user := c.String("ssh-user")
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}

	host := c.String("ssh-host")
	port := c.String("ssh-port")
	_, err = strconv.Atoi(port)
	if err != nil {
		logger.Fatalf("%s is not a valid port number", port)
	}

	connectionString := fmt.Sprintf("%s:%s", host, port)

	_, err = ssh.Dial("tcp", connectionString, config)
	if err != nil {
		logger.Fatalf("error while connecting to '%s'\n%v", connectionString, err)
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
