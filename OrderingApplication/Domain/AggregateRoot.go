package domain

import (
	"context"
	"time"

	"DDD-OrderingSystem/OrderingApplication/Domain/Event"
)

// AggregateRoot 基础聚合根接口
type AggregateRoot interface {
	// ApplyEvents 应用事件
	ApplyEvents(ctx context.Context, events ...Event.Event) error
	// GetAggregateId 获取聚合根ID
	GetAggregateId() interface{}
	// SetAggregateId 设置聚合根ID
	SetAggregateId(interface{})
	// GetCreatedTime 获取创建时间
	GetCreatedTime() time.Time
	// GetLastModifiedTime 获取最后修改时间
	GetLastModifiedTime() time.Time
}

// AggregateRootImpl 聚合根实现结构
type AggregateRootImpl struct {
	AggregateId        interface{}
	CreatedTime        time.Time
	LastModifiedTime   time.Time
	EventStore         Event.EventStore
	aggregateIdSet     bool
	createdTimeSet     bool
	lastModifiedTimeSet bool
}

// ApplyEvents 应用事件
func (a *AggregateRootImpl) ApplyEvents(ctx context.Context, events ...Event.Event) error {
	for _, event := range events {
		if err := a.EventStore.Save(ctx, event); err != nil {
			return err
		}
	}
	return nil
}

// GetAggregateId 获取聚合根ID
func (a *AggregateRootImpl) GetAggregateId() interface{} {
	return a.AggregateId
}

// SetAggregateId 设置聚合根ID
func (a *AggregateRootImpl) SetAggregateId(id interface{}) {
	a.AggregateId = id
	a.aggregateIdSet = true
}

// GetCreatedTime 获取创建时间
func (a *AggregateRootImpl) GetCreatedTime() time.Time {
	return a.CreatedTime
}

// GetLastModifiedTime 获取最后修改时间
func (a *AggregateRootImpl) GetLastModifiedTime() time.Time {
	return a.LastModifiedTime
}