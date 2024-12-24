package transport

import "net"

type Transport struct {
	Protocol, Host, Port string
}

func (t *Transport) GetProtocol() string {
	return t.Protocol
}

func (t *Transport) GetAddress() string {
	return net.JoinHostPort(t.Host, t.Port)
}
