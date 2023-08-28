//js:package http
package http

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"net/http"
	"time"

	"github.com/gosolid/solid/pkg/runtime/net"
	"github.com/gosolid/solid/pkg/runtime/stream"
	isolates "github.com/grexie/isolates"
)

type OutgoingMessage interface {
	stream.Writable

	AddTrailers(headers http.Header)
	AppendHeader(name string, values []string)
	Destroyed() bool
	Finished() bool
	FlushHeaders() error
	GetHeader(name string) []string
	GetHeaderNames() []string
	GetRawHeaderNames() []string
	Headers() map[string]any
	HasHeader(name string) bool
	HeadersSent() bool
	RemoveHeader(name string)
	SetHeader(ctx context.Context, name string, value *isolates.Value) error
	SetHeaders(headers http.Header)
	SetTimeout(time.Duration, func() error) func()
}

type ServerResponse interface {
	OutgoingMessage

	Request(context.Context) (ServerRequest, error)
	Socket() net.Socket
	SendDate() bool
	SetSendDate(bool)
	StatusCode() int
	SetStatusCode(int)
	StatusMessage() string
	SetStatusMessage(string)
	StrictContentLength() bool
	SetStrictContentLength(bool)
	WriteContinue() error
	WriteEarlyHints(http.Header) error
	WriteHead(statusCode int, statusMessage string, headers http.Header) error
	WriteProcessing() error

	Wait() error
}

type ClientRequest interface {
	OutgoingMessage

	MaxHeadersCount() int
	SetMaxHeadersCount(int)
	Path() string
	SetPath(string)
	Method() string
	SetMethod(string)
	Host() string
	SetHost(string)
	Protocol() string
	SetProtocol(string)
	ReusedSocket() bool
	SetNoDelay(bool)
	SetSocketKeepAlive(bool, time.Duration)
}

type IncomingMessage interface {
	stream.Readable

	Complete() bool
	Headers() map[string]any
	HeadersDistinct() http.Header
	HttpVersion() string

	RawHeaders() []string
	RawTrailers() []string
	SetTimeout(context.Context, time.Duration, *isolates.Value) error
	Trailers() http.Header
	TrailersDistinct() http.Header
}

type ServerRequest interface {
	IncomingMessage

	Response(context.Context) (ServerResponse, error)
	Socket() net.Socket

	IsUpgrade() bool

	Method() string
	URL() string
}

type ClientResponse interface {
	IncomingMessage

	StatusCode() int
	StatusMessage() string
}

type HttpInformation interface {
	HttpVersion() string
	HttpVersionMajor() int
	HttpVersionMinor() int
	StatusCode() int
	StatusMessage() string
	Headers() http.Header
	RawHeaders() []string
}
