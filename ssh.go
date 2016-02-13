package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"strconv"
)

var DefaultSsh = &Ssh{}

type Ssh struct{}

func (s *Ssh) Connect(keypath, username, hostname, port string) (*ssh.Client, error) {
	_, err := strconv.Atoi(port)
	if err != nil {
		logger.Fatalf("%s is not a valid port number", port)
	}

	buffer, err := ioutil.ReadFile(keypath)
	if err != nil {
		logger.Fatalf("error while opening file %v", err)
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		logger.Panicf("error while parsing key %v", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	connectionString := fmt.Sprintf("%s:%s", hostname, port)
	return ssh.Dial("tcp", connectionString, config)
}
