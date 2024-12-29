package main

import (
	"fmt"

	"github.com/ukibbb/HighPerformanceP2p/v2/p2p"
)

func main() {
	fmt.Print("Main function\n")
	opts := p2p.TCPTransportOpts{
		ListenAddr:    ":6379",
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	p2p.NewTCPTransport(opts)

	select {}
}
