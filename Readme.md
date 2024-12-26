# HighPerformanceP2P TCP Transport

## Overview

HighPerformanceP2P is a high-performance, peer-to-peer (P2P) TCP transport library written in Go. It is designed for efficient and reliable communication between peers in a distributed network. The library is optimized for low latency, high throughput, and scalability, making it suitable for real-time applications, distributed systems, and decentralized networks.

## Features

- **Efficient Connection Handling**: Supports multiplexed TCP connections with minimal overhead.
- **Peer Discovery**: Includes mechanisms for discovering and managing peers in a dynamic network.
- **Reliability**: Built-in support for retrying, error handling, and congestion control.
- **Security**: Optional TLS encryption to secure communication.
- **Custom Protocol Support**: Easy to integrate with custom application-layer protocols.
- **Scalability**: Handles thousands of concurrent connections efficiently.

## Use Cases

- Real-time messaging systems
- Decentralized applications (DApps)
- Distributed databases
- Multiplayer games
- IoT communication networks

## Installation

To install HighPerformanceP2P, use:

```bash
go get github.com/ukibbb/highperformancep2p
```

## Usage

### Setting Up a Peer

Here is a basic example of setting up a peer and connecting to another peer:

```go
package main

import (
	"fmt"
	"log"
	"github.com/yourusername/highperformancep2p"
)

func main() {
	// Initialize a new peer
	peer := highperformancep2p.NewPeer("0.0.0.0:9000")

	// Register a message handler
	peer.OnMessage(func(msg p2p.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})

	// Start the peer
	if err := peer.Start(); err != nil {
		log.Fatalf("Failed to start peer: %v", err)
	}

	// Connect to another peer
	if err := peer.Connect("192.168.1.100:9001"); err != nil {
		log.Fatalf("Failed to connect to peer: %v", err)
	}

	// Send a message to the connected peer
	peer.Send(p2p.Message{
		To:   "192.168.1.100:9001",
		Data: []byte("Hello, Peer!"),
	})
}
```

### Configuration

You can configure the peer with custom options:

```go
peer := p2p.NewPeer("0.0.0.0:9000", p2p.Options{
	TLSConfig:    yourTLSConfig, // For secure communication
	MaxRetries:   5,             // Retry count for failed connections
	BufferSize:   1024 * 64,     // Buffer size for messages
})
```

## API Reference

### Peer Methods

- `NewPeer(address string, options ...Options) *Peer`: Creates a new peer instance.
- `Start() error`: Starts the peer, allowing it to accept connections.
- `Connect(address string) error`: Connects to another peer.
- `Send(msg Message) error`: Sends a message to a specific peer.
- `OnMessage(handler func(msg Message))`: Registers a handler for incoming messages.
- `Close() error`: Stops the peer and closes all connections.

### Message Structure

```go
type Message struct {
	To   string
	Data []byte
}
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. Make sure to include tests and documentation for any new features or bug fixes.

## License

HighPerformanceP2P is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go's `net` package for efficient networking.
- Inspired by existing P2P protocols like BitTorrent and libp2p.

## Contact

For questions, suggestion
