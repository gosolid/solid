//js:package net

package net

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"errors"
	gonet "net"
	"reflect"
	"sync"

	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/grexie/isolates"
)

type HostNetOptions struct {
	__hostNetOptions bool
}

type HostSocketOptions struct {
	*stream.ReadWriteCloser
	conn   gonet.Conn
	b      []byte
	buffer *isolates.Value
}

type DeferredTCPConn struct {
	sync.Mutex
	*gonet.TCPConn
	closed bool
	unlock chan bool
}

type ITCPConn interface {
	CloseWrite() error
}

func (c *DeferredTCPConn) SetTCPConn(t *gonet.TCPConn) error {
	c.Lock()
	defer c.Unlock()
	if c.TCPConn != nil {
		return errors.New("TCPConn already bridged")
	}
	c.TCPConn = t
	c.unlock <- true
	return nil
}

func (c *DeferredTCPConn) Close() error {
	c.Lock()
	defer c.Unlock()
	c.closed = true
	if c.TCPConn != nil {
		return c.TCPConn.Close()
	}
	return nil
}

func (c *DeferredTCPConn) CloseWrite() error {
	c.Lock()
	if c.TCPConn != nil {
		c.Unlock()
		return c.TCPConn.CloseWrite()
	} else {
		c.Unlock()
		return nil
	}
}

func (c *DeferredTCPConn) Wait() {
	c.Lock()
	if c.TCPConn == nil {
		c.Unlock()
		<-c.unlock
	} else {
		c.Unlock()
	}
}

func (c *DeferredTCPConn) Read(b []byte) (int, error) {
	c.Wait()
	return c.TCPConn.Read(b)
}

func (c *DeferredTCPConn) Write(b []byte) (int, error) {
	c.Wait()
	return c.TCPConn.Write(b)
}

//js:method connect
func (n *HostNetOptions) Connect(in isolates.FunctionArgs, net Net, network string, address string) (Socket, error) {
	dconn := &DeferredTCPConn{unlock: make(chan bool)}

	if options, err := isolates.For(in.ExecutionContext).New(NewHostSocketOptions, dconn); err != nil {
		return nil, err
	} else if socketv, err := net.Socket().New(in.ExecutionContext, options); err != nil {
		return nil, err
	} else if socketrv, err := socketv.Unmarshal(in.ExecutionContext, reflect.TypeOf((*SocketBase)(nil))); err != nil {
		return nil, err
	} else {
		socket := socketrv.Interface().(*SocketBase)

		in.Context.AddMicrotask(in.ExecutionContext, func(in isolates.FunctionArgs) error {
			in.Background(func(in isolates.FunctionArgs) {
				if conn, err := gonet.Dial(network, address); err != nil {
					socket.EmitError(in.ExecutionContext, err)
				} else if err := dconn.SetTCPConn(conn.(*gonet.TCPConn)); err != nil {
					socket.EmitError(in.ExecutionContext, err)
				} else {
					socket.Emit(in.ExecutionContext, "connect")
				}
			})
			return nil
		})

		return socket, nil
	}
}

//js:method listen
func (n *HostNetOptions) Listen(in isolates.FunctionArgs, server Server, network string, address string) (Listener, error) {
	if listener, err := gonet.Listen(network, address); err != nil {
		return nil, err
	} else if wrapped, err := NewListener(server, listener); err != nil {
		return nil, err
	} else {
		in.Background(func(in isolates.FunctionArgs) {
			defer wrapped.Close(in.ExecutionContext)
			defer server.Emit(in.ExecutionContext, "close")

			for {
				if conn, err := listener.Accept(); err != nil {
					if errors.Is(err, gonet.ErrClosed) {
						return
					}

					server.EmitError(in.ExecutionContext, err)
				} else {
					in.Background(func(in isolates.FunctionArgs) {
						if options, err := isolates.For(in.ExecutionContext).New(NewHostSocketOptions, conn); err != nil {
							server.EmitError(in.ExecutionContext, err)
						} else if socket, err := server.Net().Socket().New(in.ExecutionContext, options); err != nil {
							server.EmitError(in.ExecutionContext, err)
						} else if err := server.Emit(in.ExecutionContext, "connection", socket); err != nil {
							server.EmitError(in.ExecutionContext, err)
						}
					})
				}
			}
		})

		return wrapped, nil
	}
}

//js:constructor HostNetOptions
func NewHostNetOptions(in isolates.FunctionArgs) (*HostNetOptions, error) {
	options := &HostNetOptions{}

	in.This.RebindMethod(in.ExecutionContext, "connect")
	in.This.RebindMethod(in.ExecutionContext, "listen")

	return options, nil
}

//js:method createHostNet
//js:export-instance default
func NewHostNet(ctx context.Context) (*NetBase, error) {
	if options, err := isolates.For(ctx).New(NewHostNetOptions); err != nil {
		return nil, err
	} else if n, err := isolates.For(ctx).New(NewNet, options); err != nil {
		return nil, err
	} else {
		return n.(*NetBase), nil
	}
}

//js:constructor HostSocketOptions
func NewHostSocketOptions(in isolates.FunctionArgs) (*HostSocketOptions, error) {
	if connv, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&gonet.TCPConn{})); err != nil {
		return nil, err
	} else if stream, err := stream.NewReadWriteCloser(in); err != nil {
		return nil, err
	} else {
		conn := connv.Interface().(gonet.Conn)

		options := &HostSocketOptions{
			ReadWriteCloser: stream,
			conn:            conn,
		}

		in.This.RebindMethod(in.ExecutionContext, "localAddress")
		in.This.RebindMethod(in.ExecutionContext, "remoteAddress")
		in.This.RebindMethod(in.ExecutionContext, "final")

		return options, nil
	}
}

//js:method localAddress
func (c *HostSocketOptions) LocalAddress(ctx context.Context, this *SocketBase) (*Address, error) {
	if a, err := isolates.For(ctx).New(NewAddress, c.conn.LocalAddr()); err != nil {
		return nil, err
	} else {
		return a.(*Address), nil
	}
}

//js:method remoteAddress
func (c *HostSocketOptions) RemoteAddress(ctx context.Context, this *SocketBase) (*Address, error) {
	if a, err := isolates.For(ctx).New(NewAddress, c.conn.RemoteAddr()); err != nil {
		return nil, err
	} else {
		return a.(*Address), nil
	}
}

//js:method final
func (c *HostSocketOptions) Final(in isolates.FunctionArgs, this stream.Writable, callback *isolates.Value) error {
	if _err := c.conn.(ITCPConn).CloseWrite(); _err != nil {
		if _, _err := callback.Call(in.ExecutionContext, this, _err); _err != nil {
			return this.EmitError(in.ExecutionContext, _err)
		}
		return nil
	} else if _, _err := callback.Call(in.ExecutionContext, this); _err != nil {
		return this.EmitError(in.ExecutionContext, _err)
	} else {
		return nil
	}
}
