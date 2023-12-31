//js:package stream

package stream

import (
	"context"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"

	"github.com/grexie/isolates"
)

var _ Writable = &WritableBase{}

type BufferEncoding string

var (
	BufferEncodingUTF8 BufferEncoding = "utf8"
	BufferEncodingHex  BufferEncoding = "hex"
)

type Writable interface {
	Stream
	io.Writer

	Cork(context.Context) error
	End(context.Context, ...any) error
	SetDefaultEncoding(encoding BufferEncoding)
	Uncork(context.Context) error
	Writable() bool
	WritableAborted() bool
	WritableEnded() bool
	WritableCorked() int

	WritableFinished() bool
	WritableHighWaterMark() int
	WritableLength() int
	WritableNeedDrain() bool
	WritableObjectMode() bool
	WritableWrite(context.Context, ...any) error
}

type WritableBase struct {
	*StreamBase
}

type Writer struct {
	io.Writer
	mutex sync.Mutex
}

type WriteCloser struct {
	*Closer
	*Writer
}

//js:constructor Writable
//js:export Writable
func NewWritable(in isolates.FunctionArgs) (*WritableBase, error) {
	if stream, err := NewStream(in); err != nil {
		return nil, err
	} else if writable, err := NewWritableWithStream(in, stream); err != nil {
		return nil, err
	} else {

		return writable, nil
	}
}

func NewWritableWithStream(in isolates.FunctionArgs, StreamBase *StreamBase) (*WritableBase, error) {
	writable := &WritableBase{
		StreamBase: StreamBase,
	}

	if len(in.Args) > 0 {
		options := in.Args[0]

		if write, err := options.Get(in.ExecutionContext, "write"); err != nil {
			return nil, err
		} else if writev, err := options.Get(in.ExecutionContext, "writev"); err != nil {
			return nil, err
		} else if final, err := options.Get(in.ExecutionContext, "final"); err != nil {
			return nil, err
		} else {
			if write.IsKind(isolates.KindFunction) {
				if err := in.This.Set(in.ExecutionContext, "_write", write); err != nil {
					return nil, err
				}
			}

			if writev.IsKind(isolates.KindFunction) {
				if err := in.This.Set(in.ExecutionContext, "_writev", writev); err != nil {
					return nil, err
				}
			}

			if final.IsKind(isolates.KindFunction) {
				if err := in.This.Set(in.ExecutionContext, "_final", final); err != nil {
					return nil, err
				}
			}
		}
	}

	return writable, nil
}

func NewWriter(in isolates.FunctionArgs) (*Writer, error) {
	if writer, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(io.Writer); !ok {
		return nil, fmt.Errorf("not an io.Readable implementation: %s", in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext))
	} else {
		in.This.RebindMethod(in.ExecutionContext, "write")

		return &Writer{Writer: writer}, nil
	}
}

func NewWriteCloser(in isolates.FunctionArgs) (*WriteCloser, error) {
	if closer, err := NewCloser(in); err != nil {
		return nil, err
	} else if writer, err := NewWriter(in); err != nil {
		return nil, err
	} else {
		in.This.RebindMethod(in.ExecutionContext, "final")

		return &WriteCloser{closer, writer}, nil
	}
}

//js:method
func (w *WritableBase) Cork(context.Context) error {
	return errors.New("not implemented")
}

//js:method
func (w *WritableBase) End(ctx context.Context, args ...any) error {
	var err error
	var chunk *isolates.Value
	var encoding BufferEncoding
	var callback *isolates.Value

	if len(args) >= 3 {
		if encodingv, err := isolates.For(ctx).Context().Create(ctx, args[1]); err != nil {
			return err
		} else if rv, err := encodingv.Unmarshal(ctx, reflect.TypeOf(encoding)); err != nil {
			return err
		} else {
			encoding = rv.Interface().(BufferEncoding)
		}
	} else {
		encoding = "raw"
	}

	if len(args) >= 1 {
		if chunk, err = isolates.For(ctx).Context().Create(ctx, args[0]); err != nil {
			return err
		}
	}

	if len(args) >= 1 {
		if callback, err = isolates.For(ctx).Context().Create(ctx, args[len(args)-1]); err != nil {
			return err
		}
	}

	if chunk != nil && chunk.IsKind(isolates.KindFunction) {
		chunk = nil
	}

	if callback == nil || !callback.IsKind(isolates.KindFunction) {
		callback = nil
	}

	final := func() error {
		if _, err := w.This.CallMethod(ctx, "_final", func(in isolates.FunctionArgs) (*isolates.Value, error) {
			return nil, nil
		}); err != nil {
			return err
		} else {
			return nil
		}
	}

	if chunk != nil {
		if _, err := w.This.CallMethod(ctx, "_write", chunk, encoding, func(in isolates.FunctionArgs) (*isolates.Value, error) {
			return nil, final()
		}); err != nil {
			return err
		}
	} else {
		return final()
	}

	return nil
}

//js:method
func (w *WritableBase) SetDefaultEncoding(encoding BufferEncoding) {

}

//js:method
func (w *WritableBase) Uncork(context.Context) error {
	return nil
}

//js:get
func (w *WritableBase) Writable() bool {
	return true
}

//js:get
func (w *WritableBase) WritableAborted() bool {
	return false
}

//js:get
func (w *WritableBase) WritableEnded() bool {
	return false
}

//js:get
func (w *WritableBase) WritableCorked() int {
	return 0
}

//js:get
func (w *WritableBase) WritableFinished() bool {
	return false
}

//js:get
func (w *WritableBase) WritableHighWaterMark() int {
	return 0
}

//js:get
func (w *WritableBase) WritableLength() int {
	return 0
}

//js:get
func (w *WritableBase) WritableNeedDrain() bool {
	return false
}

//js:get
func (w *WritableBase) WritableObjectMode() bool {
	return false
}

func (c *WritableBase) Write(b []byte) (int, error) {
	ctx := c.Context.GetIsolate().GetExecutionContext()

	result := make(chan struct {
		n   int
		err error
	})

	if callback, err := c.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		in.Background(func(in isolates.FunctionArgs) {
			if err := in.Arg(in.ExecutionContext, 0).ToError(in.ExecutionContext); err != nil {
				result <- struct {
					n   int
					err error
				}{
					0,
					err,
				}
			} else {
				result <- struct {
					n   int
					err error
				}{
					len(b),
					nil,
				}
			}
		})
		return nil, nil
	}); err != nil {
		return 0, err
	} else if _, err := c.This.CallMethod(ctx, "write", b, callback); err != nil {
		return 0, err
	} else {
		result := <-result
		return result.n, result.err
	}
}

//js:method write
func (w *WritableBase) WritableWrite(ctx context.Context, args ...any) error {
	var err error
	var chunk *isolates.Value
	var encoding BufferEncoding
	var callback *isolates.Value

	if len(args) >= 3 {
		if encodingv, err := isolates.For(ctx).Context().Create(ctx, args[1]); err != nil {
			return err
		} else if rv, err := encodingv.Unmarshal(ctx, reflect.TypeOf(encoding)); err != nil {
			return err
		} else {
			encoding = rv.Interface().(BufferEncoding)
		}
	} else {
		encoding = "raw"
	}

	if len(args) >= 1 {
		if chunk, err = isolates.For(ctx).Context().Create(ctx, args[0]); err != nil {
			return err
		}
	}

	if len(args) >= 2 {
		if callback, err = isolates.For(ctx).Context().Create(ctx, args[len(args)-1]); err != nil {
			return err
		}
	} else {
		if callback, err = isolates.For(ctx).Context().Create(ctx, func(isolates.FunctionArgs) (*isolates.Value, error) {
			return nil, nil
		}); err != nil {
			return err
		}
	}

	if _, err := w.This.CallMethod(ctx, "_write", chunk, encoding, callback); err != nil {
		return err
	} else {
		return nil
	}
}

//js:method
func (c *Writer) Write(in isolates.FunctionArgs, this Writable, chunk *isolates.Value, encoding string, callback *isolates.Value) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	var b []byte

	if chunk.IsKind(isolates.KindString) {
		if s, err := chunk.StringValue(in.ExecutionContext); err != nil {
			if _, err := callback.Call(in.ExecutionContext, this, err); err != nil {
				return this.EmitError(in.ExecutionContext, err)
			}
			return nil
		} else {
			b = []byte(s)
		}
	}

	if chunk.IsKind(isolates.KindArrayBuffer) {
		if bytes, err := chunk.Bytes(in.ExecutionContext); err != nil {
			if _, err := callback.Call(in.ExecutionContext, this, err); err != nil {
				return this.EmitError(in.ExecutionContext, err)
			}
			return nil
		} else {
			b = bytes
		}
	}

	if n, err := c.Writer.Write(b); err != nil {
		if _, err := callback.Call(in.ExecutionContext, this, err); err != nil {
			return this.EmitError(in.ExecutionContext, err)
		} else {
			return nil
		}
	} else if n != len(b) {
		err := errors.New("failed to write all bytes")
		if _, err := callback.Call(in.ExecutionContext, this, err); err != nil {
			return this.EmitError(in.ExecutionContext, err)
		} else {
			return nil
		}
	} else if _, err := callback.Call(in.ExecutionContext, this); err != nil {
		return this.EmitError(in.ExecutionContext, err)
	} else {
		return nil
	}
}

//js:method
func (c *WriteCloser) Final(in isolates.FunctionArgs, this Writable, callback *isolates.Value) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _err := c.Closer.Close(); _err != nil {
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
