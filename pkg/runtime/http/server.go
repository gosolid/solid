//js:package http

package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gosolid/solid/pkg/runtime/net"
	isolates "github.com/grexie/isolates"
)

var notImplemented error = errors.New("method not implemented")

var _ Server = &ServerBase{}

type Server interface {
	net.Server

	CloseAllConnections() error
	CloseIdleConnections() error
	HeadersTimeout() time.Duration
	SetHeadersTimeout(time.Duration)
	MaxHeadersCount() int
	SetMaxHeadersCount(int)
	RequestTimeout() time.Duration
	SetRequestTimeout(time.Duration)
	Timeout() time.Duration
	SetTimeout(ctx context.Context, timeout time.Duration, listener *isolates.Value) error
	MaxRequestsPerSocket() int
	SetMaxRequestsPerSocket(int)
	KeepAliveTimeout() time.Duration
	SetKeepAliveTimeout(time.Duration)
}

type ServerBase struct {
	*net.ServerBase
	http Http

	This    *isolates.Value
	context *isolates.Context

	options              HttpServerOptions
	timeout              time.Duration
	headersTimeout       time.Duration
	requestTimeout       time.Duration
	keepAliveTimeout     time.Duration
	maxHeadersCount      int
	maxRequestsPerSocket int
}

type HttpServerOptions struct {
	ConnectionsCheckingInterval time.Duration `json:"connectionsCheckingInterval"`
	HeadersTimeout              time.Duration `json:"headersTimeout"`
	HighWaterMark               int           `json:"highWaterMark"`
	JoinDuplicateHeaders        bool          `json:"joinDuplicateHeaders"`
	KeepAlive                   bool          `json:"keepAlive"`
	KeepAliveInitialDelay       time.Duration `json:"keepAliveInitialDelay"`
	KeepAliveTimeout            time.Duration `json:"keepAliveTimeout"`
	MaxHeaderSize               int           `json:"maxHeaderSize"`
	NoDelay                     bool          `json:"noDelay"`
	RequestTimeout              time.Duration `json:"requestTimeout"`
	RequireHostHeader           bool          `json:"requireHostHeader"`
	UniqueHeaders               []string      `json:"uniqueHeaders"`
}

var httpServerDefaults, _ = json.Marshal(map[string]any{
	"connectionsCheckingInterval": 30000 * time.Millisecond,
	"headersTimeout":              60000 * time.Millisecond,
	"highWaterMark":               16384,
	"joinDuplicateHeaders":        false,
	"keepAlive":                   false,
	"keepAliveInitialDelay":       0,
	"keepAliveTimeout":            5000 * time.Millisecond,
	"maxHeaderSize":               16384,
	"noDelay":                     true,
	"requestTimeout":              300000 * time.Millisecond,
	"requireHostHeader":           true,
	"uniqueHeaders":               []string{},
})

func newServer(in isolates.FunctionArgs, http Http) (*ServerBase, error) {
	var options *isolates.Value
	var listener *isolates.Value

	if len(in.Args) > 0 && in.Args[0].IsKind(isolates.KindObject) && !in.Args[0].IsKind(isolates.KindFunction) {
		options = in.Arg(in.ExecutionContext, 0)
	}

	if len(in.Args) > 0 && in.Args[len(in.Args)-1].IsKind(isolates.KindFunction) {
		listener = in.Arg(in.ExecutionContext, len(in.Args)-1)
	}

	if netServerBase, err := http.Net().Server().Call(in.ExecutionContext, in.This); err != nil {
		return nil, err
	} else {
		server := &ServerBase{
			ServerBase: netServerBase.Receiver(in.ExecutionContext).Interface().(*net.ServerBase),
			http:       http,
			This:       in.This,
			context:    in.Context,
		}

		json.Unmarshal(httpServerDefaults, &server.options)
		if options != nil && options.IsKind(isolates.KindObject) {
			in.Context.Assign(in.ExecutionContext, server.options, options)
		}

		if listener != nil && listener.IsKind(isolates.KindFunction) {
			server.AddListener(in.ExecutionContext, "request", listener)
		}

		if onConnection, err := in.Context.Create(in.ExecutionContext, server.onConnection); err != nil {
			return nil, err
		} else if err := server.AddListener(in.ExecutionContext, "connection", onConnection); err != nil {
			return nil, err
		}

		return server, nil
	}
}

func (s *ServerBase) onConnection(in isolates.FunctionArgs) (*isolates.Value, error) {
	socket := in.Args[0].Receiver(in.ExecutionContext).Interface().(net.Socket)

	in.Background(func(in isolates.FunctionArgs) {
		defer socket.Destroy(in.ExecutionContext, nil)

		for {
			if req, err := s.context.New(in.ExecutionContext, NewServerRequest, socket); err != nil {
				if err == io.EOF {
					break
				} else {
					s.EmitError(in.ExecutionContext, err)
					return
				}
			} else if req, ok := req.(ServerRequest); !ok {
				s.EmitError(in.ExecutionContext, errors.New("invalid server request"))
				return
			} else if req.IsUpgrade() {
				s.Emit(in.ExecutionContext, "upgrade", req, socket, []byte{})
				return
			} else if res, err := req.Response(in.ExecutionContext); err != nil {
				s.EmitError(in.ExecutionContext, err)
				return
			} else if _, err := s.context.Create(in.ExecutionContext, func(in isolates.FunctionArgs) (*isolates.Value, error) {
				var err error
				if len(in.Args) > 0 {
					err = in.Arg(in.ExecutionContext, 0).ToError(in.ExecutionContext)
				}

				if err == nil {
					res.WriteHead(404, "Not Found", http.Header{"Content-Type": {"text/plain"}})
					res.Write([]byte("404 Not Found\n"))
					res.End(in.ExecutionContext)
				} else {
					res.WriteHead(500, "Internal Server Error", http.Header{"Content-Type": {"text/plain"}})
					res.Write([]byte(fmt.Sprintf("500 Internal Server Error\n\n%s", err)))
					res.End(in.ExecutionContext)
				}
				return nil, nil
			}); err != nil {
				s.EmitError(in.ExecutionContext, err)
				return
			} else {
				in.Background(func(in isolates.FunctionArgs) {
					if err := s.Emit(in.ExecutionContext, "request", req, res); err != nil {
						s.EmitError(in.ExecutionContext, err)
					}
				})

				if err := res.Wait(); err != nil {
					s.EmitError(in.ExecutionContext, err)
				}
			}
		}
	})

	return nil, nil
}

//js:method
func (l *ServerBase) CloseAllConnections() error {
	return notImplemented
}

//js:method
func (l *ServerBase) CloseIdleConnections() error {
	return notImplemented
}

//js:get headersTimeout
func (l *ServerBase) HeadersTimeout() time.Duration {
	return l.headersTimeout
}

//js:set headersTimeout
func (l *ServerBase) SetHeadersTimeout(headersTimeout time.Duration) {
	l.headersTimeout = headersTimeout
}

//js:get maxHeadersCount
func (l *ServerBase) MaxHeadersCount() int {
	return l.maxHeadersCount
}

//js:set maxHeadersCount
func (l *ServerBase) SetMaxHeadersCount(maxHeadersCount int) {
	l.maxHeadersCount = maxHeadersCount
}

//js:get requestTimeout
func (l *ServerBase) RequestTimeout() time.Duration {
	return l.requestTimeout
}

//js:set requestTimeout
func (l *ServerBase) SetRequestTimeout(requestTimeout time.Duration) {
	l.requestTimeout = requestTimeout
}

//js:get timeout
func (l *ServerBase) Timeout() time.Duration {
	return l.timeout
}

//js:method setTimeout
func (l *ServerBase) SetTimeout(ctx context.Context, timeout time.Duration, listener *isolates.Value) error {
	l.timeout = timeout

	if listener != nil {
		return l.AddListener(ctx, "timeout", listener)
	} else {
		return nil
	}
}

//js:get maxRequestsPerSocket
func (l *ServerBase) MaxRequestsPerSocket() int {
	return l.maxRequestsPerSocket
}

//js:set maxRequestsPerSocket
func (l *ServerBase) SetMaxRequestsPerSocket(maxRequestsPerSocket int) {
	l.maxRequestsPerSocket = maxRequestsPerSocket
}

//js:get keepAliveTimeout
func (l *ServerBase) KeepAliveTimeout() time.Duration {
	return l.keepAliveTimeout
}

//js:set keepAliveTimeout
func (l *ServerBase) SetKeepAliveTimeout(keepAliveTimeout time.Duration) {
	l.keepAliveTimeout = keepAliveTimeout
}
