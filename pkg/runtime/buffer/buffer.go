//js:package buffer

package buffer

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"math"
	"reflect"

	"github.com/grexie/isolates"
)

//js:class
type Buffer struct {
	buffer *isolates.Value
}

type bufferStatic struct {
	__bufferStatic bool
}

type bufferModule struct {
	__bufferModule bool
}

//js:constructor Buffer
//js:export Buffer
func NewBuffer(in isolates.FunctionArgs) (*Buffer, error) {
	if global, err := in.Context.Global(in.ExecutionContext); err != nil {
		return nil, err
	} else if prototype, err := in.This.GetPrototype(in.ExecutionContext); err != nil {
		return nil, err
	} else if Uint8Array, err := global.Get(in.ExecutionContext, "Uint8Array"); err != nil {
		return nil, err
	} else if Object, err := global.Get(in.ExecutionContext, "Object"); err != nil {
		return nil, err
	} else if this, err := Uint8Array.New(in.ExecutionContext, in.Args[0]); err != nil {
		return nil, err
	} else if buffer, err := this.Get(in.ExecutionContext, "buffer"); err != nil {
		return nil, err
	} else if _, err := Object.CallMethod(in.ExecutionContext, "setPrototypeOf", this, prototype); err != nil {
		return nil, err
	} else {
		in.ReplaceThis(this)
		return &Buffer{buffer}, nil
	}
}

//js:method createBufferFunction
//js:export-instance Buffer
func createBufferFunction(in isolates.RuntimeFunctionArgs) (*isolates.Value, error) {
	if Buffer, err := in.Exports.Get(in.ExecutionContext, "Buffer"); err != nil {
		return nil, err
	} else if global, err := in.Context.Global(in.ExecutionContext); err != nil {
		return nil, err
	} else if Uint8Array, err := global.Get(in.ExecutionContext, "Uint8Array"); err != nil {
		return nil, err
	} else if Uint8ArrayPrototype, err := Uint8Array.Get(in.ExecutionContext, "prototype"); err != nil {
		return nil, err
	} else if Object, err := global.Get(in.ExecutionContext, "Object"); err != nil {
		return nil, err
	} else if prototype, err := Buffer.Get(in.ExecutionContext, "prototype"); err != nil {
		return nil, err
	} else if _, err := Object.CallMethod(in.ExecutionContext, "setPrototypeOf", prototype, Uint8ArrayPrototype); err != nil {
		return nil, err
	} else if _, err := in.Context.AssignAll(in.ExecutionContext, Buffer, &bufferStatic{}); err != nil {
		return nil, err
	} else if _, err := in.Context.AssignAll(in.ExecutionContext, in.Exports, &bufferModule{}); err != nil {
		return nil, err
	} else {
		return Buffer, nil
	}
}

func (b *Buffer) V8Construct(in isolates.FunctionArgs) (*Buffer, error) {
	return b, nil
}

//js:get buffer
func (b *Buffer) Buffer() *isolates.Value {
	return b.buffer
}

func (b *Buffer) Length(ctx context.Context) (int, error) {
	return b.buffer.GetByteLength(ctx)
}

//js:method
func (b *Buffer) Slice(ctx context.Context, start int, end int) (*Buffer, error) {
	if b1, err := b.buffer.Bytes(ctx); err != nil {
		return nil, err
	} else {
		b2 := make([]byte, end-start)
		copy(b2, b1[start:end])

		if bufferv, err := isolates.For(ctx).Context().Create(ctx, b2); err != nil {
			return nil, err
		} else if bufferv.SetBytes(ctx, b2); err != nil {
			return nil, err
		} else if bufferrv, err := bufferv.Unmarshal(ctx, reflect.TypeOf(&Buffer{})); err != nil {
			return nil, err
		} else {
			buffer := bufferrv.Interface().(*Buffer)

			return buffer, nil
		}
	}
}

//js:method
func (b *Buffer) ToString(ctx context.Context) (string, error) {
	if bytes, err := b.buffer.Bytes(ctx); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (b *bufferStatic) V8Construct(in isolates.FunctionArgs) (*bufferStatic, error) {
	return b, nil
}

//js:method
func (b *bufferStatic) Alloc(ctx context.Context, size int) (*Buffer, error) {
	if buffer, err := isolates.For(ctx).Context().Create(ctx, make([]byte, size)); err != nil {
		return nil, err
	} else {
		return &Buffer{buffer}, nil
	}
}

//js:method
func (b *bufferStatic) ByteLength(ctx context.Context, buffer *Buffer) (int, error) {
	return buffer.Length(ctx)
}

//js:method
func (s *bufferStatic) From(ctx context.Context, args ...any) (*Buffer, error) {
	if args, err := isolates.For(ctx).Context().CreateAll(ctx, args...); err != nil {
		return nil, err
	} else if len(args) >= 1 && (args[0].IsKind(isolates.KindArrayBufferView) || args[0].IsKind(isolates.KindArrayBuffer)) {
		if b, err := args[0].Bytes(ctx); err != nil {
			return nil, err
		} else if buffer, err := isolates.For(ctx).New(NewBuffer, len(b)); err != nil {
			return nil, err
		} else if err := buffer.(*Buffer).buffer.SetBytes(ctx, b); err != nil {
			return nil, err
		} else {
			return buffer.(*Buffer), nil
		}
	} else if len(args) >= 1 && args[0].IsKind(isolates.KindArray) {
		if bytes, err := args[0].Unmarshal(ctx, reflect.TypeOf([]byte{})); err != nil {
			return nil, err
		} else {
			b := bytes.Interface().([]byte)
			if buffer, err := isolates.For(ctx).New(NewBuffer, len(b)); err != nil {
				return nil, err
			} else if err := buffer.(*Buffer).Buffer().SetBytes(ctx, b); err != nil {
				return nil, err
			} else {
				return buffer.(*Buffer), nil
			}
		}
	} else if len(args) >= 1 && args[0].IsKind(isolates.KindString) {
		if s, err := args[0].StringValue(ctx); err != nil {
			return nil, err
		} else if buffer, err := isolates.For(ctx).New(NewBuffer, len([]byte(s))); err != nil {
			return nil, err
		} else if err := buffer.(*Buffer).Buffer().SetBytes(ctx, []byte(s)); err != nil {
			return nil, err
		} else {
			return buffer.(*Buffer), nil
		}
	} else {
		return nil, fmt.Errorf("%s is not a valid argument to Buffer.from", args[0])
	}
}

//js:method isBuffer
func (b *bufferStatic) IsBuffer(ctx context.Context, value any) (bool, error) {
	return false, nil
}

func (b *bufferModule) V8Construct(in isolates.FunctionArgs) (*bufferModule, error) {
	return b, nil
}

//js:get kMaxLength
func (b *bufferModule) MaxLength() int {
	return math.MaxInt
}

//js:get kMaxStringLength
func (b *bufferModule) MaxStringLength() int {
	return math.MaxInt
}

var _ = isolates.AddMarshaller(reflect.TypeOf([]byte{}), func(ctx context.Context, value []byte) (*Buffer, error) {
	if buffer, err := isolates.For(ctx).New(NewBuffer, len(value)); err != nil {
		return nil, err
	} else if err := buffer.(*Buffer).buffer.SetBytes(ctx, value); err != nil {
		return nil, err
	} else {
		return buffer.(*Buffer), nil
	}
})
