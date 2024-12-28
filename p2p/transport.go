package p2p

import "net"

// Peer interface represents remot node.
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport is everyting that hold communication between nodes
// in the network,
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Close() error
	Consume() <-chan RPC
	ConsumeError() <-chan RPCErr
}
