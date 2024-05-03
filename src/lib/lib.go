package lib

import "gym-management/src/lib/messages_broker"

func MessagesBroker() messages_broker.MessagesBroker {
	return messages_broker.NewMessagesBroker()
}
