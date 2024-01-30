//js:package stream

package stream

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/grexie/isolates"
)

var _ Readable = &ReadableBase{}
var _ io.ReadCloser = &ReadableBase{}
var _ Pipe = &PipeBase{}

type ReadableStreamState int

const (
	ReadableStreamStatePaused ReadableStreamState = iota
	ReadableStreamStateResumed
)

//js:alias ReadableBase
type Readable interface {
	Stream
	io.Reader

	IsPaused() bool
	Pause(context.Context)
	Pipe(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error)
	ReadV8(ctx context.Context, size int) error
	Readable() bool
	ReadableAborted() bool
	ReadableDidRead() bool
	ReadableEncoding() *BufferEncoding
	ReadableEnded() bool
	ReadableFlowing() bool
	ReadableHighWaterMark() int
	ReadableLength() int
	ReadableObjectMode() bool
	Resume(context.Context)
	Drain(ctx context.Context) error
	SetEncoding(BufferEncoding) (*isolates.Value, error)
	Unpipe(context.Context, *isolates.Value) (*isolates.Value, error)
	Unshift(context.Context, *isolates.Value, *BufferEncoding) error
	UnshiftBuffer() []*isolates.Value
	SetUnshiftBuffer(buffer []*isolates.Value)
	Wrap(context.Context, *isolates.Value) (*isolates.Value, error)
	Compose(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error)
	Iterator(context.Context, *isolates.Value) (*isolates.Value, error)
	Map(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error)
	Filter(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error)
	ForEach(context.Context, *isolates.Value, *isolates.Value) error
	ToArray(context.Context, *isolates.Value) (*isolates.Value, error)
	Some(context.Context, *isolates.Value, *isolates.Value) (bool, error)
	Find(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error)
	Every(context.Context, *isolates.Value, *isolates.Value) (bool, error)
	FlatMap(context.Context, *isolates.Value, *isolates.Value) (Readable, error)
	Drop(context.Context, int, *isolates.Value) (Readable, error)
	Take(context.Context, int, *isolates.Value) (Readable, error)
	AsIndexedPairs(context.Context, *isolates.Value) (Readable, error)
	Reduce(context.Context, *isolates.Value, *isolates.Value, *isolates.Value) (*isolates.Value, error)
	Push(ctx context.Context, chunk *isolates.Value, encoding *BufferEncoding) (bool, error)
}

type Pipe interface {
	Destination() *isolates.Value
	Remove(ctx context.Context) error
	OnData() *isolates.Value
	OnError() *isolates.Value
	OnEnd() *isolates.Value
}

type ReadableBase struct {
	*StreamBase
	mutex   sync.Mutex
	state   ReadableStreamState
	pipes   []Pipe
	buffer []*isolates.Value
	unshiftBuffer []*isolates.Value
	readable bool
	reading bool
	draining bool
	emittedEnd bool
	highWaterMark int
	waterMark int
	objectMode bool
}

type Reader struct {
	io.Reader
	mutex  sync.Mutex
	b      []byte
	buffer *isolates.Value
}

type ReadCloser struct {
	*Closer
	*Reader
}

type PipeBase struct {
	source      *isolates.Value
	destination *isolates.Value
	remove      *isolates.Value
	onData      *isolates.Value
	onError     *isolates.Value
	onEnd       *isolates.Value
}

//js:constructor Readable
//js:export Readable
func NewReadable(in isolates.FunctionArgs) (*ReadableBase, error) {
	if StreamBase, err := NewStream(in); err != nil {
		return nil, err
	} else if readable, err := NewReadableWithStream(in, StreamBase); err != nil {
		return nil, err
	} else {
		return readable, nil
	}
}

func NewReadableWithStream(in isolates.FunctionArgs, StreamBase *StreamBase) (*ReadableBase, error) {
	readable := &ReadableBase{
		StreamBase: 	 StreamBase,
		state:      	 ReadableStreamStatePaused,
		highWaterMark: 1024*32,
	}

	if len(in.Args) > 0 && !in.Args[0].IsNil() {
		options := in.Args[0]

		if read, err := options.Get(in.ExecutionContext, "read"); err != nil {
			return nil, err
		} else {
			if read.IsKind(isolates.KindFunction) {
				if err := in.This.Set(in.ExecutionContext, "_read", read); err != nil {
					return nil, err
				}
			}
		}

		if highWaterMarkv, err := options.Get(in.ExecutionContext, "highWaterMark"); err != nil {
			return nil, err
		} else if !highWaterMarkv.IsNil() {
			if highWaterMark, err := highWaterMarkv.Int64(in.ExecutionContext); err != nil {
				return nil, err
			} else {
				readable.highWaterMark = int(highWaterMark)
			}
		}

		if objectModev, err := options.Get(in.ExecutionContext, "objectMode"); err != nil {
			return nil, err
		} else if !objectModev.IsNil() {
			if objectMode, err := objectModev.Bool(in.ExecutionContext); err != nil {
				return nil, err
			} else {
				readable.objectMode = objectMode
			}
		}
	}

	in.Context.AddMicrotask(in.ExecutionContext, func(in isolates.FunctionArgs) error {
		return readable.ReadV8(in.ExecutionContext, 1024)
	})

	return readable, nil
}

func NewReader(in isolates.FunctionArgs) (*Reader, error) {
	if reader, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(io.Reader); !ok {
		return nil, fmt.Errorf("not an io.Reader implementation: %s", in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext))
	} else {
		in.This.RebindMethod(in.ExecutionContext, "construct")
		in.This.RebindMethod(in.ExecutionContext, "read")

		r := &Reader{Reader: reader}

		r.mutex.Lock()
		return r, nil
	}
}

func NewReadCloser(in isolates.FunctionArgs) (*ReadCloser, error) {
	if closer, err := NewCloser(in); err != nil {
		return nil, err
	} else if reader, err := NewReader(in); err != nil {
		return nil, err
	} else {
		return &ReadCloser{closer, reader}, nil
	}
}

func (s *ReadableBase) Read(b []byte) (int, error) {
	ctx := s.Context.GetIsolate().GetExecutionContext()

	result := make(chan struct {
		n   int
		err error
	})

	if dataCallback, err := s.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if bytes, err := in.Arg(in.ExecutionContext, 0).Bytes(in.ExecutionContext); err != nil {
			return nil, err
		} else {
			copy(b, bytes)
			result <- struct {
				n   int
				err error
			}{len(bytes), nil}
			return nil, nil
		}
	}); err != nil {
		return 0, err
	} else if errorCallback, err := s.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		err := in.Arg(in.ExecutionContext, 0).ToError(in.ExecutionContext)
		result <- struct {
			n   int
			err error
		}{0, err}
		return nil, nil
	}); err != nil {
		return 0, err
	} else if destroyCallback, err := s.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		result <- struct {
			n   int
			err error
		}{0, io.EOF}
		return nil, nil
	}); err != nil {
		return 0, err
	} else if err := s.AddListenerOnce(ctx, "data", dataCallback); err != nil {
		return 0, err
	} else if err := s.AddListenerOnce(ctx, "error", errorCallback); err != nil {
		return 0, err
	} else if err := s.AddListenerOnce(ctx, "destroy", destroyCallback); err != nil {
		return 0, err
	} else if _, err := s.This.CallMethod(ctx, "resume", nil); err != nil {
		return 0, err
	} else if _, err := s.This.CallMethod(ctx, "pause", nil); err != nil {
		return 0, err
	} else {
		defer s.RemoveListener(ctx, "data", dataCallback)
		defer s.RemoveListener(ctx, "error", errorCallback)
		defer s.RemoveListener(ctx, "destroy", destroyCallback)
		result := <-result
		return result.n, result.err
	}
}

func (s *ReadableBase) SetReadableState(state ReadableStreamState) bool {
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

func (s *ReadableBase) SetReadableStateConditional(fromState ReadableStreamState, toState ReadableStreamState) bool {
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

//js:get
func (r *ReadableBase) IsPaused() bool {
	return r.state == ReadableStreamStatePaused
}

//js:method
func (r *ReadableBase) Pause(context.Context) {
	r.SetReadableStateConditional(ReadableStreamStateResumed, ReadableStreamStatePaused)
}

//js:method
func (r *ReadableBase) Resume(ctx context.Context) {
	if r.SetReadableStateConditional(ReadableStreamStatePaused, ReadableStreamStateResumed) {
		isolates.For(ctx).Background(func(ctx context.Context) {
			for !r.IsPaused() && len(r.buffer) > 0 {
				chunk := r.buffer[0]
				r.buffer = r.buffer[1:]
				r.Emit(ctx, "data", chunk)
			}
			if (r.StreamBase.state == StreamStateNew || r.StreamBase.state == StreamStateReady) {
				if err := r.ReadV8(ctx, 1024); err != nil {
					r.EmitError(ctx, err)
				}
			} else if len(r.buffer) == 0 && r.Closed() && !r.emittedEnd {
				r.emittedEnd = true
				if r.IsPaused() {
					r.Emit(ctx, "readable")
				}
				r.Emit(ctx, "end")
			}
		})
	}
}

//js:method read
func (r *ReadableBase) ReadChunk(ctx context.Context) (*isolates.Value, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if null, err := isolates.For(ctx).Context().Null(ctx); err != nil {
		return nil, err
	} else if r.state == ReadableStreamStateResumed {
		return null, nil
	} else {
		if len(r.buffer) == 0 && !r.Closed() {
			if err := r.ReadV8(ctx, 1024); err != nil {
				return nil, err
			}
		}

		if len(r.buffer) != 0 {
			chunk := r.buffer[0]
			r.buffer = r.buffer[1:]
			return chunk, nil
		} else {
			return null, nil
		}
	}
}

//js:method
func (r *ReadableBase) Pipe(ctx context.Context, destination *isolates.Value, options *isolates.Value) (*isolates.Value, error) {
	end := true

	if options.IsKind(isolates.KindObject) {
		if endv, err := options.Get(ctx, "end"); err != nil {
			return nil, err
		} else if end, err = endv.Bool(ctx); err != nil {
			return nil, err
		}
	}

	var err error
	pipe := &PipeBase{
		source:      r.This,
		destination: destination,
	}

	r.pipes = append(r.pipes, pipe)

	if pipe.onData, err = r.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
		pipe.destination.CallMethod(in.ExecutionContext, "write", in.Arg(in.ExecutionContext, 0))
		return nil, nil
	}); err != nil {
		return nil, err
	} else if err := r.AddListener(ctx, "data", pipe.onData); err != nil {
		return nil, err
	} else if end {
		if pipe.onError, err = r.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
			return pipe.destination.CallMethod(in.ExecutionContext, "error", in.Arg(in.ExecutionContext, 0))
		}); err != nil {
			return nil, err
		} else if pipe.onEnd, err = r.Context.Create(ctx, func(in isolates.FunctionArgs) (*isolates.Value, error) {
			pipe.destination.CallMethod(in.ExecutionContext, "end")
			return nil, nil
		}); err != nil {
			return nil, err
		} else if err := r.AddListener(ctx, "error", pipe.onError); err != nil {
			return nil, err
		} else if err := r.AddListener(ctx, "end", pipe.onEnd); err != nil {
			return nil, err
		}
	}

	pipe.destination.CallMethod(ctx, "emit", "pipe", r)

	r.Context.AddMicrotask(ctx, func(in isolates.FunctionArgs) error {
		r.Resume(in.ExecutionContext)
		return nil
	})

	return destination, nil
}

func (r *ReadableBase) ReadV8(ctx context.Context, size int) error {
	r.mutex.Lock()
	if r.reading {
		r.mutex.Unlock()
		return nil
	} else {
		r.reading = true
		defer func() {
			r.mutex.Lock()
			defer r.mutex.Unlock()
			r.reading = false
		}()
		r.mutex.Unlock()
	}

	if _, err := r.This.CallMethod(ctx, "_read", size); err != nil {
		return err
	}

	isolates.For(ctx).Context().AddMicrotask(ctx, func(in isolates.FunctionArgs) error {
		r.mutex.Lock()
		if len(r.buffer) > 0 && !r.draining {
			r.mutex.Unlock()
			r.drainReadableBuffer(in.ExecutionContext)
			return nil
		} else {
			r.mutex.Unlock()
			return nil
		}
	})

	return nil
}

//js:get
func (r *ReadableBase) Readable() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return len(r.buffer) > 0
}

//js:get
func (r *ReadableBase) ReadableAborted() bool {
	return false
}

//js:get
func (r *ReadableBase) ReadableDidRead() bool {
	return false
}

//js:get
func (r *ReadableBase) ReadableEncoding() *BufferEncoding {
	return nil
}

//js:get
func (r *ReadableBase) ReadableEnded() bool {
	return false
}

//js:get
func (r *ReadableBase) ReadableFlowing() bool {
	return false
}

//js:get
func (r *ReadableBase) ReadableHighWaterMark() int {
	return r.highWaterMark
}

//js:set readableHighWaterMark
func (r *ReadableBase) SetReadableHighWaterMark(value int) {
	r.highWaterMark = value
}

//js:get
func (r *ReadableBase) ReadableLength() int {
	return r.waterMark
}

//js:get
func (r *ReadableBase) ReadableObjectMode() bool {
	return r.objectMode
}

//js:method
func (r *ReadableBase) SetEncoding(BufferEncoding) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:method
func (r *ReadableBase) Unpipe(ctx context.Context, destination *isolates.Value) (*isolates.Value, error) {
	for i, p := range r.pipes {
		equals := true

		if !destination.IsNil() {
			var err error
			if equals, err = p.Destination().StrictEquals(ctx, destination); err != nil {
				return nil, err
			}
		}

		if equals {
			if p.OnData() != nil {
				if err := r.RemoveListener(ctx, "data", p.OnData()); err != nil {
					return nil, err
				}
			}

			if p.OnError() != nil {
				if err := r.RemoveListener(ctx, "error", p.OnError()); err != nil {
					return nil, err
				}
			}

			if p.OnEnd() != nil {
				if err := r.RemoveListener(ctx, "end", p.OnEnd()); err != nil {
					return nil, err
				}
			}

			p.Destination().CallMethod(ctx, "emit", "unpipe", r)

			r.pipes = append(r.pipes[:i], r.pipes[i+1:]...)

			return r.This, nil
		}
	}

	return r.This, nil
}

//js:method
func (r *ReadableBase) Unshift(ctx context.Context, chunk *isolates.Value, encoding *BufferEncoding) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	length := 1

	if !r.objectMode {
		// if lengthv, err := chunk.Get(ctx, "length"); err != nil {
		// 	return err
		// } else if l, err := lengthv.Int64(ctx); err != nil {
		// 	return err
		// } else {
		// 	length = int(l)
		// }
		if l, err := chunk.GetByteLength(ctx); err != nil {
			return err
		} else {
			length = l
		}
	}

	r.waterMark += length

	r.unshiftBuffer = append([]*isolates.Value{chunk}, r.unshiftBuffer...)
	return nil
}

func (r *ReadableBase) UnshiftBuffer() []*isolates.Value {
	return r.unshiftBuffer
}

func (r *ReadableBase) SetUnshiftBuffer(buffer []*isolates.Value) {
	r.unshiftBuffer = buffer
}

//js:method
func (r *ReadableBase) Wrap(context.Context, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:method
func (r *ReadableBase) Compose(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:method
func (r *ReadableBase) Iterator(context.Context, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Map(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Filter(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) ForEach(context.Context, *isolates.Value, *isolates.Value) error {
	return errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) ToArray(context.Context, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Some(context.Context, *isolates.Value, *isolates.Value) (bool, error) {
	return false, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Find(context.Context, *isolates.Value, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Every(context.Context, *isolates.Value, *isolates.Value) (bool, error) {
	return false, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) FlatMap(context.Context, *isolates.Value, *isolates.Value) (Readable, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Drop(context.Context, int, *isolates.Value) (Readable, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Take(context.Context, int, *isolates.Value) (Readable, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) AsIndexedPairs(context.Context, *isolates.Value) (Readable, error) {
	return nil, errors.New("not implemented")
}

//js:async-method
func (r *ReadableBase) Reduce(context.Context, *isolates.Value, *isolates.Value, *isolates.Value) (*isolates.Value, error) {
	return nil, errors.New("not implemented")
}

func (r *ReadableBase) drainReadableBuffer(ctx context.Context) error {
	r.mutex.Lock()
	if r.draining {
		r.mutex.Unlock()
		return nil
	} else {
		r.draining = true
		defer func() {
			r.mutex.Lock()
			defer r.mutex.Unlock()
			r.draining = false
		}()
		r.mutex.Unlock()
	}

	for len(r.buffer) > 0 && !r.IsPaused() {
		r.mutex.Lock()
		chunk := r.buffer[0]
		r.buffer = r.buffer[1:]
		length := 1
		if !r.objectMode {
			if chunk.IsKind(isolates.KindString) {
				if s, err := chunk.StringValue(ctx); err != nil {
					return err
				} else {
					length = len(s)
				}
			} else {
				// if lengthv, err := chunk.Get(ctx, "length"); err != nil {
				// 	return err
				// } else if l, err := lengthv.Int64(ctx); err != nil {
				// 	return err
				// } else {
				// 	length = int(l)
				// }
				if l, err := chunk.GetByteLength(ctx); err != nil {
					return err
				} else {
					length = l
				}
			}
		}
		r.waterMark -= length
		r.mutex.Unlock()
		r.Emit(ctx, "data", chunk)
	}


	r.mutex.Lock()
	if len(r.buffer) == 0 && !r.IsPaused() && !r.Closed() {
		r.mutex.Unlock()
		isolates.For(ctx).Context().AddMicrotask(ctx, func(in isolates.FunctionArgs) error {
			return r.ReadV8(in.ExecutionContext, 1024)
		})
	} else if !r.emittedEnd && !r.IsPaused() && r.StreamBase.state == StreamStateClosing {
		r.emittedEnd = true
		r.SetStateConditional(StreamStateClosing, StreamStateClosed)
		r.mutex.Unlock()

		if r.IsPaused() {
			r.Emit(ctx, "readable")
		}
		r.Emit(ctx, "end")
	} else {
		r.mutex.Unlock()
	}

	return nil
}

func (r *ReadableBase) Drain(ctx context.Context) (error) {
	return isolates.For(ctx).Context().AddMicrotask(ctx, func(in isolates.FunctionArgs) error {
		r.mutex.Lock()
		if len(r.buffer) > 0 && !r.draining {
			r.mutex.Unlock()
			r.drainReadableBuffer(in.ExecutionContext)
			return nil
		} else {
			r.mutex.Unlock()
			return nil
		}
	})
}

//js:method
func (r *ReadableBase) Push(ctx context.Context, chunk *isolates.Value, encoding *BufferEncoding) (bool, error) {
	if chunk.IsNil() {
		r.SetStateConditional(StreamStateReady, StreamStateClosing)
		if len(r.buffer) == 0 && !r.emittedEnd {
			r.emittedEnd = true
			r.SetStateConditional(StreamStateClosing, StreamStateClosed)
			if r.IsPaused() {
				r.Emit(ctx, "readable")
			}
			r.Emit(ctx, "end")
		}
	} else {
		length := 1

		if !r.ReadableObjectMode() {
			if chunk.IsKind(isolates.KindString) {
				if s, err := chunk.StringValue(ctx); err != nil {
					return false, err
				} else {
					length = len(s)
				}
			} else {
				if l, err := chunk.GetByteLength(ctx); err != nil {
					return false, err
				} else {
					length = l
				}
			}
		}

		r.mutex.Lock()
		r.waterMark += length
		r.buffer = append(r.buffer, chunk)

		if r.state == ReadableStreamStateResumed {
			r.mutex.Unlock()
		} else {	
			r.mutex.Unlock()
			r.Emit(ctx, "readable")
		}
	}

	return !r.Closed() && !r.Destroyed() && !r.IsPaused() && r.Errored() == nil && r.ReadableLength() < r.ReadableHighWaterMark(), nil
}

/*****************
 *** io.Reader ***
 *****************/

//js:method
func (c *Reader) Construct(in isolates.FunctionArgs, this Stream, callback *isolates.Value) error {
	bytes := make([]byte, 1024)

	if buffer, err := isolates.For(in.ExecutionContext).Context().Create(in.ExecutionContext, bytes); err != nil {
		return this.EmitError(in.ExecutionContext, err)
	} else {
		c.b = bytes
		c.buffer = buffer
		c.mutex.Unlock()

		if _, err := callback.Call(in.ExecutionContext, this); err != nil {
			return this.EmitError(in.ExecutionContext, err)
		} else {
			return nil
		}
	}
}

//js:method
func (c *Reader) Read(in isolates.FunctionArgs, this Readable) error {
	in.Background(func(in isolates.FunctionArgs) {
		for !this.IsPaused() {
			var n int
			var err error

			c.mutex.Lock()
			buffer := this.UnshiftBuffer()
			if len(buffer) > 0 {
				chunk := buffer[0]
				buffer = buffer[1:]
				this.SetUnshiftBuffer(buffer)

				if arrayBuffer, err := chunk.Get(in.ExecutionContext, "buffer"); err != nil {
					this.EmitError(in.ExecutionContext, err)
					c.mutex.Unlock()
					break
				} else if b, err := arrayBuffer.Bytes(in.ExecutionContext); err != nil {
					this.EmitError(in.ExecutionContext, err)
					c.mutex.Unlock()
					break
				} else {
					n = copy(c.b, b)

					if n < len(b) {
						if surplus, err := in.Context.Create(in.ExecutionContext, b[n:]); err != nil {
							this.EmitError(in.ExecutionContext, err)
							c.mutex.Unlock()
							break
						} else {
							this.SetUnshiftBuffer(append([]*isolates.Value{surplus}, buffer...))
						}
					}
				}
			} else if n, err = c.Reader.Read(c.b); err != nil && n == 0 {
				if err == io.EOF {
					this.Destroy(in.ExecutionContext, nil)
				} else {
					this.EmitError(in.ExecutionContext, err)
				}
				c.mutex.Unlock()
				break
			}

			if arrayBuffer, err := c.buffer.Get(in.ExecutionContext, "buffer"); err != nil {
				this.EmitError(in.ExecutionContext, err)
				c.mutex.Unlock()
				break
			} else if err := arrayBuffer.SetBytes(in.ExecutionContext, c.b); err != nil {
				this.EmitError(in.ExecutionContext, err)
				c.mutex.Unlock()
				break
			} else if slice, err := c.buffer.CallMethod(in.ExecutionContext, "slice", 0, n); err != nil {
				this.EmitError(in.ExecutionContext, err)
				c.mutex.Unlock()
				break
			} else if err := slice.SetBytes(in.ExecutionContext, c.b[:n]); err != nil {
				this.EmitError(in.ExecutionContext, err)
				c.mutex.Unlock()
				break
			} else if cont, err := in.Sync(func(in isolates.FunctionArgs) (any, error) {
				if cont, err := this.Push(in.ExecutionContext, slice, nil); err != nil {
					return false, this.EmitError(in.ExecutionContext, err)
				} else if err := this.Drain(in.ExecutionContext); err != nil {
					return false, err
				} else {
					return cont, nil
				}
			}); err != nil || !cont.(bool) {
				c.mutex.Unlock()
				break
			}

			c.mutex.Unlock()
		}
	})

	return nil
}

/************
 *** Pipe ***
 ************/

//js:method
func (p *PipeBase) Remove(ctx context.Context) error {
	if _, err := p.source.CallMethod(ctx, "unpipe", p.destination); err != nil {
		return err
	} else {
		return nil
	}
}

//js:get
func (p *PipeBase) Source() *isolates.Value {
	return p.source
}

//js:get
func (p *PipeBase) Destination() *isolates.Value {
	return p.destination
}

//js:get
func (p *PipeBase) OnData() *isolates.Value {
	return p.onData
}

//js:get
func (p *PipeBase) OnError() *isolates.Value {
	return p.onError
}

//js:get
func (p *PipeBase) OnEnd() *isolates.Value {
	return p.onEnd
}
