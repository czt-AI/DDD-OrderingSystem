package domain

import (
	"context"
	"DDD-OrderingSystem/OrderingApplication/Domain/Command"
	"DDD-OrderingSystem/OrderingApplication/Domain/Event"
)

// CommandHandler 命令处理器接口
type CommandHandler interface {
	// Handle 处理命令
	Handle(context.Context, Command.Command) error
}

// CommandHandlerImpl 命令处理器实现
type CommandHandlerImpl struct {
	aggregateRoot AggregateRoot
}

// NewCommandHandler 创建命令处理器实例
func NewCommandHandler(aggregateRoot AggregateRoot) *CommandHandlerImpl {
	return &CommandHandlerImpl{
		aggregateRoot: aggregateRoot,
	}
}

// Handle 处理命令
func (h *CommandHandlerImpl) Handle(ctx context.Context, cmd Command.Command) error {
	// 应用命令转换的事件
	events, err := cmd.ToDomainEvents()
	if err != nil {
		return err
	}

	// 应用事件到聚合根
	if err := h.aggregateRoot.ApplyEvents(ctx, events...); err != nil {
		return err
	}

	return nil
}

// ToDomainEvents 将命令转换为领域事件
type ToDomainEvents interface {
	ToDomainEvents() ([]Event.Event, error)
}

// CommandWithEvents 命令带事件结构
type CommandWithEvents struct {
	Command Command.Command
	Events  []Event.Event
}

// ToDomainEvents 实现ToDomainEvents接口
func (c CommandWithEvents) ToDomainEvents() ([]Event.Event, error) {
	return c.Events, nil
}