package rabbitmq

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Infrastructure/Adapters/MessageQueue"
	"github.com/streadway/amqp"
)

// RabbitMQUserCreatedConsumer RabbitMQ用户创建消费者
type RabbitMQUserCreatedConsumer struct {
	connection *RabbitMQConnection
}

// NewRabbitMQUserCreatedConsumer 创建新的RabbitMQ用户创建消费者实例
func NewRabbitMQUserCreatedConsumer(connection *RabbitMQConnection) *RabbitMQUserCreatedConsumer {
	return &RabbitMQUserCreatedConsumer{
		connection: connection,
	}
}

// Consume 消费用户创建消息
func (c *RabbitMQUserCreatedConsumer) Consume() error {
	// 创建队列
	_, err := c.connection.Channel.QueueDeclare(
		"user_created_queue", // 队列名称
		true,                 // 队列持久化
		false,                // 队列非自动删除
		false,                // 队列非独占队列
		false,                // 队列不等待消息
		nil,                  // 队列参数
	)
	if err != nil {
		return err
	}

	// 设置消费者
	messages, err := c.connection.Channel.Consume(
		"user_created_queue", // 队列名称
		"",                   // 消费者标签
		true,                 // 自动确认消息
		false,                // 非独占队列
		false,                // 不等待消息
		false,                // 不排除该消费者
		nil,                  // 参数
	)
	if err != nil {
		return err
	}

	// 处理消息
	go func() {
		for d := range messages {
			var event MessageQueue.UserCreatedIntegrationEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				continue // 忽略无法解析的消息
			}

			// 转换为领域命令并处理
			createUserCommand := Command.NewUserRegisterCommand(event.UserID, event.Username, event.Email, event.Role)
			if err := Command.Handle(context.Background(), createUserCommand); err != nil {
				continue // 忽略处理失败的消息
			}
		}
	}()

	return nil
}
