package messages_broker

var broker MessagesBroker = nil

func NewMessagesBroker() MessagesBroker {
	//uri, ok := os.LookupEnv("RABBITMQ_URI")
	//if !ok {
	//	panic("RABBITMQ_URI env var is required")
	//}

	if broker == nil {
		broker = NewInMemoryMessagesBroker(InMemoryMessagesBrokerConfig{AsyncEvents: false})
		//broker = NewRabbitMQMessagesBroker(RabbitMQMessagesBrokerConfig{
		//	Uri:     uri,
		//	Durable: true,
		//})
	}

	return broker
}
