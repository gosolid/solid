// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package stream

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("stream", "writable.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "Writable", func (in isolates.FunctionArgs) (*WritableBase, error) {
    return NewWritable(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "Writable", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (w *WritableBase) V8FuncCork(in isolates.FunctionArgs) (*isolates.Value, error) {
  w.Cork()
  return nil, nil
}

func (w *WritableBase) V8FuncEnd(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  if err := w.End(in.ExecutionContext, args...); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (w *WritableBase) V8FuncSetDefaultEncoding(in isolates.FunctionArgs) (*isolates.Value, error) {
  var encoding BufferEncoding
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&encoding).Elem()); __err != nil {
    return nil, __err
  } else {
    encoding = v.Interface().(BufferEncoding)
  }

  w.SetDefaultEncoding(encoding)
  return nil, nil
}

func (w *WritableBase) V8FuncUncork(in isolates.FunctionArgs) (*isolates.Value, error) {
  w.Uncork()
  return nil, nil
}

func (w *WritableBase) V8GetWritable(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.Writable()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableAborted(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableAborted()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableEnded(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableEnded()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableCorked(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableCorked()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableFinished(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableFinished()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableHighWaterMark(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableHighWaterMark()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableLength(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableLength()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableNeedDrain(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableNeedDrain()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8GetWritableObjectMode(in isolates.GetterArgs) (*isolates.Value, error) {
  result := w.WritableObjectMode()
  return in.Context.Create(in.ExecutionContext, result)
}

func (w *WritableBase) V8FuncWrite(in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]any, len(in.Args)-0)
  for i, arg := range in.Args[0:] {
    args[i] = arg
  }

  if result, err := w.WritableWrite(in.ExecutionContext, args...); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (c *Writer) V8FuncWrite(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this Writable
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else {
    this = v.Interface().(Writable)
  }

  chunk := in.Arg(in.ExecutionContext, 1)
  var encoding string
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&encoding).Elem()); __err != nil {
    return nil, __err
  } else {
    encoding = v.Interface().(string)
  }

  callback := in.Arg(in.ExecutionContext, 3)
  if err := c.Write(in, this, chunk, encoding, callback); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (c *WriteCloser) V8FuncFinal(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this Writable
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else {
    this = v.Interface().(Writable)
  }

  callback := in.Arg(in.ExecutionContext, 1)
  if err := c.Final(in, this, callback); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}