package simple_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessage(t *testing.T) {
	as := assert.New(t)
	rocketMsg := NewMessage(MessageTypeRocketMQ)
	rabbitMsg := NewMessage(MessageTypeRabbitMQ)

	as.Equal(rocketMsg.MessageType(), MessageTypeRocketMQ)
	as.Equal(rabbitMsg.MessageType(), MessageTypeRabbitMQ)

	rocketMsg.Send("")
	rocketMsg.Receive()

	rabbitMsg.Send("")
	rabbitMsg.Receive()
}
