package events

import "errors"

var ErrhandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispacher struct{
	handlers map[string][]EventHandlerInterface

}

func NewEventDispacher() *EventDispacher {
	return &EventDispacher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispacher) Register(eventName string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName];ok {
		for _, h := range ed.handlers[eventName]{
			if h == handler{
				return ErrhandlerAlreadyRegistered
			}
		}

	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
