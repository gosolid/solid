//js:package tty

package tty

import (
	"os"
	"reflect"

	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/grexie/isolates"
	"golang.org/x/term"
)

var _ ReadStream = &ReadStreamBase{}

type ReadStream interface {
	stream.Readable
	IsRaw() bool
	IsTTY() bool
	SetRawMode(in isolates.FunctionArgs, mode bool) (*isolates.Value, error)
}

type ReadStreamBase struct {
	*stream.ReadableBase
	fd    int
	state *term.State
}

//js:constructor
//js:export ReadStream
func NewReadStream(in isolates.FunctionArgs) (*ReadStreamBase, error) {
	var fd int

	if f, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&os.File{})); err != nil {
		return nil, err
	} else if file, ok := f.Interface().(*os.File); ok {
		if term.IsTerminal(int(file.Fd())) {
			fd = int(file.Fd())
		}
	}

	if options, err := in.Context.New(in.ExecutionContext, stream.NewReadCloser, in.Arg(in.ExecutionContext, 0)); err != nil {
		return nil, err
	} else if args, err := in.WithArgs(options); err != nil {
		return nil, err
	} else if ReadableBase, err := stream.NewReadable(args); err != nil {
		return nil, err
	} else {
		self := in.This

		in.Context.AddMicrotask(in.ExecutionContext, func(in isolates.FunctionArgs) error {
			if _, err := self.CallMethod(in.ExecutionContext, "resume"); err != nil {
				return err
			} else {
				return nil
			}
		})

		return &ReadStreamBase{
			ReadableBase: ReadableBase,
			fd:           fd,
		}, nil
	}
}

//js:get isRaw
func (r *ReadStreamBase) IsRaw() bool {
	return r.state != nil
}

//js:get isTTY
func (r *ReadStreamBase) IsTTY() bool {
	return true
}

//js:method
func (r *ReadStreamBase) SetRawMode(in isolates.FunctionArgs, mode bool) (*isolates.Value, error) {
	if r.state == nil && mode {
		var err error
		if r.state, err = term.MakeRaw(r.fd); err != nil {
			return nil, err
		}
	} else if r.state != nil && !mode {
		if err := term.Restore(r.fd, r.state); err != nil {
			return nil, err
		}
		r.state = nil
	}

	return in.This, nil
}
