# Golang P2P Transport

A simple and extensible implementation of a peer-to-peer (P2P) transport layer using TCP in Go. This library provides core features for managing connections, sending and receiving data, and implementing custom handshake protocols.

## Features

- TCP-based P2P transport
- Simple interface for handling connections
- Customizable handshake functionality
- Efficient message decoding using pluggable decoders
- Concurrency-safe design with `sync.WaitGroup`

## Installation

To use this library in your Go project, install it via `go get`:

```bash
go get github.com/yourusername/p2p
```

Replace `yourusername` with your GitHub username if hosted there.

## Usage

### Create a TCP Transport

You can create a TCP transport using the `NewTCPTransport` constructor. Pass in options to configure the listening address, handshake function, decoder, and peer connection handler.

```go
package main

import (
    "log"
    "github.com/yourusername/p2p"
)

func main() {
    opts := p2p.TCPTransportOpts{
        ListenAddr: ":8080",
        HandshakeFunc: func(peer p2p.Peer) error {
            log.Printf("New peer connected: %s", peer.Conn.RemoteAddr())
            return nil
        },
        Decoder: myCustomDecoder{}, // Implement the `Decoder` interface
        OnPeer: func(peer p2p.Peer) error {
            log.Printf("Peer ready: %s", peer.Conn.RemoteAddr())
            return nil
        },
    }

    transport := p2p.NewTCPTransport(opts)

    if err := transport.ListenAndAccept(); err != nil {
        log.Fatalf("Failed to start transport: %v", err)
    }
}
```

### Sending Data

To send data to a peer, use the `Send` method of the `TCPPeer` struct:

```go
peer.Send([]byte("Hello, Peer!"))
```

### Receiving Data

Consume incoming RPC messages from the transport's `Consume` method:

```go
for rpc := range transport.Consume() {
    log.Printf("Received message: %+v from %s", rpc.Data, rpc.From)
}
```

## Interfaces

### `Peer`

Represents a connected peer. The `TCPPeer` struct implements this interface.

```go
type Peer interface {
    Send(data []byte) error
    Conn() net.Conn
}
```

### `Decoder`

Defines the interface for decoding incoming messages.

```go
type Decoder interface {
    Decode(conn net.Conn, out *RPC) error
}
```

## Example

Here’s a complete example of creating a TCP transport and handling peer connections:

```go
package main

import (
    "log"
    "github.com/yourusername/p2p"
)

func main() {
    opts := p2p.TCPTransportOpts{
        ListenAddr: ":9000",
        HandshakeFunc: func(peer p2p.Peer) error {
            log.Printf("Handshake complete with peer: %s", peer.Conn().RemoteAddr())
            return nil
        },
        Decoder: MyDecoder{},
        OnPeer: func(peer p2p.Peer) error {
            log.Printf("Peer connected: %s", peer.Conn().RemoteAddr())
            return nil
        },
    }

    transport := p2p.NewTCPTransport(opts)

    if err := transport.ListenAndAccept(); err != nil {
        log.Fatalf("Error starting transport: %v", err)
    }

    for rpc := range transport.Consume() {
        log.Printf("Message received: %s", string(rpc.Data))
    }
}

// MyDecoder implements the Decoder interface
struct MyDecoder {}

func (d MyDecoder) Decode(conn net.Conn, out *p2p.RPC) error {
    // Implement your decoding logic here
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
