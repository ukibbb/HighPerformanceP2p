package p2p

import (
	"fmt"
	"log"
	"net"
	"time"
)

type TCPListenerOpts struct {
	ListenAddr string
}

type TCPListener struct {
	TCPListenerOpts
	listener net.Listener

	// peers map[net.Addr]Peer
}

// TCP server contructor
func NewTCPListener(opts TCPListenerOpts) *TCPListener {
	return &TCPListener{
		TCPListenerOpts: opts,
	}
}

func (s *TCPListener) ListenAndAccept() error {
	if s.listener != nil {
		log.Printf(
			"[INFO]: Listener already provided %v created in memory address %v",
			s.listener,
			&s.listener,
		)
		return nil
	}
	l, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return fmt.Errorf(
			"[ERROR]: Failed to bind to %s, error: %v",
			s.listener.Addr(),
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
		"[INFO]: Server is running on `%s`",
		s.listener.Addr(),
	)

	go s.AcceptAndHandle()

	return nil
}

func (s *TCPListener) AcceptAndHandle() error {
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

func (s *TCPListener) Handle(connection net.Conn) {
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
