// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package url

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("url", "url.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "URL", func (in isolates.FunctionArgs) (*URL, error) {
    return NewURL(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "URL", constructor); err != nil {
    return nil, err
  }

  {
    fnName := "pathToFileURL"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
  var path string
      if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
        return nil, __err
      } else {
        path = v.Interface().(string)
      }

      if result, err := PathToFileURL(in.ExecutionContext, path); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "pathToFileURL", fn); err != nil {
      return nil, err
    }
  }

  return nil, nil
})

func (u *URL) V8FuncToString(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := u.String()
  return in.Context.Create(in.ExecutionContext, result)
}