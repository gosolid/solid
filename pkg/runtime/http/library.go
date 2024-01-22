//js:package http

package http

//go:generate go run github.com/grexie/isolates/codegen

import (
	"reflect"

	"github.com/gosolid/solid/pkg/runtime/net"
	isolates "github.com/grexie/isolates"
)

//js:alias HttpBase
type Http interface {
	Net() net.Net
	Server() *isolates.Value
	Agent() *isolates.Value
}

type HttpBase struct {
	net         net.Net
	server      *isolates.Value
	agent       *isolates.Value
	globalAgent Agent
}

//js:constructor Http
//js:export Http
func NewHttp(in isolates.FunctionArgs) (*HttpBase, error) {
	options := in.Arg(in.ExecutionContext, 0)
	var n net.Net
	var ok bool

	if netv, err := options.Get(in.ExecutionContext, "net"); err != nil {
		return nil, err
	} else if n, ok = netv.Receiver(in.ExecutionContext).Interface().(net.Net); !ok {
		return nil, net.ErrInvalidNet
	}

	http := &HttpBase{
		net: n,
	}

	in.This.SetReceiver(in.ExecutionContext, reflect.ValueOf(http))
	in.This.RebindAll(in.ExecutionContext)

	if Server, err := in.Context.CreateConstructorWithName(in.ExecutionContext, "Server", func(in isolates.FunctionArgs) (*ServerBase, error) {
		return newServer(in, http)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "Server", Server); err != nil {
		return nil, err
	} else if createServer, err := in.Context.CreateWithName(in.ExecutionContext, "createServer", func(in isolates.FunctionArgs) (*isolates.Value, error) {
		return Server.NewValue(in.ExecutionContext, in.Args...)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "createServer", createServer); err != nil {
		return nil, err
	} else {
		http.server = Server
	}

	if _Agent, err := in.Context.CreateConstructorWithName(in.ExecutionContext, "Agent", func(in isolates.FunctionArgs) (*AgentBase, error) {
		return newAgent(in, http)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "Agent", _Agent); err != nil {
		return nil, err
	} else if globalAgentv, err := _Agent.New(in.ExecutionContext); err != nil {
		return nil, err
	} else if globalAgentrv, err := globalAgentv.Unmarshal(in.ExecutionContext, reflect.TypeOf(&AgentBase{})); err != nil {
		return nil, err
	} else {
		http.agent = _Agent
		http.globalAgent = globalAgentrv.Interface().(Agent)
	}

	// if request, err := in.This.BindMethod(in.ExecutionContext, "request"); err != nil {
	// 	return nil, err
	// } else if err := in.This.Set(in.ExecutionContext, "request", request); err != nil {
	// 	return nil, err
	// }

	return http, nil
}

//js:method createDefaultHttp
//js:export-instance default
func NewDefaultHttp(in isolates.RuntimeFunctionArgs) (Http, error) {
	if options, err := in.Context.NewObject(in.ExecutionContext); err != nil {
		return nil, err
	} else if net, err := in.Require.Call(in.ExecutionContext, nil, "net"); err != nil {
		return nil, err
	} else if err := options.Set(in.ExecutionContext, "net", net); err != nil {
		return nil, err
	} else if n, err := in.Context.New(in.ExecutionContext, NewHttp, options); err != nil {
		return nil, err
	} else {
		return n.(Http), nil
	}
}

//js:export net
//js:get
func (h *HttpBase) Net() net.Net {
	return h.net
}

//js:export Server
//js:get
func (h *HttpBase) Server() *isolates.Value {
	return h.server
}

//js:export Agent
//js:get
func (h *HttpBase) Agent() *isolates.Value {
	return h.agent
}

//js:get
func (h *HttpBase) GlobalAgent() Agent {
	return h.globalAgent
}
