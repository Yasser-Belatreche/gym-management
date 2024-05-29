package messages_broker

var broker MessagesBroker = nil

func NewMessagesBroker() MessagesBroker {
	if broker == nil {
		broker = NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{Async: false})
	}

	return broker
}
