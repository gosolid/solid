//js:package native:@grexie/workers/stream
package stream

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"runtime/debug"

	"github.com/gosolid/solid/pkg/runtime/events"
	"github.com/grexie/isolates"
)

type EventedCloser interface {
	io.Closer
	OnClose(func() error) func()
}

type EventedReadCloser interface {
	io.ReadCloser
	EventedCloser
}

type EventedWriteCloser interface {
	io.WriteCloser
	EventedCloser
}

type EventedReadWriteCloser interface {
	io.ReadWriteCloser
	EventedCloser
}

//js:class
type ReadableStream struct {
	stream  io.ReadCloser
	onClose events.EventHandler
}

func NewReadableStream(stream io.ReadCloser) *ReadableStream {
	return &ReadableStream{
		stream: stream,
	}
}

func (s *ReadableStream) Read(p []byte) (int, error) {
	return s.stream.Read(p)
}

func (s *ReadableStream) V8FuncRead(in isolates.FunctionArgs) (*isolates.Value, error) {
	if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		go func() {
			debug.SetPanicOnFault(true)
			defer func() {
				if p := recover(); p != nil {
					err := fmt.Errorf("recover panic: sigseg")
					isolates.For(in.ExecutionContext).Debug(err)
					return
				}
			}()

			ctx := isolates.WithContext(context.Background())

			var p []byte
			if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(p)); err != nil {
				resolver.Reject(ctx, err)
			} else {
				p = v.Interface().([]byte)
			}

			if result, err := s.Read(p); err != nil && result <= 0 {
				resolver.Reject(ctx, err)
			} else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
				resolver.Reject(ctx, err)
			} else if err := in.Arg(in.ExecutionContext, 0).SetBytes(in.ExecutionContext, p); err != nil {
				resolver.Reject(ctx, err)
			} else {
				resolver.Resolve(ctx, result)
			}
		}()

		return resolver.Promise(in.ExecutionContext)
	}
}

//js:async-method
func (s *ReadableStream) Close() error {
	if err := s.stream.Close(); err != nil {
		return err
	} else {
		return s.onClose.Emit()
	}
}

//js:event close
func (s *ReadableStream) OnClose(listener func() error) func() {
	return s.onClose.Add(func(...any) error {
		return listener()
	})
}

//js:class
type WritableStream struct {
	stream  io.WriteCloser
	onClose events.EventHandler
}

func NewWritableStream(stream io.WriteCloser) *WritableStream {
	return &WritableStream{
		stream: stream,
	}
}

//js:async-method
func (s *WritableStream) Write(p []byte) (int, error) {
	return s.stream.Write(p)
}

//js:async-method
func (s *WritableStream) Close() error {
	if err := s.stream.Close(); err != nil {
		return err
	} else {
		return s.onClose.Emit()
	}
}

//js:event close
func (s *WritableStream) OnClose(listener func() error) func() {
	return s.onClose.Add(func(...any) error {
		return listener()
	})
}

//js:class
type StreamOld struct {
	stream  io.ReadWriteCloser
	onClose events.EventHandler
}

func NewStreamOld(stream io.ReadWriteCloser) *StreamOld {
	return &StreamOld{
		stream: stream,
	}
}

func (s *StreamOld) Read(p []byte) (int, error) {
	return s.stream.Read(p)
}

func (s *StreamOld) V8FuncRead(in isolates.FunctionArgs) (*isolates.Value, error) {
	if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		go func() {
			debug.SetPanicOnFault(true)
			ctx := isolates.WithContext(context.Background())

			defer func() {
				if p := recover(); p != nil {
					err := fmt.Errorf("recover panic: sigseg")
					isolates.For(ctx).Debug(err)
					return
				}
			}()

			var p []byte
			if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(p)); err != nil {
				resolver.Reject(ctx, err)
			} else {
				p = v.Interface().([]byte)
			}

			if result, err := s.Read(p); err != nil && result <= 0 {
				resolver.Reject(ctx, err)
			} else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
				resolver.Reject(ctx, err)
			} else if err := in.Arg(in.ExecutionContext, 0).SetBytes(in.ExecutionContext, p); err != nil {
				resolver.Reject(ctx, err)
			} else {
				resolver.Resolve(ctx, result)
			}
		}()

		return resolver.Promise(in.ExecutionContext)
	}
}

//js:async-method
func (s *StreamOld) Write(p []byte) (int, error) {
	return s.stream.Write(p)
}

//js:async-method
func (s *StreamOld) Close() error {
	if err := s.stream.Close(); err != nil {
		return err
	} else {
		return s.onClose.Emit()
	}
}

//js:event close
func (s *StreamOld) OnClose(listener func() error) func() {
	return s.onClose.Add(func(...any) error {
		return listener()
	})
}
