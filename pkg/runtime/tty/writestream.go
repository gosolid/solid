//js:package tty

package tty

import (
	"os"
	"reflect"

	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/grexie/isolates"
	"golang.org/x/term"
)

var _ WriteStream = &WriteStreamBase{}

type WriteStream interface {
	stream.Writable
	IsTTY() bool
	GetWindowSize() ([]int, error)
	Rows() (int, error)
	Columns() (int, error)
}

type WriteStreamBase struct {
	*stream.WritableBase
	fd int
}

//js:constructor
//js:export WriteStream
func NewWriteStream(in isolates.FunctionArgs) (*WriteStreamBase, error) {
	var fd int

	if f, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&os.File{})); err != nil {
		return nil, err
	} else if file, ok := f.Interface().(*os.File); ok {
		if term.IsTerminal(int(file.Fd())) {
			fd = int(file.Fd())
		}
	}

	if options, err := in.Context.New(in.ExecutionContext, stream.NewWriteCloser, in.Arg(in.ExecutionContext, 0)); err != nil {
		return nil, err
	} else if args, err := in.WithArgs(options); err != nil {
		return nil, err
	} else if WritableBase, err := stream.NewWritable(args); err != nil {
		return nil, err
	} else {
		return &WriteStreamBase{
			WritableBase: WritableBase,
			fd:           fd,
		}, nil
	}
}

//js:get isTTY
func (w *WriteStreamBase) IsTTY() bool {
	return true
}

//js:method
func (w *WriteStreamBase) Clear() error {
	if _, err := w.Write([]byte("\033[H\033[2J")); err != nil {
		return err
	} else {
		return nil
	}
}

//js:method
func (w *WriteStreamBase) GetWindowSize() ([]int, error) {
	if width, height, err := term.GetSize(w.fd); err != nil {
		return nil, err
	} else {
		return []int{width, height}, nil
	}
}

//js:get
func (w *WriteStreamBase) Rows() (int, error) {
	if _, height, err := term.GetSize(w.fd); err != nil {
		return 0, err
	} else {
		return height, nil
	}
}

//js:get
func (w *WriteStreamBase) Columns() (int, error) {
	if width, _, err := term.GetSize(w.fd); err != nil {
		return 0, err
	} else {
		return width, nil
	}
}
