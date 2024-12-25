package tcp

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

type MockNetAddr struct{}

func (MockNetAddr) Network() string {
	return "tcp"
}
func (MockNetAddr) String() string {
	return "192.0.0.1:6379"
}

type MockNetConn struct{}

func (MockNetConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (MockNetConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (MockNetConn) Close() error { return nil }

func (MockNetConn) LocalAddr() net.Addr { return &MockNetAddr{} }

func (MockNetConn) RemoteAddr() net.Addr { return &MockNetAddr{} }

func (MockNetConn) SetDeadline(t time.Time) error { return nil }

func (MockNetConn) SetReadDeadline(t time.Time) error { return nil }

func (MockNetConn) SetWriteDeadline(t time.Time) error { return nil }

type MockNetListener struct{}

func (MockNetListener) Accept() (net.Conn, error) {
	log.Println("[INFO]: Coming from MockNetListener.Accept()")
	time.Sleep(time.Second * 3)
	return &MockNetConn{}, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (MockNetListener) Close() error {
	log.Println("[INFO]: Coming from MockNetListener.Close()")
	return nil
}

// Addr returns the listener's network address.
func (MockNetListener) Addr() net.Addr { return &MockNetAddr{} }

func NewMockListener() net.Listener {
	return &MockNetListener{}
}

func TestListenerStart(t *testing.T) {
	tests := struct {
		name        string
		host        string
		protocol    string
		port        string
		shouldError bool
	}{}
	fmt.Println(tests)
}

func TestListenerAcceptAndHandle(t *testing.T) {
	// mockListener := NewMockListener().(*MockNetListener)
	// listener := Listener{
	// 	Transport: &transport.Transport{
	// 		Host:     "localhost",
	// 		Port:     "6379",
	// 		Protocol: "tcp",
	// 	},
	// 	listener: mockListener,
	// }

	// listener.Start()
	// listener.AcceptAndHandle()
}
