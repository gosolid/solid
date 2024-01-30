// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package fs

import (
  embed "embed"
  isolates "github.com/grexie/isolates"
  reflect "reflect"
)

var _ = isolates.RegisterRuntime("fs", "embed.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "EmbeddedFileSystem", func (in isolates.FunctionArgs) (*embedfs, error) {
    var _embed *embed.FS
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&_embed).Elem()); err != nil {
      return nil, err
    } else if v != nil {
      _embed = v.Interface().(*embed.FS)
    }

    return NewEmbedFS(_embed), nil
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "EmbeddedFileSystem", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (fs *embedfs) V8FuncToString(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := fs.String()
  return in.Context.Create(in.ExecutionContext, result)
}

func (fs *embedfs) V8FuncCreateReadStream(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncReaddirSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncReaddir(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncReadFileSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncReadFile(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncOpenSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncOpen(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncCloseSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  var fd FileDescriptor
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&fd).Elem()); __err != nil {
    return nil, __err
  } else {
    fd = v.Interface().(FileDescriptor)
  }

  if err := fs.Close(in.ExecutionContext, fd); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (fs *embedfs) V8FuncClose(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var fd FileDescriptor
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&fd).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      fd = v.Interface().(FileDescriptor)
    }

      if err := fs.Close(in.ExecutionContext, fd); err != nil {
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

func (fs *embedfs) V8FuncFile(in isolates.FunctionArgs) (*isolates.Value, error) {
  var fd FileDescriptor
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&fd).Elem()); __err != nil {
    return nil, __err
  } else {
    fd = v.Interface().(FileDescriptor)
  }

  if result, err := fs.File(in.ExecutionContext, fd); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (fs *embedfs) V8FuncExistsSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncExists(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncStatSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncStat(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncLstatSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (fs *embedfs) V8FuncLstat(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (f *embedfs) V8FuncRealpath(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (f *embedfs) V8FuncRealpathSync(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (f *embedfile) V8FuncClose(in isolates.FunctionArgs) (*isolates.Value, error) {
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

func (f *embedfile) V8FuncCloseSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := f.Close(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (f *embedfile) V8FuncReadAllSync(in isolates.FunctionArgs) (*isolates.Value, error) {
  if result, err := f.ReadAll(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (f *embedfile) V8FuncReadAll(in isolates.FunctionArgs) (*isolates.Value, error) {
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