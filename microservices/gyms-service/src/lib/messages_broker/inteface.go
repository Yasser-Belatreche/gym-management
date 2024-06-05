package messages_broker

import "gym-management-gyms/src/lib/primitives/application_specific"

type MessagesBroker interface {
	Publish(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException

	Subscribe(subscribers ...*Subscriber)

	GetReply(message string, params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)

	RegisterReply(replies ...*Reply)

	HealthCheck() *Health
}

type Health struct {
	Provider string `json:"provider"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type Subscriber struct {
	Event     string
	Component string
	Handler   func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException
}

type Reply struct {
	Message string
	Handler func(params map[string]string, session *application_specific.Session) (map[string]string, *application_specific.ApplicationException)
}
