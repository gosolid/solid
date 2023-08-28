//js:package net
package net

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gosolid/solid/pkg/runtime/events"
	"github.com/grexie/isolates"
)

type Server interface {
	events.EventEmitter
	Listener
	Net() Net
}

type ServerBase struct {
	*events.EventEmitterBase
	net      Net
	listener *isolates.Value
}

func newServer(in isolates.FunctionArgs, net Net) (*ServerBase, error) {
	return &ServerBase{
		EventEmitterBase: events.NewEventEmitter(in),
		net:              net,
	}, nil
}

func (s *ServerBase) Net() Net {
	return s.net
}

//js:method
func (s *ServerBase) Close(ctx context.Context) error {
	if s.listener == nil {
		return nil
	}

	if _, err := s.listener.CallMethod(ctx, "close"); err != nil {
		return err
	} else {
		return nil
	}
}

//js:method address
func (s *ServerBase) Address(ctx context.Context) (*Address, error) {
	if s.listener == nil {
		return nil, nil
	}

	if addrv, err := s.listener.CallMethod(ctx, "address"); err != nil {
		return nil, err
	} else if rv, err := addrv.Unmarshal(ctx, reflect.TypeOf(&Address{})); err != nil {
		return nil, err
	} else {
		return rv.Interface().(*Address), nil
	}
}

//js:get _listener
func (s *ServerBase) Listener() *isolates.Value {
	return s.listener
}

//js:method listen
func (s *ServerBase) Listen(in isolates.FunctionArgs) (*isolates.Value, error) {
	if callback := in.Arg(in.ExecutionContext, 2); callback.IsKind(isolates.KindFunction) {
		s.AddListener(in.ExecutionContext, "listening", callback)
	}

	if host, err := in.Arg(in.ExecutionContext, 1).StringValue(in.ExecutionContext); err != nil {
		return nil, err
	} else if port, err := in.Arg(in.ExecutionContext, 0).Int64(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		in.Background(func(in isolates.FunctionArgs) {
			if listener, err := s.net.Listen(in.ExecutionContext, s, "tcp", fmt.Sprintf("%s:%d", host, port)); err != nil {
				s.EmitError(in.ExecutionContext, err)
			} else {
				s.listener = listener
				s.Emit(in.ExecutionContext, "listening")
			}
		})
		return in.This, nil
	}
}
