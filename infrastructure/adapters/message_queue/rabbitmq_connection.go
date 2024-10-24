package rabbitmq

import (
	"github.com/streadway/amqp"
)

// RabbitMQConnection RabbitMQ连接配置
type RabbitMQConnection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbitMQConnection 创建新的RabbitMQ连接实例
func NewRabbitMQConnection(connectionConfig amqp.Config) (*RabbitMQConnection, error) {
	connection, err := amqp.DialConfig(connectionConfig)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQConnection{
		Connection: connection,
		Channel:    channel,
	}, nil
}

// Close 关闭连接和通道
func (r *RabbitMQConnection) Close() error {
	if err := r.Channel.Close(); err != nil {
		return err
	}
	return r.Connection.Close()
}
