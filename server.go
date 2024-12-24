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
	*Transport
	listener net.Listener
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
	l, err := net.Listen(s.GetProtocol(), s.GetAddress())
	if err != nil {
		return fmt.Errorf(
			"[ERROR]: Failed to bind to %s:%s, error: %v",
			s.Transport.Host,
			s.Transport.Port,
			err,
		)
	}
	s.listener = l
	log.Printf(
		"[INFO]: Listener.Start() %v created in memory address %v",
		s.listener,
		&s.listener,
	)
	log.Printf(
		"[INFO]: Server is running on host `%s` port `%s`",
		s.Transport.Host,
		s.Transport.Port,
	)
	return nil
}

func (s *Listener) AcceptAndHandle() error {
	log.Printf(
		"[INFO]: Listener.AcceptAndHandle() %v created in memory address %v",
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

func main() {
	server := Listener{Transport: &Transport{
		Host:     "0.0.0.0",
		Port:     "6379",
		Protocol: "tcp",
	}}
	err := server.Start()
	if err != nil {
	}
	err = server.AcceptAndHandle()
	if err != nil {
	}
}
