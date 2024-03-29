// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package fs

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("fs", "local.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "LocalFileSystem", func (in isolates.FunctionArgs) (*localfs, error) {
    var _path string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&_path).Elem()); err != nil {
      return nil, err
    } else {
      _path = v.Interface().(string)
    }

    return NewLocalFS(_path), nil
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "LocalFileSystem", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (fs *localfs) V8FuncToString(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := fs.String()
  return in.Context.Create(in.ExecutionContext, result)
}

func (fs *localfs) V8FuncMount(in isolates.FunctionArgs) (*isolates.Value, error) {
  var mountPoint string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&mountPoint).Elem()); __err != nil {
    return nil, __err
  } else {
    mountPoint = v.Interface().(string)
  }

  var mount FS
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&mount).Elem()); __err != nil {
    return nil, __err
  } else {
    mount = v.Interface().(FS)
  }

  var path string
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
    return nil, __err
  } else {
    path = v.Interface().(string)
  }

  if err := fs.Mount(mountPoint, mount, path); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (fs *localfs) V8FuncReaddirSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var path string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
    return nil, __err
  } else {
    path = v.Interface().(string)
  }

  if result, err := fs.ReadDir(in.ExecutionContext, path); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncReaddir(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var path string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      path = v.Interface().(string)
    }

      if result, err := fs.ReadDir(in.ExecutionContext, path); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (fs *localfs) V8FuncReadFileSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var path string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
    return nil, __err
  } else {
    path = v.Interface().(string)
  }

  options := make([]any, len(in.Args)-1)
  for i, arg := range in.Args[1:] {
    options[i] = arg
  }

  if result, err := fs.ReadFile(in.ExecutionContext, path, options...); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncReadFile(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var path string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      path = v.Interface().(string)
    }

  options := make([]any, len(in.Args)-1)
  for i, arg := range in.Args[1:] {
    options[i] = arg
  }

      if result, err := fs.ReadFile(in.ExecutionContext, path, options...); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (fs *localfs) V8FuncCreateReadStream(in isolates.FunctionArgs) (*isolates.Value, error) {
  var path string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
    return nil, __err
  } else {
    path = v.Interface().(string)
  }

  if result, err := fs.ReadStream(in.ExecutionContext, path); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncOpenSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var path string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); __err != nil {
    return nil, __err
  } else {
    path = v.Interface().(string)
  }

  if result, err := fs.Open(in.ExecutionContext, path); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncOpen(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var path string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&path).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      path = v.Interface().(string)
    }

      if result, err := fs.Open(in.ExecutionContext, path); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (fs *localfs) V8FuncExistsSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var p string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); __err != nil {
    return nil, __err
  } else {
    p = v.Interface().(string)
  }

  if result, err := fs.Exists(in.ExecutionContext, p); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncExists(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var p string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      p = v.Interface().(string)
    }

      if result, err := fs.Exists(in.ExecutionContext, p); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (fs *localfs) V8FuncStatSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var p string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); __err != nil {
    return nil, __err
  } else {
    p = v.Interface().(string)
  }

  if result, err := fs.Stat(in.ExecutionContext, p); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncStat(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var p string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      p = v.Interface().(string)
    }

      if result, err := fs.Stat(in.ExecutionContext, p); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (fs *localfs) V8FuncLstatSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var p string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); __err != nil {
    return nil, __err
  } else {
    p = v.Interface().(string)
  }

  if result, err := fs.LStat(in.ExecutionContext, p); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *localfs) V8FuncLstat(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var p string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      p = v.Interface().(string)
    }

      if result, err := fs.LStat(in.ExecutionContext, p); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (f *localfs) V8FuncRealpath(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var p string
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      p = v.Interface().(string)
    }

      if result, err := f.RealPath(in.ExecutionContext, p); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (f *localfs) V8FuncRealpathSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var p string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&p).Elem()); __err != nil {
    return nil, __err
  } else {
    p = v.Interface().(string)
  }

  if result, err := f.RealPath(in.ExecutionContext, p); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (f *localfile) V8FuncCloseSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := f.Close(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (f *localfile) V8FuncClose(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
      if err := f.Close(in.ExecutionContext); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Undefined(in.ExecutionContext); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}

func (f *localfile) V8FuncReadAllSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  if result, err := f.ReadAll(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (f *localfile) V8FuncReadAll(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
      if result, err := f.ReadAll(in.ExecutionContext); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    if len(in.Args) > 0 {
      callback := in.Arg(in.ExecutionContext, len(in.Args) - 1)
      if callback.IsKind(isolates.KindFunction) {
        return nil, resolver.ToCallback(in.ExecutionContext, callback)
      }
    }
    return nil, nil
  }
}