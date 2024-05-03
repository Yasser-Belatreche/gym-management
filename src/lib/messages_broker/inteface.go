package messages_broker

import "gym-management/src/lib/primitives/application_specific"

type MessagesBroker interface {
	Publish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException

	Subscribe(subscriber *Subscriber)

	Ask(question *Question, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)

	RegisterAnswer(answer *Answer)
}

type Subscriber struct {
	event   string
	handler func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException
}

type Question struct {
	Question string
	Params   map[string]string
}

type Answer struct {
	Question string
	Answer   func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)
}
