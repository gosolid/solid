// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package events

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("events", "/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "EventEmitter", func (in isolates.FunctionArgs) (*EventEmitterBase, error) {
    return NewEventEmitter(in), nil
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "default", constructor); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "EventEmitter", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (e *EventEmitterBase) V8FuncAddListener(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.AddListener(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncOn(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.AddListener(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncOnce(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.AddListenerOnce(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncEmit(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  args := make([]any, len(in.Args)-1)
  for i, arg := range in.Args[1:] {
    args[i] = arg
  }

  if err := e.Emit(in.ExecutionContext, eventName, args...); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncEventNames(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EventNames()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *EventEmitterBase) V8FuncGetMaxListeners(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.GetMaxListeners()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *EventEmitterBase) V8FuncSetMaxListeners(in isolates.FunctionArgs) (*isolates.Value, error) {
  var maxListeners int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&maxListeners).Elem()); __err != nil {
    return nil, __err
  } else {
    maxListeners = v.Interface().(int)
  }

  e.SetMaxListeners(maxListeners)
  return nil, nil
}

func (e *EventEmitterBase) V8FuncListenerCount(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if result, err := e.ListenerCount(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (e *EventEmitterBase) V8FuncListeners(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  result := e.Listeners(eventName)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *EventEmitterBase) V8FuncRawListeners(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  result := e.Listeners(eventName)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *EventEmitterBase) V8FuncRemoveListener(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.RemoveListener(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncOff(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.RemoveListener(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncPrependListener(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.PrependListener(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncPrependListenerOnce(in isolates.FunctionArgs) (*isolates.Value, error) {
  var eventName string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&eventName).Elem()); __err != nil {
    return nil, __err
  } else {
    eventName = v.Interface().(string)
  }

  listener := in.Arg(in.ExecutionContext, 1)
  if err := e.PrependListenerOnce(in.ExecutionContext, eventName, listener); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (e *EventEmitterBase) V8FuncRemoveAllListeners(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := e.RemoveAllListeners(in); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}