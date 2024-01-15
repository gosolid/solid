//js:package net

package net

//go:generate go run github.com/grexie/isolates/codegen

import (
	"errors"
	"fmt"
	"net"
	"reflect"

	"github.com/grexie/isolates"
)

type Address struct {
	family  string
	address string
	port    int
}

//js:constructor SocketAddress
func NewAddress(in isolates.FunctionArgs) (*Address, error) {
	var addr net.Addr
	var ok bool

	if len(in.Args) < 1 {
		return nil, errors.New("SocketAddress: requires constructor options")
	} else if rv, err := in.Args[0].Unmarshal(in.ExecutionContext, reflect.TypeOf(&struct{}{})); err == nil && !rv.IsNil() {
		if addr, ok = rv.Interface().(net.Addr); ok {

		} else if addr, ok := rv.Interface().(*Address); ok {
			return &Address{
				family:  addr.Family(),
				address: addr.Address(),
				port:    addr.Port(),
			}, nil
		} else {
			return nil, errors.New("invalid type")
		}
	}

	tcpAddr := addr.(*net.TCPAddr)

	if ip := tcpAddr.IP.To4(); ip != nil {
		return &Address{
			family:  "ipv4",
			address: ip.String(),
			port:    tcpAddr.Port,
		}, nil
	} else if ip := tcpAddr.IP.To16(); ip != nil {
		return &Address{
			family:  "ipv6",
			address: ip.String(),
			port:    tcpAddr.Port,
		}, nil
	} else {
		return nil, fmt.Errorf("invalid address: %s", addr)
	}
}

//js:get family
func (a *Address) Family() string {
	return a.family
}

//js:get address
func (a *Address) Address() string {
	return a.address
}

//js:get port
func (a *Address) Port() int {
	return a.port
}
