package simple_factory

import "log"

type (
	IMessage interface {
		Send(msg string) error
		Receive() string
		MessageType() MessageType
	}

	rabbitMQ struct {
		messageType MessageType
	}

	rocketMQ struct {
		messageType MessageType
	}

	MessageType int
)

const (
	MessageTypeRabbitMQ MessageType = iota + 1
	MessageTypeRocketMQ
)

func NewMessage(mt MessageType) IMessage {
	switch mt {
	case MessageTypeRabbitMQ:
		return &rabbitMQ{
			messageType: MessageTypeRabbitMQ,
		}
	case MessageTypeRocketMQ:
		return &rocketMQ{
			messageType: MessageTypeRocketMQ,
		}
	default:
		return nil
	}
}

// interface implementation
func (r rabbitMQ) Send(msg string) error {
	log.Println("rabbit mq send a message")
	return nil
}

func (r rabbitMQ) Receive() string {
	log.Println("rabbit mq receive a message")
	return ""
}

func (r rabbitMQ) MessageType() MessageType {
	return r.messageType
}

func (r rocketMQ) Send(msg string) error {
	log.Println("rocket mq send a message")
	return nil
}

func (r rocketMQ) Receive() string {
	log.Println("rocket mq receive a message")
	return ""
}

func (r rocketMQ) MessageType() MessageType {
	return r.messageType
}
