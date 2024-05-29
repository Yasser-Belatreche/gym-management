package messages_broker

import "gym-management/src/lib/primitives/application_specific"

type MessagesBroker interface {
	Publish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException

	Subscribe(subscribers ...*Subscriber)

	Ask(question string, params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)

	RegisterAnswer(answers ...*Answer)
}

type Subscriber struct {
	Event   string
	Handler func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException
}

type Answer struct {
	Question string
	Answer   func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)
}
