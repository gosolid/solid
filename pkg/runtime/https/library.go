//js:package https

package https

import (
	"reflect"

	"github.com/gosolid/solid/pkg/runtime/net"
	"github.com/grexie/isolates"
)

//go:generate go run github.com/grexie/isolates/codegen

type Import interface{}

var _ Https = &HttpsBase{}

//js:alias HttpsBase
type Https interface {
	Net() net.Net
	Agent() *isolates.Value
	GlobalAgent() Agent
}

//js:class Https
type HttpsBase struct {
	net         net.Net
	agent       *isolates.Value
	globalAgent Agent
}

//js:constructor Https
//js:export Https
func NewHttps(in isolates.FunctionArgs) (*HttpsBase, error) {
	options := in.Arg(in.ExecutionContext, 0)
	var n net.Net
	var ok bool

	if netv, err := options.Get(in.ExecutionContext, "net"); err != nil {
		return nil, err
	} else if n, ok = netv.Receiver(in.ExecutionContext).Interface().(net.Net); !ok {
		return nil, net.ErrInvalidNet
	}

	https := &HttpsBase{
		net: n,
	}

	in.This.SetReceiver(in.ExecutionContext, reflect.ValueOf(https))
	in.This.RebindAll(in.ExecutionContext)

	if _Agent, err := in.Context.CreateConstructorWithName(in.ExecutionContext, "Agent", func(in isolates.FunctionArgs) (*AgentBase, error) {
		return newAgent(in, https)
	}); err != nil {
		return nil, err
	} else if err := in.This.Set(in.ExecutionContext, "Agent", _Agent); err != nil {
		return nil, err
	} else if globalAgentv, err := _Agent.New(in.ExecutionContext); err != nil {
		return nil, err
	} else if globalAgentrv, err := globalAgentv.Unmarshal(in.ExecutionContext, reflect.TypeOf(&AgentBase{})); err != nil {
		return nil, err
	} else {
		https.agent = _Agent
		https.globalAgent = globalAgentrv.Interface().(Agent)
	}

	// if request, err := in.This.BindMethod(in.ExecutionContext, "request"); err != nil {
	// 	return nil, err
	// } else if err := in.This.Set(in.ExecutionContext, "request", request); err != nil {
	// 	return nil, err
	// }

	return https, nil
}

//js:method createDefaultHttp
//js:export-instance default
func NewDefaultHttps(in isolates.RuntimeFunctionArgs) (Https, error) {
	if options, err := in.Context.NewObject(in.ExecutionContext); err != nil {
		return nil, err
	} else if net, err := in.Require.Call(in.ExecutionContext, nil, "net"); err != nil {
		return nil, err
	} else if err := options.Set(in.ExecutionContext, "net", net); err != nil {
		return nil, err
	} else if n, err := in.Context.New(in.ExecutionContext, NewHttps, options); err != nil {
		return nil, err
	} else {
		return n.(Https), nil
	}
}

//js:export net
//js:get
func (h *HttpsBase) Net() net.Net {
	return h.net
}

//js:export Agent
//js:get
func (h *HttpsBase) Agent() *isolates.Value {
	return h.agent
}

//js:get
func (h *HttpsBase) GlobalAgent() Agent {
	return h.globalAgent
}
