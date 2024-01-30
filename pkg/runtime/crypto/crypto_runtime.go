// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package crypto

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
  buffer "github.com/gosolid/solid/pkg/runtime/buffer"
)

var _ = isolates.RegisterRuntime("crypto", "crypto.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  {
    fnName := "createHash"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
  var algorithm string
      if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&algorithm).Elem()); __err != nil {
        return nil, __err
      } else {
        algorithm = v.Interface().(string)
      }

      if result, err := CreateHash(in.ExecutionContext, algorithm); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "createHash", fn); err != nil {
      return nil, err
    }
  }

  {
    fnName := "createHmac"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
  var algorithm string
      if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&algorithm).Elem()); __err != nil {
        return nil, __err
      } else {
        algorithm = v.Interface().(string)
      }

      key := in.Arg(in.ExecutionContext, 1)
      if result, err := CreateHmac(in.ExecutionContext, algorithm, key); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "createHmac", fn); err != nil {
      return nil, err
    }
  }

  {
    fnName := "randomBytes"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
      if result, err := RandomBytes(in.ExecutionContext); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "randomBytes", fn); err != nil {
      return nil, err
    }
  }

  return nil, nil
})

func (h *Hash) V8GetAlgorithm(in isolates.GetterArgs) (*isolates.Value, error) {
  result := h.Algorithm()
  return in.Context.Create(in.ExecutionContext, result)
}

func (h *Hash) V8FuncUpdate(in isolates.FunctionArgs) (*isolates.Value, error) {
  b := in.Arg(in.ExecutionContext, 0)
  if result, err := h.Update(in.ExecutionContext, b); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (h *Hash) V8FuncDigest(in isolates.FunctionArgs) (*isolates.Value, error) {
  var encoding *buffer.BufferEncoding
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&encoding).Elem()); __err != nil {
    return nil, __err
  } else if v != nil {
    encoding = v.Interface().(*buffer.BufferEncoding)
  }

  if result, err := h.Sum(encoding); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}