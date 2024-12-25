package tcp

import (
	"fmt"
	"github/ukibbb/tcp-server/transport"
	"log"
	"net"
	"time"
)

type Listener struct {
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

	go s.AcceptAndHandle()

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
			"[INFO]: Connection `AcceptAndHandle()` %v",
			connection,
		)

		if err != nil {

		}
		go s.Handle(connection)
	}
}

func (s *Listener) Handle(connection net.Conn) {
	log.Printf(
		"[INFO]: Handling connection `Handle()` %v\n",
		connection,
	)
	defer func() {
		log.Printf("[INFO]: Dropping connection %v\n", connection)

		defer (connection).Close()
	}()
	(connection).SetDeadline(time.Now().Add(time.Second * 3))

}

func main() {
	server := Listener{Transport: &transport.Transport{
		Host:     "0.0.0.0",
		Port:     "6379",
		Protocol: "tcp",
	}}
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	select {}
}
