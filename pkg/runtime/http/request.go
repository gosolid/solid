//js:package http
package http

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gosolid/solid/pkg/runtime/net"
	"github.com/grexie/isolates"
)

//go:generate go run github.com/grexie/isolates/codegen

//js:method
func (h *HttpBase) Request(ctx context.Context, args ...any) (ClientRequest, error) {
	var _url string
	var callback *isolates.Value

	if u, ok := args[0].(string); ok {
		_url = u
	} else if uv, ok := args[0].(*isolates.Value); ok && uv.IsKind(isolates.KindString) {
		_url = uv.String()
	} else {
		// options
		return nil, fmt.Errorf("not implemented")
	}

	if cv, ok := args[len(args)-1].(*isolates.Value); ok && cv.IsKind(isolates.KindFunction) {
		callback = cv
	}

	if urlParsed, err := url.Parse(_url); err != nil {
		return nil, err
	} else {
		hostname := urlParsed.Hostname()
		port, _ := strconv.ParseInt(urlParsed.Port(), 10, 16)
		if socket, err := h.GlobalAgent().CreateConnection(ctx, net.ConnectOptions{
			Host: &hostname,
			Port: int(port),
		}, nil); err != nil {
			return nil, err
		} else if reqrv, err := isolates.For(ctx).Context().New(ctx, NewClientRequest, socket, _url); err != nil {
			return nil, err
		} else {
			req := reqrv.(ClientRequest)

			if callback != nil {
				if err := req.AddListener(ctx, "response", callback); err != nil {
					return nil, err
				}
			}

			isolates.For(ctx).Isolate().Background(ctx, func(ctx context.Context) {
				if _, err := isolates.For(ctx).Context().New(ctx, NewClientResponse, req); err != nil {
					req.EmitError(ctx, err)
					return
				}
			})

			isolates.For(ctx).Isolate().Background(ctx, func(ctx context.Context) {
				if err := req.Wait(); err != nil {
					req.EmitError(ctx, err)
				}
			})

			return req, nil
		}
	}
}
