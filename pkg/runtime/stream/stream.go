//js:package stream

package stream

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/gosolid/solid/pkg/runtime/events"
	isolates "github.com/grexie/isolates"
)

var _ Stream = &StreamBase{}

type StreamState int

const (
	StreamStateNew StreamState = iota
	StreamStateReady
	StreamStateErrored
	StreamStateDestroying
	StreamStateDestroyed
	StreamStateClosing
	StreamStateClosed
)

type Stream interface {
	events.EventEmitter
	io.Closer
	Destroy(context.Context, *isolates.Value) error
	Closed() bool
	Destroyed() bool
	Errored() *isolates.Value
}

type StreamBase struct {
	*events.EventEmitterBase
	mutex         sync.Mutex
	Context       *isolates.Context
	This          *isolates.Value
	Options       *isolates.Value
	onStateChange events.EventHandler
	state         StreamState
	errored       *isolates.Value
}

type Closer struct {
	io.Closer
}

//js:method
//js:export pipeline
func Pipeline(ctx context.Context, args ...*isolates.Value) *isolates.Value {
	isolates.For(ctx).Debug("pipeline", args)
	return nil
}

//js:constructor Stream
//js:export Stream
//js:export default
func NewStream(in isolates.FunctionArgs) (*StreamBase, error) {
	stream := &StreamBase{
		EventEmitterBase: events.NewEventEmitter(in),
		Context:          in.Context,
		This:             in.This,
		Options:          in.Arg(in.ExecutionContext, 0),
		state:            StreamStateNew,
	}

	hasConstruct := false

	if len(in.Args) > 0 && !in.Args[0].IsNil() {
		options := in.Args[0]

		if construct, err := options.Get(in.ExecutionContext, "construct"); err != nil {
			return nil, err
		} else if construct != nil {
			if construct.IsKind(isolates.KindFunction) {
				if err := in.This.Set(in.ExecutionContext, "_construct", construct); err != nil {
					return nil, err
				} else {
					hasConstruct = true
				}
			}
		}

		if destroy, err := options.Get(in.ExecutionContext, "destroy"); err != nil {
			return nil, err
		} else if destroy != nil {
			if destroy.IsKind(isolates.KindFunction) {
				if err := in.This.Set(in.ExecutionContext, "_destroy", destroy); err != nil {
					return nil, err
				}
			}
		}
	}

	if hasConstruct {
		in.Context.AddMicrotask(in.ExecutionContext, func(in isolates.FunctionArgs) error {
			if construct, err := stream.This.Get(in.ExecutionContext, "_construct"); err != nil {
				return err
			} else if !construct.IsNil() {
				if construct.IsKind(isolates.KindFunction) {
					if _, err := construct.Call(in.ExecutionContext, stream, func(in isolates.FunctionArgs) (*isolates.Value, error) {
						if errorv := in.Arg(in.ExecutionContext, 0); !(errorv.IsKind(isolates.KindNull) || errorv.IsKind(isolates.KindUndefined)) {
							stream.EmitErrorValue(in.ExecutionContext, errorv)
						} else {
							stream.SetState(StreamStateReady)
						}
						return nil, nil
					}); err != nil {
						return err
					} else {
						return nil
					}
				}
			}

			return nil
		})
	} else {
		stream.SetState(StreamStateReady)
	}

	return stream, nil
}

func NewCloser(in isolates.FunctionArgs) (*Closer, error) {
	if closer, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(io.Closer); !ok {
		return nil, fmt.Errorf("not an io.Closer implementation: %s", in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext))
	} else {
		in.This.RebindMethod(in.ExecutionContext, "destroy")

		return &Closer{Closer: closer}, nil
	}
}

func (s *StreamBase) Close() error {
	ctx := s.Context.GetIsolate().GetExecutionContext()

	close := make(chan error)

	if callback, err := s.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		close <- in.Arg(in.ExecutionContext, 0).ToError(in.ExecutionContext)
		return nil, nil
	}); err != nil {
		return err
	} else if err := s.AddListenerOnce(ctx, "close", callback); err != nil {
		return err
	} else if _, err := s.This.CallMethod(ctx, "destroy", nil); err != nil {
		return err
	} else {
		return <-close
	}
}

//js:method
func (s *StreamBase) Destroy(ctx context.Context, err *isolates.Value) error {
	if err == nil || err.IsKind(isolates.KindUndefined) || err.IsKind(isolates.KindNull) {
		s.SetState(StreamStateDestroying)
	} else {
		s.SetState(StreamStateErrored)
	}

	if errv, err := s.Context.Create(ctx, err); err != nil {
		return err
	} else if callback, err := s.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if len(in.Args) > 0 && !in.Args[0].IsKind(isolates.KindUndefined) && !in.Args[0].IsKind(isolates.KindNull) {
			s.SetState(StreamStateErrored)
			s.EmitErrorValue(in.ExecutionContext, in.Args[0])
			return nil, nil
		} else {
			s.SetState(StreamStateDestroyed)
			s.Emit(in.ExecutionContext, "close")
			return nil, nil
		}
	}); err != nil {
		return err
	} else if destroy, err := s.This.Get(ctx, "_destroy"); err != nil {
		if errv != nil {
			return s.EmitErrorValue(ctx, errv)
		} else {
			return nil
		}
	} else if destroy != nil && destroy.IsKind(isolates.KindFunction) {
		if _, err := destroy.Call(ctx, s.This, errv, callback); err != nil {
			return err
		} else {
			return nil
		}
	} else if _, err := callback.Call(ctx, s.This); err != nil {
		return err
	} else {
		return nil
	}
}

//js:get
func (s *StreamBase) Closed() bool {
	return s.state == StreamStateErrored || s.state == StreamStateDestroying || s.state == StreamStateDestroyed || s.state == StreamStateClosing || s.state == StreamStateClosed
}

//js:get
func (s *StreamBase) Destroyed() bool {
	return s.state == StreamStateErrored || s.state == StreamStateDestroying || s.state == StreamStateDestroyed
}

//js:get
func (s *StreamBase) Errored() *isolates.Value {
	if s.errored != nil {
		return s.errored
	} else {
		return nil
	}
}

func (s *StreamBase) SetState(state StreamState) bool {
	s.mutex.Lock()
	if s.state != state {
		s.state = state
		s.mutex.Unlock()
		s.onStateChange.Emit(state)
		return true
	} else {
		s.mutex.Unlock()
		return false
	}
}

func (s *StreamBase) SetStateConditional(fromState StreamState, toState StreamState) bool {
	if fromState == toState {
		return false
	}

	s.mutex.Lock()
	if s.state == fromState {
		s.state = toState
		s.mutex.Unlock()
		s.onStateChange.Emit(toState)
		return true
	} else {
		s.mutex.Unlock()
		return false
	}
}

func (s *StreamBase) OnStateChange(callback func(s StreamState) error) func() {
	return s.onStateChange.Add(func(args ...any) error {
		s := args[0].(StreamState)
		return callback(s)
	})
}

func (s *StreamBase) EmitError(ctx context.Context, err error) error {
	if errv, err := isolates.For(ctx).Context().Create(ctx, err); err != nil {
		return err
	} else {
		return s.EmitErrorValue(ctx, errv)
	}
}

func (s *StreamBase) EmitErrorValue(ctx context.Context, error *isolates.Value) error {
	if s.errored != nil {
		return nil
	}
	s.errored = error
	s.SetState(StreamStateErrored)
	return s.Emit(ctx, "error", error)
}

//js:method
func (c *Closer) Destroy(in isolates.FunctionArgs, this Stream, err *isolates.Value, callback *isolates.Value) error {
	if _err := c.Closer.Close(); _err != nil {
		if callback != nil && callback.IsKind(isolates.KindFunction) {
			if _, _err := callback.Call(in.ExecutionContext, this, _err); _err != nil {
				return this.EmitError(in.ExecutionContext, _err)
			}
		}
		return nil
	} else if callback != nil && callback.IsKind(isolates.KindFunction) {
		if _, _err := callback.Call(in.ExecutionContext, this, err); _err != nil {
			return this.EmitError(in.ExecutionContext, _err)
		} else {
			return nil
		}
	} else {
		return nil
	}
}
