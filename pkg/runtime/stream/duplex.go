//js:package stream

package stream

import (
	isolates "github.com/grexie/isolates"
)

var _ Duplex = &DuplexBase{}

type Duplex interface {
	Readable
	Writable
}

type DuplexBase struct {
	*StreamBase
	*WritableBase
	*ReadableBase
}

type ReadWriteCloser struct {
	*Closer
	*Reader
	*Writer
}

//js:constructor Duplex
//js:export Duplex
func NewDuplex(in isolates.FunctionArgs) (*DuplexBase, error) {
	if StreamBase, err := NewStream(in); err != nil {
		return nil, err
	} else if ReadableBase, err := NewReadableWithStream(in, StreamBase); err != nil {
		return nil, err
	} else if WritableBase, err := NewWritableWithStream(in, StreamBase); err != nil {
		return nil, err
	} else {
		return &DuplexBase{
			StreamBase:   StreamBase,
			ReadableBase: ReadableBase,
			WritableBase: WritableBase,
		}, nil
	}
}

func NewReadWriteCloser(in isolates.FunctionArgs) (*ReadWriteCloser, error) {
	if closer, err := NewCloser(in); err != nil {
		return nil, err
	} else if reader, err := NewReader(in); err != nil {
		return nil, err
	} else if writer, err := NewWriter(in); err != nil {
		return nil, err
	} else {
		// in.This.RebindMethod(in.ExecutionContext, "final")

		return &ReadWriteCloser{Closer: closer, Reader: reader, Writer: writer}, nil
	}
}

// //js:method
// func (c *ReadWriteCloser) Destroy(in isolates.FunctionArgs, this Stream, err *isolates.Value, callback *isolates.Value) error {
// 	return c.Closer.Destroy(in, this, err, callback)
// }

//js:method
func (c *ReadWriteCloser) Final(in isolates.FunctionArgs, this Writable, callback *isolates.Value) error {
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
