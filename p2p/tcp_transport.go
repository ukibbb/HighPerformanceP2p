package p2p

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

// Peer represents remote node over established connection
// in this case TCP
type TCPPeer struct {
	// Underlying connection of the peer in this case tcp connection.
	net.Conn

	// if we dial and receive connection => true
	// if we listen and receive connection => false
	outbound bool

	wg *sync.WaitGroup
}

type RPCErr struct {
	net.Conn
	err error
}

func NewRPCErr(conn net.Conn, err error) *RPCErr {
	return &RPCErr{
		Conn: conn,
		err:  err,
	}
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn:     conn,
		outbound: outbound,
		wg:       &sync.WaitGroup{},
	}
}

func (p *TCPPeer) Send(data []byte) error {
	_, err := p.Conn.Write(data)
	return err
}

func (p *TCPPeer) CloseStream() {
	p.wg.Done()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

// TCPTransport will implement Transport interface
type TCPTransport struct {
	TCPTransportOpts
	listener  net.Listener
	rpcch     chan RPC
	connerrch chan *RPCErr
}

// TCP server contructor
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC, 1024),
		connerrch:        make(chan *RPCErr),
	}
}

// Addr implements the Transport interface return the address
// the transport is accepting connections.
func (t *TCPTransport) Addr() string {
	return t.ListenAddr
}

// Consume implements the Tranport interface, which will return read-only channel
// for reading the incoming messages received from another peer in the network.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ConsumeError() <-chan *RPCErr {
	return t.connerrch
}

// Close implements the Transport interface.
func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

// Dial implements the Transport interface.
func (t *TCPTransport) Dial(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	go t.handleConn(conn, true)

	return nil
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)

	if err != nil {
		return fmt.Errorf(
			"[ERROR]: Failed to listen on %s, error: %w",
			t.ListenAddr,
			err,
		)
	}

	go t.startAcceptLoop()

	log.Printf(
		"[INFO]: Server is running on `%s`",
		t.listener.Addr(),
	)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return
		}
		if err != nil {
			log.Printf("[WARNING]: Failed to accept connection `startAcceptLoop()` %s\n", err)
		}
		go func() {
			err = t.handleConn(conn, false)
			if err != nil {
				t.connerrch <- NewRPCErr(conn, err)
			}
		}()

	}
}

func (t *TCPTransport) handleConn(conn net.Conn, outbound bool) error {
	var err error

	peer := NewTCPPeer(conn, outbound)

	if err = t.HandshakeFunc(peer); err != nil {
		return err
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return err
		}
	}

	// Read loop from connection
	for {
		rpc := RPC{}
		err = t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return err
		}
		rpc.From = conn.RemoteAddr().String()

		if rpc.Stream {
			peer.wg.Add(1)
			fmt.Printf("[%s] incoming stream, waiting...\n", conn.RemoteAddr())
			peer.wg.Wait()
			fmt.Printf("[%s] stream closed, resuming read loop\n", conn.RemoteAddr())
			continue
		}
		t.rpcch <- rpc
	}
}
