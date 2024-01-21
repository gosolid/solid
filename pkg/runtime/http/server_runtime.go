// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package http

import (
  time "time"
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("http", "server.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  return nil, nil
})

func (l *ServerBase) V8FuncCloseAllConnections(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := l.CloseAllConnections(); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (l *ServerBase) V8FuncCloseIdleConnections(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := l.CloseIdleConnections(); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (l *ServerBase) V8GetHeadersTimeout(in isolates.GetterArgs) (*isolates.Value, error) {
  result := l.HeadersTimeout()
  return in.Context.Create(in.ExecutionContext, result)
}

func (l *ServerBase) V8GetMaxHeadersCount(in isolates.GetterArgs) (*isolates.Value, error) {
  result := l.MaxHeadersCount()
  return in.Context.Create(in.ExecutionContext, result)
}

func (l *ServerBase) V8GetRequestTimeout(in isolates.GetterArgs) (*isolates.Value, error) {
  result := l.RequestTimeout()
  return in.Context.Create(in.ExecutionContext, result)
}

func (l *ServerBase) V8GetTimeout(in isolates.GetterArgs) (*isolates.Value, error) {
  result := l.Timeout()
  return in.Context.Create(in.ExecutionContext, result)
}

func (l *ServerBase) V8FuncSetTimeout(in isolates.FunctionArgs) (*isolates.Value, error) {
  var timeout time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&timeout).Elem()); __err != nil {
    return nil, __err
  } else {
    timeout = v.Interface().(time.Duration)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := l.SetTimeout(in.ExecutionContext, timeout, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (l *ServerBase) V8GetMaxRequestsPerSocket(in isolates.GetterArgs) (*isolates.Value, error) {
  result := l.MaxRequestsPerSocket()
  return in.Context.Create(in.ExecutionContext, result)
}

func (l *ServerBase) V8GetKeepAliveTimeout(in isolates.GetterArgs) (*isolates.Value, error) {
  result := l.KeepAliveTimeout()
  return in.Context.Create(in.ExecutionContext, result)
}