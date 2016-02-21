package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net"
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

func (s *Ssh) Tunnel(data *ConnectionData, tunnel *TunnelPorts) error {
	logger.Printf("tunnelling localhost:%s to %s:%s...", tunnel.local, data.host, tunnel.remote)

	client, err := DefaultSsh.Connect(data)
	if err != nil {
		logger.Fatalf("error while connecting to %s\n%v", data.String(), err)
	}

	listener, err := client.Listen("tcp", tunnel.RemoteConnectionString())
	if err != nil {
		logger.Fatalf("error while opening port %s on remote host\n%v", tunnel.remote, err)
	}

	logger.Printf("tunnel established (CTRL+C to quit)")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatal(err)
		}
		go handleConnection(conn, tunnel)
	}
}

func handleConnection(conn net.Conn, tunnel *TunnelPorts) {
	local, err := net.Dial("tcp", tunnel.LocalConnectionString())
	if err != nil {
		logger.Fatalf("error while connecting to localhost:%s\n%v", tunnel.local, err)
	}

	go copyConnection(local, conn)
	go copyConnectionAndClose(conn, local)
}

func copyConnection(writer, reader net.Conn) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		logger.Fatalf("io.Copy error: %v", err)
	}
}

func copyConnectionAndClose(writer, reader net.Conn) {
	copyConnection(writer, reader)
	writer.Close()
}
