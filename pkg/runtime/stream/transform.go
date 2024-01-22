//js:package stream

package stream

//go:generate go run github.com/grexie/isolates/codegen

import (
	"fmt"
	"log"

	"github.com/grexie/isolates"
)

var _ Transform = &TransformBase{}

//js:alias TransformBase
type Transform interface {
	Duplex
}

type TransformBase struct {
	*DuplexBase
}

//js:constructor Transform
//js:export Transform
func NewTransform(in isolates.FunctionArgs) (*TransformBase, error) {
	TransformStream := &TransformBase{}

	if options, err := in.Context.Create(in.ExecutionContext, map[string]any{}); err != nil {
		return nil, err
	} else if err := options.Set(in.ExecutionContext, "read", TransformStream.read); err != nil {
		return nil, err
	} else if err := options.Set(in.ExecutionContext, "write", TransformStream.write); err != nil {
		return nil, err
	} else if args, err := in.WithArgs(options); err != nil {
		return nil, err
	} else if DuplexStream, err := NewDuplex(args); err != nil {
		return nil, err
	} else {
		if len(in.Args) > 0 && !in.Args[0].IsNil() {
			options := in.Args[0]

			if transform, err := options.Get(in.ExecutionContext, "transform"); err != nil {
				return nil, err
			} else if transform != nil {
				if transform.IsKind(isolates.KindFunction) {
					if err := in.This.Set(in.ExecutionContext, "_transform", transform); err != nil {
						return nil, err
					}
				}
			}

			if flush, err := options.Get(in.ExecutionContext, "flush"); err != nil {
				return nil, err
			} else if flush != nil {
				if flush.IsKind(isolates.KindFunction) {
					if err := in.This.Set(in.ExecutionContext, "_flush", flush); err != nil {
						return nil, err
					}
				}
			}
		}

		TransformStream.DuplexBase = DuplexStream
		return TransformStream, nil
	}
}

func (t *TransformBase) read(in isolates.FunctionArgs) (*isolates.Value, error) {
	log.Println("READ NOT IMPLEMENTED")
	return nil, fmt.Errorf("read not implemented")
}

func (t *TransformBase) write(in isolates.FunctionArgs) (*isolates.Value, error) {
	log.Println("WRITE NOT IMPLEMENTED")
	return nil, fmt.Errorf("write not implemented")
}
