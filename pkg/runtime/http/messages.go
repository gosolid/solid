//js:package http

package http

//go:generate go run github.com/grexie/isolates/codegen

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gosolid/solid/pkg/runtime/events"
	"github.com/gosolid/solid/pkg/runtime/net"
	"github.com/gosolid/solid/pkg/runtime/stream"
	isolates "github.com/grexie/isolates"
)

var _ IncomingMessage = &IncomingMessageBase{}
var _ ServerRequest = &ServerRequestBase{}
var _ ClientResponse = &ClientResponseBase{}
var _ OutgoingMessage = &OutgoingMessageBase{}
var _ ServerResponse = &ServerResponseBase{}
var _ ClientRequest = &ClientRequestBase{}

type IncomingMessageBase struct {
	*stream.ReadableBase
	context *isolates.Context

	timeout time.Duration

	onTimeout events.EventHandler

	header     http.Header
	trailer    http.Header
	protoMajor int
	protoMinor int

	complete bool
}

type ServerRequestBase struct {
	*IncomingMessageBase
	*events.EventEmitterBase

	parsed   bool
	socket   net.Socket
	request  *http.Request
	response ServerResponse
}

type ClientResponseBase struct {
	*IncomingMessageBase

	parsed   bool
	response *http.Response
	request  ClientRequest
}

type OutgoingMessageBase struct {
	*stream.WritableBase
	context isolates.Context
	mutex   sync.Mutex

	headersWriter  io.WriteCloser
	trailersWriter io.WriteCloser
	finished       chan error
	isFinished     bool

	header     http.Header
	trailer    http.Header
	protoMajor int
	protoMinor int

	headersSent bool
}

type ServerResponseBase struct {
	*OutgoingMessageBase

	request ServerRequest

	statusCode int
	status     string
}

type ClientRequestBase struct {
	*OutgoingMessageBase

	socket   net.Socket
	response ClientResponse

	host   string
	path   string
	method string
}

var _ io.ReadCloser = &ServerRequestBase{}
var _ io.WriteCloser = &ServerResponseBase{}

type readCloser struct {
	io.ReadCloser
}

//js:constructor IncomingMessage
//js:export IncomingMessage
func NewIncomingMessage(in isolates.FunctionArgs) (*IncomingMessageBase, error) {
	if ReadableBase, err := stream.NewReadable(in); err != nil {
		return nil, err
	} else {
		m := &IncomingMessageBase{
			ReadableBase: ReadableBase,
			context:      in.Context,
		}

		return m, nil
	}
}

type messageReader struct {
	sync.Mutex
	onDataCh       chan bool
	context        *isolates.Context
	bufioReader    *bufio.Reader
	socket         stream.Duplex
	chunks         [][]byte
	onDataHandler  *isolates.Value
	onCloseHandler *isolates.Value
}

func newMessageReader(ctx context.Context, socket stream.Duplex) (*messageReader, error) {
	reader := &messageReader{
		onDataCh: make(chan bool),
		context:  isolates.For(ctx).Context(),
		socket:   socket,
		chunks:   [][]byte{},
	}

	reader.bufioReader = bufio.NewReader(reader)

	if onData, err := isolates.For(ctx).Context().Create(ctx, reader.onData); err != nil {
		return nil, err
	} else if err := socket.AddListener(ctx, "data", onData); err != nil {
		return nil, err
	} else if onClose, err := isolates.For(ctx).Context().Create(ctx, reader.onClose); err != nil {
		return nil, err
	} else if err := socket.AddListener(ctx, "close", onClose); err != nil {
		return nil, err
	} else {
		reader.onDataHandler = onData
		reader.onCloseHandler = onClose
		return reader, nil
	}
}

func (r *messageReader) BufioReader() *bufio.Reader {
	return r.bufioReader
}

func (r *messageReader) Finish(ctx context.Context) error {
	if err := r.socket.RemoveListener(ctx, "data", r.onDataHandler); err != nil {
		return err
	} else if err := r.socket.RemoveListener(ctx, "close", r.onCloseHandler); err != nil {
		return err
	}

	n := r.bufioReader.Buffered()
	b := make([]byte, n)
	n, _ = r.bufioReader.Read(b)
	chunks := append([][]byte{b[:n]}, r.chunks...)

	for i := len(chunks) - 1; i >= 0; i-- {
		if buffer, err := isolates.For(ctx).Context().Create(ctx, chunks[i]); err != nil {
			return err
		} else if err := r.socket.Unshift(ctx, buffer, nil); err != nil {
			return err
		}
	}

	return nil
}

func (r *messageReader) onClose(in isolates.FunctionArgs) (*isolates.Value, error) {
	r.onDataCh <- true
	return nil, nil
}

func (r *messageReader) onData(in isolates.FunctionArgs) (*isolates.Value, error) {
	if buffer, err := in.Arg(in.ExecutionContext, 0).Get(in.ExecutionContext, "buffer"); err != nil {
		return nil, err
	} else if b, err := buffer.Bytes(in.ExecutionContext); err != nil {
		return nil, err
	} else {
		if len(b) > 0 {
			r.Lock()
			r.chunks = append(r.chunks, b)
			r.socket.Pause(in.ExecutionContext)
			r.Unlock()
			go func() { r.onDataCh <- false }()
		}
		return nil, nil
	}
}

func (r *messageReader) Read(b []byte) (int, error) {
	r.Lock()
	r.socket.Resume(r.context.GetIsolate().GetExecutionContext())
	if r.socket.Closed() {
		r.Unlock()
		return 0, io.EOF
	}
	r.Unlock()

	closed := <-r.onDataCh

	if closed {
		return 0, io.EOF
	}

	r.Lock()
	defer r.Unlock()

	if len(r.chunks) == 0 {
		return 0, nil
	} else {
		chunk := r.chunks[0]
		n := copy(b, chunk)
		if n < len(chunk) {
			r.chunks = append([][]byte{chunk[n:]}, r.chunks[1:]...)
		} else {
			r.chunks = r.chunks[1:]
		}
		return n, nil
	}
}

//js:constructor ServerRequest
//js:export ServerRequest
func NewServerRequest(in isolates.FunctionArgs) (*ServerRequestBase, error) {
	if socket, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(stream.Duplex); !ok {
		return nil, fmt.Errorf("ServerRequest constructor requires a duplex stream / socket as first argument")
	} else if reader, err := newMessageReader(in.ExecutionContext, socket); err != nil {
		return nil, err
	} else {
		req := &ServerRequestBase{EventEmitterBase: events.NewEventEmitter(in), socket: socket}

		in.Background(func(in isolates.FunctionArgs) {
			if request, err := http.ReadRequest(reader.BufioReader()); err != nil {
				req.EmitError(in.ExecutionContext, err)
			} else if err := reader.Finish(in.ExecutionContext); err != nil {
				req.EmitError(in.ExecutionContext, err)
			} else if options, err := in.Context.New(in.ExecutionContext, stream.NewReadCloser, &readCloser{request.Body}); err != nil {
				req.EmitError(in.ExecutionContext, err)
			} else if args, err := in.WithArgs(options); err != nil {
				req.EmitError(in.ExecutionContext, err)
			} else if IncomingMessageBase, err := NewIncomingMessage(args); err != nil {
				req.EmitError(in.ExecutionContext, err)
			} else {
				IncomingMessageBase.StreamBase.EventEmitterBase = req.EventEmitterBase
				req.IncomingMessageBase = IncomingMessageBase
				req.request = request
				req.header = request.Header.Clone()
				req.protoMajor = request.ProtoMajor
				req.protoMinor = request.ProtoMinor
				req.parsed = true
				req.Emit(in.ExecutionContext, "parsed")
			}
		})

		return req, nil
	}
}

//js:constructor ClientResponse
//js:export ClientResponse
func NewClientResponse(in isolates.FunctionArgs) (*ClientResponseBase, error) {
	if request, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(*ClientRequestBase); !ok {
		return nil, fmt.Errorf("ClientResponse constructor requires a request as first argument")
	} else if reader, err := newMessageReader(in.ExecutionContext, request.socket); err != nil {
		return nil, err
	} else {
		res := &ClientResponseBase{request: request}

		in.Background(func(in isolates.FunctionArgs) {
			if response, err := http.ReadResponse(reader.BufioReader(), nil); err != nil {
				res.EmitError(in.ExecutionContext, err)
				// } else if err := reader.Finish(in.ExecutionContext); err != nil {
				// 	res.EmitError(in.ExecutionContext, err)
			} else if options, err := in.Context.New(in.ExecutionContext, stream.NewReadCloser, &readCloser{response.Body}); err != nil {
				res.EmitError(in.ExecutionContext, err)
			} else if args, err := in.WithArgs(options); err != nil {
				res.EmitError(in.ExecutionContext, err)
			} else if IncomingMessageBase, err := NewIncomingMessage(args); err != nil {
				res.EmitError(in.ExecutionContext, err)
			} else {
				res.IncomingMessageBase = IncomingMessageBase
				res.response = response
				res.header = response.Header.Clone()
				res.protoMajor = response.ProtoMajor
				res.protoMinor = response.ProtoMinor
				request.Emit(in.ExecutionContext, "response", res)
			}
		})

		return res, nil
	}

}

type writeCloserWithSignal struct {
	io.WriteCloser
	sync.Mutex
	writing bool
	signal  chan bool
}

func newWriteCloserWithSignal(writeCloser io.WriteCloser, signal chan bool) io.WriteCloser {
	return &writeCloserWithSignal{
		WriteCloser: writeCloser,
		signal:      signal,
	}
}

func (w *writeCloserWithSignal) Write(b []byte) (int, error) {
	w.Lock()
	if !w.writing {
		w.writing = true
		w.Unlock()
		go func() { w.signal <- true }()
	} else {
		w.Unlock()
	}
	return w.WriteCloser.Write(b)
}

//js:constructor OutgoingMessage
//js:export OutgoingMessage
func NewOutgoingMessage(in isolates.FunctionArgs) (*OutgoingMessageBase, error) {
	if writable, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(stream.Writable); !ok {
		return nil, fmt.Errorf("OutgoingMessage constructor requires a stream.Writable as first argument")
	} else {
		body, bodyWriter := io.Pipe()
		headers, headersWriter := io.Pipe()
		trailers, trailersWriter := io.Pipe()
		finished := make(chan error)

		trailersWriter.Close()

		writing := make(chan bool)

		if options, err := in.Context.New(in.ExecutionContext, stream.NewWriteCloser, newWriteCloserWithSignal(bodyWriter, writing)); err != nil {
			return nil, err
		} else if args, err := in.WithArgs(options); err != nil {
			return nil, err
		} else if WritableBase, err := stream.NewWritable(args); err != nil {
			return nil, err
		} else {
			msg := &OutgoingMessageBase{
				WritableBase:   WritableBase,
				headersWriter:  headersWriter,
				trailersWriter: trailersWriter,
				finished:       finished,
				header:         http.Header{},
				trailer:        http.Header{},
			}

			in.Background(func(in isolates.FunctionArgs) {
				writer := bufio.NewWriterSize(writable, 10240)
				_, err := io.Copy(writer, io.MultiReader(headers, body, trailers))
				writer.Flush()
				writable.End(in.ExecutionContext, func(in isolates.FunctionArgs) (*isolates.Value, error) {
					msg.isFinished = true
					finished <- err
					return nil, nil
				})
			})

			in.Background(func(in isolates.FunctionArgs) {
				if <-writing {
					in.Background(func(in isolates.FunctionArgs) {
						if _, err := msg.This.CallMethod(in.ExecutionContext, "flushHeaders"); err != nil {
							isolates.For(in.ExecutionContext).Debug(err)
						}
					})
				}
			})

			return msg, nil
		}
	}
}

//js:constructor ServerResponse
//js:export ServerResponse
func NewServerResponse(in isolates.FunctionArgs) (*ServerResponseBase, error) {
	if request, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(*ServerRequestBase); !ok {
		return nil, fmt.Errorf("ServerResponse constructor requires a request as first argument")
	} else {
		if args, err := in.WithArgs(request.Socket()); err != nil {
			return nil, err
		} else if OutgoingMessageBase, err := NewOutgoingMessage(args); err != nil {
			return nil, err
		} else {
			res := &ServerResponseBase{
				OutgoingMessageBase: OutgoingMessageBase,
				request:             request,
			}

			res.protoMajor = request.protoMajor
			res.protoMinor = request.protoMinor
			res.statusCode = 200
			res.status = "OK"

			return res, nil
		}
	}
}

//js:constructor ClientRequest
//js:export ClientRequest
func NewClientRequest(in isolates.FunctionArgs) (*ClientRequestBase, error) {
	if socket, ok := in.Arg(in.ExecutionContext, 0).Receiver(in.ExecutionContext).Interface().(net.Socket); !ok {
		return nil, fmt.Errorf("ClientRequest constructor requires a socket as first argument")
	} else if u, err := in.Arg(in.ExecutionContext, 1).StringValue(in.ExecutionContext); err != nil {
		return nil, err
	} else if urlParsed, err := url.Parse(u); err != nil {
		return nil, err
	} else {
		if args, err := in.WithArgs(socket); err != nil {
			return nil, err
		} else if OutgoingMessageBase, err := NewOutgoingMessage(args); err != nil {
			return nil, err
		} else {
			req := &ClientRequestBase{
				OutgoingMessageBase: OutgoingMessageBase,
				socket:              socket,
			}

			req.protoMajor = 1
			req.protoMinor = 1
			req.host = urlParsed.Host
			req.path = urlParsed.RequestURI()
			req.method = "GET"

			return req, nil
		}
	}
}

/***********************
 *** IncomingMessage ***
 ***********************/

//js:get complete
func (r *IncomingMessageBase) Complete() bool {
	return r.complete
}

//js:get headers
func (r *IncomingMessageBase) Headers() map[string]any {
	out := map[string]any{}

	for k, v := range r.header {
		out[strings.ToLower(k)] = v[0]
	}

	return out
}

//js:get headersDistinct
func (r *IncomingMessageBase) HeadersDistinct() http.Header {
	return r.header
}

//js:get httpVersion
func (r *IncomingMessageBase) HttpVersion() string {
	return fmt.Sprintf("%d.%d", r.protoMajor, r.protoMinor)
}

//js:get rawHeaders
func (r *IncomingMessageBase) RawHeaders() []string {
	return nil
}

//js:get rawTrailers
func (r *IncomingMessageBase) RawTrailers() []string {
	return nil
}

//js:method
func (r *IncomingMessageBase) SetTimeout(ctx context.Context, timeout time.Duration, listener *isolates.Value) error {
	r.timeout = timeout

	if listener != nil {
		return r.AddListener(ctx, "timeout", listener)
	} else {
		return nil
	}
}

//js:get trailers
func (r *IncomingMessageBase) Trailers() http.Header {
	return r.trailer
}

//js:get trailersDistinct
func (r *IncomingMessageBase) TrailersDistinct() http.Header {
	return r.trailer
}

/***********************
 *** OutgoingMessage ***
 ***********************/

func (r *OutgoingMessageBase) FlushHeaders() error {
	return errors.New("not implemented")
}

func (r *OutgoingMessageBase) WriteHead(statusCode int, statusMessage string, headers http.Header) error {
	return errors.New("not implemented")
}

//js:method
func (r *OutgoingMessageBase) AddTrailers(headers http.Header) {
	for k, values := range headers {
		for _, v := range values {
			r.trailer.Add(k, v)
		}
	}
}

//js:method
func (r *OutgoingMessageBase) AppendHeader(name string, values []string) {
	for _, v := range values {
		r.header.Add(name, v)
	}
}

//js:get destroyed
func (r *OutgoingMessageBase) Destroyed() bool {
	return false
}

//js:get finished
func (r *OutgoingMessageBase) Finished() bool {
	return r.isFinished
}

//js:method
func (r *OutgoingMessageBase) GetHeader(name string) []string {
	return r.header.Values(name)
}

//js:method
func (r *OutgoingMessageBase) GetHeaderNames() []string {
	names := []string{}

	for name := range r.header {
		names = append(names, name)
	}

	return names
}

//js:method
func (r *OutgoingMessageBase) GetRawHeaderNames() []string {
	return nil
}

//js:get headers
func (r *OutgoingMessageBase) Headers() map[string]any {
	out := map[string]any{}

	for k, v := range r.header {
		out[strings.ToLower(k)] = v[0]
	}

	return out
}

//js:method
func (r *OutgoingMessageBase) HasHeader(name string) bool {
	return len(r.header.Values(name)) > 0
}

//js:get headersSent
func (r *OutgoingMessageBase) HeadersSent() bool {
	return r.headersSent
}

//js:method
func (r *OutgoingMessageBase) RemoveHeader(name string) {
	r.header.Del(name)
}

//js:method
func (r *OutgoingMessageBase) SetHeader(ctx context.Context, name string, value *isolates.Value) error {
	if value.IsKind(isolates.KindString) {
		if v, err := value.StringValue(ctx); err != nil {
			return err
		} else {
			r.header.Set(name, v)
			return nil
		}
	} else if value.IsKind(isolates.KindNumber) {
		if v, err := value.StringValue(ctx); err != nil {
			return err
		} else {
			r.header.Set(name, v)
			return nil
		}
	} else if value.IsKind(isolates.KindArray) {
		r.header.Del(name)

		if v, err := value.Unmarshal(ctx, reflect.TypeOf([]string{})); err != nil {
			return err
		} else {
			for _, v := range v.Interface().([]string) {
				r.header.Add(name, v)
			}
		}

		return nil
	} else {
		return fmt.Errorf("invalid header value")
	}
}

//js:method
func (r *OutgoingMessageBase) SetHeaders(headers http.Header) {
	for name, values := range headers {
		r.header.Del(name)
		for _, v := range values {
			r.header.Add(name, v)
		}
	}
}

//js:method
//ts:setTimeout(timeout: number, listener: () => void): void
func (r *OutgoingMessageBase) SetTimeout(timeout time.Duration, listener func() error) func() {
	return nil
}

//js:get writableEnded
func (r *OutgoingMessageBase) WritableEnded() bool {
	return false
}

//js:get writableFinished
func (r *OutgoingMessageBase) WritableFinished() bool {
	return false
}

//js:get sendDate
func (r *OutgoingMessageBase) SendDate() bool {
	return false
}

//js:set sendDate
func (r *OutgoingMessageBase) SetSendDate(bool) {

}

//js:get strictContentLength
func (r *OutgoingMessageBase) StrictContentLength() bool {
	return false
}

//js:set strictContentLength
func (r *OutgoingMessageBase) SetStrictContentLength(bool) {

}

//js:method
func (r *OutgoingMessageBase) WriteContinue() error {
	return notImplemented
}

//js:method
func (r *OutgoingMessageBase) WriteEarlyHints(header http.Header) error {
	return notImplemented
}

//js:method
func (r *OutgoingMessageBase) WriteProcessing() error {
	return notImplemented
}

//js:async-method
func (r *OutgoingMessageBase) Wait() error {
	return <-r.finished
}

/*********************
 *** ServerRequest ***
 *********************/

func (r *ServerRequestBase) Response(ctx context.Context) (ServerResponse, error) {
	if r.response == nil {
		if response, err := r.context.New(ctx, NewServerResponse, r); err != nil {
			return nil, err
		} else {
			response := response.(*ServerResponseBase)
			r.response = response
		}
	}

	return r.response, nil
}

//js:get socket
func (r *ServerRequestBase) Socket() net.Socket {
	return r.socket
}

func (r *ServerRequestBase) IsUpgrade() bool {
	return len(r.header.Values("Upgrade")) > 0
}

//js:get method
func (r *ServerRequestBase) Method() string {
	return r.request.Method
}

//js:get url
func (r *ServerRequestBase) URL() string {
	return r.request.URL.String()
}

/**********************
 *** ClientResponse ***
 **********************/

//js:get statusCode
func (r *ClientResponseBase) StatusCode() int {
	return r.response.StatusCode
}

//js:get statusMessage
func (r *ClientResponseBase) StatusMessage() string {
	return strings.TrimPrefix(r.response.Status, fmt.Sprintf("%d ", r.StatusCode()))
}

/**********************
 *** ServerResponse ***
 **********************/

func (r *ServerResponseBase) Request(context.Context) (ServerRequest, error) {
	return r.request, nil
}

//js:get socket
func (r *ServerResponseBase) Socket() net.Socket {
	return r.request.Socket()
}

//js:method
func (r *ServerResponseBase) FlushHeaders() error {
	r.mutex.Lock()

	if r.headersSent {
		r.mutex.Unlock()
		return nil
	}

	r.headersSent = true
	r.mutex.Unlock()

	if _, err := r.headersWriter.Write([]byte(fmt.Sprintf("HTTP/%d.%d %d %s\r\n", r.protoMajor, r.protoMinor, r.statusCode, r.status))); err != nil {
		return err
	}
	if err := r.header.Write(r.headersWriter); err != nil {
		return err
	}
	if _, err := r.headersWriter.Write([]byte("\r\n")); err != nil {
		return err
	}
	if err := r.headersWriter.Close(); err != nil {
		return err
	}

	return nil
}

//js:method
func (r *ServerResponseBase) WriteHead(statusCode int, statusMessage string, headers http.Header) error {
	if r.headersSent {
		return fmt.Errorf("headers already written")
	}

	r.SetStatusCode(statusCode)
	r.SetStatusMessage(statusMessage)
	r.SetHeaders(headers)
	return r.FlushHeaders()
}

//js:get statusCode
func (r *ServerResponseBase) StatusCode() int {
	return r.statusCode
}

//js:set statusCode
func (r *ServerResponseBase) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}

//js:get statusMessage
func (r *ServerResponseBase) StatusMessage() string {
	return r.status
}

//js:set statusMessage
func (r *ServerResponseBase) SetStatusMessage(statusMessage string) {
	r.status = statusMessage
}

/*********************
 *** ClientRequest ***
 *********************/

//js:get
func (r *ClientRequestBase) Socket() net.Socket {
	return r.socket
}

//js:method
func (r *ClientRequestBase) FlushHeaders() error {
	r.mutex.Lock()

	if r.headersSent {
		r.mutex.Unlock()
		return nil
	}

	r.headersSent = true
	r.mutex.Unlock()

	if _, err := r.headersWriter.Write([]byte(fmt.Sprintf("%s %s HTTP/%d.%d\r\n", r.method, r.path, r.protoMajor, r.protoMinor))); err != nil {
		return err
	}
	if err := r.header.Write(r.headersWriter); err != nil {
		return err
	}
	if _, err := r.headersWriter.Write([]byte("\r\n")); err != nil {
		return err
	}
	if err := r.headersWriter.Close(); err != nil {
		return err
	}

	return nil
}

//js:get host
func (r *ClientRequestBase) Host() string {
	return r.host
}

//js:get maxHeadersCount
func (r *ClientRequestBase) MaxHeadersCount() int {
	return 0
}

//js:get method
func (r *ClientRequestBase) Method() string {
	return r.method
}

//js:get path
func (r *ClientRequestBase) Path() string {
	return r.path
}

//js:get protocol
func (r *ClientRequestBase) Protocol() string {
	return ""
}

//js:get reusedSocket
func (r *ClientRequestBase) ReusedSocket() bool {
	return false
}

//js:set host
func (r *ClientRequestBase) SetHost(host string) {
	r.host = host
}

//js:set maxHeadersCount
func (r *ClientRequestBase) SetMaxHeadersCount(int) {

}

//js:set method
func (r *ClientRequestBase) SetMethod(method string) {
	r.method = method
}

//js:method setNoDelay
func (r *ClientRequestBase) SetNoDelay(bool) {

}

//js:set path
func (r *ClientRequestBase) SetPath(path string) {
	r.path = path
}

//js:set protocol
func (r *ClientRequestBase) SetProtocol(string) {

}

//js:method setSocketKeepAlive
func (r *ClientRequestBase) SetSocketKeepAlive(bool, time.Duration) {

}
