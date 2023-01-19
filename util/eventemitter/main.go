package eventemitter

type Handler func(string)

type Event struct {
	handlers map[string][]Handler
}

func (e Event) On(event string, h Handler) {
	e.handlers[event] = append(e.handlers[event], h)
}

func (e Event) Emit(event string, value string) {
	handlers, ok := e.handlers[event]
	if !ok {
		return
	}
	for _, handler := range handlers {
		handler(value)
	}
}

func New() Event {
	return Event{
		handlers: make(map[string][]Handler),
	}
}
