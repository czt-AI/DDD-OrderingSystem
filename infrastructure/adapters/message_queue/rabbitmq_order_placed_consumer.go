package rabbitmq

import (
	"DDD-OrderingSystem/Domain/Command"
	"DDD-OrderingSystem/Infrastructure/Adapters/MessageQueue"
	"github.com/streadway/amqp"
)

// RabbitMQOrderPlacedConsumer RabbitMQ订单已下单消费者
type RabbitMQOrderPlacedConsumer struct {
	connection *RabbitMQConnection
}

// NewRabbitMQOrderPlacedConsumer 创建新的RabbitMQ订单已下单消费者实例
func NewRabbitMQOrderPlacedConsumer(connection *RabbitMQConnection) *RabbitMQOrderPlacedConsumer {
	return &RabbitMQOrderPlacedConsumer{
		connection: connection,
	}
}

// Consume 消费订单已下单消息
func (c *RabbitMQOrderPlacedConsumer) Consume() error {
	// 创建队列
	_, err := c.connection.Channel.QueueDeclare(
		"order_placed_queue", // 队列名称
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
		"order_placed_queue", // 队列名称
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
			var event MessageQueue.OrderPlacedIntegrationEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				continue // 忽略无法解析的消息
			}

			// 转换为领域命令并处理
			createOrderCommand := Command.NewOrderCreateCommand(event.CustomerID, event.Items)
			if err := Command.Handle(context.Background(), createOrderCommand); err != nil {
				continue // 忽略处理失败的消息
			}
		}
	}()

	return nil
}
