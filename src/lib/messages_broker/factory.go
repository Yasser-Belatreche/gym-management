package messages_broker

func NewMessagesBroker() MessagesBroker {
	return NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{Async: true})
}
