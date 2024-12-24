package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var _ = net.Listen
var _ = os.Exit

type Listener struct {
	protocol, host, port string
	listener             net.Listener
}

func (s *Listener) Start() error {
	if s.listener != nil {
		log.Printf(
			"[INFO]: Listener already provided %v created in memory address %v",
			s.listener,
			&s.listener,
		)
		return nil
	}
	l, err := net.Listen(s.getProtocol(), s.getAddress())
	if err != nil {
		return fmt.Errorf(
			"[ERROR]: Failed to bind to %s:%s, error: %v",
			s.host,
			s.port,
			err,
		)
	}
	s.listener = l
	log.Printf(
		"[INFO]: Listener %v created in memory address %v",
		s.listener,
		&s.listener,
	)
	return nil
}

func (s *Listener) AcceptAndHandle() error {
	log.Printf(
		"[INFO]: Listener %v created in memory address %v",
		s.listener,
		&s.listener,
	)
	defer s.listener.Close()
	for {
		connection, err := s.listener.Accept()
		log.Printf(
			"[INFO]: Connection %v",
			&connection,
		)

		if err != nil {

		}
		go s.Handle(&connection)
	}
}

func (s *Listener) Handle(connection *net.Conn) {
	log.Printf(
		"[INFO]: Handling connection %v",
		connection,
	)
	defer (*connection).Close()
	(*connection).SetDeadline(time.Now().Add(time.Second * 3))

}

func (s *Listener) getProtocol() string {
	return s.protocol
}

func (s *Listener) getAddress() string {
	return net.JoinHostPort(s.host, s.port)
}

func main() {
	redis := Listener{host: "localhost", port: "6379", protocol: "tcp"}
	err := redis.Start()
	if err != nil {
	}
	err = redis.AcceptAndHandle()
	if err != nil {
	}
}
