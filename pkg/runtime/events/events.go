package events

import "sync"

type EventHandler struct {
	mutex    sync.Mutex
	handlers []*struct{ callback func(...any) error }
}

func (h *EventHandler) Add(callback func(...any) error) func() {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	slot := &struct{ callback func(...any) error }{
		callback: callback,
	}

	if h.handlers == nil {
		h.handlers = []*struct{ callback func(...any) error }{}
	}

	h.handlers = append(h.handlers, slot)

	return func() {
		h.mutex.Lock()
		defer h.mutex.Unlock()

		for i, v := range h.handlers {
			if v == slot {
				h.handlers = append(h.handlers[:i], h.handlers[i+1:]...)
				return
			}
		}
	}
}

func (h *EventHandler) Emit(args ...any) error {
	h.mutex.Lock()
	handlers := make([]*struct{ callback func(...any) error }, len(h.handlers))
	copy(handlers, h.handlers)
	h.mutex.Unlock()

	for _, slot := range handlers {
		if err := slot.callback(args...); err != nil {
			return err
		}
	}

	return nil
}
