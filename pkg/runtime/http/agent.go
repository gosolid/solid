//js:package http
package http

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gosolid/solid/pkg/runtime/net"
	"github.com/grexie/isolates"
)

var _ Agent = &AgentBase{}

//js:alias AgentBase
type Agent interface {
	Http() Http
	CreateConnection(ctx context.Context, options net.ConnectOptions, connectListener *isolates.Value) (net.Socket, error)
	GetName(ctx context.Context, options AgentRequestOptions) (string, error)
}

//js:class
type AgentBase struct {
	http Http
}

type IPFamily int

const (
	IPv4 IPFamily = 4
	IPv6 IPFamily = 6
)

type AgentRequestOptions struct {
	Host         string    `v8:"host"`
	Port         int       `v8:"port"`
	LocalAddress string    `v8:"localAddress"`
	Family       *IPFamily `v8:"family"`
}

func newAgent(in isolates.FunctionArgs, http Http) (*AgentBase, error) {
	agent := &AgentBase{
		http: http,
	}

	return agent, nil
}

func (a *AgentBase) Http() Http {
	return a.http
}

//js:method
func (a *AgentBase) CreateConnection(ctx context.Context, options net.ConnectOptions, connectListener *isolates.Value) (net.Socket, error) {
	if options.Host == nil {
		h := "localhost"
		options.Host = &h
	}

	if socketv, err := a.http.Net().Connect(ctx, "tcp", fmt.Sprintf("%s:%d", *options.Host, options.Port)); err != nil {
		return nil, err
	} else if socketrv, err := socketv.Unmarshal(ctx, reflect.TypeOf(&net.SocketBase{})); err != nil {
		return nil, err
	} else {
		socket := socketrv.Interface().(net.Socket)
		return socket, nil
	}
}

//js:method
func (a *AgentBase) GetName(ctx context.Context, options AgentRequestOptions) (string, error) {
	if options.Family == nil {
		ipv4 := IPv4
		options.Family = &ipv4
	} else if *options.Family != IPv4 && *options.Family != IPv6 {
		return "", fmt.Errorf("invalid ip address family: %d", *options.Family)
	}

	if *options.Family == IPv4 {
		return fmt.Sprintf("%s:%d:%s", options.Host, options.Port, options.LocalAddress), nil
	} else {
		return fmt.Sprintf("%s:%d:%s:%d", options.Host, options.Port, options.LocalAddress, *options.Family), nil
	}
}
