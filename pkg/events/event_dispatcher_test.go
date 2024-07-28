package events

import (
	"time"

	"github.com/tiago-g-sales/tgsutils/pkg/events"
)

type TestEvent struct{
	Name 		string
	Payload 	interface{}
}

func (e *TestEvent) GetName() string{
	return e.Name
}

func (e *TestEvent) GetPayload() interface{}{
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time{
	return time.Now()
}

type TestEventHandler struct {}

func (h *TestEventHandler) Handle(event EventInterface){

}