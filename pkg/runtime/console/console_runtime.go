// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package console

import (
  isolates "github.com/grexie/isolates"
  reflect "reflect"
)

var _ = isolates.RegisterRuntime("console", "console.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "Console", func (in isolates.FunctionArgs) (*ConsoleBase, error) {
    return NewConsole(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "Console", constructor); err != nil {
    return nil, err
  }

var Module *isolates.Module
Exports := in.Args[1]
Require := in.Args[2]
if m, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(Module)); err != nil {
  return nil, err
} else {
  Module = m.Interface().(*isolates.Module)
}


  {
    fnName := "createDefaultConsole"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
rin := isolates.RuntimeFunctionArgs{FunctionArgs: in, Module: Module, Exports: Exports, Require: Require}

      if result, err := newDefaultConsole(rin); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if instance, err := fn.Call(in.ExecutionContext, nil); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "default", instance); err != nil {
      return nil, err
    }
  }

  return nil, nil
})

func (c *ConsoleBase) V8FuncAssert(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Assert(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncClear(in isolates.FunctionArgs) (*isolates.Value, error) {
  c.Clear()
  return nil, nil
}

func (c *ConsoleBase) V8FuncCount(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Count(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncCountReset(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.CountReset(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncDebug(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Debug(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncDir(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Dir(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncDirXML(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.DirXML(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncError(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Error(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncGroup(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Group(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncGroupCollapsed(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.GroupCollapsed(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncGroupEnd(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.GroupEnd(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncInfo(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Info(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncLog(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Log(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncTable(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Table(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncTime(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Time(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncTimeEnd(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.TimeEnd(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncTimeLog(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.TimeLog(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncTrace(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Trace(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncWarn(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Warn(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncProfile(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.Profile(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncProfileEnd(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.ProfileEnd(args...)
  return nil, nil
}

func (c *ConsoleBase) V8FuncTimeStamp(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  c.TimeStamp(args...)
  return nil, nil
}