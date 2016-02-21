package main

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

var DefaultSsh = &Ssh{}

type Ssh struct{}

func (s *Ssh) Connect(data *ConnectionData) (*ssh.Client, error) {
	buffer, err := ioutil.ReadFile(data.key)
	if err != nil {
		logger.Fatalf("error while opening file %v", err)
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		logger.Panicf("error while parsing key %v", err)
	}

	config := &ssh.ClientConfig{
		User: data.user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	client, err := ssh.Dial("tcp", data.String(), config)
	return client, err
}

func (s *Ssh) Ping(data *ConnectionData) error {
	_, err := s.Connect(data)
	return err
}
