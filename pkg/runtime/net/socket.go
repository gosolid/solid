//js:package net
package net

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"

	"github.com/gosolid/solid/pkg/runtime/stream"
	isolates "github.com/grexie/isolates"
)

//js:alias SocketBase
type Socket interface {
	stream.Duplex
}

//js:class Socket
type SocketBase struct {
	*stream.DuplexBase
}

func newSocket(in isolates.FunctionArgs, net *NetBase) (*SocketBase, error) {
	if args, err := in.WithArgs(); err != nil {
		return nil, err
	} else if DuplexBase, err := stream.NewDuplex(args); err != nil {
		return nil, err
	} else {
		socket := &SocketBase{DuplexBase: DuplexBase}
		return socket, nil
	}
}

func (s *SocketBase) localAddress(ctx context.Context) (*Address, error) {
	if localAddress, err := s.Options.Get(ctx, "localAddress"); err != nil {
		return nil, err
	} else if addr, err := localAddress.Call(ctx, s, "localAddress"); err != nil {
		return nil, err
	} else if addr, err := isolates.For(ctx).New(NewAddress, addr); err != nil {
		return nil, err
	} else {
		return addr.(*Address), nil
	}
}

func (s *SocketBase) remoteAddress(ctx context.Context) (*Address, error) {
	if localAddress, err := s.Options.Get(ctx, "remoteAddress"); err != nil {
		return nil, err
	} else if addr, err := localAddress.Call(ctx, s, "remoteAddress"); err != nil {
		return nil, err
	} else if addr, err := isolates.For(ctx).New(NewAddress, addr); err != nil {
		return nil, err
	} else {
		return addr.(*Address), nil
	}
}

//js:method
func (s *SocketBase) Address(ctx context.Context) (*Address, error) {
	return s.localAddress(ctx)
}

//js:get
func (s *SocketBase) LocalAddress(ctx context.Context) (*string, error) {
	if address, err := s.localAddress(ctx); err != nil {
		return nil, nil
	} else {
		return &address.address, nil
	}
}

//js:get
func (s *SocketBase) LocalPort(ctx context.Context) (*int, error) {
	if address, err := s.localAddress(ctx); err != nil {
		return nil, nil
	} else {
		return &address.port, nil
	}
}

//js:get
func (s *SocketBase) LocalFamily(ctx context.Context) (*string, error) {
	if address, err := s.localAddress(ctx); err != nil {
		return nil, nil
	} else {
		return &address.family, nil
	}
}

//js:get
func (s *SocketBase) RemoteAddress(ctx context.Context) (*string, error) {
	if address, err := s.remoteAddress(ctx); err != nil {
		return nil, nil
	} else {
		return &address.address, nil
	}
}

//js:get
func (s *SocketBase) RemotePort(ctx context.Context) (*int, error) {
	if address, err := s.remoteAddress(ctx); err != nil {
		return nil, nil
	} else {
		return &address.port, nil
	}
}

//js:get
func (s *SocketBase) RemoteFamily(ctx context.Context) (*string, error) {
	if address, err := s.remoteAddress(ctx); err != nil {
		return nil, nil
	} else {
		return &address.family, nil
	}
}

//js:method
func (s *SocketBase) SetTimeout(ctx context.Context) error {
	return nil
}

//js:method
func (s *SocketBase) SetNoDelay(ctx context.Context) error {
	return nil
}
