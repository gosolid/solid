// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package util

import (
  isolates "github.com/grexie/isolates"
  reflect "reflect"
)

var _ = isolates.RegisterRuntime("util", "debug.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  {
    fnName := "debuglog"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
  var section string
      if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&section).Elem()); __err != nil {
        return nil, __err
      } else {
        section = v.Interface().(string)
      }

      callback := in.Arg(in.ExecutionContext, 1)
      if result, err := Debuglog(in.ExecutionContext, section, callback); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "debuglog", fn); err != nil {
      return nil, err
    }
  }

  return nil, nil
})

