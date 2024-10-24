package domain

import (
	"DDD-OrderingSystem/OrderingApplication/Domain/Event"
	"DDD-OrderingSystem/OrderingApplication/Domain/ValueObject"
	"context"
	"time"
)

// EventSourcing 事件溯源接口
type EventSourcing interface {
	// LoadAggregates 加载聚合根
	LoadAggregates(ctx context.Context, aggregateIds ...interface{}) ([]AggregateRoot, error)
	// SaveAggregate 保存聚合根
	SaveAggregate(ctx context.Context, aggregate AggregateRoot) error
}

// EventSourcingImpl 事件溯源实现
type EventSourcingImpl struct {
	eventStore Event.EventStore
}

// LoadAggregates 加载聚合根
func (e *EventSourcingImpl) LoadAggregates(ctx context.Context, aggregateIds ...interface{}) ([]AggregateRoot, error) {
	// 获取所有相关事件
	events, err := e.eventStore.LoadEvents(ctx, aggregateIds...)
	if err != nil {
		return nil, err
	}

	// 根据事件创建聚合根
	aggregates := make([]AggregateRoot, 0, len(aggregateIds))
	for _, id := range aggregateIds {
		aggregate, err := e.createAggregateFromEvents(id, events)
		if err != nil {
			return nil, err
		}
		aggregates = append(aggregates, aggregate)
	}
	return aggregates, nil
}

// SaveAggregate 保存聚合根
func (e *EventSourcingImpl) SaveAggregate(ctx context.Context, aggregate AggregateRoot) error {
	// 保存事件
	return e.eventStore.SaveEvents(ctx, aggregate.GetEvents()...)
}

// createAggregateFromEvents 从事件创建聚合根
func (e *EventSourcingImpl) createAggregateFromEvents(aggregateId interface{}, events []Event.Event) (AggregateRoot, error) {
	// 创建聚合根实例
	aggregate := NewAggregateRootImpl(aggregateId)

	// 应用事件到聚合根
	for _, event := range events {
		if err := aggregate.ApplyEvents(ctx, event); err != nil {
			return nil, err
		}
	}

	return aggregate, nil
}

// NewAggregateRootImpl 创建聚合根实例
func NewAggregateRootImpl(aggregateId interface{}) *AggregateRootImpl {
	return &AggregateRootImpl{
		AggregateId: aggregateId,
		CreatedTime: time.Now(),
		LastModifiedTime: time.Now(),
	}
}

// GetEvents 获取聚合根的所有事件
func (a *AggregateRootImpl) GetEvents() []Event.Event {
	// 这里应该包含所有已应用的事件
	// 以下代码仅为示例，实际应用中需要根据具体聚合根逻辑来获取事件
	var events []Event.Event
	// events = append(events, a.LastModifiedTime) // 示例事件
	return events
}