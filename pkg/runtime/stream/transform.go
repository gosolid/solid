//js:package stream

package stream

//go:generate go run github.com/grexie/isolates/codegen

import "github.com/grexie/isolates"

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
	if DuplexStream, err := NewDuplex(in); err != nil {
		return nil, err
	} else {
		return &TransformBase{
			DuplexBase: DuplexStream,
		}, nil
	}
}
