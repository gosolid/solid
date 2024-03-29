// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package path

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("path", "posix.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  {
    fnName := "toNamespacedPath"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
  var path string
      if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
        return nil, __err
      } else {
        path = v.Interface().(string)
      }

      result := ToNamespacedPath(path)
      return in.Context.Create(in.ExecutionContext, result)
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "toNamespacedPath", fn); err != nil {
      return nil, err
    }
  }

  return nil, nil
})

