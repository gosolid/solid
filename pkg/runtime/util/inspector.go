//js:package util

package util

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"io"

	"github.com/gosolid/solid/pkg/runtime/stream"
	"github.com/grexie/isolates"
)

var _ io.ReadWriteCloser = &Inspector{}

type Inspector struct {
	*stream.DuplexBase
	context   *isolates.Context
	inspector *isolates.Inspector
}

type inspectorStream struct {
	stream    *Inspector
	inspector *isolates.Inspector
	reader    chan []byte
}

//js:constructor Inspector
//js:export Inspector
func NewInspector(in isolates.FunctionArgs) (*Inspector, error) {
	i := &Inspector{
		context: in.Context,
	}

	s := &inspectorStream{
		reader: make(chan []byte),
		stream: i,
	}
	inspector := in.Context.GetIsolate().NewInspector(s)
	s.inspector = inspector
	i.inspector = inspector
	inspector.AddContext(in.Context, "solid")

	if options, err := in.Context.New(in.ExecutionContext, stream.NewReadWriteCloser, s); err != nil {
		return nil, err
	} else if args, err := in.WithArgs(options); err != nil {
		return nil, err
	} else if DuplexBase, err := stream.NewDuplex(args); err != nil {
		return nil, err
	} else {
		i.DuplexBase = DuplexBase

		return i, nil
	}
}

func (i *inspectorStream) Read(b []byte) (int, error) {
	return copy(b, <-i.reader), nil
}

func (i *inspectorStream) Write(b []byte) (int, error) {
	i.inspector.DispatchMessage(string(b))
	return len(b), nil
}

func (i *inspectorStream) Close() error {
	i.inspector.RemoveContext(i.stream.context)
	i.inspector = nil
	return nil
}

func (i *inspectorStream) V8InspectorSendResponse(callId int, message string) {
	i.reader <- []byte(message)
}

func (i *inspectorStream) V8InspectorSendNotification(message string) {
	i.reader <- []byte(message)
}

func (i *inspectorStream) V8InspectorFlushProtocolNotifications() {
	ctx := isolates.WithContext(context.Background())
	isolates.For(ctx).SetContext(i.stream.context)

	i.stream.Emit(ctx, "flush")
}
