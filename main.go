package main

import (
	"log"

	"github.com/ukibbb/tcp-server/tcp"
)

func main() {
	opts := tcp.TCPListenerOpts{
		ListenAddr: ":6379",
	}
	server := tcp.NewTCPListener(opts)

	if err := server.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
