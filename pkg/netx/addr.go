package netx

import "fmt"

type Addr struct {
	host string
	port int
}

// String returns host:port.
func (a Addr) String() string { return fmt.Sprintf("%s:%d", a.host, a.port) }

// NewAddr creates new Addr.
func NewAddr(host string, port int) Addr { return Addr{host: host, port: port} }

// Host returns host.
func (a Addr) Host() string { return a.host }

// Port returns port.
func (a Addr) Port() int { return a.port }
