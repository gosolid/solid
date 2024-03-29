// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package http

import (
  time "time"
  reflect "reflect"
  isolates "github.com/grexie/isolates"
  http "net/http"
)

var _ = isolates.RegisterRuntime("http", "messages.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "IncomingMessage", func (in isolates.FunctionArgs) (*IncomingMessageBase, error) {
    return NewIncomingMessage(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "IncomingMessage", constructor); err != nil {
    return nil, err
  }

  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "ServerRequest", func (in isolates.FunctionArgs) (*ServerRequestBase, error) {
    return NewServerRequest(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "ServerRequest", constructor); err != nil {
    return nil, err
  }

  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "ClientResponse", func (in isolates.FunctionArgs) (*ClientResponseBase, error) {
    return NewClientResponse(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "ClientResponse", constructor); err != nil {
    return nil, err
  }

  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "OutgoingMessage", func (in isolates.FunctionArgs) (*OutgoingMessageBase, error) {
    return NewOutgoingMessage(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "OutgoingMessage", constructor); err != nil {
    return nil, err
  }

  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "ServerResponse", func (in isolates.FunctionArgs) (*ServerResponseBase, error) {
    return NewServerResponse(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "ServerResponse", constructor); err != nil {
    return nil, err
  }

  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "ClientRequest", func (in isolates.FunctionArgs) (*ClientRequestBase, error) {
    return NewClientRequest(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "ClientRequest", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (r *IncomingMessageBase) V8GetComplete(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Complete()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8GetHeaders(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Headers()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8GetHeadersDistinct(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.HeadersDistinct()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8GetHttpVersion(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.HttpVersion()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8GetRawHeaders(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.RawHeaders()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8GetRawTrailers(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.RawTrailers()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8FuncSetTimeout(in isolates.FunctionArgs) (*isolates.Value, error) {
  var timeout time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&timeout).Elem()); __err != nil {
    return nil, __err
  } else {
    timeout = v.Interface().(time.Duration)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := r.SetTimeout(in.ExecutionContext, timeout, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *IncomingMessageBase) V8GetTrailers(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Trailers()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *IncomingMessageBase) V8GetTrailersDistinct(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.TrailersDistinct()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncAddTrailers(in isolates.FunctionArgs) (*isolates.Value, error) {
  var headers http.Header
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&headers).Elem()); __err != nil {
    return nil, __err
  } else {
    headers = v.Interface().(http.Header)
  }

  r.AddTrailers(headers)
  return nil, nil
}

func (r *OutgoingMessageBase) V8FuncAppendHeader(in isolates.FunctionArgs) (*isolates.Value, error) {
  var name string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&name).Elem()); __err != nil {
    return nil, __err
  } else {
    name = v.Interface().(string)
  }

  var values []string
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&values).Elem()); __err != nil {
    return nil, __err
  } else if v != nil {
    values = v.Interface().([]string)
  }

  r.AppendHeader(name, values)
  return nil, nil
}

func (r *OutgoingMessageBase) V8GetDestroyed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Destroyed()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetFinished(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Finished()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncGetHeader(in isolates.FunctionArgs) (*isolates.Value, error) {
  var name string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&name).Elem()); __err != nil {
    return nil, __err
  } else {
    name = v.Interface().(string)
  }

  result := r.GetHeader(name)
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncGetHeaderNames(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := r.GetHeaderNames()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncGetRawHeaderNames(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := r.GetRawHeaderNames()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetHeaders(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Headers()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncHasHeader(in isolates.FunctionArgs) (*isolates.Value, error) {
  var name string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&name).Elem()); __err != nil {
    return nil, __err
  } else {
    name = v.Interface().(string)
  }

  result := r.HasHeader(name)
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetHeadersSent(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.HeadersSent()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncRemoveHeader(in isolates.FunctionArgs) (*isolates.Value, error) {
  var name string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&name).Elem()); __err != nil {
    return nil, __err
  } else {
    name = v.Interface().(string)
  }

  r.RemoveHeader(name)
  return nil, nil
}

func (r *OutgoingMessageBase) V8FuncSetHeader(in isolates.FunctionArgs) (*isolates.Value, error) {
  var name string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&name).Elem()); __err != nil {
    return nil, __err
  } else {
    name = v.Interface().(string)
  }

  value := in.Arg(in.ExecutionContext, 1)
  if err := r.SetHeader(in.ExecutionContext, name, value); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *OutgoingMessageBase) V8FuncSetHeaders(in isolates.FunctionArgs) (*isolates.Value, error) {
  var headers http.Header
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&headers).Elem()); __err != nil {
    return nil, __err
  } else {
    headers = v.Interface().(http.Header)
  }

  r.SetHeaders(headers)
  return nil, nil
}

func (r *OutgoingMessageBase) V8FuncSetTimeout(in isolates.FunctionArgs) (*isolates.Value, error) {
  var timeout time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&timeout).Elem()); __err != nil {
    return nil, __err
  } else {
    timeout = v.Interface().(time.Duration)
  }

  var listener func() error
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&listener).Elem()); __err != nil {
    return nil, __err
  } else {
    listener = v.Interface().(func() error)
  }

  result := r.SetTimeout(timeout, listener)
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetWritableEnded(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.WritableEnded()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetWritableFinished(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.WritableFinished()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetSendDate(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.SendDate()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8GetStrictContentLength(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.StrictContentLength()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *OutgoingMessageBase) V8FuncWriteContinue(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := r.WriteContinue(); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *OutgoingMessageBase) V8FuncWriteEarlyHints(in isolates.FunctionArgs) (*isolates.Value, error) {
  var header http.Header
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&header).Elem()); __err != nil {
    return nil, __err
  } else {
    header = v.Interface().(http.Header)
  }

  if err := r.WriteEarlyHints(header); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *OutgoingMessageBase) V8FuncWriteProcessing(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := r.WriteProcessing(); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *OutgoingMessageBase) V8FuncWait(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
      if err := r.Wait(); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Undefined(in.ExecutionContext); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ServerRequestBase) V8GetSocket(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Socket()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ServerRequestBase) V8GetMethod(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Method()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ServerRequestBase) V8GetUrl(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.URL()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientResponseBase) V8GetStatusCode(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.StatusCode()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientResponseBase) V8GetStatusMessage(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.StatusMessage()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ServerResponseBase) V8GetSocket(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Socket()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ServerResponseBase) V8FuncFlushHeaders(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := r.FlushHeaders(); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *ServerResponseBase) V8FuncWriteHead(in isolates.FunctionArgs) (*isolates.Value, error) {
  var statusCode int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&statusCode).Elem()); __err != nil {
    return nil, __err
  } else {
    statusCode = v.Interface().(int)
  }

  var statusMessage string
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&statusMessage).Elem()); __err != nil {
    return nil, __err
  } else {
    statusMessage = v.Interface().(string)
  }

  var headers http.Header
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&headers).Elem()); __err != nil {
    return nil, __err
  } else {
    headers = v.Interface().(http.Header)
  }

  if err := r.WriteHead(statusCode, statusMessage, headers); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *ServerResponseBase) V8GetStatusCode(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.StatusCode()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ServerResponseBase) V8GetStatusMessage(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.StatusMessage()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8GetSocket(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Socket()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8FuncFlushHeaders(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := r.FlushHeaders(); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *ClientRequestBase) V8GetHost(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Host()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8GetMaxHeadersCount(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.MaxHeadersCount()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8GetMethod(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Method()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8GetPath(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Path()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8GetProtocol(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Protocol()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8GetReusedSocket(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReusedSocket()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ClientRequestBase) V8FuncSetNoDelay(in isolates.FunctionArgs) (*isolates.Value, error) {
  var args0 bool
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); __err != nil {
    return nil, __err
  } else {
    args0 = v.Interface().(bool)
  }

  r.SetNoDelay(args0)
  return nil, nil
}

func (r *ClientRequestBase) V8FuncSetSocketKeepAlive(in isolates.FunctionArgs) (*isolates.Value, error) {
  var args0 bool
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); __err != nil {
    return nil, __err
  } else {
    args0 = v.Interface().(bool)
  }

  var args1 time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); __err != nil {
    return nil, __err
  } else {
    args1 = v.Interface().(time.Duration)
  }

  r.SetSocketKeepAlive(args0, args1)
  return nil, nil
}