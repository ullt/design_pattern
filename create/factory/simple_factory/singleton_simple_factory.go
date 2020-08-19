package simple_factory

import (
	"sync"
)

var (
	raMQ   IMessage
	raOnce sync.Once
	roMQ   IMessage
	roOnce sync.Once
)

func newRabbitMQInstance() IMessage {
	if raMQ == nil {
		raOnce.Do(func() {
			raMQ = &rabbitMQ{
				messageType: MessageTypeRabbitMQ,
			}
		})
	}
	return raMQ
}

func newRocketMQInstance() IMessage {
	if roMQ == nil {
		roOnce.Do(func() {
			roMQ = &rocketMQ{
				messageType: MessageTypeRocketMQ,
			}
		})
	}
	return roMQ
}

func NewMessageWithSingleton(mt MessageType) IMessage {
	switch mt {
	case MessageTypeRabbitMQ:
		return newRabbitMQInstance()
	case MessageTypeRocketMQ:
		return newRocketMQInstance()
	default:
		return nil
	}
}
