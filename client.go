package main

import (
	"log"
	"net"
)

type Sender struct {
	*Transport
}

func (s *Sender) CreateConnection() {
	conn, _ := net.Dial(s.GetProtocol(), s.GetAddress())
	log.Printf("[INFO]: Connection %v to address %s created.", conn, s.GetAddress())

}
