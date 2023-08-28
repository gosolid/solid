//js:package events
package events

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"sync"

	"github.com/grexie/isolates"
)

var _ EventEmitter = &EventEmitterBase{}

type registration struct {
	listener *isolates.Value
	once     bool
}

type EventEmitter interface {
	AddListener(ctx context.Context, eventName string, listener *isolates.Value) error
	AddListenerOnce(ctx context.Context, eventName string, listener *isolates.Value) error
	Emit(ctx context.Context, eventName string, args ...any) error
	EmitError(ctx context.Context, err error) error
	EmitErrorValue(ctx context.Context, err *isolates.Value) error
	EventNames() []string
	GetMaxListeners() int
	SetMaxListeners(maxListeners int)
	ListenerCount(ctx context.Context, eventName string, listener *isolates.Value) (int, error)
	Listeners(eventName string) []*isolates.Value
	RemoveListener(ctx context.Context, eventName string, listener *isolates.Value) error
	RemoveAllListeners(in isolates.FunctionArgs) error
	PrependListener(ctx context.Context, eventName string, listener *isolates.Value) error
	PrependListenerOnce(ctx context.Context, eventName string, listener *isolates.Value) error
}

type EventEmitterBase struct {
	context      *isolates.Context
	mutex        sync.Mutex
	listeners    map[string][]*registration
	maxListeners int
}

//js:constructor EventEmitter
//js:export default
//js:export EventEmitter
func NewEventEmitter(in isolates.FunctionArgs) *EventEmitterBase {
	return &EventEmitterBase{
		context:      in.Context,
		maxListeners: 20,
	}
}

func (e *EventEmitterBase) V8Construct(in isolates.FunctionArgs) (*EventEmitterBase, error) {
	e.context = in.Context
	e.maxListeners = 20
	return e, nil
}

//js:method
//js:method on
func (e *EventEmitterBase) AddListener(ctx context.Context, eventName string, listener *isolates.Value) error {
	if err := e.Emit(ctx, "newListener", eventName, listener); err != nil {
		return err
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		e.listeners = map[string][]*registration{}
	}

	if _, ok := e.listeners[eventName]; !ok {
		e.listeners[eventName] = []*registration{}
	}

	e.listeners[eventName] = append(e.listeners[eventName], &registration{listener, false})

	return nil
}

//js:method once
func (e *EventEmitterBase) AddListenerOnce(ctx context.Context, eventName string, listener *isolates.Value) error {
	if err := e.Emit(ctx, "newListener", eventName, listener); err != nil {
		return err
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		e.listeners = map[string][]*registration{}
	}

	if _, ok := e.listeners[eventName]; !ok {
		e.listeners[eventName] = []*registration{}
	}

	e.listeners[eventName] = append(e.listeners[eventName], &registration{listener, true})

	return nil
}

//js:method emit
func (e *EventEmitterBase) Emit(ctx context.Context, eventName string, args ...any) error {
	e.mutex.Lock()
	var registrations []*registration
	if l, ok := e.listeners[eventName]; !ok {
		e.mutex.Unlock()
		return nil
	} else {
		registrations = make([]*registration, len(l))
		copy(registrations, l)
	}
	e.mutex.Unlock()

	for _, registration := range registrations {

		if registration.once {
			e.mutex.Lock()
			if l, ok := e.listeners[eventName]; ok {
				for i, v := range l {
					if v == registration {
						e.listeners[eventName] = append(e.listeners[eventName][:i], e.listeners[eventName][i+1:]...)
						break
					}
				}
			}
			e.mutex.Unlock()
		}

		if global, err := registration.listener.GetContext().Global(ctx); err != nil {
			return err
		} else {
			if _, err := registration.listener.Call(ctx, global, args...); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *EventEmitterBase) EmitError(ctx context.Context, err error) error {
	isolates.For(ctx).Debug("emit error", err)
	return e.Emit(ctx, "error", err)
}

func (e *EventEmitterBase) EmitErrorValue(ctx context.Context, err *isolates.Value) error {
	isolates.For(ctx).Debug("emit error", err)
	return e.Emit(ctx, "error", err)
}

//js:method
func (e *EventEmitterBase) EventNames() []string {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	names := []string{}

	if e.listeners == nil {
		return names
	}

	for name := range e.listeners {
		names = append(names, name)
	}

	return names
}

//js:method
func (e *EventEmitterBase) GetMaxListeners() int {
	return e.maxListeners
}

//js:method
func (e *EventEmitterBase) SetMaxListeners(maxListeners int) {
	e.maxListeners = maxListeners
}

//js:method
func (e *EventEmitterBase) ListenerCount(ctx context.Context, eventName string, listener *isolates.Value) (int, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		return 0, nil
	}

	if listeners, ok := e.listeners[eventName]; !ok {
		return 0, nil
	} else if listener == nil {
		return len(listeners), nil
	} else {
		count := 0
		for _, h := range listeners {
			if equals, err := h.listener.StrictEquals(ctx, listener); err != nil {
				return 0, err
			} else if equals {
				count++
			}
		}
		return count, nil
	}
}

//js:method
//js:method rawListeners
func (e *EventEmitterBase) Listeners(eventName string) []*isolates.Value {
	listeners := []*isolates.Value{}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		return listeners
	}

	if registrations, ok := e.listeners[eventName]; !ok {
		return listeners
	} else {
		for _, registration := range registrations {
			listeners = append(listeners, registration.listener)
		}
		return listeners
	}
}

//js:method
//js:method off
func (e *EventEmitterBase) RemoveListener(ctx context.Context, eventName string, listener *isolates.Value) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		return nil
	}

	if registrations, ok := e.listeners[eventName]; !ok {
		return nil
	} else {
		for i, registration := range registrations {
			if equals, err := registration.listener.StrictEquals(ctx, listener); err != nil {
				return err
			} else if equals {
				e.listeners[eventName] = append(e.listeners[eventName][:i], e.listeners[eventName][i+1:]...)
				break
			}
		}

		if len(e.listeners[eventName]) == 0 {
			delete(e.listeners, eventName)
		}

		if len(e.listeners) == 0 {
			e.listeners = nil
		}

		return nil
	}
}

//js:method
func (e *EventEmitterBase) PrependListener(ctx context.Context, eventName string, listener *isolates.Value) error {
	if err := e.Emit(ctx, "newListener", listener); err != nil {
		return err
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		e.listeners = map[string][]*registration{}
	}

	if _, ok := e.listeners[eventName]; !ok {
		e.listeners[eventName] = []*registration{}
	}

	e.listeners[eventName] = append([]*registration{{listener, false}}, e.listeners[eventName]...)

	return nil
}

//js:method
func (e *EventEmitterBase) PrependListenerOnce(ctx context.Context, eventName string, listener *isolates.Value) error {
	if err := e.Emit(ctx, "newListener", listener); err != nil {
		return err
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.listeners == nil {
		e.listeners = map[string][]*registration{}
	}

	if _, ok := e.listeners[eventName]; !ok {
		e.listeners[eventName] = []*registration{}
	}

	e.listeners[eventName] = append([]*registration{{listener, true}}, e.listeners[eventName]...)

	return nil
}

//js:method
func (e *EventEmitterBase) RemoveAllListeners(in isolates.FunctionArgs) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	var eventName *string

	if len(in.Args) > 0 {
		if e, err := in.Arg(in.ExecutionContext, 0).StringValue(in.ExecutionContext); err != nil {
			return err
		} else {
			eventName = &e
		}
	}

	if eventName == nil {
		e.listeners = nil
	} else {
		delete(e.listeners, *eventName)
	}
	return nil
}
