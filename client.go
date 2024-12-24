package main

import (
	"github/ukibbb/tcp-server/transport"
	"log"
	"net"
)

type Sender struct {
	*transport.Transport
	conn net.Conn
}

func (s *Sender) CreateConnection() {
	conn, _ := net.Dial(s.Transport.GetProtocol(), s.Transport.GetAddress())
	log.Printf(
		"[INFO]: Connection %v to address %s created.",
		conn,
		s.Transport.GetAddress(),
	)
	s.conn = conn
}

func (s *Sender) Send(data []byte) {
	if s.conn == nil {
		log.Fatalln("[ERROR]: Connection is not established...")
	}
	defer s.conn.Close()
	s.conn.Write(data)
}
func main() {
	sender := &Sender{
		Transport: &transport.Transport{
			Host:     "0.0.0.0",
			Port:     "6379",
			Protocol: "tcp",
		},
	}
	sender.CreateConnection()
	sender.Send([]byte("Hello World"))
}
