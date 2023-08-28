//js:package fs

package fs

import (
	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/grexie/isolates"
)

type ReadStream interface {
	stream.Readable
}

type ReadStreamBase struct {
	*stream.ReadableBase
}

//js:constructor ReadStream
//js:export ReadStream
func NewReadStream(in isolates.FunctionArgs) (*ReadStreamBase, error) {
	if options, err := in.Context.New(in.ExecutionContext, stream.NewReadCloser, in.Args[0]); err != nil {
		return nil, err
	} else if args, err := in.WithArgs(options); err != nil {
		return nil, err
	} else if ReadableBase, err := stream.NewReadable(args); err != nil {
		return nil, err
	} else {
		stream := &ReadStreamBase{
			ReadableBase: ReadableBase,
		}

		return stream, nil
	}
}
