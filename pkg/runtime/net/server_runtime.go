// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package net

import (
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("net", "server.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  return nil, nil
})

func (s *ServerBase) V8FuncClose(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := s.Close(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (s *ServerBase) V8FuncAddress(in isolates.FunctionArgs) (*isolates.Value, error) {
  if result, err := s.Address(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (s *ServerBase) V8Get_listener(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Listener()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *ServerBase) V8FuncListen(in isolates.FunctionArgs) (*isolates.Value, error) {
  if result, err := s.Listen(in); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}