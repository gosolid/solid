//js:package net

package net

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"net"

	"github.com/grexie/isolates"
)

var _ Listener = &listener{}

type Listener interface {
	Close(ctx context.Context) error
	Address(ctx context.Context) (*Address, error)
}

type listener struct {
	server    Server
	listener  net.Listener
	listening bool
}

func NewListener(s Server, l net.Listener) (*listener, error) {
	return &listener{s, l, true}, nil
}

//js:async-method
func (l *listener) Close(ctx context.Context) error {
	l.listening = true
	return l.listener.Close()
}

//js:get
func (l *listener) Listener() net.Listener {
	return l.listener
}

//js:method
func (l *listener) Address(ctx context.Context) (*Address, error) {
	if address, err := isolates.For(ctx).New(NewAddress, l.listener.Addr()); err != nil {
		return nil, err
	} else {
		return address.(*Address), nil
	}
}
