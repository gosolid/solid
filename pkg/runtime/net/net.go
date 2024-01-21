//js:package net
package net

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"errors"
	"fmt"

	"github.com/grexie/isolates"
)

var ErrInvalidNet = errors.New("receiver does not implement net.Net")

var _ Net = &NetBase{}

type Net interface {
	Server() *isolates.Value
	Socket() *isolates.Value
	Connect(ctx context.Context, network string, address string) (*isolates.Value, error)
	Listen(ctx context.Context, server Server, network string, address string) (*isolates.Value, error)
}

type NetBase struct {
	This    *isolates.Value
	connect *isolates.Value
	listen  *isolates.Value
	server  *isolates.Value
	socket  *isolates.Value
}

//js:constructor Net
//js:export Net
func NewNet(in isolates.FunctionArgs) (*NetBase, error) {
	options := in.Arg(in.ExecutionContext, 0)
	var connect *isolates.Value
	var listen *isolates.Value
	var err error

	if connect, err = options.Get(in.ExecutionContext, "connect"); err != nil {
		return nil, err
	}

	if listen, err = options.Get(in.ExecutionContext, "listen"); err != nil {
		return nil, err
	}

	net := &NetBase{
		connect: connect,
		listen:  listen,
	}

	if Server, err := in.Context.CreateConstructorWithName(in.ExecutionContext, "Server", func(in isolates.FunctionArgs) (*ServerBase, error) {
		return newServer(in, net)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "Server", Server); err != nil {
		return nil, err
	} else if createServer, err := in.Context.CreateWithName(in.ExecutionContext, "createServer", func(in isolates.FunctionArgs) (*isolates.Value, error) {
		return Server.NewValue(in.ExecutionContext, in.Args...)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "createServer", createServer); err != nil {
		return nil, err
	} else {
		net.server = Server
	}

	if Socket, err := in.Context.CreateConstructorWithName(in.ExecutionContext, "Socket", func(in isolates.FunctionArgs) (*SocketBase, error) {
		return newSocket(in, net)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "Socket", Socket); err != nil {
		return nil, err
	} else if createConnection, err := in.Context.CreateWithName(in.ExecutionContext, "createConnection", func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if port, err := in.Arg(in.ExecutionContext, 0).Int64(in.ExecutionContext); err != nil {
			return nil, err
		} else if host, err := in.Arg(in.ExecutionContext, 1).StringValue(in.ExecutionContext); err != nil {
			return nil, err
		} else {
			address := fmt.Sprintf("%s:%d", host, port)
			return net.connect.Call(in.ExecutionContext, net, "tcp", address)
		}
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "createConnection", createConnection); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "connect", createConnection); err != nil {
		return nil, err
	} else {
		net.socket = Socket
	}

	return net, nil
}

func (n *NetBase) Server() *isolates.Value {
	return n.server
}

func (n *NetBase) Socket() *isolates.Value {
	return n.socket
}

func (n *NetBase) Listen(ctx context.Context, server Server, network, address string) (*isolates.Value, error) {
	return n.listen.Call(ctx, server, network, address)
}

func (n *NetBase) Connect(ctx context.Context, network, address string) (*isolates.Value, error) {
	return n.connect.Call(ctx, n, network, address)
}
